-- create a default subject and a default state for that subject.
insert into bbq.subjects (id, name, description) values (uuid_nil(), 'tbd', 'To Be Determined - Not yet classified');
insert into bbq.subject_states (id, subject_id, state, temperature) values (uuid_nil(), uuid_nil(), 'tbd', 500);

alter table bbq.sessions add subject_id uuid not null default(uuid_nil());
alter table bbq.sessions add constraint fk_sessions_subjects foreign key (subject_id) references bbq.subjects (id);

create index ix_session_subject on bbq.sessions (subject_id);

alter table bbq.sessions alter column desired_state set default(uuid_nil());