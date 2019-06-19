package session

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
)

type sessionRepository struct {
	database *sql.DB
}

// NewMonitorRepository will return a repo for MonitorRepository
func NewSessionRepository(database *sql.DB) bbq.SessionRepository {
	return &sessionRepository{database: database}
}

func (s *sessionRepository) GetByTenantID(tenantID uuid.UUID) ([]bbq.Session, error) {
	panic("not implemented")
}

func (s *sessionRepository) GetByID(tenantID uuid.UUID, id uuid.UUID) (bbq.Session, error) {
	panic("not implemented")
}

func (s *sessionRepository) GetByMonitorAddress(tenantID uuid.UUID, address string) (bbq.Session, error) {
	panic("not implemented")
}

func (s *sessionRepository) Create(tenantID uuid.UUID, entity bbq.Session) (bbq.Session, error) {
	panic("not implemented")
}

func (s *sessionRepository) Update(tenantID uuid.UUID, entity bbq.Session) (bbq.Session, error) {
	panic("not implemented")
}

func (s *sessionRepository) Delete(tenantID uuid.UUID, entity bbq.Session) error {
	panic("not implemented")
}
