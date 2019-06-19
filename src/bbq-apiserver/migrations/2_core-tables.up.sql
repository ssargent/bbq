begin;

create table datacollection.bbq_temp_readings (
	id bigserial not null,
	probe0 numeric(5,2) NOT NULL,
    probe1 numeric(5,2) NOT NULL,
    probe2 numeric(5,2) NOT NULL,
    probe3 numeric(5,2) NOT NULL,
    recordedat timestamptz not null default now(),
	constraint pk_bbq_temp_readings_id primary key (id)
);

create table datacollection.weather_readings (
	id bigserial not null,
	location point not null,
	main text not null,
	description text,
	temp numeric(6,3) not null,
	pressure int not null,
	humidity int not null,
	wind_speed numeric(5,2) not null,
	wind_direction int not null,
	wind_gust numeric(5,2) not null,
	recordedat timestamptz not null default now(),
	constraint pk_weather_readings_id primary key (id)
);

commit;