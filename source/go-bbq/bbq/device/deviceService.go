package device

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ssargent/go-bbq/internal/infrastructure"
	"github.com/ssargent/go-bbq/bbq"
)

type deviceService struct {
	repository bbq.DeviceRepository
	cache      infrastructure.CacheService
}

// NewAccountService will create an AccountService
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


func(d *deviceService) GetDevices(tenantId uuid.UUID) ([]Device, error) {
	return d.repository.GetByTenantId(tenantId)
}

func(d *deviceService) GetDevice(tenantId uuid.UUID, deviceName string) (Device, error) {
	return d.repository.GetDevice(tenantId, deviceName)
}

func(d *deviceService) CreateDevice(tenantId uuid.UUID, newDevice Device) (Device, error) {
	newDevice.TenantID = tenantId
	return d.repository.Create(newDevice)
}

func(d *deviceService) UpdateDevice(tenantId uuid.UUID, existingDevice Device) (Device, error) {
	existingDevice.TenantID = tenantId
	return d.repository.Update(existingDevice)
}

func (d *deviceService) DeleteDevice(tenantId uuid.UUID, existingDevice Device) error {
	return d.repository.Delete(existingDevice)
}