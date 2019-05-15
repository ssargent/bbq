package device

import (
	"fmt"
	"time"

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

func (d *deviceService) GetDevices(tenantID uuid.UUID) ([]bbq.Device, error) {
	cacheKey := fmt.Sprintf("bbq$devices$%s", tenantID.String())

	var devices []bbq.Device

	if err := d.cache.GetItem(cacheKey, &devices); err == nil {
		return devices, nil
	} else {

		devices, err := d.repository.GetByTenantId(tenantID)
		if err != nil {
			return []bbq.Device{}, err
		}

		d.cache.SetItem(cacheKey, devices, time.Minute*10)

		return devices, nil
	}
}

func (d *deviceService) GetDevice(tenantID uuid.UUID, deviceName string) (bbq.Device, error) {
	cacheKey := fmt.Sprintf("bbq$devices$%s$%s", tenantID.String(), deviceName)
	var device bbq.Device

	if err := d.cache.GetItem(cacheKey, &device); err == nil {
		return device, nil
	} else {

		device, err := d.repository.GetDevice(tenantID, deviceName)
		if err != nil {
			return bbq.Device{}, err
		}

		d.cache.SetItem(cacheKey, device, time.Minute*10)

		return device, nil
	}
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
