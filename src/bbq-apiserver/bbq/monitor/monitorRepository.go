package monitor

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
)

type monitorRepository struct {
	database *sql.DB
}

// NewMonitorRepository will return a repo for MonitorRepository
func NewMonitorRepository(database *sql.DB) bbq.MonitorRepository {
	return &monitorRepository{database: database}
}

func (m *monitorRepository) GetByID(tenantID uuid.UUID, id uuid.UUID) (bbq.Monitor, error) {
	var mon bbq.Monitor
	query := "select id, uid, address, name, description, tenantid from bbq.devices where Uid = $1 AND tenantid = $2"

	err := m.database.QueryRow(query, id, tenantID).Scan(&mon.ID, &mon.Uid, &mon.Address, &mon.Name, &mon.Description, &mon.TenantID)

	if err != nil {
		return bbq.Monitor{}, err
	}

	return mon, nil

}

func (m *monitorRepository) GetByTenantID(tenantID uuid.UUID) ([]bbq.Monitor, error) {
	var monitors []bbq.Monitor
	rows, err := m.database.Query(
		"select id, uid, address, name, description, tenantid from bbq.devices where tenantid = $1", tenantID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var mon bbq.Monitor
		if err := rows.Scan(&mon.ID, &mon.Uid, &mon.Address, &mon.Name, &mon.Description, &mon.TenantID); err != nil {
			return nil, err
		}
		monitors = append(monitors, mon)
	}

	return monitors, nil
}

func (m *monitorRepository) GetByName(tenantID uuid.UUID, name string) (bbq.Monitor, error) {
	var mon bbq.Monitor
	query := "select id, uid, address, name, description, tenantid from bbq.devices where Name = $1 AND tenantid = $2"

	err := m.database.QueryRow(query, name, tenantID).Scan(&mon.ID, &mon.Uid, &mon.Address, &mon.Name, &mon.Description, &mon.TenantID)

	if err != nil {
		return bbq.Monitor{}, err
	}

	return mon, nil

}

func (m *monitorRepository) GetByAddress(tenantID uuid.UUID, address string) (bbq.Monitor, error) {
	var mon bbq.Monitor
	query := "select id, uid, address, name, description, tenantid from bbq.devices where Address = $1 AND tenantid = $2"

	err := m.database.QueryRow(query, address, tenantID).Scan(&mon.ID, &mon.Uid, &mon.Address, &mon.Name, &mon.Description, &mon.TenantID)

	if err != nil {
		return bbq.Monitor{}, err
	}

	return mon, nil

}

func (m *monitorRepository) Create(entity bbq.Monitor) (bbq.Monitor, error) {
	insertStatement := "insert into bbq.monitors (address, name, description, tenantid) values ($1, $2, $3, $4) returning *"

	var mon bbq.Monitor
	err := m.database.QueryRow(insertStatement, entity.Address, entity.Name, entity.Description, entity.TenantID).Scan(&mon.ID, &mon.Uid, &mon.Address, &mon.Name, &mon.Description, &mon.TenantID)

	if err != nil {
		// There must be a more elegant way of doing this...  but for now...
		if err, ok := err.(*pq.Error); ok {
			// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
			if err.Code.Name() == "unique_violation" {
				return bbq.Monitor{}, errors.New("a monitor with that name already exists for your account, please choose a different name")
			}
		}
		return bbq.Monitor{}, err
	}

	return mon, nil
}

func (m *monitorRepository) Update(entity bbq.Monitor) (bbq.Monitor, error) {
	var mon bbq.Monitor
	query := `update bbq.Monitors set name = $3, description = $4
			  where Uid = $1 and TenantID = $2
			  returning *`

	err := m.database.
		QueryRow(query, entity.Uid, entity.TenantID, entity.Name, entity.Description).
		Scan(&mon.ID, &mon.Uid, &mon.Address, &mon.Name, &mon.Description, &mon.TenantID)

	if err != nil {
		// There must be a more elegant way of doing this...  but for now...
		if err, ok := err.(*pq.Error); ok {
			// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
			if err.Code.Name() == "unique_violation" {
				return bbq.Monitor{}, errors.New("a monitor with that name already exists for your account, please choose a different name")
			}
		}
		return bbq.Monitor{}, err
	}

	return mon, nil
}

func (m *monitorRepository) Delete(entity bbq.Monitor) error {
	result, err := m.database.Exec("delete from bbq.Monitors where Uid = $1 and tenantid = $2", entity.Uid, entity.TenantID)

	if rows, afferr := result.RowsAffected(); rows == 0 || afferr != nil {
		return errors.New("not-found")
	}

	return err
}
