package data

import (
	"time"

	"github.com/google/uuid"
)

//Device is
// SensorReadings represents the structure of arrays datastructure for a cooking session.
type SensorReadings struct {
	SessionId uuid.UUID `json:"sessionid"`
	TenantId  uuid.UUID `json:"tenantid"`
	Probe0    []float32 `json:"probe0"`
	Probe1    []float32 `json:"probe1"`
	Probe2    []float32 `json:"probe2"`
	Probe3    []float32 `json:"probe3"`
}

// ThermalSensorRecord represents the raw structure of a measure at a point in time.
type ThermalSensorRecord struct {
	SessionId  uuid.UUID `json:"sessionid"`
	TenantID   uuid.UUID `json:"tenantid"`
	Probe0     float32   `json:"probe0"`
	Probe1     float32   `json:"probe1"`
	Probe2     float32   `json:"probe2"`
	Probe3     float32   `json:"probe3"`
	RecordedAt time.Time `json:"recordedat"`
}

// SensorReadingService gets structure-of-arrays data about a cooking session
type SensorReadingService interface {
	GetReadings(tenantID uuid.UUID, sessionid uuid.UUID) (SensorReadings, error)
}

//SensorReadingRepository gets the raw array-of-structures data about a cooking sessin
type SensorReadingRepository interface {
	GetThermalReadings(tenantID uuid.UUID, sessionid uuid.UUID) ([]ThermalSensorRecord, error)
}
