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

	return bbq.Session{
		ID:          id,
		Name:        "My Session",
		Description: "My Session",
		Subject:     "Pulled Pork",
		Type:        "Pulled Pork",
		Weight:      9.2,
		Device:      "Large Big Green Egg",
		Monitor:     "My Great Monitor",
		StartTime:   time.Now(),
		EndTime:     nt,
		TenantID:    tenant,
		UID:         uid,
	}
}

func getSessionRecord(id int, tenant uuid.UUID, uid uuid.UUID) bbq.SessionRecord {
	nt := pq.NullTime{}

	return bbq.SessionRecord{
		ID:          id,
		DeviceID:    2,
		MonitorID:   2,
		Name:        "Pulled Pork",
		Description: "Pulled Pork",
		StartTime:   time.Now(),
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

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	sessionid, err := uuid.NewUUID()

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	session := getSession(1, tenant, sessionid)

	var returnedSessions []bbq.Session

	cacheKey := fmt.Sprintf("bbq$sessions$%s", tenant.String())

	mockRepo.EXPECT().GetByTenantID(tenant).Return([]bbq.Session{session}, nil).Times(1)
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

	var returnedSessions []bbq.Session

	cacheKey := fmt.Sprintf("bbq$sessions$%s", tenant.String())

	mockRepo.EXPECT().GetByTenantID(tenant).Return([]bbq.Session{session}, nil).Times(0)
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

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	sessionid, err := uuid.NewUUID()

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	session := getSession(1, tenant, sessionid)

	var returnedSession bbq.Session

	cacheKey := fmt.Sprintf("bbq$sessions$%s$%s", tenant.String(), sessionid.String())

	mockRepo.EXPECT().GetByID(tenant, sessionid).Return(session, nil).Times(1)
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

	var returnedSession bbq.Session

	cacheKey := fmt.Sprintf("bbq$sessions$%s$%s", tenant.String(), sessionid.String())

	mockRepo.EXPECT().GetByID(tenant, sessionid).Return(session, nil).Times(0)
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

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	sessionid, err := uuid.NewUUID()

	if err != nil {
		assert.Fail(t, "Failed to get UUID")
	}

	address := "deadbeefdeadbeef"
	session := getSession(1, tenant, sessionid)

	mockRepo.EXPECT().GetByMonitorAddress(tenant, address).Return(session, nil).Times(1)

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

	cacheKey := fmt.Sprintf("bbq$sessions$%s$%s", tenant.String(), sessionid.String())

	session := getSession(42, tenant, sessionid)
	sessionRecord := getSessionRecord(42, tenant, sessionid)

	unitOfWork := createUnitOfWork(mockCtrl)
	mockRepo := unitOfWork.Session.(*mock_bbq.MockSessionRepository)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	mockDeviceService := mock_bbq.NewMockDeviceService(mockCtrl)
	mockMonitorService := mock_bbq.NewMockMonitorService(mockCtrl)
	mockSubjectService := mock_bbq.NewMockSubjectService(mockCtrl)
	sessionService := NewSessionService(mockCacheService, unitOfWork, mockDeviceService, mockMonitorService, mockSubjectService)

	mockRepo.EXPECT().Create(tenant, sessionRecord).Return(sessionRecord, nil).Times(1)
	mockCacheService.EXPECT().SetItem(cacheKey, session, time.Minute*10).Return(nil).Times(1)

	createdSession, err := sessionService.CreateSession(tenant, session)

	assert.Nil(t, err)
	assert.NotNil(t, createdSession)
	assert.Equal(t, session, createdSession)
}
