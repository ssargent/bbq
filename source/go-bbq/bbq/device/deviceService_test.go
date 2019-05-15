package device

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/golang/mock/gomock"
	"github.com/ssargent/go-bbq/bbq"
	mock_bbq "github.com/ssargent/go-bbq/bbq/mocks"
	mock_infrastructure "github.com/ssargent/go-bbq/internal/infrastructure/mocks"
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

	mockRepo.EXPECT().GetByTenantId(tenant).Return([]bbq.Device{dev}, nil).Times(1)
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
	mockRepo.EXPECT().GetByTenantId(tenant).Return([]bbq.Device{dev}, nil).Times(0)
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
	var returnedLogin bbq.Device

	cacheKey := fmt.Sprintf("bbq$devices$%s$%s", tenant.String(), "My device")

	mockRepo.EXPECT().GetDevice(tenant, "My device").Return(dev, nil).Times(1)
	mockCacheService.EXPECT().GetItem(cacheKey, &returnedLogin).Return(errors.New("not found")).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, dev, time.Minute*10).Return(nil).Times(1)

	deviceService.GetDevice(tenant, "My device")
}
