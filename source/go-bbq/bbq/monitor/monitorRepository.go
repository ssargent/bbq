package monitor

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/ssargent/go-bbq/bbq"
)

type monitorRepository struct {
	database *sql.DB
}

// NewMonitorRepository will return a repo for MonitorRepository
func NewMonitorRepository(database *sql.DB) bbq.MonitorRepository {
	return &monitorRepository{database: database}
}

func (m *monitorRepository) GetById(tenantId uuid.UUID, id uuid.UUID) (bbq.Monitor, error) {
	return bbq.Monitor{}, nil
}

func (m *monitorRepository) GetByTenantID(tenantID uuid.UUID) ([]bbq.Monitor, error) {
	return nil, nil
}

func (m *monitorRepository) GetByName(tenantID uuid.UUID, name string) (bbq.Monitor, error) {
	return bbq.Monitor{}, nil
}

func (m *monitorRepository) GetByAddress(tenantID uuid.UUID, address string) (bbq.Monitor, error) {
	return bbq.Monitor{}, nil
}

func (m *monitorRepository) Create(entity bbq.Monitor) (bbq.Monitor, error) {
	return bbq.Monitor{}, nil
}

func (m *monitorRepository) Update(entity bbq.Monitor) (bbq.Monitor, error) {
	return bbq.Monitor{}, nil
}

func (m *monitorRepository) Delete(entity bbq.Monitor) error {
	return nil
}
