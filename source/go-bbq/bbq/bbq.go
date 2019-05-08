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
	GetDevices(tenantId uuid.UUID) ([]Device, error)
	GetDevice(tenantId uuid.UUID, deviceName string) (Device, error)
	CreateDevice(tenantId uuid.UUID, newDevice Device) (Device, error)
	UpdateDevice(tenantId uuid.UUID, existingDevice Device) (Device, error)
	DeleteDevice(tenantId uuid.UUID, existingDevice Device) error
}

type DeviceRepository interface {
	GetByTenantId(tenantId uuid.UUID) ([]Device, error)
	GetDevice(tenantId uuid.UUID, deviceName string) (Device, error)
	Create(newDevice Device) (Device, error)
	Update(device Device) (Device, error)
	Delete(device Device) error
}
