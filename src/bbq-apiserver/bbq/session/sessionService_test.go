package session

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	mock_bbq "github.com/ssargent/bbq/bbq-apiserver/bbq/mocks"
	mock_infrastructure "github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure/mocks"
	"github.com/stretchr/testify/assert"
)

func getSession(id int, tenant uuid.UUID, uid uuid.UUID) bbq.Session {
	nt := pq.NullTime{}

	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	return bbq.Session{
		ID:          id,
		Name:        "Pulled Pork",
		Description: "Pulled Pork",
		Subject:     "Pulled Pork",
		//Type:        "Pulled Pork",
		Weight:    9.2,
		Device:    "Large Big Green Egg",
		Monitor:   "My Great Monitor",
		StartTime: t1,
		EndTime:   nt,
		TenantID:  tenant,
		UID:       uid,
	}
}

func getSessionRecord(id int, tenant uuid.UUID, uid uuid.UUID) bbq.SessionRecord {
	nt := pq.NullTime{}

	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	return bbq.SessionRecord{
		ID:          id,
		DeviceID:    2,
		MonitorID:   2,
		Name:        "Pulled Pork",
		Description: "Pulled Pork",
		StartTime:   t1,
		SubjectID:   2,
		Weight:      9.2,
		TenantID:    tenant,
		UID:         uid,
		EndTime:     nt,
	}
}

func createUnitOfWork(c *gomock.Controller) bbq.BBQUnitOfWork {
	var unitofwork bbq.BBQUnitOfWork

	unitofwork.Monitor = mock_bbq.NewMockMonitorRepository(c)
	unitofwork.Device = mock_bbq.NewMockDeviceRepository(c)
	unitofwork.Subject = mock_bbq.NewMockSubjectRepository(c)
	unitofwork.Session = mock_bbq.NewMockSessionRepository(c)

	return unitofwork
}

func TestGetSessions(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	unitOfWork := createUnitOfWork(mockCtrl)

	mockRepo := unitOfWork.Session.(*mock_bbq.MockSessionRepository)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	mockDeviceService := mock_bbq.NewMockDeviceService(mockCtrl)
	mockMonitorService := mock_bbq.NewMockMonitorService(mockCtrl)
	mockSubjectService := mock_bbq.NewMockSubjectService(mockCtrl)
	sessionService := NewSessionService(mockCacheService, unitOfWork, mockDeviceService, mockMonitorService, mockSubjectService)

	tenant, err := uuid.NewUUID()
	dev := bbq.Device{
		Name:        "Large Big Green Egg",
		Description: "My Device",
		TenantID:    tenant,
		ID:          2,
	}

	mon := bbq.Monitor{
		Name:        "My Great Monitor",
		Description: "My Great Monitor",
		TenantID:    tenant,
		ID:          2,
	}

	sub := bbq.Subject{
		Name:        "Pulled Pork",
		Description: "Pulled Pork",
		TenantID:    tenant,
		ID:          2,
	}
	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	sessionid, err := uuid.NewUUID()

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	session := getSession(1, tenant, sessionid)
	sessionRecord := getSessionRecord(1, tenant, sessionid)

	var returnedSessions []bbq.Session

	cacheKey := fmt.Sprintf("bbq$sessions$%s", tenant.String())

	mockRepo.EXPECT().GetByTenantID(tenant).Return([]bbq.SessionRecord{sessionRecord}, nil).Times(1)
	mockDeviceService.EXPECT().GetDeviceByID(tenant, dev.Uid).Return(dev, nil).Times(1)
	mockMonitorService.EXPECT().GetMonitorByID(tenant, mon.Uid).Return(mon, nil).Times(1)
	mockSubjectService.EXPECT().GetSubjectByID(tenant, sub.Uid).Return(sub, nil).Times(1)

	mockCacheService.EXPECT().GetItem(cacheKey, &returnedSessions).Return(errors.New("not found")).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, []bbq.Session{session}, time.Minute*10).Return(nil).Times(1)

	sessions, err := sessionService.GetSessions(tenant)

	assert.NotNil(t, sessions)
	assert.NotEmpty(t, sessions)
	assert.ElementsMatch(t, []bbq.Session{session}, sessions)
	assert.Nil(t, err)

}

