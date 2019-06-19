package bbq

import "github.com/google/uuid"

//go:generate mockgen  -destination=./mocks/bbq.go -package=mock_bbq github.com/ssargent/bbq/bbq-apiserver/bbq DeviceRepository,MonitorRepository,DeviceService,MonitorService

//Device is
type Device struct {
	ID          int       `json:"id"`
	Uid         uuid.UUID `json:"uid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TenantID    uuid.UUID `json:"tenantid"`
}

//Monitor is
type Monitor struct {
	ID          int       `json:"id"`
	Uid         uuid.UUID `json:"uid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	TenantID    uuid.UUID `json:"tenantid"`
}

// DeviceService is the service for devices
type DeviceService interface {
	GetDevices(tenantID uuid.UUID) ([]Device, error)
	GetDeviceByName(tenantID uuid.UUID, deviceName string) (Device, error)
	GetDeviceByID(tenantId uuid.UUID, id uuid.UUID) (Device, error)
	CreateDevice(tenantID uuid.UUID, newDevice Device) (Device, error)
	UpdateDevice(tenantID uuid.UUID, existingDevice Device) (Device, error)
	DeleteDevice(tenantID uuid.UUID, existingDevice Device) error
}

// DeviceRepository is the repo for Devices
type DeviceRepository interface {
	GetByTenantID(tenantID uuid.UUID) ([]Device, error)
	GetByID(tenantID uuid.UUID, id uuid.UUID) (Device, error)
	GetByName(tenantID uuid.UUID, deviceName string) (Device, error)
	Create(newDevice Device) (Device, error)
	Update(device Device) (Device, error)
	Delete(device Device) error
}

// MonitorService is the service for devices
type MonitorService interface {
	GetMonitors(tenantID uuid.UUID) ([]Monitor, error)
	GetMonitorByID(tenantID uuid.UUID, monitorId uuid.UUID) (Monitor, error)
	GetMonitorByName(tenantID uuid.UUID, name string) (Monitor, error)
	GetMonitorByAddress(tenantID uuid.UUID, address string) (Monitor, error)
	CreateMonitor(tenantID uuid.UUID, entity Monitor) (Monitor, error)
	UpdateMonitor(tenantID uuid.UUID, entity Monitor) (Monitor, error)
	DeleteMonitor(tenantID uuid.UUID, entity Monitor) error
}

// MonitorRepository is the repo for Devices
type MonitorRepository interface {
	GetByTenantID(tenantID uuid.UUID) ([]Monitor, error)
	GetByID(tenantID uuid.UUID, monitorId uuid.UUID) (Monitor, error)
	GetByName(tenantID uuid.UUID, name string) (Monitor, error)
	GetByAddress(tenantID uuid.UUID, address string) (Monitor, error)
	Create(entity Monitor) (Monitor, error)
	Update(entity Monitor) (Monitor, error)
	Delete(entity Monitor) error
}
