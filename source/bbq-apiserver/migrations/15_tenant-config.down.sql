begin;

    alter table sys.config drop canoverride;
    drop table sys.tenant_config;

commit;