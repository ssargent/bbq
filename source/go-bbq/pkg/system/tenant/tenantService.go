package tenant

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/cache"
	"github.com/ssargent/go-bbq/pkg/config"
	"github.com/ssargent/go-bbq/pkg/system"
)

/*

type TenantRepository interface {
	GetByKey(key string) (Tenant, error)
	Create(tenant Tenant) (Tenant, error)
	//Update(tenant Tenant) (Tenant, error)
	//Delete(tenant Tenant) (Tenant, error)
}

*/

type tenantService struct {
	repository system.TenantRepository
	config     *config.Config
}

// NewTenantService will create an TenantService
func NewTenantService(config *config.Config, repository system.TenantRepository) system.TenantService {
	return &tenantService{repository: repository, config: config}
}

func (t *tenantService) GetByKey(key string) (system.Tenant, error) {
	var theTenant system.Tenant
	cacheKey := fmt.Sprintf("sys$tenant-v2$%s", key)
	if err := t.config.Cache.Get(cacheKey, &theTenant); err == nil {
		return theTenant, nil
	}

	theTenant, err := t.repository.GetByKey(key)

	if err != nil {
		return system.Tenant{}, err
	}

	t.config.Cache.Set(&cache.Item{
		Key:        cacheKey,
		Object:     theTenant,
		Expiration: time.Minute * 10,
	})

	return theTenant, nil
}

func (t *tenantService) CreateTenant(entity system.Tenant) (system.Tenant, error) {

	_, tenantExistsErr := t.repository.GetByKey(entity.URLKey)

	if tenantExistsErr != nil {
		if tenantExistsErr != sql.ErrNoRows {
			return system.Tenant{}, tenantExistsErr
		}
	}

	// create account
	createdTenant, err := t.repository.Create(entity)

	if err != nil {
		return system.Tenant{}, err
	}

	return createdTenant, nil

}

func (t *tenantService) DeleteTenant(entity system.Tenant) error {
	_, tenantExistsErr := t.repository.GetByKey(entity.URLKey)

	if tenantExistsErr != nil {
		if tenantExistsErr == sql.ErrNoRows {
			return errors.New("Tenant not found")
		}
	}
	return t.repository.Delete(entity)
}
