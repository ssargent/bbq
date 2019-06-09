package device

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/ssargent/go-bbq/bbq"
	mock_bbq "github.com/ssargent/go-bbq/bbq/mocks"
	mock_infrastructure "github.com/ssargent/go-bbq/internal/infrastructure/mocks"
	"github.com/stretchr/testify/assert"
)

/*type Device struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TenantID    uuid.UUID `json:"tenantid"`
}*/
func TestGetDevices(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockDeviceRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	deviceService := NewDeviceService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	dev := bbq.Device{
		ID:          1,
		Name:        "My device",
		Description: "My device",
		TenantID:    tenant,
	}
	var returnedDevices []bbq.Device

	cacheKey := fmt.Sprintf("bbq$devices$%s", tenant.String())

	mockRepo.EXPECT().GetByTenantID(tenant).Return([]bbq.Device{dev}, nil).Times(1)
	mockCacheService.EXPECT().GetItem(cacheKey, &returnedDevices).Return(errors.New("not found")).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, []bbq.Device{dev}, time.Minute*10).Return(nil).Times(1)

	deviceService.GetDevices(tenant)
}

func TestGetDevicesWhenCached(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockDeviceRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	deviceService := NewDeviceService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	dev := bbq.Device{
		ID:          1,
		Name:        "My device",
		Description: "My device",
		TenantID:    tenant,
	}
	var returnedDevices []bbq.Device

	cacheKey := fmt.Sprintf("bbq$devices$%s", tenant.String())

	mockCacheService.EXPECT().GetItem(cacheKey, &returnedDevices).Return(nil).Times(1)
	mockRepo.EXPECT().GetByTenantID(tenant).Return([]bbq.Device{dev}, nil).Times(0)
	mockCacheService.EXPECT().SetItem(cacheKey, []bbq.Device{dev}, time.Minute*10).Return(nil).Times(0)

	deviceService.GetDevices(tenant)
}

func TestGetDevice(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockDeviceRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	deviceService := NewDeviceService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	dev := bbq.Device{
		ID:          1,
		Name:        "My device",
		Description: "My device",
		TenantID:    tenant,
	}
	var returnedDevice bbq.Device

	cacheKey := fmt.Sprintf("bbq$devices$%s$%s", tenant.String(), "My device")

	mockRepo.EXPECT().GetDevice(tenant, "My device").Return(dev, nil).Times(1)
	mockCacheService.EXPECT().GetItem(cacheKey, &returnedDevice).Return(errors.New("not found")).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, dev, time.Minute*10).Return(nil).Times(1)

	deviceService.GetDevice(tenant, "My device")
}

func TestCreateDevice(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockDeviceRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	deviceService := NewDeviceService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	dev := bbq.Device{
		Name:        "My Device",
		Description: "My Device",
		TenantID:    tenant,
	}

	notFoundErr := sql.ErrNoRows
	//var returnedDevice bbq.Device

	cacheKey := fmt.Sprintf("bbq$devices$%s$%s", tenant.String(), "My Device")

	mockRepo.EXPECT().GetDevice(tenant, "My Device").Return(bbq.Device{}, notFoundErr).Times(1)
	mockRepo.EXPECT().Create(dev).Return(dev, nil).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, dev, time.Minute*10).Return(nil).Times(1)

	returnedDevice, err := deviceService.CreateDevice(tenant, dev)

	assert.Nil(t, err)
	assert.NotNil(t, returnedDevice)

}

func TestCreateDeviceWhenItAlreadyExists(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockDeviceRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	deviceService := NewDeviceService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	dev := bbq.Device{
		Name:        "My Device",
		Description: "My Device",
		TenantID:    tenant,
	}

	//var returnedDevice bbq.Device

	cacheKey := fmt.Sprintf("bbq$devices$%s$%s", tenant.String(), "My Device")

	mockRepo.EXPECT().GetDevice(tenant, "My Device").Return(dev, nil).Times(1)
	mockRepo.EXPECT().Create(dev).Return(dev, nil).Times(0)
	mockCacheService.EXPECT().SetItem(cacheKey, dev, time.Minute*10).Return(nil).Times(0)

	returnedDevice, err := deviceService.CreateDevice(tenant, dev)

	assert.NotNil(t, err)
	assert.Equal(t, returnedDevice, bbq.Device{})

}

func TestUpdateDevice(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockDeviceRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	deviceService := NewDeviceService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	dev := bbq.Device{
		Name:        "My Device",
		Description: "My Device",
		TenantID:    tenant,
	}

	cacheKey := fmt.Sprintf("bbq$devices$%s$%s", tenant.String(), "My Device")

	mockRepo.EXPECT().GetDevice(tenant, "My Device").Return(dev, nil).Times(1)
	mockRepo.EXPECT().Update(dev).Return(dev, nil).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, dev, time.Minute*10).Return(nil).Times(1)

	returnedDevice, err := deviceService.UpdateDevice(tenant, dev)

	assert.Nil(t, err)
	assert.NotNil(t, returnedDevice)

}

func TestUpdateDeviceWhenDeviceDoesntExist(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockDeviceRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	deviceService := NewDeviceService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	dev := bbq.Device{
		Name:        "My Device",
		Description: "My Device",
		TenantID:    tenant,
	}
	notFoundErr := sql.ErrNoRows
	cacheKey := fmt.Sprintf("bbq$devices$%s$%s", tenant.String(), "My Device")

	mockRepo.EXPECT().GetDevice(tenant, "My Device").Return(bbq.Device{}, notFoundErr).Times(1)
	mockRepo.EXPECT().Update(dev).Return(dev, nil).Times(0)
	mockCacheService.EXPECT().SetItem(cacheKey, dev, time.Minute*10).Return(nil).Times(0)

	returnedDevice, err := deviceService.UpdateDevice(tenant, dev)

	assert.NotNil(t, err)
	assert.Equal(t, returnedDevice, bbq.Device{})

}
