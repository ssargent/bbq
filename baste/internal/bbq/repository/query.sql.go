// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const deleteSubjectState = `-- name: DeleteSubjectState :exec
delete from bbq.subject_states where id = $1
`

func (q *Queries) DeleteSubjectState(ctx context.Context, db DBTX, id uuid.UUID) error {
	_, err := db.Exec(ctx, deleteSubjectState, id)
	return err
}

const getDefaultDevice = `-- name: GetDefaultDevice :one
select id, name, location, is_default from bbq.devices where is_default = true
`

func (q *Queries) GetDefaultDevice(ctx context.Context, db DBTX) (*BbqDevice, error) {
	row := db.QueryRow(ctx, getDefaultDevice)
	var i BbqDevice
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.IsDefault,
	)
	return &i, err
}

const getDefaultSensor = `-- name: GetDefaultSensor :one
select id, name, description, is_default from bbq.sensors where is_default = true
`

func (q *Queries) GetDefaultSensor(ctx context.Context, db DBTX) (*BbqSensor, error) {
	row := db.QueryRow(ctx, getDefaultSensor)
	var i BbqSensor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.IsDefault,
	)
	return &i, err
}

const getDeviceByID = `-- name: GetDeviceByID :one
select id, name, location, is_default from bbq.devices where id = $1
`

func (q *Queries) GetDeviceByID(ctx context.Context, db DBTX, id uuid.UUID) (*BbqDevice, error) {
	row := db.QueryRow(ctx, getDeviceByID, id)
	var i BbqDevice
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.IsDefault,
	)
	return &i, err
}

const getDeviceByName = `-- name: GetDeviceByName :one
select id, name, location, is_default from bbq.devices where name = $1
`

func (q *Queries) GetDeviceByName(ctx context.Context, db DBTX, name string) (*BbqDevice, error) {
	row := db.QueryRow(ctx, getDeviceByName, name)
	var i BbqDevice
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.IsDefault,
	)
	return &i, err
}

const getDevices = `-- name: GetDevices :many
select id, name, location, is_default from bbq.devices
`

func (q *Queries) GetDevices(ctx context.Context, db DBTX) ([]*BbqDevice, error) {
	rows, err := db.Query(ctx, getDevices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*BbqDevice
	for rows.Next() {
		var i BbqDevice
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Location,
			&i.IsDefault,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReadingsBySessionID = `-- name: GetReadingsBySessionID :many
select id, session_id, probe_number, temperature, reading_occurred from bbq.sensor_readings 
where session_id = $1
order by reading_occurred desc, probe_number desc
`

func (q *Queries) GetReadingsBySessionID(ctx context.Context, db DBTX, sessionID uuid.UUID) ([]*BbqSensorReading, error) {
	rows, err := db.Query(ctx, getReadingsBySessionID, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*BbqSensorReading
	for rows.Next() {
		var i BbqSensorReading
		if err := rows.Scan(
			&i.ID,
			&i.SessionID,
			&i.ProbeNumber,
			&i.Temperature,
			&i.ReadingOccurred,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSensorByID = `-- name: GetSensorByID :one
select id, name, description, is_default from bbq.sensors where id = $1
`

func (q *Queries) GetSensorByID(ctx context.Context, db DBTX, id uuid.UUID) (*BbqSensor, error) {
	row := db.QueryRow(ctx, getSensorByID, id)
	var i BbqSensor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.IsDefault,
	)
	return &i, err
}

const getSensorByName = `-- name: GetSensorByName :one
select id, name, description, is_default from bbq.sensors where name = $1
`

func (q *Queries) GetSensorByName(ctx context.Context, db DBTX, name string) (*BbqSensor, error) {
	row := db.QueryRow(ctx, getSensorByName, name)
	var i BbqSensor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.IsDefault,
	)
	return &i, err
}

const getSensors = `-- name: GetSensors :many
select id, name, description, is_default from bbq.sensors
`

func (q *Queries) GetSensors(ctx context.Context, db DBTX) ([]*BbqSensor, error) {
	rows, err := db.Query(ctx, getSensors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*BbqSensor
	for rows.Next() {
		var i BbqSensor
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.IsDefault,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSessionByID = `-- name: GetSessionByID :one
select id, device_id, desired_state, start_time, end_time, sensor_id, session_type, subject_id, description from bbq.sessions
where id = $1
`

func (q *Queries) GetSessionByID(ctx context.Context, db DBTX, id uuid.UUID) (*BbqSession, error) {
	row := db.QueryRow(ctx, getSessionByID, id)
	var i BbqSession
	err := row.Scan(
		&i.ID,
		&i.DeviceID,
		&i.DesiredState,
		&i.StartTime,
		&i.EndTime,
		&i.SensorID,
		&i.SessionType,
		&i.SubjectID,
		&i.Description,
	)
	return &i, err
}

const getSessions = `-- name: GetSessions :many
select id, device_id, desired_state, start_time, end_time, sensor_id, session_type, subject_id, description from bbq.sessions
order by start_time desc
`

func (q *Queries) GetSessions(ctx context.Context, db DBTX) ([]*BbqSession, error) {
	rows, err := db.Query(ctx, getSessions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*BbqSession
	for rows.Next() {
		var i BbqSession
		if err := rows.Scan(
			&i.ID,
			&i.DeviceID,
			&i.DesiredState,
			&i.StartTime,
			&i.EndTime,
			&i.SensorID,
			&i.SessionType,
			&i.SubjectID,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubjectByID = `-- name: GetSubjectByID :one
