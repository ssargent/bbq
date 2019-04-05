package bbq

import "database/sql"

// DeviceRepository is ...
type DeviceRepository struct {
	database *sql.DB
}

type deviceService struct {
}

// NewDeviceRepository creates a fully instantiated
func NewDeviceRepository(db *sql.DB) *DeviceRepository {
	repo := &DeviceRepository{db}
	return repo
}
