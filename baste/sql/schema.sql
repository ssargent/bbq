--
-- PostgreSQL database dump
--

-- Dumped from database version 14.0
-- Dumped by pg_dump version 15.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: bbq; Type: SCHEMA; Schema: -; Owner: smokey
--

CREATE SCHEMA bbq;


ALTER SCHEMA bbq OWNER TO smokey;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO postgres;

--
-- Name: tablefunc; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS tablefunc WITH SCHEMA public;


--
-- Name: EXTENSION tablefunc; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION tablefunc IS 'functions that manipulate whole tables, including crosstab';


--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: devices; Type: TABLE; Schema: bbq; Owner: smokey
--

CREATE TABLE bbq.devices (
    id uuid NOT NULL,
    name character varying(128) NOT NULL,
    location character varying(128) NOT NULL,
    is_default boolean DEFAULT false NOT NULL
);


ALTER TABLE bbq.devices OWNER TO smokey;

--
-- Name: sensor_alert_rules; Type: TABLE; Schema: bbq; Owner: smokey
--

CREATE TABLE bbq.sensor_alert_rules (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    alert_id uuid NOT NULL,
    rule_name character varying(128) NOT NULL,
    rule_min double precision NOT NULL,
    rule_max double precision NOT NULL
);


ALTER TABLE bbq.sensor_alert_rules OWNER TO smokey;

--
-- Name: sensor_alerts; Type: TABLE; Schema: bbq; Owner: smokey
--

CREATE TABLE bbq.sensor_alerts (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    alert_name character varying(128) NOT NULL,
    description text NOT NULL,
    notification_rules text NOT NULL
);


ALTER TABLE bbq.sensor_alerts OWNER TO smokey;

--
-- Name: sensor_readings; Type: TABLE; Schema: bbq; Owner: smokey
--

CREATE TABLE bbq.sensor_readings (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    session_id uuid NOT NULL,
    probe_number integer NOT NULL,
    temperature double precision NOT NULL,
    reading_occurred bigint NOT NULL
);


ALTER TABLE bbq.sensor_readings OWNER TO smokey;

--
-- Name: sensors; Type: TABLE; Schema: bbq; Owner: smokey
--

CREATE TABLE bbq.sensors (
    id uuid NOT NULL,
    name character varying(128) NOT NULL,
    description text NOT NULL,
    is_default boolean DEFAULT false NOT NULL
);


ALTER TABLE bbq.sensors OWNER TO smokey;

--
-- Name: session_alerts; Type: TABLE; Schema: bbq; Owner: smokey
--

CREATE TABLE bbq.session_alerts (
    session_id uuid NOT NULL,
    alert_id uuid NOT NULL,
    alert_suspended boolean NOT NULL,
    alert_start bigint NOT NULL,
    alert_end bigint
);


ALTER TABLE bbq.session_alerts OWNER TO smokey;

--
-- Name: sessions; Type: TABLE; Schema: bbq; Owner: smokey
--

