begin;
    alter table bbq.subjects add tenantid uuid null;
    alter table bbq.subjects add constraint fk_subjects_tenants foreign key (tenantid) references sys.tenants (id);
commit;