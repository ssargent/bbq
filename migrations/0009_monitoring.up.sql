create table bbq.sensor_alerts (
    id uuid not null default uuid_generate_v4(),
    alert_name varchar(128) not null,
    description text not null,
    notification_rules text not null,
    constraint pk_sensoralerts_id primary key (id)
);

create table bbq.sensor_alert_rules (
    id uuid not null default uuid_generate_v4(),
    alert_id uuid not null,
    rule_name varchar(128) not null,
    rule_min float8 not null,
    rule_max float8 not null,
    constraint pk_sensoralertrules_id primary key (id),
    constraint fk_sensoralertrules_alerts foreign key (alert_id) references bbq.sensor_alerts (id)
);

create index ix_sensor_alert_rules_alerts on bbq.sensor_alert_rules (alert_id);

create table bbq.session_alerts (
    session_id uuid not null,
    alert_id uuid not null,
    alert_suspended boolean not null,
    alert_start bigint not null,
    alert_end bigint,
    constraint pk_sessionalerts_id primary key (session_id, alert_id),
    constraint fk_sessionalerts_sessions foreign key (session_id) references bbq.sessions (id),
    constraint fk_sessionalerts_alerts foreign key (alert_id) references bbq.sensor_alerts (id)
)