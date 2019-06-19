package monitors

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"
	"github.com/ssargent/bbq/bbq-apiserver/internal/apis/system/tenants"
)

func getTenantMonitors(db *sql.DB, tenantName string) ([]Monitor, error) {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found Tenant: ", tenant.ID, tenant.Name, tenant.URLKey)
	rows, err := db.Query(
		"SELECT id, name, description, address, tenantid FROM bbq.monitors where tenantid = $1", tenant.ID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	monitors := []Monitor{}

	for rows.Next() {
		var m Monitor
		if err := rows.Scan(&m.ID, &m.Name, &m.Description, &m.Address, &m.TenantID); err != nil {
			return nil, err
		}
		monitors = append(monitors, m)
	}

	return monitors, nil
}

// GetTenantMonitorByName returns a tenant monitor given its name
func GetTenantMonitorByName(db *sql.DB, tenantName string, monitorName string) (Monitor, error) {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found Tenant: ", tenant.ID, tenant.Name, tenant.URLKey)

	var m Monitor
	err = db.QueryRow("select id, name, description, address, tenantid from bbq.monitors where Name = $1 and tenantid = $2", monitorName, tenant.ID).Scan(&m.ID, &m.Name, &m.Description, &m.Address, &m.TenantID)

	if err != nil {
		return Monitor{}, err
	}

	return m, nil
}

func getTenantMonitor(db *sql.DB, tenantName string, monitorID int) (Monitor, error) {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found Tenant: ", tenant.ID, tenant.Name, tenant.URLKey)

	var m Monitor
	err = db.QueryRow("select id, name, description, address, tenantid from bbq.monitors where id = $1 and tenantid = $2", monitorID, tenant.ID).Scan(&m.ID, &m.Name, &m.Description, &m.Address, &m.TenantID)

	if err != nil {
		return Monitor{}, err
	}

	return m, nil
}

// CreateTenantDevice creates a tenant device
func createTenantMonitor(db *sql.DB, tenantName string, monitor Monitor) (Monitor, error) {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	return createMonitorInternal(db, tenant, monitor)
}

func createMonitorInternal(db *sql.DB, tenant tenants.Tenant, monitor Monitor) (Monitor, error) {
	insertStatement := "insert into bbq.monitors (name, description, address, tenantid) values ($1, $2, $3, $4) returning *"

	var createdMonitor Monitor
	err := db.QueryRow(insertStatement, monitor.Name, monitor.Description, monitor.Address, tenant.ID).Scan(&createdMonitor.ID, &createdMonitor.Name, &createdMonitor.Description, &createdMonitor.Address, &createdMonitor.TenantID)

	if err != nil {
		// There must be a more elegant way of doing this...  but for now...
		if err, ok := err.(*pq.Error); ok {
			// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
			if err.Code.Name() == "unique_violation" {
				return Monitor{}, errors.New("a monitor with that name already exists for your account, please choose a different name")
			}
		}
		return Monitor{}, err
	}

	return createdMonitor, nil
}

func deleteTenantMonitor(db *sql.DB, tenantName string, monitorID int) error {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		return err
	}

	result, err2 := db.Exec("delete from bbq.monitors where id = $1 and tenantid = $2", monitorID, tenant.ID)

	if rows, afferr := result.RowsAffected(); rows == 0 || afferr != nil {
		return errors.New("not-found")
	}

	return err2
}
