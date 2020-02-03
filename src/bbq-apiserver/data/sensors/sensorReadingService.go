package sensors

import (
	"github.com/google/uuid"

	"github.com/ssargent/bbq/bbq-apiserver/data"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
)

type sensorReadingService struct {
	repository data.SensorReadingRepository
	cache      infrastructure.CacheService
}

// NewSensorReadingService will create an DeviceService
func NewSensorReadingService(cache infrastructure.CacheService, repository data.SensorReadingRepository) data.SensorReadingService {
	return &sensorReadingService{repository: repository, cache: cache}
}

func (s *sensorReadingService) GetReadings(tenantID uuid.UUID, sessionid uuid.UUID) (data.SensorReadings, error) {
	rawReadingData, error := s.repository.GetThermalReadings(tenantID, sessionid)

	if error != nil {
		return data.SensorReadings{}, error
	}

	sensorReadings := data.SensorReadings{}

	sensorReadings.SessionId = sessionid
	sensorReadings.TenantId = tenantID

	for _, reading := range rawReadingData {
		sensorReadings.Probe0 = append(sensorReadings.Probe0, reading.Probe0)
		sensorReadings.Probe1 = append(sensorReadings.Probe1, reading.Probe1)
		sensorReadings.Probe2 = append(sensorReadings.Probe2, reading.Probe2)
		sensorReadings.Probe3 = append(sensorReadings.Probe3, reading.Probe3)
	}

	return sensorReadings, nil
}

func (s *sensorReadingService) GetRawReadings(tenantID uuid.UUID, sessionid uuid.UUID) ([]data.ThermalSensorRecord, error) {
	rawReadingData, error := s.repository.GetThermalReadings(tenantID, sessionid)

	if error != nil {
		return []data.ThermalSensorRecord{}, error
	}

	return rawReadingData, nil
}
