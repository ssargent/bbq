package session

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
)

type sessionService struct {
	unitOfWork bbq.BBQUnitOfWork
	cache      infrastructure.CacheService
}

// NewSessionService will create an SessionService
func NewSessionService(cache infrastructure.CacheService, unitOfWork bbq.BBQUnitOfWork) bbq.SessionService {
	return &sessionService{unitOfWork: unitOfWork, cache: cache}
}

func (s *sessionService) GetSessions(tenantID uuid.UUID) ([]bbq.Session, error) {
	cacheKey := fmt.Sprintf("bbq$sessions$%s", tenantID.String())

	var sessions []bbq.Session

	if err := s.cache.GetItem(cacheKey, &sessions); err == nil {
		return sessions, nil
	}

	sessions, err := s.unitOfWork.Session.GetByTenantID(tenantID)
	if err != nil {
		return []bbq.Session{}, err
	}

	s.cache.SetItem(cacheKey, sessions, time.Minute*10)

	return sessions, nil
}

func (s *sessionService) GetSessionByID(tenantID uuid.UUID, id uuid.UUID) (bbq.Session, error) {
	cacheKey := fmt.Sprintf("bbq$sessions$%s$%s", tenantID.String(), id.String())

	var session bbq.Session

	if err := s.cache.GetItem(cacheKey, &session); err == nil {
		return session, nil
	}

	session, err := s.unitOfWork.Session.GetByID(tenantID, id)
	if err != nil {
		return bbq.Session{}, err
	}

	s.cache.SetItem(cacheKey, session, time.Minute*10)

	return session, nil
}

func (s *sessionService) GetSessionByMonitorAddress(tenantID uuid.UUID, address string) (bbq.Session, error) {
	session, err := s.unitOfWork.Session.GetByMonitorAddress(tenantID, address)

	if err != nil {
		return bbq.Session{}, err
	}

	return session, nil
}

func (s *sessionService) CreateSession(tenantID uuid.UUID, entity bbq.Session) (bbq.Session, error) {
	panic("not implemented")
}

func (s *sessionService) UpdateSession(tenantID uuid.UUID, entity bbq.Session) (bbq.Session, error) {
	panic("not implemented")
}

func (s *sessionService) DeleteSession(tenantID uuid.UUID, entity bbq.Session) error {
	panic("not implemented")
}