CREATE TABLE bbq.sessions (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    device_id uuid NOT NULL,
    desired_state uuid DEFAULT public.uuid_nil() NOT NULL,
    start_time bigint NOT NULL,
    end_time bigint,
    sensor_id uuid DEFAULT public.uuid_nil() NOT NULL,
    session_type integer DEFAULT 1 NOT NULL,
    subject_id uuid DEFAULT public.uuid_nil() NOT NULL,
    description character varying(128) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE bbq.sessions OWNER TO smokey;

--
-- Name: subject_states; Type: TABLE; Schema: bbq; Owner: smokey
--

CREATE TABLE bbq.subject_states (
    id uuid NOT NULL,
    subject_id uuid NOT NULL,
    state character varying(64) NOT NULL,
    temperature integer NOT NULL
);


ALTER TABLE bbq.subject_states OWNER TO smokey;

--
-- Name: subjects; Type: TABLE; Schema: bbq; Owner: smokey
--

CREATE TABLE bbq.subjects (
    id uuid NOT NULL,
    name character varying(128) NOT NULL,
    description text NOT NULL
);


ALTER TABLE bbq.subjects OWNER TO smokey;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: smokey
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO smokey;

--
-- Name: devices pk_devices_id; Type: CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.devices
    ADD CONSTRAINT pk_devices_id PRIMARY KEY (id);


--
-- Name: sensor_readings pk_readings_id; Type: CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.sensor_readings
    ADD CONSTRAINT pk_readings_id PRIMARY KEY (id);


--
-- Name: sensor_alert_rules pk_sensoralertrules_id; Type: CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.sensor_alert_rules
    ADD CONSTRAINT pk_sensoralertrules_id PRIMARY KEY (id);


--
-- Name: sensor_alerts pk_sensoralerts_id; Type: CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.sensor_alerts
    ADD CONSTRAINT pk_sensoralerts_id PRIMARY KEY (id);


--
-- Name: sensors pk_sensors_id; Type: CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.sensors
    ADD CONSTRAINT pk_sensors_id PRIMARY KEY (id);


--
-- Name: session_alerts pk_sessionalerts_id; Type: CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.session_alerts
    ADD CONSTRAINT pk_sessionalerts_id PRIMARY KEY (session_id, alert_id);


--
-- Name: sessions pk_sessions_id; Type: CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.sessions
    ADD CONSTRAINT pk_sessions_id PRIMARY KEY (id);


--
-- Name: subject_states pk_states_id; Type: CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.subject_states
    ADD CONSTRAINT pk_states_id PRIMARY KEY (id);


--
-- Name: subjects pk_subjects_id; Type: CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.subjects
    ADD CONSTRAINT pk_subjects_id PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: smokey
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: ix_readings_sessions; Type: INDEX; Schema: bbq; Owner: smokey
--

CREATE INDEX ix_readings_sessions ON bbq.sensor_readings USING btree (session_id);


--
-- Name: ix_sensor_alert_rules_alerts; Type: INDEX; Schema: bbq; Owner: smokey
--

CREATE INDEX ix_sensor_alert_rules_alerts ON bbq.sensor_alert_rules USING btree (alert_id);


--
-- Name: ix_session_subject; Type: INDEX; Schema: bbq; Owner: smokey
--

CREATE INDEX ix_session_subject ON bbq.sessions USING btree (subject_id);


--
-- Name: ix_session_type; Type: INDEX; Schema: bbq; Owner: smokey
--

CREATE INDEX ix_session_type ON bbq.sessions USING btree (session_type);


--
-- Name: ix_sessions_devices; Type: INDEX; Schema: bbq; Owner: smokey
--

CREATE INDEX ix_sessions_devices ON bbq.sessions USING btree (device_id);


--
-- Name: ix_sessions_sensors; Type: INDEX; Schema: bbq; Owner: smokey
--

CREATE INDEX ix_sessions_sensors ON bbq.sessions USING btree (sensor_id);


--
-- Name: ix_sessions_states; Type: INDEX; Schema: bbq; Owner: smokey
--

CREATE INDEX ix_sessions_states ON bbq.sessions USING btree (desired_state);


--
-- Name: ix_states_subject; Type: INDEX; Schema: bbq; Owner: smokey
--

CREATE INDEX ix_states_subject ON bbq.subject_states USING btree (subject_id);


--
-- Name: sensor_readings fk_readings_sessions; Type: FK CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.sensor_readings
    ADD CONSTRAINT fk_readings_sessions FOREIGN KEY (session_id) REFERENCES bbq.sessions(id);


--
-- Name: sensor_alert_rules fk_sensoralertrules_alerts; Type: FK CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.sensor_alert_rules
    ADD CONSTRAINT fk_sensoralertrules_alerts FOREIGN KEY (alert_id) REFERENCES bbq.sensor_alerts(id);


--
-- Name: session_alerts fk_sessionalerts_alerts; Type: FK CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.session_alerts
    ADD CONSTRAINT fk_sessionalerts_alerts FOREIGN KEY (alert_id) REFERENCES bbq.sensor_alerts(id);


--
-- Name: session_alerts fk_sessionalerts_sessions; Type: FK CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.session_alerts
    ADD CONSTRAINT fk_sessionalerts_sessions FOREIGN KEY (session_id) REFERENCES bbq.sessions(id);


--
-- Name: sessions fk_sessions_devices; Type: FK CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.sessions
    ADD CONSTRAINT fk_sessions_devices FOREIGN KEY (device_id) REFERENCES bbq.devices(id);


--
-- Name: sessions fk_sessions_sensors; Type: FK CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.sessions
    ADD CONSTRAINT fk_sessions_sensors FOREIGN KEY (sensor_id) REFERENCES bbq.sensors(id);


--
-- Name: sessions fk_sessions_states; Type: FK CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.sessions
    ADD CONSTRAINT fk_sessions_states FOREIGN KEY (desired_state) REFERENCES bbq.subject_states(id);


--
-- Name: sessions fk_sessions_subjects; Type: FK CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.sessions
    ADD CONSTRAINT fk_sessions_subjects FOREIGN KEY (subject_id) REFERENCES bbq.subjects(id);


--
-- Name: subject_states fk_states_subjects; Type: FK CONSTRAINT; Schema: bbq; Owner: smokey
--

ALTER TABLE ONLY bbq.subject_states
    ADD CONSTRAINT fk_states_subjects FOREIGN KEY (subject_id) REFERENCES bbq.subjects(id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

