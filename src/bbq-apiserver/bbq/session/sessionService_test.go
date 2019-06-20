package session

import (
	
	"testing"
	"time"
	"fmt"
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/google/uuid"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	mock_bbq "github.com/ssargent/bbq/bbq-apiserver/bbq/mocks"
	mock_infrastructure "github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure/mocks"
	"github.com/stretchr/testify/assert"
)

func getSession(id int, tenant uuid.UUID, uid uuid.UUID) bbq.Session {
	nt := pq.NullTime{}

	return bbq.Session{
		ID:	id,
		Name: "My Session",
		Description: "My Session",
		Subject: "Pulled Pork",
		Type: 	"Pulled Pork",
		Weight: 9.2,
		Device: "Large Big Green Egg",
		Monitor: "My Great Monitor",
		StartTime: time.Now(),
		EndTime: nt,
		TenantID: tenant,
		UID: uid,
	}
}

func TestGetSessions(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_bbq.NewMockSessionRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	sessionService := NewSessionService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return;
	}
 
	sessionid, err := uuid.NewUUID()

	if err != nil {
		return;
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

	mockRepo := mock_bbq.NewMockSessionRepository(mockCtrl)
	mockCacheService := mock_infrastructure.NewMockCacheService(mockCtrl)
	sessionService := NewSessionService(mockCacheService, mockRepo)

	tenant, err := uuid.NewUUID()

	if err != nil {
		return;
	}
 
	sessionid, err := uuid.NewUUID()

	if err != nil {
		return;
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
