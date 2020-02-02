package sensors

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/ssargent/bbq/bbq-apiserver/data"
)

type sensorReadingRepository struct {
	database *sql.DB
}

// NewDeviceRepository will return a repo for DeviceRepository
func NewSensorReadingRepository(database *sql.DB) data.SensorReadingRepository {
	return &sensorReadingRepository{database: database}
}

func (s *sensorReadingRepository) GetThermalReadings(tenantID uuid.UUID, sessionid uuid.UUID) ([]data.ThermalSensorRecord, error) {

	rows, err := s.database.Query(
		"select max(probe0) as probe0, "+
			"max(probe1) as probe1, "+
			"max(probe2) as probe2, "+
			"max(probe3) as probe3, date_trunc('minute', recordedat) as recordedat, sessionid "+
			"from data.bbq_temp_readings "+
			"where sessionid = $1 "+
			"group by sessionid, date_trunc('minute', recordedat) "+
			"order by date_trunc('minute', recordedat)", sessionid)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	readings := []data.ThermalSensorRecord{}

	for rows.Next() {
		s, err := scanReadingRows(rows)

		if err != nil {
			return nil, err
		}

		readings = append(readings, s)
	}

	return readings, nil
}

func scanReadingRows(row *sql.Rows) (data.ThermalSensorRecord, error) {
	var t data.ThermalSensorRecord
	if err := row.Scan(&t.Probe0, &t.Probe1, &t.Probe2, &t.Probe3, &t.RecordedAt, &t.SessionId); err != nil {
		return data.ThermalSensorRecord{}, err
	}

	return t, nil
}
