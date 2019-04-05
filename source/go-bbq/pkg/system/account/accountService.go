package account

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/ssargent/go-bbq/pkg/config"
	"github.com/ssargent/go-bbq/pkg/system"
)

type accountService struct {
	repository system.AccountRepository
}

// NewAccountService will create an AccountService
func NewAccountService(config *config.Config, repository system.AccountRepository) system.AccountService {
	return &accountService{repository: repository}
}

func (a *accountService) GetAccount(loginName string) (system.Account, error) {
	login, err := a.repository.GetByLogin(loginName)

	if err != nil {
		return system.Account{}, err
	}

	return login, nil
}

func (a *accountService) CreateToken(account system.Account) string {
	return "" // finish with token
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
	_, loginExistsErr := a.repository.GetByLogin(account.LoginName)

	if loginExistsErr != nil {
		if loginExistsErr != sql.ErrNoRows {
			return system.Account{}, loginExistsErr
		}

	}

	_, emailExistsErr := a.repository.GetByEmail(account.Email)

	if emailExistsErr != nil {
		if emailExistsErr != sql.ErrNoRows {
			return system.Account{}, emailExistsErr
		}
	}

	fmt.Println("About to encrypt pw")
	// encrypt password
	account.LoginPassword = hashAndSalt([]byte(account.LoginPassword))
	fmt.Println("Finished with encrypt pw")
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
	_, emailExistsErr := a.repository.GetByEmail(account.Email)

	if emailExistsErr != nil {
		if emailExistsErr == sql.ErrNoRows {
			return system.Account{}, errors.New("a login with that email already exists.  please choose another")
		}

		return system.Account{}, emailExistsErr
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
