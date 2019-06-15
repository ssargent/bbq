BEGIN;
    create table bbq.ingredients (
        id bigserial not null,
        type text not null,
        name text not null,
        description text not null,
        constraint pk_ingredients_id primary key (id)
    );

    insert into bbq.ingredients (id, type, name, description) VALUES (-1, 'N/A', 'Not Specified', 'Not Specified');

    alter table bbq.sessions add ingredientid bigint not null default(-1);
    alter table bbq.sessions add constraint pk_sessions_ingredients foreign key (ingredientid) references bbq.ingredients (id);
commit;