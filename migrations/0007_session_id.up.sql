alter table bbq.sessions alter column id set default(uuid_generate_v4());