begin;
    create table sys.toggles (
        id uuid not null default(uuid_generate_v4()),
        togglekey varchar(32) not null,
        toggleenabled bit not null default(false),
        overrideallowed bit not null default(false),
		constraint pk_toggles_id primary key (id)
    );

    create table sys.tenant_toggles (
        id uuid not null default(uuid_generate_v4()),
        toggleid uuid not null, 
        toggleenabled bit not null default(false),
		constraint pk_tenanttoggles_id primary key (id),
        constraint fk_tenanttoggles_toggles foreign key (toggleid) references sys.toggles (id)
    );

--   drop table sys.tenant_config;
--   drop table sys.config;

commit;