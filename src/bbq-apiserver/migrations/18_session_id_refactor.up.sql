begin;
    alter table bbq.sessions add deviceuid uuid null;
    alter table bbq.sessions add monitoruid uuid null;
    alter table bbq.sessions add subjectuid uuid null;

    update bbq.sessions set deviceuid = d.uid from bbq.devices d inner join bbq.sessions s on d.id = s.deviceid;
    update bbq.sessions set monitoruid = m.uid from bbq.monitors m inner join bbq.sessions s on m.id = s.monitorid;
    update bbq.sessions set subjectuid = sb.uid from bbq.subjects sb inner join bbq.sessions s on sb.id = s.subjectid;

commit;