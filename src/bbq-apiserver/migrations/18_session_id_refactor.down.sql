begin;
    alter table bbq.sessions drop deviceuid;
    alter table bbq.sessions drop monitoruid;
    alter table bbq.sessions drop subjectuid;

commit;