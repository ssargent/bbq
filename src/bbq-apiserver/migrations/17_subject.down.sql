begin;
    alter table bbq.subjects drop constraint fk_subjects_tenants;
    alter table bbq.subjects drop column tenantid;
commit;