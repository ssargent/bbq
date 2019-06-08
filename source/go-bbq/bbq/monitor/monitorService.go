package monitor

import (
	"github.com/google/uuid"
	"github.com/ssargent/go-bbq/bbq"
	"github.com/ssargent/go-bbq/internal/infrastructure"
)

type monitorService struct {
	repository bbq.MonitorRepository
	cache      infrastructure.CacheService
}

// NewMonitorService will create an MonitorService
func NewMonitorService(cache infrastructure.CacheService, repository bbq.MonitorRepository) bbq.MonitorService {
	return &monitorService{repository: repository, cache: cache}
}

func (m *monitorService) GetMonitors(tenantID uuid.UUID) ([]bbq.Monitor, error) {
	return nil, nil
}
func (m *monitorService) GetMonitor(tenantID uuid.UUID, name string) (bbq.Monitor, error) {
	return bbq.Monitor{}, nil
}
func (m *monitorService) CreateMonitor(tenantID uuid.UUID, entity bbq.Monitor) (bbq.Monitor, error) {
	return bbq.Monitor{}, nil
}
func (m *monitorService) UpdateMonitor(tenantID uuid.UUID, entity bbq.Monitor) (bbq.Monitor, error) {
	return bbq.Monitor{}, nil
}
func (m *monitorService) DeleteMonitor(tenantID uuid.UUID, entity bbq.Monitor) error {
	return nil
}
