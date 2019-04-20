package devices

import (
	//	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/cache"
	"github.com/lib/pq"
	"github.com/ssargent/go-bbq/internal/apis/system/tenants"
	"github.com/ssargent/go-bbq/config"
)

// GetAllDevices  returns all devices
func GetAllDevices(config *config.Config, count int, start int) ([]Device, error) {
	rows, err := config.Database.Query(
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
func GetDevice(config *config.Config, deviceID int) (Device, error) {
	var d Device
	config.Database.QueryRow("select id, name, description, tenantid from bbq.devices where id = $1", deviceID).Scan(&d.ID, &d.Name, &d.Description, &d.TenantID)
	return d, nil
}

// CreateDevice creates and returns a device
func CreateDevice(config *config.Config, device Device) (Device, error) {
	// validate that the tenant is ok
	tenant, err := tenants.GetTenantByID2(config, device.TenantID)

	if err != nil {
		return Device{}, err
	}

	// now create the device.
	return createDeviceInternal(config, tenant, device)
}

func createDeviceInternal(config *config.Config, tenant tenants.Tenant, device Device) (Device, error) {
	insertStatement := "insert into bbq.devices (name, description, tenantid) values ($1, $2, $3) returning *"

	var createdDevice Device
	err := config.Database.QueryRow(insertStatement, device.Name, device.Description, tenant.ID).Scan(&createdDevice.ID, &createdDevice.Name, &createdDevice.Description, &createdDevice.TenantID)

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
func DeleteDevice(config *config.Config, deviceID int) error {

	result, err := config.Database.Exec("delete from bbq.devices where id = $1", deviceID)

	if err != nil {
		return err
	}

	if rows, afferr := result.RowsAffected(); rows == 0 || afferr != nil {
		return errors.New("not-found")
	}

	return err
}

// GetTenantDevices returns devices for a given tenant
func GetTenantDevices(config *config.Config, tenantName string) ([]Device, error) {
	tenant, err := tenants.GetTenantByKey2(config, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	devices := []Device{}

	cacheKey := fmt.Sprintf("bbq$devices$%s", tenantName)

	if err := config.Cache.Get(cacheKey, &devices); err == nil {
		return devices, nil
	} else {

		fmt.Println("Found Tenant: ", tenant.ID, tenant.Name, tenant.URLKey)
		rows, err := config.Database.Query(
			"SELECT id, name, description, tenantid FROM bbq.devices  where tenantid = $1", tenant.ID)

		if err != nil {
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			var d Device
			if err := rows.Scan(&d.ID, &d.Name, &d.Description, &d.TenantID); err != nil {
				return nil, err
			}
			devices = append(devices, d)
		}

		config.Cache.Set(&cache.Item{
			Key:        cacheKey,
			Object:     devices,
			Expiration: time.Minute * 10,
		})

		return devices, nil
	}
}

// GetTenantDeviceByName returns a device object given its name.
func GetTenantDeviceByName(config *config.Config, tenantName string, deviceName string) (Device, error) {
	tenant, err := tenants.GetTenantByKey2(config, tenantName)

	if err != nil {
		log.Fatal(err)
	}
	var d Device
	cacheKey := fmt.Sprintf("bbq$devices$%s-%s", tenantName, deviceName)

	if err := config.Cache.Get(cacheKey, &d); err == nil {
		return d, nil
	} else {

		config.Database.QueryRow("select id, name, description, tenantid from bbq.devices where Name = $1 AND tenantid = $2", deviceName, tenant.ID).Scan(&d.ID, &d.Name, &d.Description, &d.TenantID)

		config.Cache.Set(&cache.Item{
			Key:        cacheKey,
			Object:     d,
			Expiration: time.Minute * 10,
		})

		return d, nil

	}
}

// GetTenantDevice gets a specific device for a tenant
func GetTenantDevice(config *config.Config, tenantName string, deviceID int) (Device, error) {
	tenant, err := tenants.GetTenantByKey2(config, tenantName)

	if err != nil {
		log.Fatal(err)
	}
	var d Device

	cacheKey := fmt.Sprintf("bbq$device$%s-%d", tenantName, deviceID)
	if err := config.Cache.Get(cacheKey, &d); err == nil {
		return d, nil
	} else {

		fmt.Println("Found Tenant: ", tenant.ID, tenant.Name, tenant.URLKey, "Device ID", deviceID)

		config.Database.QueryRow("select id, name, description, tenantid from bbq.devices where id = $1 AND tenantid = $2", deviceID, tenant.ID).Scan(&d.ID, &d.Name, &d.Description, &d.TenantID)

		config.Cache.Set(&cache.Item{
			Key:        cacheKey,
			Object:     d,
			Expiration: time.Minute * 10,
		})
		return d, nil
	}
}

// CreateTenantDevice creates a tenant device
func CreateTenantDevice(config *config.Config, tenantName string, device Device) (Device, error) {
	tenant, err := tenants.GetTenantByKey2(config, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	return createDeviceInternal(config, tenant, device)
}

// DeleteTenantDevice deletes a device from a specific tenant.  IT will not delete devices outside that tenant.
func DeleteTenantDevice(config *config.Config, tenantName string, deviceID int) error {
	tenant, err := tenants.GetTenantByKey2(config, tenantName)

	if err != nil {
		return err
	}

	result, err2 := config.Database.Exec("delete from bbq.devices where id = $1 and tenantid = $2", deviceID, tenant.ID)

	if rows, afferr := result.RowsAffected(); rows == 0 || afferr != nil {
		return errors.New("not-found")
	}

	return err2
}
