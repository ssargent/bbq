begin;

	alter table bbq.monitors drop uid;
	alter table bbq.devices drop uid;
	alter table bbq.subjects drop uid;
	
commit;