package system

import "github.com/google/uuid"

type Tenant struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	URLKey    string    `json:"urlkey"`
	IsEnabled bool      `json:"isenabled"`
}

type Account struct {
	ID            uuid.UUID `json:"id"`
	LoginName     string    `json:"loginName"`
	LoginPassword string    `json:"loginPassword"`
	FullName      string    `json:"fullName"`
	Email         string    `json:"email"`
	IsEnabled     bool      `json:"isenabled"`
	TenantID      uuid.UUID `json:"tenantid"`
}

type TenantService interface {
	GetTenant(tenantName string) (*Tenant, error)
	GetTenants() ([]*Tenant, error)
	CreateTenant(tenant *Tenant) (*Tenant, error)
	UpdateTenant(tenant *Tenant) (*Tenant, error)
	DeleteTenant(tenant *Tenant) error
}

type AccountService interface {
	GetAccount(accountName string) (Account, error)
	Login(login string, password string) (Account, error)
	GetAccounts() ([]Account, error)
	CreateAccount(account Account) (Account, error)
	UpdateAccount(account Account) (Account, error)
	DeleteAccount(account Account) error
}

type AccountRepository interface {
	GetByEmail(email string) (Account, error)
	GetByID(id uuid.UUID) (Account, error)
	GetByLogin(accountName string) (Account, error)
	GetAll() ([]Account, error)
	Create(account Account) (Account, error)
	Update(account Account) (Account, error)
	Delete(account Account) error
}
