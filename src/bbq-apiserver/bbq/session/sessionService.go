package session

import (
	"errors"
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

	sessionRecs, err := s.unitOfWork.Session.GetByTenantID(tenantID)
	if err != nil {
		fmt.Println("No Sessions", tenantID)
		return []bbq.Session{}, err
	}

	for _, element := range sessionRecs {
		session, err := s.convertToSession(element)

		if err != nil {
			fmt.Println("Cannot Convert Session ", element.UID)

			return []bbq.Session{}, err
		}

		sessions = append(sessions, session)
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

	sessionRec, err := s.unitOfWork.Session.GetByID(tenantID, id)
	if err != nil {
		return bbq.Session{}, err
	}

	session, err = s.convertToSession(sessionRec)

	if err != nil {
		return bbq.Session{}, err
	}

	s.cache.SetItem(cacheKey, session, time.Minute*10)

	return session, nil
}

func (s *sessionService) GetSessionByMonitorAddress(tenantID uuid.UUID, address string) (bbq.Session, error) {
	sessionRec, err := s.unitOfWork.Session.GetByMonitorAddress(tenantID, address)

	if err != nil {
		return bbq.Session{}, err
	}

	session, err := s.convertToSession(sessionRec)

	return session, nil
}

func (s *sessionService) convertToSession(record bbq.SessionRecord) (bbq.Session, error) {

	device, err := s.deviceService.GetDeviceByID(record.TenantID, record.DeviceUID)

	if err != nil {
		return bbq.Session{}, errors.New("Cannot get device for session")
	}

	monitor, err := s.monitorService.GetMonitorByID(record.TenantID, record.MonitorUID)

	if err != nil {
		return bbq.Session{}, errors.New("Cannot get monitor for session")
	}

	subject, err := s.subjectService.GetSubjectByID(record.TenantID, record.SubjectUID)

	if err != nil {
		return bbq.Session{}, errors.New("Cannot get subject for sesson")
	}

	return bbq.Session{
		ID:          record.ID,
		UID:         record.UID,
		Name:        record.Name,
		Description: record.Description,
		Subject:     subject.Name,
		Device:      device.Name,
		Monitor:     monitor.Name,
		Weight:      record.Weight,
		StartTime:   record.StartTime,
		EndTime:     record.EndTime,
		TenantID:    record.TenantID,
		//	Type:        record.Type
		//	Subject: ,
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
		MonitorUID:  monitor.Uid,
		DeviceID:    device.ID,
		DeviceUID:   device.Uid,
		Name:        record.Name,
		Description: record.Description,
		StartTime:   record.StartTime,
		SubjectID:   subject.ID,
		SubjectUID:  subject.Uid,
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
