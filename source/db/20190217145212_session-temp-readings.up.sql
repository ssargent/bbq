begin;

alter table bbq.sessions drop constraint pk_sessions_id;
alter table bbq.sessions add uid uuid default(uuid_generate_v4());
alter table bbq.sessions add constraint pk_session_uid primary key (uid);

CREATE OR REPLACE VIEW bbq.vw_sessions AS
 SELECT s.id,
    s.name,
    s.description,
    sub.name AS subject,
    it.name AS type,
    s.weight,
    d.name AS device,
    m.name AS monitor,
    s.starttime,
    s.tenantid,
    s.uid
   FROM bbq.sessions s
     JOIN bbq.subjects sub ON s.subjectid = sub.id
     JOIN bbq.ingredient_types it ON sub.typeid = it.id
     JOIN bbq.monitors m ON s.monitorid = m.id
     JOIN bbq.devices d ON s.deviceid = d.id;;

alter table data.bbq_temp_readings add sessionid uuid null;
alter table data.bbq_temp_readings add constraint fk_temp_session foreign key (sessionid) references bbq.sessions(uid);

commit;