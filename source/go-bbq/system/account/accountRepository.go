package account

import (
	"github.com/google/uuid"
	"github.com/ssargent/go-bbq/internal/config"
	"github.com/ssargent/go-bbq/system"
)

type pgAccountRepository struct {
	config *config.Config
}

// NewAccountRepository will create an AccountService
func NewAccountRepository(config *config.Config) system.AccountRepository {
	return &pgAccountRepository{config: config}
}

func (a *pgAccountRepository) GetByID(id uuid.UUID) (*system.Account, error) {
	return nil, nil
}

func (a *pgAccountRepository) GetByLogin(accountName string) (*system.Account, error) {

	var account system.Account
	err := a.config.Database.QueryRow("select id, loginname, loginpassword, fullname, email, isenabled, tenantid from sys.accounts where loginname = $1", accountName).Scan(&account.ID, &account.LoginName, &account.LoginPassword, &account.FullName, &account.Email, &account.IsEnabled, &account.TenantID)

	if err != nil {
		return &system.Account{}, err
	}

	return &account, nil
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
