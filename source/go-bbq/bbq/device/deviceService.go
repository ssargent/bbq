package device

import (
	"github.com/google/uuid"

	"github.com/ssargent/go-bbq/bbq"
	"github.com/ssargent/go-bbq/internal/infrastructure"
)

type deviceService struct {
	repository bbq.DeviceRepository
	cache      infrastructure.CacheService
}

// NewDeviceService will create an DeviceService
func NewDeviceService(cache infrastructure.CacheService, repository bbq.DeviceRepository) bbq.DeviceService {
	return &deviceService{repository: repository, cache: cache}
}

/*type DeviceService interface {
	GetDevices(tenantId uuid.UUID) ([]Device, error)
	GetDevice(tenantId uuid.UUID, deviceName string) (Device, error)
	CreateDevice(tenantId uuid.UUID, newDevice Device) (Device, error)
	UpdateDevice(tenantId uuid.UUID, existingDevice Device) (Device, error)
	DeleteDevice(tenantId uuid.UUID, existingDevice Device) error
}*/

func (d *deviceService) GetDevices(tenantID uuid.UUID) ([]bbq.Device, error) {
	devices, err := d.repository.GetByTenantId(tenantID)
	if err != nil {
		return []bbq.Device{}, err
	}

	return devices, nil
}

func (d *deviceService) GetDevice(tenantID uuid.UUID, deviceName string) (bbq.Device, error) {
	device, err := d.repository.GetDevice(tenantID, deviceName)
	if err != nil {
		return bbq.Device{}, err
	}

	return device, nil
}

func (d *deviceService) CreateDevice(tenantID uuid.UUID, newDevice bbq.Device) (bbq.Device, error) {
	newDevice.TenantID = tenantID
	device, err := d.repository.Create(newDevice)
	if err != nil {
		return bbq.Device{}, err
	}

	return device, nil
}

func (d *deviceService) UpdateDevice(tenantID uuid.UUID, existingDevice bbq.Device) (bbq.Device, error) {
	existingDevice.TenantID = tenantID
	device, err := d.repository.Update(existingDevice)
	if err != nil {
		return bbq.Device{}, err
	}

	return device, nil

}

func (d *deviceService) DeleteDevice(tenantID uuid.UUID, existingDevice bbq.Device) error {
	d.repository.Delete(existingDevice)
	return nil
}
