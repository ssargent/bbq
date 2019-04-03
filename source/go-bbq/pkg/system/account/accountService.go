package account

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/ssargent/go-bbq/internal/config"
	"github.com/ssargent/go-bbq/pkg/system"
)

type accountService struct {
	repository system.AccountRepository
}

// NewAccountService will create an AccountService
func NewAccountService(config *config.Config, repository system.AccountRepository) system.AccountService {
	return &accountService{repository: repository}
}

func (a *accountService) GetAccount(accountName string) (system.Account, error) {
	return system.Account{}, nil
}

func (a *accountService) Login(login string, password string) (system.Account, error) {
	return system.Account{}, nil
}

func (a *accountService) GetAccounts() ([]system.Account, error) {
	return []system.Account{}, nil
}

func (a *accountService) CreateAccount(account system.Account) (system.Account, error) {
	loginAccount, err := a.repository.GetByLogin(account.LoginName)

	if loginAccount != nil {
		return system.Account{}, errors.New("a login with that loginname already exists. please choose another")
	}

	emailAccount, err := a.repository.GetByEmail(account.Email)

	if emailAccount != nil {
		return system.Account{}, errors.New("a login with that email already exists.  please choose another")
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
	return system.Account{}, nil
}

func (a *accountService) DeleteAccount(account system.Account) error {
	return system.Account{}
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MaxCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
