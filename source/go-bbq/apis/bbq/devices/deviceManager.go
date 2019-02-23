package devices

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"
	"github.com/ssargent/go-bbq/apis/system/tenants"
)

// GetAllDevices  returns all devices
func GetAllDevices(db *sql.DB, count int, start int) ([]Device, error) {
	rows, err := db.Query(
		"SELECT id, name, description, tenantid FROM bbq.devices LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	devices := []Device{}

	for rows.Next() {
		var d Device
		if err := rows.Scan(&d.ID, &d.Name, &d.Description, &d.TenantID); err != nil {
			return nil, err
		}
		devices = append(devices, d)
	}

	return devices, nil
}

// GetDevice returns a specific device
func GetDevice(db *sql.DB, deviceID int) (Device, error) {
	var d Device
	db.QueryRow("select id, name, description, tenantid from bbq.devices where id = $1", deviceID).Scan(&d.ID, &d.Name, &d.Description, &d.TenantID)
	return d, nil
}

// CreateDevice creates and returns a device
func CreateDevice(db *sql.DB, device Device) (Device, error) {
	// validate that the tenant is ok
	tenant, err := tenants.GetTenantByID(db, device.TenantID)

	if err != nil {
		return Device{}, err
	}

	// now create the device.
	return createDeviceInternal(db, tenant, device)
}

func createDeviceInternal(db *sql.DB, tenant tenants.Tenant, device Device) (Device, error) {
	insertStatement := "insert into bbq.devices (name, description, tenantid) values ($1, $2, $3) returning *"

	var createdDevice Device
	err := db.QueryRow(insertStatement, device.Name, device.Description, tenant.ID).Scan(&createdDevice.ID, &createdDevice.Name, &createdDevice.Description, &createdDevice.TenantID)

	if err != nil {
		// There must be a more elegant way of doing this...  but for now...
		if err, ok := err.(*pq.Error); ok {
			// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
			if err.Code.Name() == "unique_violation" {
				return Device{}, errors.New("a device with that name already exists for your account, please choose a different name")
			}
		}
		return Device{}, err
	}

	return createdDevice, nil
}

// DeleteDevice deletes a device WTSE=1
func DeleteDevice(db *sql.DB, deviceID int) error {

	result, err := db.Exec("delete from bbq.devices where id = $1", deviceID)

	if err != nil {
		return err
	}

	if rows, afferr := result.RowsAffected(); rows == 0 || afferr != nil {
		return errors.New("not-found")
	}

	return err
}

// GetTenantDevices returns devices for a given tenant
func GetTenantDevices(db *sql.DB, tenantName string) ([]Device, error) {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found Tenant: ", tenant.ID, tenant.Name, tenant.URLKey)
	rows, err := db.Query(
		"SELECT id, name, description, tenantid FROM bbq.devices  where tenantid = $1", tenant.ID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	devices := []Device{}

	for rows.Next() {
		var d Device
		if err := rows.Scan(&d.ID, &d.Name, &d.Description, &d.TenantID); err != nil {
			return nil, err
		}
		devices = append(devices, d)
	}

	return devices, nil
}

// GetTenantDeviceByName returns a device object given its name.
func GetTenantDeviceByName(db *sql.DB, tenantName string, deviceName string) (Device, error) {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	var d Device
	db.QueryRow("select id, name, description, tenantid from bbq.devices where Name = $1 AND tenantid = $2", deviceName, tenant.ID).Scan(&d.ID, &d.Name, &d.Description, &d.TenantID)

	return d, nil
}

// GetTenantDevice gets a specific device for a tenant
func GetTenantDevice(db *sql.DB, tenantName string, deviceID int) (Device, error) {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found Tenant: ", tenant.ID, tenant.Name, tenant.URLKey, "Device ID", deviceID)

	var d Device
	db.QueryRow("select id, name, description, tenantid from bbq.devices where id = $1 AND tenantid = $2", deviceID, tenant.ID).Scan(&d.ID, &d.Name, &d.Description, &d.TenantID)

	return d, nil
}

// CreateTenantDevice creates a tenant device
func CreateTenantDevice(db *sql.DB, tenantName string, device Device) (Device, error) {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	return createDeviceInternal(db, tenant, device)
}

// DeleteTenantDevice deletes a device from a specific tenant.  IT will not delete devices outside that tenant.
func DeleteTenantDevice(db *sql.DB, tenantName string, deviceID int) error {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		return err
	}

	result, err2 := db.Exec("delete from bbq.devices where id = $1 and tenantid = $2", deviceID, tenant.ID)

	if rows, afferr := result.RowsAffected(); rows == 0 || afferr != nil {
		return errors.New("not-found")
	}

	return err2
}
