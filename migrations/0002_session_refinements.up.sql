-- create a default sensor so data can flow in.
insert into bbq.sensors (id, name, description) values (uuid_nil(), 'default-sensor', 'Default Sensor');
-- create a new sensor field on sessions 
alter table bbq.sessions add sensor_id uuid not null default(uuid_nil());
-- make that field a foreign key
alter table bbq.sessions add constraint fk_sessions_sensors foreign key (sensor_id) references bbq.sensors (id);
-- create index on the fk field.
create index ix_sessions_sensors on bbq.sessions (sensor_id);

-- at runtime we'll pull the default sensor from the db.
alter table bbq.sensors add is_default boolean not null default(FALSE);
alter table bbq.devices add is_default boolean not null default(FALSE);

alter table bbq.sensor_readings drop sensor_id; 