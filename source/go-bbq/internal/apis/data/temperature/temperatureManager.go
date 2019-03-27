package temperature

import (
	"database/sql"
	"log"

	"github.com/google/uuid"

	"github.com/ssargent/go-bbq/internal/apis/bbq/sessions"
)

func getTenantSessionReadings(db *sql.DB, tenantName string, sessionid uuid.UUID) ([]Reading, error) {
	session, err := sessions.GetTenantSession(db, tenantName, sessionid)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(
		"select max(probe0) as probe0, "+
			"max(probe1) as probe1, "+
			"max(probe2) as probe2, "+
			"max(probe3) as probe3, date_trunc('minute', recordedat) as recordedat, sessionid "+
			"from data.bbq_temp_readings "+
			"where sessionid = $1 "+
			"group by sessionid, date_trunc('minute', recordedat) "+
			"order by date_trunc('minute', recordedat)", session.UID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	readings := []Reading{}

	for rows.Next() {
		s, err := scanReadingRows(rows)

		if err != nil {
			return nil, err
		}

		readings = append(readings, s)
	}

	return readings, nil
}

func scanReadingRows(row *sql.Rows) (Reading, error) {
	var t Reading
	if err := row.Scan(&t.Probe0, &t.Probe1, &t.Probe2, &t.Probe3, &t.RecordedAt, &t.SessionID); err != nil {
		return Reading{}, err
	}

	return t, nil
}

/*
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

	createdSession, err = getTenantSession(db, tenantName, createdSession.UID)

	if err != nil {
		return Session{}, err
	}

	return createdSession, nil
}
*/
