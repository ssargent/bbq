package bbq

import "github.com/ssargent/bbq/internal/bbq/repository"

type Device struct {
	repository.BbqDevice
}

type Sensor struct {
	repository.BbqSensor
}

type SensorReading struct {
	repository.BbqSensorReading
}

type Session struct {
	repository.BbqSession

	DeviceName *string `json:"device_name,omitempty"`
	SensorName *string `json:"sensor_name,omitempty"`
}

type Subject struct {
	repository.BbqSubject
}

type SubjectState struct {
	repository.BbqSubjectState
}
