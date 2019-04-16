package tenant

import (
	"github.com/ssargent/go-bbq/pkg/config"
	"github.com/ssargent/go-bbq/pkg/system"
)

type pgTenantRepository struct {
	config *config.Config
}

// NewTenantRepository will create an TenantRepository
func NewTenantRepository(config *config.Config) system.TenantRepository {
	return &pgTenantRepository{config: config}
}

func (t *pgTenantRepository) GetByKey(key string) (system.Tenant, error) {

	var tenant system.Tenant
	err := t.config.Database.QueryRow("select id, name, urlkey, isenabled from sys.tenants where urlkey = $1", key).Scan(&tenant.ID, &tenant.Name, &tenant.URLKey, &tenant.IsEnabled)

	if err != nil {
		return system.Tenant{}, err
	}

	return tenant, nil
}

func (t *pgTenantRepository) Create(tenant system.Tenant) (system.Tenant, error) {
	query := "insert into sys.tenants (name, urlkey, isenabled) values ($1, $2, $3) returning *"

	var createdTenant system.Tenant
	err := t.config.Database.QueryRow(query, tenant.Name, tenant.URLKey, tenant.IsEnabled).Scan(&createdTenant.ID, &createdTenant.Name, &createdTenant.URLKey, &createdTenant.IsEnabled)

	if err != nil {
		return system.Tenant{}, err
	}

	return createdTenant, nil
}
