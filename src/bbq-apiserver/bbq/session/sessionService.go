package session

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
)

type sessionService struct {
	unitOfWork     bbq.BBQUnitOfWork
	cache          infrastructure.CacheService
	deviceService  bbq.DeviceService
	monitorService bbq.MonitorService
	subjectService bbq.SubjectService
}

// NewSessionService will create an SessionService
func NewSessionService(cache infrastructure.CacheService, unitOfWork bbq.BBQUnitOfWork, deviceService bbq.DeviceService, monitorService bbq.MonitorService, subjectService bbq.SubjectService) bbq.SessionService {
	return &sessionService{unitOfWork: unitOfWork, cache: cache, subjectService: subjectService, monitorService: monitorService, deviceService: deviceService}
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

func (s *sessionService) convertToSession(record bbq.SessionRecord) (bbq.Session, error) {
/*
			type Session struct {
			ID          int         `json:"id"`
			Name        string      `json:"name"`
			Description string      `json:"description"`
			Subject     string      `json:"subject"`
			Type        string      `json:"type"`
			Weight      float64     `json:"weight"`
			Device      string      `json:"device"`
			Monitor     string      `json:"monitor"`
			StartTime   time.Time   `json:"starttime"`
			EndTime     pq.NullTime `json:"endtime"`
			TenantID    uuid.UUID   `json:"tenantid"`
			UID         uuid.UUID   `json:"uid"`
		}

		type SessionRecord struct {
			ID          int         `json:"id"`
			DeviceID    int         `json:"deviceid"`
			MonitorID   int         `json:"monitorid"`
			Name        string      `json:"name"`
			Description string      `json:"description"`
			StartTime   time.Time   `json:"starttime"`
			SubjectID   int         `json:"subjectid"`
			Weight      float64     `json:"weight"`
			TenantID    uuid.UUID   `json:"tenantid"`
			UID         uuid.UUID   `json:"uid"`
			EndTime     pq.NullTime `json:"endtime"`
		}
	*/

	return bbq.Session{
		ID:  record.ID,
		UID: record.UID,
		Name: record.Name,
		Description: record.Description,
		Subject: ,
	}, nil
}

func (s *sessionService) convertToRecord(tenantID uuid.UUID, record bbq.Session) (bbq.SessionRecord, error) {
	device, err := s.deviceService.GetDeviceByName(tenantID, record.Device)

	if err != nil {
		return bbq.SessionRecord{}, err
	}

	monitor, err := s.monitorService.GetMonitorByName(tenantID, record.Monitor)

	if err != nil {
		return bbq.SessionRecord{}, err
	}

	subject, err := s.subjectService.GetOrCreateSubject(tenantID, record.Subject, record.Subject)

	if err != nil {
		return bbq.SessionRecord{}, err
	}

	
	return bbq.SessionRecord{
		MonitorID:   monitor.ID,
		DeviceID:    device.ID,
		Name:        record.Name,
		Description: record.Description,
		StartTime:   record.StartTime,
		SubjectID:   subject.ID,
		Weight:      record.Weight,
		TenantID:    tenantID,
		EndTime:     record.EndTime,
		ID:          record.ID,
		UID:         record.UID,
	}, nil

}

func (s *sessionService) CreateSession(tenantID uuid.UUID, entity bbq.Session) (bbq.Session, error) {
	record, err := s.convertToRecord(tenantID, entity)

	if err != nil {
		return bbq.Session{}, err
	}

	createdRecord, err := s.unitOfWork.Session.Create(tenantID, record)

	if err != nil {
		return bbq.Session{}, err
	}

	createdSession, err := s.convertToSession(createdRecord)

	if err != nil {
		return bbq.Session{}, err
	}

	cacheKey := fmt.Sprintf("bbq$sessions$%s$%s", tenantID.String(), createdSession.UID.String())

	s.cache.SetItem(cacheKey, createdSession, time.Minute*10)

	return createdSession, nil
}

func (s *sessionService) UpdateSession(tenantID uuid.UUID, entity bbq.Session) (bbq.Session, error) {
	panic("not implemented")
}

func (s *sessionService) DeleteSession(tenantID uuid.UUID, entity bbq.Session) error {
	panic("not implemented")
}
