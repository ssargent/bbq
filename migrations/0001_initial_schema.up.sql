--//Modify this code to update the diagram
create schema if not exists bbq;

create table bbq.devices (
    id uuid not null,
    name varchar(128) not null,
    location varchar(128) not null,
    constraint pk_devices_id primary key (id)
);

create table bbq.subjects (
    id uuid not null,
    name varchar(128) not null,
    description text not null,
    constraint pk_subjects_id primary key (id)
);

create table bbq.subject_states (
    id uuid not null,
    subject_id uuid not null,
    state varchar(64) not null,
    temperature int not null,
    constraint pk_states_id primary key (id),
    constraint fk_states_subjects foreign key (subject_id) references bbq.subjects (id)
);

create index ix_states_subject on bbq.subject_states (subject_id);


create table bbq.sessions (
    id uuid not null,
    device_id uuid not null,
    desired_state uuid not null,
    description uuid not null,
    start_time bigint not null,
    end_time bigint,
    constraint pk_sessions_id primary key (id),
    constraint fk_sessions_devices foreign key (device_id) references bbq.devices (id),
    constraint fk_sessions_states foreign key (desired_state) references bbq.subject_states (id)
);

create index ix_sessions_devices on bbq.sessions (device_id);
create index ix_sessions_states on bbq.sessions (desired_state);


create table bbq.sensors (
    id uuid not null,
    name varchar(128) not null,
    description text not null,
    constraint pk_sensors_id primary key (id)
);

create table bbq.sensor_readings (
    id uuid not null,
    session_id uuid not null,
    sensor_id uuid not null,
    probe_number int not null,
    temperature float8 not null,
    reading_occurred bigint not null,
    constraint pk_readings_id primary key (id),
    constraint fk_readings_sessions foreign key (session_id) references bbq.sessions (id),
    constraint fk_readings_sensors foreign key (sensor_id) references bbq.sensors (id)
);

create index ix_readings_sessions on bbq.sensor_readings (session_id);
create index ix_readings_sensors on bbq.sensor_readings (sensor_id);
 