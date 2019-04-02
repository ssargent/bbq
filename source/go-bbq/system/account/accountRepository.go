package account

import (
	"github.com/google/uuid"
	"github.com/ssargent/go-bbq/internal/config"
	"github.com/ssargent/go-bbq/system"
)

type pgAccountRepository struct {
	config *config.Config
}

/*
	type AccountRepository interface {
	GetById(id uuid.UUID) (*Account, error)
	GetByLogin(accountName string) (*Account, error)
	GetAll() ([]*Account, error)
	Create(account *Account) (*Account, error)
	Update(account *Account) (*Account, error)
	Delete(account *Account) error
}
*/

// NewAccountRepository will create an AccountService
func NewAccountRepository(config *config.Config) system.AccountRepository {
	return &pgAccountRepository{config: config}
}

func (a *pgAccountRepository) GetByID(id uuid.UUID) (*system.Account, error) {
	return nil, nil
}

func (a *pgAccountRepository) GetByLogin(accountName string) (*system.Account, error) {
	return nil, nil
}

func (a *pgAccountRepository) GetAll() ([]*system.Account, error) {
	return nil, nil
}

func (a *pgAccountRepository) Create(account *system.Account) (*system.Account, error) {
	return nil, nil
}

func (a *pgAccountRepository) Update(account *system.Account) (*system.Account, error) {
	return nil, nil
}

func (a *pgAccountRepository) Delete(account *system.Account) error {
	return nil
}
