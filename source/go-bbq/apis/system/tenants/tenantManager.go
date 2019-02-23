package tenants

import (
	"database/sql"
	"github.com/google/uuid"
)

// GetTenantByKey gets a tenant by the urlkey 
func GetTenantByKey(db *sql.DB, name string) (Tenant,error) {
	// CACHING is desparately needed here.
	var t Tenant
	db.QueryRow("select id, name, urlkey, isenabled from sys.tenants where urlkey = $1", name).Scan(&t.ID, &t.Name, &t.URLKey, &t.IsEnabled)

	return t,nil
}

// GetTenantByID returns a full tenant object from its guid id.
func GetTenantByID(db *sql.DB, id uuid.UUID) (Tenant, error) {
	var t Tenant
	db.QueryRow("select id, name, urlkey, isenabled from sys.tenants where id = $1", id).Scan(&t.ID, &t.Name, &t.URLKey, &t.IsEnabled)

	return t,nil
}