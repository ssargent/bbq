-- name: GetDevices :many
select * from bbq.devices;

-- name: GetDeviceByID :one
select * from bbq.devices where id = $1;

-- name: GetDeviceByName :one 
select * from bbq.devices where name = $1;

-- name: GetDefaultDevice :one
select * from bbq.devices where is_default = true;

-- name: InsertDevice :one 
insert into bbq.devices
(name, location)
values
($1, $2)
returning *;

-- name: UpdateDevice :one
update bbq.devices
set name = $2, location = $3
where id = $1
returning *;

-- name: GetSensors :many
select * from bbq.sensors;

-- name: GetSensorByID :one
select * from bbq.sensors where id = $1;

-- name: GetSensorByName :one
select * from bbq.sensors where name = $1;

-- name: GetDefaultSensor :one
select * from bbq.sensors where is_default = true;

-- name: InsertSensor :one 
insert into bbq.sensors
(name, description)
values
($1, $2)
returning *;

-- name: UpdateSensor :one
update bbq.sensors
set name = $2, description = $3
where id = $1
returning *;

-- name: InsertSensorReading :exec
insert into bbq.sensor_readings
(session_id, probe_number, temperature, reading_occurred)
values
($1, $2, $3, $4);

-- name: GetReadingsBySessionID :many
select * from bbq.sensor_readings 
where session_id = $1
order by reading_occurred desc, probe_number desc;

-- name: InsertSession :one
insert into bbq.sessions
(device_id, desired_state, description, start_time, sensor_id, session_type, subject_id)
values 
($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: GetSessions :many
select * from bbq.sessions
order by start_time desc;

-- name: GetSessionByID :one
select * from bbq.sessions
where id = $1;

-- name: UpdateSession :one
update bbq.sessions
set device_id = $2, desired_state = $3, description = $4, start_time = $5, end_time = $6
where id = $1
returning *;

-- name: GetSubjects :many
select * from bbq.subjects;

-- name: GetSubjectByID :one
select * from bbq.subjects where id = $1;

-- name: InsertSubject :one 
insert into bbq.subjects
(name, description)
values
($1, $2)
returning *;

-- name: UpdateSubject :one
update bbq.subjects
set name = $2, description = $3
where id = $1
returning *;

-- name: GetSubjectStatesBySubjectID :many
select * from bbq.subject_states where subject_id = $1;

-- name: InsertSubjectState :one
insert into bbq.subject_states
(subject_id, state, temperature)
values
($1, $2, $3)
returning *;

-- name: DeleteSubjectState :exec
delete from bbq.subject_states where id = $1;
