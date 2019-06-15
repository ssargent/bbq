begin;

	alter table bbq.monitors add uid uuid not null default(uuid_generate_v4());
	alter table bbq.devices add uid uuid not null default(uuid_generate_v4());
	alter table bbq.subjects add uid uuid not null default(uuid_generate_v4());
	
commit;