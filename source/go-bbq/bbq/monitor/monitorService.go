package monitor

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

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
	cacheKey := fmt.Sprintf("bbq$monitors$%s", tenantID.String())

	var monitors []bbq.Monitor

	if err := m.cache.GetItem(cacheKey, &monitors); err == nil {
		return monitors, nil
	}

	monitors, err := m.repository.GetByTenantID(tenantID)
	if err != nil {
		return []bbq.Monitor{}, err
	}

	m.cache.SetItem(cacheKey, monitors, time.Minute*10)

	return monitors, nil

}
func (m *monitorService) GetMonitorById(tenantID uuid.UUID, id uuid.UUID) (bbq.Monitor, error) {
	return bbq.Monitor{}, nil
}

func (m *monitorService) GetMonitorByName(tenantID uuid.UUID, name string) (bbq.Monitor, error) {
	cacheKey := fmt.Sprintf("bbq$monitors$%s$%s", tenantID.String(), name)
	var monitor bbq.Monitor

	if err := m.cache.GetItem(cacheKey, &monitor); err == nil {
		return monitor, nil
	}

	monitor, err := m.repository.GetByName(tenantID, name)
	if err != nil {
		return bbq.Monitor{}, err
	}

	m.cache.SetItem(cacheKey, monitor, time.Minute*10)

	return monitor, nil
}

func (m *monitorService) GetMonitorByAddress(tenantID uuid.UUID, address string) (bbq.Monitor, error) {

	monitor, err := m.repository.GetByAddress(tenantID, address)
	if err != nil {
		return bbq.Monitor{}, err
	}

	return monitor, nil
}

func (m *monitorService) CreateMonitor(tenantID uuid.UUID, entity bbq.Monitor) (bbq.Monitor, error) {
	entity.TenantID = tenantID
	cacheKey := fmt.Sprintf("bbq$monitors$%s$%s", tenantID.String(), entity.Name)

	_, err := m.repository.GetByName(tenantID, entity.Name)

	if err == nil {
		return bbq.Monitor{}, errors.New("A monitor with that name already exists for your tenant")
	}

	monitor, err := m.repository.Create(entity)
	if err != nil {
		return bbq.Monitor{}, err
	}

	m.cache.SetItem(cacheKey, monitor, time.Minute*10)

	return monitor, nil
}
func (m *monitorService) UpdateMonitor(tenantID uuid.UUID, entity bbq.Monitor) (bbq.Monitor, error) {
	entity.TenantID = tenantID
	cacheKey := fmt.Sprintf("bbq$monitors$%s$%s", tenantID.String(), entity.Name)

	_, err := m.repository.GetByName(tenantID, entity.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return bbq.Monitor{}, errors.New("Monitor not found.  You must create it first.")
		}

		return bbq.Monitor{}, err
	}

	monitor, err := m.repository.Update(entity)
	if err != nil {
		return bbq.Monitor{}, err
	}

	m.cache.SetItem(cacheKey, monitor, time.Minute*10)

	return monitor, nil

}
func (m *monitorService) DeleteMonitor(tenantID uuid.UUID, entity bbq.Monitor) error {
	return nil
}
