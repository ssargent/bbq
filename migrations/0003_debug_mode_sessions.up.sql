alter table bbq.sessions add session_type int not null default(1);

create index ix_session_type on bbq.sessions (session_type);