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
	GetDevices(tenantName string) ([]*Device, error)
	GetDevice(tenantName string, deviceName string) (*Device, error)
	CreateDevice(tenantName string, newDevice *Device) (*Device, error)
	UpdateDevice(tenantName string, existingDevice *Device) (*Device, error)
	DeleteDevice(tenantName string, existingDevice *Device) error
}
