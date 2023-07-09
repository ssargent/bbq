package bbq

import "github.com/ssargent/bbq/internal/repository"

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
}

type Subject struct {
	repository.BbqSubject
}

type SubjectState struct {
	repository.BbqSubjectState
}
