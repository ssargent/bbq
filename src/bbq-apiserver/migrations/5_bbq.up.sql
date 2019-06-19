begin;

	create schema
	if not exists bbq;

create table bbq.devices
(
	id bigserial not null,
	name text not null,
	description text not null,
	constraint pk_devices_id primary key (id)
);

create table bbq.monitors
(
	id bigserial not null,
	name text not null,
	description text not null,
	address varchar(32) not null,
	constraint pk_monitors_id primary key (id)
);


create table bbq.sessions
(
	id bigserial not null,
	deviceid bigint not null,
	monitorid bigint not null,
	name text not null,
	description text not null,
	starttime timestamptz not null default now(),
	constraint pk_sessions_id primary key (id),
	constraint fk_sessions_devices foreign key (deviceid) references bbq.devices (id),
	constraint fk_sessions_monitors foreign key (monitorid) references bbq.monitors (id)
);

commit;