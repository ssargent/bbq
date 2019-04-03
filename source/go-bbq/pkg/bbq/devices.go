package bbq

import "database/sql"

type DeviceRepository struct {
	database *sql.DB	
}

type DeviceService struct {

}

// NewDeviceRepository creates a fully instantiated 
func NewDeviceRepository(db *sql.DB) *DeviceRepository {
	repo := &DeviceRepository{db}
	return repo
}

func GetDevices()