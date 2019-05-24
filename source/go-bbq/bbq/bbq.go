package bbq

import "github.com/google/uuid"

//Device is
type Device struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TenantID    uuid.UUID `json:"tenantid"`
}

// DeviceService is the service for devices
type DeviceService interface {
	GetDevices(tenantID uuid.UUID) ([]Device, error)
	GetDevice(tenantID uuid.UUID, deviceName string) (Device, error)
	CreateDevice(tenantID uuid.UUID, newDevice Device) (Device, error)
	UpdateDevice(tenantID uuid.UUID, existingDevice Device) (Device, error)
	DeleteDevice(tenantID uuid.UUID, existingDevice Device) error
}

// DeviceRepository is the repo for Devices
type DeviceRepository interface {
	GetByTenantID(tenantID uuid.UUID) ([]Device, error)
	GetDevice(tenantID uuid.UUID, deviceName string) (Device, error)
	Create(newDevice Device) (Device, error)
	Update(device Device) (Device, error)
	Delete(device Device) error
}
