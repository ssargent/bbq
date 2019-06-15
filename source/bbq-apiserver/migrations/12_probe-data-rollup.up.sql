begin;

CREATE TABLE data.bbq_temp_rollup
(
    id bigserial not null,
    probe0 numeric(5,2) NOT NULL,
    probe1 numeric(5,2) NOT NULL,
    probe2 numeric(5,2) NOT NULL,
    probe3 numeric(5,2) NOT NULL,
    recordedat timestamp with time zone NOT NULL DEFAULT now(),
    sessionid uuid,
    CONSTRAINT pk_bbq_temp_rollup_id PRIMARY KEY (id),
    CONSTRAINT fk_temp_session FOREIGN KEY (sessionid)
        REFERENCES bbq.sessions (uid) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

commit;