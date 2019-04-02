begin;

	create table sys.accounts (
		id uuid not null default(uuid_generate_v4()),
		loginname text not null,
		loginpassword text not null,
		fullname text not null,
		email text not null,
		isenabled boolean not null default(TRUE),
		tenantid uuid not null,
		constraint pk_accounts_id primary key (id),
		constraint fk_accounts_tenants foreign key (tenantid) references sys.tenants (id)
	);

	
commit;