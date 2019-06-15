begin;
    create table bbq.ingredient_types (
        id bigserial not null,
        name text not null,
        urlkey text not null,
        description text not null,
        constraint pk_ingredient_types_id primary key (id)
    );

    insert into bbq.ingredient_types (id, name, urlkey, description) values (-1, 'Uncategorized', 'na', 'Uncategorized Ingredients - Use this temporarily');

    alter table bbq.ingredients add typeid bigint not null default (-1);
    alter table bbq.ingredients add constraint fk_ingredients_types foreign key (typeid) references bbq.ingredient_types (id);
    alter table bbq.ingredients drop type;  

commit;
