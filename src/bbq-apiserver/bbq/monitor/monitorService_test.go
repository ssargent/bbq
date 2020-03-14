package monitor

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	mock_bbq "github.com/ssargent/bbq/bbq-apiserver/bbq/mocks"
	mock_infrastructure "github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetMonitors(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockMonitorRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	monitorService := NewMonitorService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	mon := bbq.Monitor{
		ID:          1,
		Name:        "My Monitor",
		Description: "My Monitor",
		Address:     "deadbeef",
		TenantID:    tenant,
	}
	var returnedMonitors []bbq.Monitor

	cacheKey := fmt.Sprintf("bbq$monitors$%s", tenant.String())

	mockRepo.EXPECT().GetByTenantID(tenant).Return([]bbq.Monitor{mon}, nil).Times(1)
	mockCacheService.EXPECT().GetItem(cacheKey, &returnedMonitors).Return(errors.New("not found")).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, []bbq.Monitor{mon}, time.Minute*10).Return(nil).Times(1)

	monitorService.GetMonitors(tenant)
}

func TestGetMonitorsWhenCached(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockMonitorRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	monitorService := NewMonitorService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	mon := bbq.Monitor{
		ID:          1,
		Name:        "My Monitor",
		Description: "My Monitor",
		Address:     "deadbeef",
		TenantID:    tenant,
	}
	var returnedMonitors []bbq.Monitor

	cacheKey := fmt.Sprintf("bbq$monitors$%s", tenant.String())

	mockCacheService.EXPECT().GetItem(cacheKey, &returnedMonitors).Return(nil).Times(1)
	mockRepo.EXPECT().GetByTenantID(tenant).Return([]bbq.Monitor{mon}, nil).Times(0)
	mockCacheService.EXPECT().SetItem(cacheKey, []bbq.Monitor{mon}, time.Minute*10).Return(nil).Times(0)

	monitorService.GetMonitors(tenant)
}

func TestGetMonitorByName(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockMonitorRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	monitorService := NewMonitorService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	mon := bbq.Monitor{
		ID:          1,
		Name:        "My Monitor",
		Description: "My Monitor",
		Address:     "deadbeef",
		TenantID:    tenant,
	}
	var returnedMonitor bbq.Monitor

	cacheKey := fmt.Sprintf("bbq$monitors$%s$%s", tenant.String(), "My Monitor")

	mockRepo.EXPECT().GetByName(tenant, "My Monitor").Return(mon, nil).Times(1)
	mockCacheService.EXPECT().GetItem(cacheKey, &returnedMonitor).Return(errors.New("not found")).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, mon, time.Minute*10).Return(nil).Times(1)

	monitorService.GetMonitorByName(tenant, "My Monitor")
}

func TestCreateMonitor(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockMonitorRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	monitorService := NewMonitorService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	mon := bbq.Monitor{
		Name:        "My Monitor",
		Description: "My Monitor",
		Address:     "deadbeef",
		TenantID:    tenant,
	}

	cacheKey := fmt.Sprintf("bbq$monitors$%s$%s", tenant.String(), "My Monitor")
	tenantMonitorsCacheKey := fmt.Sprintf("bbq$monitors$%s", tenant.String())

	notFoundErr := sql.ErrNoRows
	//var returnedDevice bbq.Device

	mockRepo.EXPECT().GetByName(tenant, "My Monitor").Return(bbq.Monitor{}, notFoundErr).Times(1)
	mockRepo.EXPECT().Create(mon).Return(mon, nil).Times(1)
	mockCacheService.EXPECT().RemoveItem(tenantMonitorsCacheKey).Return(nil).Times(1)

	mockCacheService.EXPECT().SetItem(cacheKey, mon, time.Minute*10).Return(nil).Times(1)

	returnedMonitor, err := monitorService.CreateMonitor(tenant, mon)

	assert.Nil(t, err)
	assert.NotNil(t, returnedMonitor)

}

func TestCreateMonitorWhenItAlreadyExists(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockMonitorRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	monitorService := NewMonitorService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	mon := bbq.Monitor{
		Name:        "My Monitor",
		Description: "My Monitor",
		Address:     "deadbeef",
		TenantID:    tenant,
	}
	//var returnedMonitor bbq.Monitor

	cacheKey := fmt.Sprintf("bbq$monitors$%s$%s", tenant.String(), "My Monitor")

	//var returnedDevice bbq.Device
	mockRepo.EXPECT().GetByName(tenant, "My Monitor").Return(mon, nil).Times(1)
	mockRepo.EXPECT().Create(mon).Return(mon, nil).Times(0)
	mockCacheService.EXPECT().SetItem(cacheKey, mon, time.Minute*10).Return(nil).Times(0)

	returnedMonitor, err := monitorService.CreateMonitor(tenant, mon)

	assert.NotNil(t, err)
	assert.Equal(t, returnedMonitor, bbq.Monitor{})
	assert.Equal(t, err, errors.New("A monitor with that name already exists for your tenant"))

}

func TestUpdateMonitor(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockMonitorRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	monitorService := NewMonitorService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	mon := bbq.Monitor{
		Name:        "My Monitor",
		Description: "My Monitor",
		Address:     "deadbeef",
		TenantID:    tenant,
	}

	cacheKey := fmt.Sprintf("bbq$monitors$%s$%s", tenant.String(), "My Monitor")

	mockRepo.EXPECT().GetByName(tenant, "My Monitor").Return(mon, nil).Times(1)
	mockRepo.EXPECT().Update(mon).Return(mon, nil).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, mon, time.Minute*10).Return(nil).Times(1)

	returnedMonitor, err := monitorService.UpdateMonitor(tenant, mon)

	assert.Nil(t, err)
	assert.NotNil(t, returnedMonitor)

}

func TestUpdateMonitorWhenMonitorDoesntExist(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockMonitorRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	monitorService := NewMonitorService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return
	}

	mon := bbq.Monitor{
		Name:        "My Monitor",
		Description: "My Monitor",
		Address:     "deadbeef",
		TenantID:    tenant,
	}

	notFoundErr := sql.ErrNoRows
	cacheKey := fmt.Sprintf("bbq$monitors$%s$%s", tenant.String(), "My Monitor")

	mockRepo.EXPECT().GetByName(tenant, "My Monitor").Return(bbq.Monitor{}, notFoundErr).Times(1)
	mockRepo.EXPECT().Update(mon).Return(bbq.Monitor{}, nil).Times(0)
	mockCacheService.EXPECT().SetItem(cacheKey, mon, time.Minute*10).Return(nil).Times(0)

	returnedMonitor, err := monitorService.UpdateMonitor(tenant, mon)

	assert.NotNil(t, err)
	assert.Equal(t, bbq.Monitor{}, returnedMonitor)

}
