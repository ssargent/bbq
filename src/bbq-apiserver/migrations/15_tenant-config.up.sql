begin;

	alter table sys.config add canoverride boolean not null default(FALSE);

	create table sys.tenant_config (
		id uuid not null default(uuid_generate_v4()),
		tenantid uuid not null,
		settingkey varchar(64) not null,
		settingvalue varchar(512) not null,
		constraint pk_tenantsettings_id primary key (id),
		constraint fk_tenantsettings_tenants foreign key (tenantid) references sys.tenants (id)
	);

	
	
commit;