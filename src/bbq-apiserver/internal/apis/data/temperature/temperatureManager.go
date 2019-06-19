package temperature

import (
	"database/sql"
	"log"

	"github.com/google/uuid"

	"github.com/ssargent/bbq/bbq-apiserver/internal/apis/bbq/sessions"
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

func createTenantSessionReading(db *sql.DB, tenantName string, tempReading Reading) error {
	//fmt.Println("Found Tenant: ", tenant.ID, tenant.Name, tenant.URLKey)

	insertStatement := `
	
insert into data.bbq_temp_readings
(probe0, probe1, probe2, probe3, recordedat, sessionid)
values 
($1, $2, $3, $4, now(), $5) returning id
	`

	_, err := db.Exec(insertStatement, tempReading.Probe0, tempReading.Probe1, tempReading.Probe2, tempReading.Probe3, tempReading.SessionID)

	if err != nil {
		return err
	}

	return nil
}