func TestGetCachedSessions(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	unitOfWork := createUnitOfWork(mockCtrl)

	mockRepo := unitOfWork.Session.(*mock_bbq.MockSessionRepository)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	mockDeviceService := mock_bbq.NewMockDeviceService(mockCtrl)
	mockMonitorService := mock_bbq.NewMockMonitorService(mockCtrl)
	mockSubjectService := mock_bbq.NewMockSubjectService(mockCtrl)
	sessionService := NewSessionService(mockCacheService, unitOfWork, mockDeviceService, mockMonitorService, mockSubjectService)

	tenant, err := uuid.NewUUID()

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	sessionid, err := uuid.NewUUID()

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	session := getSession(1, tenant, sessionid)
	sessionRecord := getSessionRecord(1, tenant, sessionid)

	var returnedSessions []bbq.Session

	cacheKey := fmt.Sprintf("bbq$sessions$%s", tenant.String())

	mockRepo.EXPECT().GetByTenantID(tenant).Return([]bbq.SessionRecord{sessionRecord}, nil).Times(0)
	mockCacheService.EXPECT().GetItem(cacheKey, &returnedSessions).Return(nil).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, []bbq.Session{session}, time.Minute*10).Return(nil).Times(0)
	//   returnedSessions = []bbq.Session{session}
	sessionService.GetSessions(tenant)

}

func TestGetSessionByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	unitOfWork := createUnitOfWork(mockCtrl)

	mockRepo := unitOfWork.Session.(*mock_bbq.MockSessionRepository)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	mockDeviceService := mock_bbq.NewMockDeviceService(mockCtrl)
	mockMonitorService := mock_bbq.NewMockMonitorService(mockCtrl)
	mockSubjectService := mock_bbq.NewMockSubjectService(mockCtrl)
	sessionService := NewSessionService(mockCacheService, unitOfWork, mockDeviceService, mockMonitorService, mockSubjectService)

	tenant, err := uuid.NewUUID()
	dev := bbq.Device{
		Name:        "Large Big Green Egg",
		Description: "My Device",
		TenantID:    tenant,
		ID:          2,
	}

	mon := bbq.Monitor{
		Name:        "My Great Monitor",
		Description: "My Great Monitor",
		TenantID:    tenant,
		ID:          2,
	}

	sub := bbq.Subject{
		Name:        "Pulled Pork",
		Description: "Pulled Pork",
		TenantID:    tenant,
		ID:          2,
	}
	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	sessionid, err := uuid.NewUUID()

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	session := getSession(1, tenant, sessionid)
	sessionRecord := getSessionRecord(1, tenant, sessionid)

	var returnedSession bbq.Session

	cacheKey := fmt.Sprintf("bbq$sessions$%s$%s", tenant.String(), sessionid.String())

	mockRepo.EXPECT().GetByID(tenant, sessionid).Return(sessionRecord, nil).Times(1)
	mockDeviceService.EXPECT().GetDeviceByID(tenant, dev.Uid).Return(dev, nil).Times(1)
	mockMonitorService.EXPECT().GetMonitorByID(tenant, mon.Uid).Return(mon, nil).Times(1)
	mockSubjectService.EXPECT().GetSubjectByID(tenant, sub.Uid).Return(sub, nil).Times(1)

	mockCacheService.EXPECT().GetItem(cacheKey, &returnedSession).Return(errors.New("not found")).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, session, time.Minute*10).Return(nil).Times(1)

	mySession, err := sessionService.GetSessionByID(tenant, sessionid)

	assert.NotNil(t, mySession)
	assert.Equal(t, session, mySession)
	assert.Nil(t, err)
}

func TestGetCachedSessionByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	unitOfWork := createUnitOfWork(mockCtrl)

	mockRepo := unitOfWork.Session.(*mock_bbq.MockSessionRepository)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	mockDeviceService := mock_bbq.NewMockDeviceService(mockCtrl)
	mockMonitorService := mock_bbq.NewMockMonitorService(mockCtrl)
	mockSubjectService := mock_bbq.NewMockSubjectService(mockCtrl)
	sessionService := NewSessionService(mockCacheService, unitOfWork, mockDeviceService, mockMonitorService, mockSubjectService)

	tenant, err := uuid.NewUUID()

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	sessionid, err := uuid.NewUUID()

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	session := getSession(1, tenant, sessionid)
	sessionRecord := getSessionRecord(1, tenant, sessionid)

	var returnedSession bbq.Session

	cacheKey := fmt.Sprintf("bbq$sessions$%s$%s", tenant.String(), sessionid.String())

	mockRepo.EXPECT().GetByID(tenant, sessionid).Return(sessionRecord, nil).Times(0)
	mockCacheService.EXPECT().GetItem(cacheKey, &returnedSession).Return(nil).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, session, time.Minute*10).Return(nil).Times(0)

	sessionService.GetSessionByID(tenant, sessionid)

}

