begin;

	create schema
	if not exists system;

create table system.config
(
	id bigserial not null,
	settingkey varchar(64) not null,
	settingvalue varchar(512) not null,
	constraint pk_config_id primary key (id)
);

create table system.config_changes
(
	id bigserial not null,
	settingkey varchar(64) not null,
	settingvalue varchar(512) not null,
	change_date timestamptz not null,
	constraint pk_changes_id primary key (id)
);

commit;