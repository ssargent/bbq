begin;

    alter table bbq.sessions drop constraint pk_sessions_ingredients;
    alter table bbq.sessions rename ingredientid to subjectid;

    alter table bbq.sessions add weight numeric(5,2) not null default(-1.0);

    alter table bbq.ingredients rename to subjects;

    alter table bbq.sessions add constraint fk_sessions_subjects foreign key (subjectid) references bbq.subjects (id);

commit;