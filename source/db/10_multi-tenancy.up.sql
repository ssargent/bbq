begin;

create table sys.tenants (
	id uuid not null default(uuid_generate_v4()),
	name text not null,
	urlkey text not null,
	isenabled boolean not null default(TRUE),
	constraint pk_tenants_id primary key (id)
);
	
insert into sys.tenants (id, name, urlkey) values (uuid_generate_v5(uuid_ns_dns(), 'development.myfamilycooks.com'), 'development', 'development');

alter table bbq.monitors add tenantid uuid not null default(uuid_generate_v5(uuid_ns_dns(), 'development.myfamilycooks.com'));
alter table bbq.devices add tenantid uuid not null default(uuid_generate_v5(uuid_ns_dns(), 'development.myfamilycooks.com'));
alter table bbq.sessions add tenantid uuid not null default(uuid_generate_v5(uuid_ns_dns(), 'development.myfamilycooks.com'));

alter table bbq.monitors add constraint fk_monitors_tenants foreign key (tenantid) references sys.tenants (id);
alter table bbq.devices add constraint fk_devices_tenants foreign key (tenantid) references sys.tenants (id);
alter table bbq.sessions add constraint fk_sessions_tenants foreign key (tenantid) references sys.tenants (id);
																			 
CREATE UNIQUE INDEX  devices_tenant_name ON bbq.devices (tenantid, name);
CREATE UNIQUE INDEX  monitors_tenant_name ON bbq.monitors (tenantid, name); 

ALTER TABLE bbq.devices ADD CONSTRAINT devices_unique_tenant_name UNIQUE USING INDEX devices_tenant_name;
ALTER TABLE bbq.monitors ADD CONSTRAINT monitors_unique_tenant_name UNIQUE USING INDEX monitors_tenant_name;

commit;