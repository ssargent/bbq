alter table bbq.sessions drop column description;
alter table bbq.sessions add description varchar(128) not null default('');