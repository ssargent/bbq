package tenants

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/go-redis/cache"
	"github.com/ssargent/go-bbq/internal/config"
)

// Config is internal config representational

// GetTenantByKey gets a tenant by the urlkey
func GetTenantByKey(db *sql.DB, name string) (Tenant, error) {
	// CACHING is desparately needed here.
	var t Tenant
	db.QueryRow("select id, name, urlkey, isenabled from sys.tenants where urlkey = $1", name).Scan(&t.ID, &t.Name, &t.URLKey, &t.IsEnabled)

	return t, nil
}

// GetTenantByID returns a full tenant object from its guid id.
func GetTenantByID(db *sql.DB, id uuid.UUID) (Tenant, error) {
	var t Tenant
	db.QueryRow("select id, name, urlkey, isenabled from sys.tenants where id = $1", id).Scan(&t.ID, &t.Name, &t.URLKey, &t.IsEnabled)

	return t, nil
}

// These methods are temporary while the caching strategy is validated.

// GetTenantByKey2 gets a tenant by the urlkey
func GetTenantByKey2(config *config.Config, name string) (Tenant, error) {
	// CACHING is desparately needed here.
	var t Tenant
	cacheKey := fmt.Sprintf("sys$tenant$%s", name)
	if err := config.Cache.Get(cacheKey, &t); err == nil {
		return t, nil
	}

	config.Database.QueryRow("select id, name, urlkey, isenabled from sys.tenants where urlkey = $1", name).Scan(&t.ID, &t.Name, &t.URLKey, &t.IsEnabled)
	config.Cache.Set(&cache.Item{
		Key:        cacheKey,
		Object:     t,
		Expiration: time.Minute * 10,
	})
	return t, nil

}

// GetTenantByID2 returns a full tenant object from its guid id.
func GetTenantByID2(config *config.Config, id uuid.UUID) (Tenant, error) {
	var t Tenant
	config.Database.QueryRow("select id, name, urlkey, isenabled from sys.tenants where id = $1", id).Scan(&t.ID, &t.Name, &t.URLKey, &t.IsEnabled)

	return t, nil
}
