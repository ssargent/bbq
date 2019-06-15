begin;

	alter table data.bbq_temp_readings drop constraint fk_temp_session;
	alter table data.bbq_temp_readings drop sessionid;

	drop view bbq.vw_sessions;

	alter table bbq.sessions drop constraint pk_session_uid;
	alter table bbq.sessions drop uid;
	alter table bbq.sessions add constraint pk_session_id primary key (id);

	create or replace view bbq.vw_sessions as
	select
		s.id,
		s.name,
		s.description,
		sub.name as subject,
		it.name as type,
		s.weight,
		d.name as  device,
		m.name as  monitor,
		s.starttime
	from bbq.sessions s inner join bbq.subjects sub on s.subjectid = sub.id
		inner join bbq.ingredient_types it on sub.typeid = it.id
		inner join bbq.monitors m on s.monitorid = m.id
		inner join bbq.devices d on s.deviceid = d.id;

commit;
