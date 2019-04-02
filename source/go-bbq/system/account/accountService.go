package account

import (
	"github.com/ssargent/go-bbq/internal/config"
	"github.com/ssargent/go-bbq/system"
)

type accountService struct {
	repository *system.AccountRepository
}

// NewAccountService will create an AccountService
func NewAccountService(config *config.Config, repository *system.AccountRepository) system.AccountService {
	return &accountService{repository: repository}
}

func (a *accountService) GetAccount(accountName string) (*system.Account, error) {
	return nil, nil
}

func (a *accountService) Login(login string, password string) (*system.Account, error) {
	return nil, nil
}

func (a *accountService) GetAccounts() ([]*system.Account, error) {
	return nil, nil
}

func (a *accountService) CreateAccount(account *system.Account) (*system.Account, error) {
	return nil, nil
}

func (a *accountService) UpdateAccount(account *system.Account) (*system.Account, error) {
	return nil, nil
}

func (a *accountService) DeleteAccount(account *system.Account) error {
	return nil
}
