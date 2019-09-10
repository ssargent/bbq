begin;

alter table bbq.devices drop constraint devices_unique_tenant_name;
alter table bbq.monitors drop constraint monitors_unique_tenant_name;

drop index devices_tenant_name;
drop index monitors_tenant_name;

alter table bbq.monitors drop constraint fk_monitors_tenants;
alter table bbq.devies drop constraint fk_devices_tenants;
alter table bbq.sessions drop constraint fk_sessions_tenants;

alter table bbq.monitors drop tenantid;
alter table bbq.devices drop tenantid;
alter table bbq.sessions drop tenantid;

drop table sys.tenants;

commit;