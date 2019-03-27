package sessions

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"

	//"errors"

	//"github.com/ssargent/go-bbq/apis/bbq/devices"
	//"github.com/ssargent/go-bbq/apis/bbq/monitors"
	//"github.com/ssargent/go-bbq/apis/bbq/subjects"
	"github.com/ssargent/go-bbq/internal/apis/system/tenants"
)

func getTenantSessions(db *sql.DB, tenantName string) ([]Session, error) {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found Tenant: ", tenant.ID, tenant.Name, tenant.URLKey)
	rows, err := db.Query(
		"SELECT * FROM bbq.vw_sessions where tenantid = $1", tenant.ID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sessions := []Session{}

	for rows.Next() {
		s, err := scanSessionRows(rows)

		if err != nil {
			return nil, err
		}

		sessions = append(sessions, s)
	}

	return sessions, nil
}

func scanSessionRows(row *sql.Rows) (Session, error) {
	var s Session
	if err := row.Scan(&s.ID, &s.Name, &s.Description, &s.Subject, &s.Type, &s.Weight, &s.Device, &s.Monitor, &s.StartTime, &s.TenantID, &s.UID, &s.EndTime); err != nil {
		return Session{}, err
	}

	return s, nil
}

func scanSession(row *sql.Row) (Session, error) {
	var s Session
	if err := row.Scan(&s.ID, &s.Name, &s.Description, &s.Subject, &s.Type, &s.Weight, &s.Device, &s.Monitor, &s.StartTime, &s.TenantID, &s.UID, &s.EndTime); err != nil {
		return Session{}, err
	}

	return s, nil
}

// GetTenantSession returns a session for a given tenant...
func GetTenantSession(db *sql.DB, tenantName string, sessionid uuid.UUID) (Session, error) {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found Tenant: ", tenant.ID, tenant.Name, tenant.URLKey)

	row := db.QueryRow("select * from bbq.vw_sessions where uid = $1 and tenantid = $2", sessionid, tenant.ID)

	s, scanErr := scanSession(row)

	if scanErr != nil {
		return Session{}, scanErr
	}

	return s, nil
}

func translateSessionRecord(db *sql.DB, tenantName string, session Session) (sessionRecord, error) {
	var record sessionRecord

	/*	device, err := devices.GetTenantDeviceByName(db, tenantName, session.Device)

		if err != nil {
			return sessionRecord{}, err
		}

		monitor, err := monitors.GetTenantMonitorByName(db, tenantName, session.Monitor)

		if err != nil {
			return sessionRecord{}, err
		}

		subject, err := subjects.GetSubjectByName(db, session.Subject)

		if err != nil {
			return sessionRecord{}, err
		}

		record.DeviceID = device.ID
		record.MonitorID = monitor.ID
		record.SubjectID = subject.ID
		record.Name = session.Name
		record.Description = session.Description
		record.StartTime = session.StartTime
		record.TenantID = session.TenantID
		record.EndTime = session.EndTime
	*/
	return record, nil
}

func createTenantSession(db *sql.DB, tenantName string, session Session) (Session, error) {
	tenant, err := tenants.GetTenantByKey(db, tenantName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found Tenant: ", tenant.ID, tenant.Name, tenant.URLKey)

	insertStatement := "insert into bbq.sessions " +
		"(deviceid, monitorid, name, description, starttime, subjectid, weight, tenantid) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8) returning id"

	translatedRecord, err := translateSessionRecord(db, tenantName, session)

	if err != nil {
		return Session{}, err
	}

	var createdSession Session
	err = db.QueryRow(insertStatement, translatedRecord.DeviceID, translatedRecord.MonitorID, translatedRecord.Name, translatedRecord.Description, translatedRecord.StartTime, translatedRecord.SubjectID, translatedRecord.Weight, tenant.ID).Scan(&createdSession.ID)

	if err != nil {
		return Session{}, err
	}

	createdSession, err = GetTenantSession(db, tenantName, createdSession.UID)

	if err != nil {
		return Session{}, err
	}

	return createdSession, nil
}

func deleteTenantSession(db *sql.DB, tenantid string, sessionid int) error {
	return nil
}