select id, name, description from bbq.subjects where id = $1
`

func (q *Queries) GetSubjectByID(ctx context.Context, db DBTX, id uuid.UUID) (*BbqSubject, error) {
	row := db.QueryRow(ctx, getSubjectByID, id)
	var i BbqSubject
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return &i, err
}

const getSubjectStatesBySubjectID = `-- name: GetSubjectStatesBySubjectID :many
select id, subject_id, state, temperature from bbq.subject_states where subject_id = $1
`

func (q *Queries) GetSubjectStatesBySubjectID(ctx context.Context, db DBTX, subjectID uuid.UUID) ([]*BbqSubjectState, error) {
	rows, err := db.Query(ctx, getSubjectStatesBySubjectID, subjectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*BbqSubjectState
	for rows.Next() {
		var i BbqSubjectState
		if err := rows.Scan(
			&i.ID,
			&i.SubjectID,
			&i.State,
			&i.Temperature,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubjects = `-- name: GetSubjects :many
select id, name, description from bbq.subjects
`

func (q *Queries) GetSubjects(ctx context.Context, db DBTX) ([]*BbqSubject, error) {
	rows, err := db.Query(ctx, getSubjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*BbqSubject
	for rows.Next() {
		var i BbqSubject
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertDevice = `-- name: InsertDevice :one
insert into bbq.devices
(name, location)
values
($1, $2)
returning id, name, location, is_default
`

type InsertDeviceParams struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (q *Queries) InsertDevice(ctx context.Context, db DBTX, arg *InsertDeviceParams) (*BbqDevice, error) {
	row := db.QueryRow(ctx, insertDevice, arg.Name, arg.Location)
	var i BbqDevice
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.IsDefault,
	)
	return &i, err
}

const insertSensor = `-- name: InsertSensor :one
insert into bbq.sensors
(name, description)
values
($1, $2)
returning id, name, description, is_default
`

type InsertSensorParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) InsertSensor(ctx context.Context, db DBTX, arg *InsertSensorParams) (*BbqSensor, error) {
	row := db.QueryRow(ctx, insertSensor, arg.Name, arg.Description)
	var i BbqSensor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.IsDefault,
	)
	return &i, err
}

const insertSensorReading = `-- name: InsertSensorReading :exec
insert into bbq.sensor_readings
(session_id, probe_number, temperature, reading_occurred)
values
($1, $2, $3, $4)
`

type InsertSensorReadingParams struct {
	SessionID       uuid.UUID `json:"session_id"`
	ProbeNumber     int32     `json:"probe_number"`
	Temperature     float64   `json:"temperature"`
	ReadingOccurred int64     `json:"reading_occurred"`
}

func (q *Queries) InsertSensorReading(ctx context.Context, db DBTX, arg *InsertSensorReadingParams) error {
	_, err := db.Exec(ctx, insertSensorReading,
		arg.SessionID,
		arg.ProbeNumber,
		arg.Temperature,
		arg.ReadingOccurred,
	)
	return err
}

const insertSession = `-- name: InsertSession :one
insert into bbq.sessions
(device_id, desired_state, description, start_time, sensor_id, session_type, subject_id)
values 
($1, $2, $3, $4, $5, $6, $7)
returning id, device_id, desired_state, start_time, end_time, sensor_id, session_type, subject_id, description
`

type InsertSessionParams struct {
	DeviceID     uuid.UUID `json:"device_id"`
	DesiredState uuid.UUID `json:"desired_state"`
	Description  string    `json:"description"`
	StartTime    int64     `json:"start_time"`
	SensorID     uuid.UUID `json:"sensor_id"`
	SessionType  int32     `json:"session_type"`
	SubjectID    uuid.UUID `json:"subject_id"`
}

func (q *Queries) InsertSession(ctx context.Context, db DBTX, arg *InsertSessionParams) (*BbqSession, error) {
	row := db.QueryRow(ctx, insertSession,
		arg.DeviceID,
		arg.DesiredState,
		arg.Description,
		arg.StartTime,
		arg.SensorID,
		arg.SessionType,
		arg.SubjectID,
	)
	var i BbqSession
	err := row.Scan(
		&i.ID,
		&i.DeviceID,
		&i.DesiredState,
		&i.StartTime,
		&i.EndTime,
		&i.SensorID,
		&i.SessionType,
		&i.SubjectID,
		&i.Description,
	)
	return &i, err
}

const insertSubject = `-- name: InsertSubject :one
insert into bbq.subjects
(name, description)
values
($1, $2)
returning id, name, description
`

type InsertSubjectParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) InsertSubject(ctx context.Context, db DBTX, arg *InsertSubjectParams) (*BbqSubject, error) {
	row := db.QueryRow(ctx, insertSubject, arg.Name, arg.Description)
	var i BbqSubject
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return &i, err
}

const insertSubjectState = `-- name: InsertSubjectState :one
insert into bbq.subject_states
(subject_id, state, temperature)
values
($1, $2, $3)
returning id, subject_id, state, temperature
`

type InsertSubjectStateParams struct {
	SubjectID   uuid.UUID `json:"subject_id"`
	State       string    `json:"state"`
	Temperature int32     `json:"temperature"`
}

func (q *Queries) InsertSubjectState(ctx context.Context, db DBTX, arg *InsertSubjectStateParams) (*BbqSubjectState, error) {
	row := db.QueryRow(ctx, insertSubjectState, arg.SubjectID, arg.State, arg.Temperature)
	var i BbqSubjectState
	err := row.Scan(
		&i.ID,
		&i.SubjectID,
		&i.State,
		&i.Temperature,
	)
	return &i, err
}

const updateDevice = `-- name: UpdateDevice :one
update bbq.devices
set name = $2, location = $3
where id = $1
returning id, name, location, is_default
`

type UpdateDeviceParams struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
}

func (q *Queries) UpdateDevice(ctx context.Context, db DBTX, arg *UpdateDeviceParams) (*BbqDevice, error) {
	row := db.QueryRow(ctx, updateDevice, arg.ID, arg.Name, arg.Location)
	var i BbqDevice
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.IsDefault,
	)
	return &i, err
}

const updateSensor = `-- name: UpdateSensor :one
update bbq.sensors
set name = $2, description = $3
where id = $1
returning id, name, description, is_default
`

type UpdateSensorParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (q *Queries) UpdateSensor(ctx context.Context, db DBTX, arg *UpdateSensorParams) (*BbqSensor, error) {
	row := db.QueryRow(ctx, updateSensor, arg.ID, arg.Name, arg.Description)
	var i BbqSensor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.IsDefault,
	)
	return &i, err
}

const updateSession = `-- name: UpdateSession :one
update bbq.sessions
set device_id = $2, desired_state = $3, description = $4, start_time = $5, end_time = $6
where id = $1
returning id, device_id, desired_state, start_time, end_time, sensor_id, session_type, subject_id, description
`

type UpdateSessionParams struct {
	ID           uuid.UUID `json:"id"`
	DeviceID     uuid.UUID `json:"device_id"`
	DesiredState uuid.UUID `json:"desired_state"`
	Description  string    `json:"description"`
	StartTime    int64     `json:"start_time"`
	EndTime      *int64    `json:"end_time"`
}

func (q *Queries) UpdateSession(ctx context.Context, db DBTX, arg *UpdateSessionParams) (*BbqSession, error) {
	row := db.QueryRow(ctx, updateSession,
		arg.ID,
		arg.DeviceID,
		arg.DesiredState,
		arg.Description,
		arg.StartTime,
		arg.EndTime,
	)
	var i BbqSession
	err := row.Scan(
		&i.ID,
		&i.DeviceID,
		&i.DesiredState,
		&i.StartTime,
		&i.EndTime,
		&i.SensorID,
		&i.SessionType,
		&i.SubjectID,
		&i.Description,
	)
	return &i, err
}

const updateSubject = `-- name: UpdateSubject :one
update bbq.subjects
set name = $2, description = $3
where id = $1
returning id, name, description
`

type UpdateSubjectParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (q *Queries) UpdateSubject(ctx context.Context, db DBTX, arg *UpdateSubjectParams) (*BbqSubject, error) {
	row := db.QueryRow(ctx, updateSubject, arg.ID, arg.Name, arg.Description)
	var i BbqSubject
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return &i, err
}