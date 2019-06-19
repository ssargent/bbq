package account

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
	"github.com/ssargent/bbq/bbq-apiserver/system"
)

type accountService struct {
	repository system.AccountRepository
	cache      infrastructure.CacheService
}

// NewAccountService will create an AccountService
func NewAccountService(cache infrastructure.CacheService, repository system.AccountRepository) system.AccountService {
	return &accountService{repository: repository, cache: cache}
}

func (a *accountService) GetAccount(loginName string) (system.Account, error) {
	cacheKey := fmt.Sprintf("system$accounts$%s", loginName)
	var account system.Account

	if err := a.cache.GetItem(cacheKey, &account); err == nil {
		return account, nil
	} else {
		login, err := a.repository.GetByLogin(loginName)

		if err != nil {
			return system.Account{}, err
		}

		a.cache.SetItem(cacheKey, login, time.Minute*10)

		return login, nil
	}
}

func (a *accountService) Login(login string, password string) (system.Account, error) {
	account, err := a.GetAccount(login)

	if err != nil || !comparePasswords(account.LoginPassword, []byte(password)) {
		return system.Account{}, errors.New("account not found")
	}

	account.LoginPassword = ""
	return account, nil
}

func (a *accountService) GetAccounts() ([]system.Account, error) {
	return []system.Account{}, nil
}

func (a *accountService) CreateAccount(account system.Account) (system.Account, error) {
	existingLogin, err := a.repository.GetByLogin(account.LoginName)
	if err != nil {
		if err != sql.ErrNoRows {
			return system.Account{}, err
		}
	}
	if !existingLogin.Empty() {
		return system.Account{}, errors.New("LoginName already exists")
	}

	existingEmail, err := a.repository.GetByEmail(account.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			return system.Account{}, err
		}
	}
	if !existingEmail.Empty() {
		return system.Account{}, errors.New("Email Already exists")
	}

	// encrypt password
	account.LoginPassword = hashAndSalt([]byte(account.LoginPassword))
	// create account
	createdAccount, err := a.repository.Create(account)
	if err != nil {
		return system.Account{}, err
	}

	// clear password before sending or caching
	createdAccount.LoginPassword = ""

	return createdAccount, nil
}

func (a *accountService) UpdateAccount(account system.Account) (system.Account, error) {
	existingEmailAccount, err := a.repository.GetByEmail(account.Email)

	if err != nil {
		if err != sql.ErrNoRows {
			return system.Account{}, err
		}
	}

	if !existingEmailAccount.Empty() {
		return system.Account{}, errors.New("a login with that email already exists.  please choose another")
	}

	// encrypt password
	account.LoginPassword = hashAndSalt([]byte(account.LoginPassword))

	// create account
	updatedAccount, err := a.repository.Update(account)

	if err != nil {
		return system.Account{}, err
	}

	// clear password before sending or caching
	updatedAccount.LoginPassword = ""

	return updatedAccount, nil
}

func (a *accountService) DeleteAccount(account system.Account) error {
	//todo check permissions...
	return a.repository.Delete(account)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
