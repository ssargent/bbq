package account

import (
	"github.com/google/uuid"
	"github.com/ssargent/bbq/bbq-apiserver/config"
	"github.com/ssargent/bbq/bbq-apiserver/system"
)

type pgAccountRepository struct {
	config *config.Config
}

// NewAccountRepository will create an AccountRepository
func NewAccountRepository(config *config.Config) system.AccountRepository {
	return &pgAccountRepository{config: config}
}

func (a *pgAccountRepository) GetByID(id uuid.UUID) (system.Account, error) {
	var account system.Account
	err := a.config.Database.QueryRow("select id, loginname, loginpassword, fullname, email, isenabled, tenantid from sys.accounts where id = $1", id).Scan(&account.ID, &account.LoginName, &account.LoginPassword, &account.FullName, &account.Email, &account.IsEnabled, &account.TenantID)

	if err != nil {
		return system.Account{}, err
	}

	return account, nil
}

func (a *pgAccountRepository) GetByEmail(email string) (system.Account, error) {

	var account system.Account
	err := a.config.Database.QueryRow("select id, loginname, loginpassword, fullname, email, isenabled, tenantid from sys.accounts where email = $1", email).Scan(&account.ID, &account.LoginName, &account.LoginPassword, &account.FullName, &account.Email, &account.IsEnabled, &account.TenantID)

	if err != nil {
		return system.Account{}, err
	}

	return account, nil
}

func (a *pgAccountRepository) GetByLogin(accountName string) (system.Account, error) {

	var account system.Account
	err := a.config.Database.QueryRow("select id, loginname, loginpassword, fullname, email, isenabled, tenantid from sys.accounts where loginname = $1", accountName).Scan(&account.ID, &account.LoginName, &account.LoginPassword, &account.FullName, &account.Email, &account.IsEnabled, &account.TenantID)

	if err != nil {
		return system.Account{}, err
	}

	return account, nil
}

func (a *pgAccountRepository) GetAll() ([]system.Account, error) {
	return []system.Account{}, nil
}

func (a *pgAccountRepository) Create(account system.Account) (system.Account, error) {
	var createdAccount system.Account
	query := `insert into sys.accounts 
			  (loginname, loginpassword, fullname, email, isenabled, tenantid) 
			  values ($1, $2, $3, $4, $5, $6)
			  returning *`

	err := a.config.Database.
		QueryRow(query, account.LoginName, account.LoginPassword, account.FullName, account.Email, account.IsEnabled, account.TenantID).
		Scan(&createdAccount.ID, &createdAccount.LoginName, &createdAccount.LoginPassword, &createdAccount.FullName, &createdAccount.Email, &createdAccount.IsEnabled, &createdAccount.TenantID)

	if err != nil {
		return system.Account{}, err
	}

	return createdAccount, nil
}

func (a *pgAccountRepository) Update(account system.Account) (system.Account, error) {
	var updatedAccount system.Account
	query := `update sys.accounts set loginpassword = $2, fullname = $3, email = $4, isenabled = $5
			  where id = $1
			  returning *`

	err := a.config.Database.
		QueryRow(query, account.ID, account.LoginPassword, account.FullName, account.Email, account.IsEnabled).
		Scan(&updatedAccount.ID, &updatedAccount.LoginName, &updatedAccount.LoginPassword, &updatedAccount.FullName, &updatedAccount.Email, &updatedAccount.IsEnabled, &updatedAccount.TenantID)

	if err != nil {
		return system.Account{}, err
	}

	return updatedAccount, nil
}

func (a *pgAccountRepository) Delete(account system.Account) error {
	return nil
}
