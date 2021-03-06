package device

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
)

type deviceRepository struct {
	database *sql.DB
}

// NewDeviceRepository will return a repo for DeviceRepository
func NewDeviceRepository(database *sql.DB) bbq.DeviceRepository {
	return &deviceRepository{database: database}
}

func (d *deviceRepository) GetByTenantID(tenantID uuid.UUID) ([]bbq.Device, error) {
	var devices []bbq.Device
	rows, err := d.database.Query(
		"SELECT id, name, description, tenantid, uid FROM bbq.devices  where tenantid = $1", tenantID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var dev bbq.Device
		if err := rows.Scan(&dev.ID, &dev.Name, &dev.Description, &dev.TenantID, &dev.Uid); err != nil {
			return nil, err
		}
		devices = append(devices, dev)
	}

	return devices, nil
}

func (d *deviceRepository) GetByName(tenantID uuid.UUID, deviceName string) (bbq.Device, error) {
	var dev bbq.Device
	query := "select id, name, description, tenantid, uid from bbq.devices where Name = $1 AND tenantid = $2"

	err := d.database.QueryRow(query, deviceName, tenantID).Scan(&dev.ID, &dev.Name, &dev.Description, &dev.TenantID, &dev.Uid)

	if err != nil {
		return bbq.Device{}, fmt.Errorf("Cannot find device %s (tenant: %s) - %s", deviceName, tenantID, err.Error())
	}

	return dev, nil

}

func (d *deviceRepository) GetByID(tenantID uuid.UUID, id uuid.UUID) (bbq.Device, error) {
	var dev bbq.Device
	query := "select id, name, description, tenantid, Uid from bbq.devices where uid = $1 AND tenantid = $2"

	err := d.database.QueryRow(query, id, tenantID).Scan(&dev.ID, &dev.Name, &dev.Description, &dev.TenantID, &dev.Uid)

	if err != nil {
		return bbq.Device{}, err
	}

	return dev, nil

}

func (d *deviceRepository) Create(device bbq.Device) (bbq.Device, error) {
	insertStatement := "insert into bbq.devices (name, description, tenantid) values ($1, $2, $3) returning *"

	var createdDevice bbq.Device
	err := d.database.QueryRow(insertStatement, device.Name, device.Description, device.TenantID).Scan(&createdDevice.ID, &createdDevice.Name, &createdDevice.Description, &createdDevice.TenantID, &createdDevice.Uid)

	if err != nil {
		// There must be a more elegant way of doing this...  but for now...
		if err, ok := err.(*pq.Error); ok {
			// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
			if err.Code.Name() == "unique_violation" {
				return bbq.Device{}, errors.New("a device with that name already exists for your account, please choose a different name")
			}
		}
		return bbq.Device{}, err
	}

	return createdDevice, nil
}

func (d *deviceRepository) Update(device bbq.Device) (bbq.Device, error) {
	var updatedDevice bbq.Device
	query := `update bbq.devices set name = $3, description = $4
			  where id = $1 and TenantID = $2
			  returning *`

	err := d.database.
		QueryRow(query, device.ID, device.TenantID, device.Name, device.Description).
		Scan(&updatedDevice.ID, &updatedDevice.Name, &updatedDevice.Description, &updatedDevice.TenantID)

	if err != nil {
		return bbq.Device{}, err
	}

	return updatedDevice, nil
}

func (d *deviceRepository) Delete(device bbq.Device) error {
	result, err := d.database.Exec("delete from bbq.devices where id = $1 and tenantid = $2", device.ID, device.TenantID)

	if rows, afferr := result.RowsAffected(); rows == 0 || afferr != nil {
		return errors.New("not-found")
	}

	return err
}
