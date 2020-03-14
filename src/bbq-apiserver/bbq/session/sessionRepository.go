package session

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ssargent/bbq/bbq-apiserver/bbq"
)

type sessionRepository struct {
	database *sql.DB
}

// NewSessionRepository will return a repo for MonitorRepository
func NewSessionRepository(database *sql.DB) bbq.SessionRepository {
	return &sessionRepository{database: database}
}

func (s *sessionRepository) GetByTenantID(tenantID uuid.UUID) ([]bbq.SessionRecord, error) {
	query := "select id, deviceid, monitorid, name, description, starttime, subjectid, weight, tenantid, uid, endtime, deviceuid, monitoruid, subjectuid from bbq.sessions where tenantid = $1"

	var records []bbq.SessionRecord
	rows, err := s.database.Query(
		query, tenantID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var session bbq.SessionRecord
		if err := rows.Scan(&session.ID, &session.DeviceID, &session.MonitorID, &session.Name, &session.Description, &session.StartTime, &session.SubjectID, &session.Weight, &session.TenantID, &session.UID, &session.EndTime, &session.DeviceUID, &session.MonitorUID, &session.SubjectUID); err != nil {
			return nil, err
		}
		records = append(records, session)
	}

	return records, nil

}

func (s *sessionRepository) GetByID(tenantID uuid.UUID, id uuid.UUID) (bbq.SessionRecord, error) {
	var session bbq.SessionRecord
	query := "select id, deviceid, monitorid, name, description, starttime, subjectid, weight, tenantid, uid, endtime, deviceuid, monitoruid, subjectuid from bbq.sessions where Uid = $1 AND tenantid = $2"

	err := s.database.QueryRow(query, id, tenantID).Scan(&session.ID, &session.DeviceID, &session.MonitorID, &session.Name, &session.Description, &session.StartTime, &session.SubjectID, &session.Weight, &session.TenantID, &session.UID, &session.EndTime, &session.DeviceUID, &session.MonitorUID, &session.SubjectUID)

	if err != nil {
		return bbq.SessionRecord{}, err
	}

	return session, nil
}

func (s *sessionRepository) GetByMonitorAddress(tenantID uuid.UUID, address string) (bbq.SessionRecord, error) {
	var session bbq.SessionRecord
	query := `select s.id, s.deviceid, s.monitorid, s.name, s.description, s.starttime, s.subjectid, s.weight, s.tenantid, s.uid, s.endtime, s.deviceuid, s.monitoruid, s.subjectuid 
	from bbq.sessions s inner join bbq.monitors m on s.tenantid = m.tenantid and s.monitoruid = m.uid
	where s.tenantid = $1 and m.address = $2`

	err := s.database.QueryRow(query, tenantID, address).Scan(&session.ID, &session.DeviceID, &session.MonitorID, &session.Name, &session.Description, &session.StartTime, &session.SubjectID, &session.Weight, &session.TenantID, &session.UID, &session.EndTime, &session.DeviceUID, &session.MonitorUID, &session.SubjectUID)

	if err != nil {
		return bbq.SessionRecord{}, err
	}

	return session, nil
}

func (s *sessionRepository) Create(tenantID uuid.UUID, entity bbq.SessionRecord) (bbq.SessionRecord, error) {
	insertStatement := `INSERT INTO bbq.sessions
	(deviceid, monitorid, name, description, starttime, subjectid, weight, tenantid, deviceuid, monitoruid, subjectuid)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning *`

	var session bbq.SessionRecord
	err := s.database.QueryRow(insertStatement, entity.DeviceID, entity.MonitorID, entity.Name, entity.Description, entity.StartTime, entity.SubjectID, entity.Weight, entity.TenantID, entity.DeviceUID, entity.MonitorUID, entity.SubjectUID).Scan(&session.ID, &session.DeviceID, &session.MonitorID, &session.Name, &session.Description, &session.StartTime, &session.SubjectID, &session.Weight, &session.TenantID, &session.UID, &session.EndTime, &session.DeviceUID, &session.MonitorUID, &session.SubjectUID)

	if err != nil {
		// There must be a more elegant way of doing this...  but for now...
		if err, ok := err.(*pq.Error); ok {
			// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
			if err.Code.Name() == "unique_violation" {
				return bbq.SessionRecord{}, errors.New("a session with that name already exists for your account, please choose a different name")
			}
		}
		return bbq.SessionRecord{}, err
	}

	return session, nil
}

func (s *sessionRepository) Update(tenantID uuid.UUID, entity bbq.SessionRecord) (bbq.SessionRecord, error) {
	updateStatement := `update bbq.sessions 
		set deviceid = $1, monitorid = $2, name = $3, description = $4, starttime = $5, subjectid = $6, weight = $7, tenantid = $8, deviceuid = $9, monitoruid = $10, subjectuid = $11
		where uid = $12 and tenantid = $8  returning *`

	var session bbq.SessionRecord
	err := s.database.QueryRow(updateStatement, entity.DeviceID, entity.MonitorID, entity.Name, entity.Description, entity.StartTime, entity.SubjectID, entity.Weight, entity.TenantID, entity.DeviceUID, entity.MonitorUID, entity.SubjectUID, entity.UID).Scan(&session.ID, &session.DeviceID, &session.MonitorID, &session.Name, &session.Description, &session.StartTime, &session.SubjectID, &session.Weight, &session.TenantID, &session.UID, &session.EndTime, &session.DeviceUID, &session.MonitorUID, &session.SubjectUID)

	if err != nil {
		// There must be a more elegant way of doing this...  but for now...
		if err, ok := err.(*pq.Error); ok {
			// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
			if err.Code.Name() == "unique_violation" {
				return bbq.SessionRecord{}, errors.New("a session with that name already exists for your account, please choose a different name")
			}
		}
		return bbq.SessionRecord{}, err
	}

	return session, nil

}

func (s *sessionRepository) Delete(tenantID uuid.UUID, entity bbq.SessionRecord) error {
	result, err := s.database.Exec("delete from bbq.sessions where Uid = $1 and tenantid = $2", entity.UID, entity.TenantID)

	if rows, afferr := result.RowsAffected(); rows == 0 || afferr != nil {
		return errors.New("not-found")
	}

	return err
}
