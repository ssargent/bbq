// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repository

import (
	"github.com/google/uuid"
)

type BbqDevice struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	IsDefault bool      `json:"is_default"`
}

type BbqSensor struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsDefault   bool      `json:"is_default"`
}

type BbqSensorAlert struct {
	ID                uuid.UUID `json:"id"`
	AlertName         string    `json:"alert_name"`
	Description       string    `json:"description"`
	NotificationRules string    `json:"notification_rules"`
}

type BbqSensorAlertRule struct {
	ID       uuid.UUID `json:"id"`
	AlertID  uuid.UUID `json:"alert_id"`
	RuleName string    `json:"rule_name"`
	RuleMin  float64   `json:"rule_min"`
	RuleMax  float64   `json:"rule_max"`
}

type BbqSensorReading struct {
	ID              uuid.UUID `json:"id"`
	SessionID       uuid.UUID `json:"session_id"`
	ProbeNumber     int32     `json:"probe_number"`
	Temperature     float64   `json:"temperature"`
	ReadingOccurred int64     `json:"reading_occurred"`
}

type BbqSession struct {
	ID           uuid.UUID `json:"id"`
	DeviceID     uuid.UUID `json:"device_id"`
	DesiredState uuid.UUID `json:"desired_state"`
	StartTime    int64     `json:"start_time"`
	EndTime      *int64    `json:"end_time"`
	SensorID     uuid.UUID `json:"sensor_id"`
	SessionType  int32     `json:"session_type"`
	SubjectID    uuid.UUID `json:"subject_id"`
	Description  string    `json:"description"`
}

type BbqSessionAlert struct {
	SessionID      uuid.UUID `json:"session_id"`
	AlertID        uuid.UUID `json:"alert_id"`
	AlertSuspended bool      `json:"alert_suspended"`
	AlertStart     int64     `json:"alert_start"`
	AlertEnd       *int64    `json:"alert_end"`
}

type BbqSubject struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type BbqSubjectState struct {
	ID          uuid.UUID `json:"id"`
	SubjectID   uuid.UUID `json:"subject_id"`
	State       string    `json:"state"`
	Temperature int32     `json:"temperature"`
}

type SchemaMigration struct {
	Version int64 `json:"version"`
	Dirty   bool  `json:"dirty"`
}