func TestGetSessionByAddress(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	unitOfWork := createUnitOfWork(mockCtrl)

	mockRepo := unitOfWork.Session.(*mock_bbq.MockSessionRepository)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	mockDeviceService := mock_bbq.NewMockDeviceService(mockCtrl)
	mockMonitorService := mock_bbq.NewMockMonitorService(mockCtrl)
	mockSubjectService := mock_bbq.NewMockSubjectService(mockCtrl)
	sessionService := NewSessionService(mockCacheService, unitOfWork, mockDeviceService, mockMonitorService, mockSubjectService)

	tenant, err := uuid.NewUUID()

	dev := bbq.Device{
		Name:        "Large Big Green Egg",
		Description: "My Device",
		TenantID:    tenant,
		ID:          2,
	}

	mon := bbq.Monitor{
		Name:        "My Great Monitor",
		Description: "My Great Monitor",
		TenantID:    tenant,
		ID:          2,
	}

	sub := bbq.Subject{
		Name:        "Pulled Pork",
		Description: "Pulled Pork",
		TenantID:    tenant,
		ID:          2,
	}

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	sessionid, err := uuid.NewUUID()

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	address := "deadbeefdeadbeef"
	session := getSession(1, tenant, sessionid)
	sessionRecord := getSessionRecord(1, tenant, sessionid)

	mockRepo.EXPECT().GetByMonitorAddress(tenant, address).Return(sessionRecord, nil).Times(1)
	mockDeviceService.EXPECT().GetDeviceByID(tenant, dev.Uid).Return(dev, nil).Times(1)
	mockMonitorService.EXPECT().GetMonitorByID(tenant, mon.Uid).Return(mon, nil).Times(1)
	mockSubjectService.EXPECT().GetSubjectByID(tenant, sub.Uid).Return(sub, nil).Times(1)

	mySession, err := sessionService.GetSessionByMonitorAddress(tenant, address)

	assert.NotNil(t, mySession)
	assert.Equal(t, session, mySession)
	assert.Nil(t, err)
}

func TestCreateSession(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tenant, err := uuid.NewUUID()

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	sessionid, err := uuid.NewUUID()

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	dev := bbq.Device{
		Name:        "Large Big Green Egg",
		Description: "My Device",
		TenantID:    tenant,
		ID:          2,
	}

	mon := bbq.Monitor{
		Name:        "My Great Monitor",
		Description: "My Great Monitor",
		TenantID:    tenant,
		ID:          2,
	}

	sub := bbq.Subject{
		Name:        "Pulled Pork",
		Description: "Pulled Pork",
		TenantID:    tenant,
		ID:          2,
	}

	cacheKey := fmt.Sprintf("bbq$sessions$%s$%s", tenant.String(), sessionid.String())
	tenantSessionCacheKey := fmt.Sprintf("bbq$sessions$%s", tenant.String())

	session := getSession(42, tenant, sessionid)
	sessionRecord := getSessionRecord(42, tenant, sessionid)

	// make sure times match for test records.
	sessionRecord.StartTime = session.StartTime

	unitOfWork := createUnitOfWork(mockCtrl)
	mockRepo := unitOfWork.Session.(*mock_bbq.MockSessionRepository)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	mockDeviceService := mock_bbq.NewMockDeviceService(mockCtrl)
	mockMonitorService := mock_bbq.NewMockMonitorService(mockCtrl)
	mockSubjectService := mock_bbq.NewMockSubjectService(mockCtrl)
	sessionService := NewSessionService(mockCacheService, unitOfWork, mockDeviceService, mockMonitorService, mockSubjectService)

	mockDeviceService.EXPECT().GetDeviceByName(tenant, dev.Name).Return(dev, nil).Times(1)
	mockDeviceService.EXPECT().GetDeviceByID(tenant, dev.Uid).Return(dev, nil).Times(1)
	mockMonitorService.EXPECT().GetMonitorByName(tenant, mon.Name).Return(mon, nil).Times(1)
	mockMonitorService.EXPECT().GetMonitorByID(tenant, mon.Uid).Return(mon, nil).Times(1)
	mockSubjectService.EXPECT().GetOrCreateSubject(tenant, sub.Name, sub.Description).Return(sub, nil).Times(1)
	mockSubjectService.EXPECT().GetSubjectByID(tenant, sub.Uid).Return(sub, nil).Times(1)
	mockRepo.EXPECT().Create(tenant, sessionRecord).Return(sessionRecord, nil).Times(1)
	mockCacheService.EXPECT().RemoveItem(tenantSessionCacheKey).Return(nil).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, session, time.Minute*10).Return(nil).Times(1)

	createdSession, err := sessionService.CreateSession(tenant, session)

	assert.Nil(t, err)
	assert.NotNil(t, createdSession)
	assert.Equal(t, session, createdSession)
}
