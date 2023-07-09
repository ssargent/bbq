--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2
-- Dumped by pg_dump version 15.2

-- Started on 2023-06-27 21:04:58 EDT

/**********************************************************************
    IMPORTANT: DO NOT USE THIS SCRIPT TO CREATE DATABASE.  
    Use the migrations here: https://github.com/ssargent/bbq/tree/develop/migrations
    instead.  This uses gomigrate to safely update database schemas.

***********************************************************************/

--
-- TOC entry 6 (class 2615 OID 16714)
-- Name: bbq; Type: SCHEMA; Schema: -; Owner: scott
--

create schema bbq;
 
--
-- TOC entry 215 (class 1259 OID 16715)
-- Name: devices; Type: TABLE; Schema: bbq; Owner: scott
--

CREATE TABLE bbq.devices (
    id uuid NOT NULL,
    name character varying(128) NOT NULL,
    location character varying(128) NOT NULL
);


ALTER TABLE bbq.devices OWNER TO scott;

--
-- TOC entry 220 (class 1259 OID 16762)
-- Name: sensor_readings; Type: TABLE; Schema: bbq; Owner: scott
--

CREATE TABLE bbq.sensor_readings (
    id uuid NOT NULL,
    session_id uuid NOT NULL,
    sensor_id uuid NOT NULL,
    probe_number integer NOT NULL,
    temperature double precision NOT NULL,
    reading_occurred bigint NOT NULL
);


ALTER TABLE bbq.sensor_readings OWNER TO scott;

--
-- TOC entry 219 (class 1259 OID 16755)
-- Name: sensors; Type: TABLE; Schema: bbq; Owner: scott
--

CREATE TABLE bbq.sensors (
    id uuid NOT NULL,
    name character varying(128) NOT NULL,
    description text NOT NULL
);


ALTER TABLE bbq.sensors OWNER TO scott;

--
-- TOC entry 218 (class 1259 OID 16738)
-- Name: sessions; Type: TABLE; Schema: bbq; Owner: scott
--

CREATE TABLE bbq.sessions (
    id uuid NOT NULL,
    device_id uuid NOT NULL,
    desired_state uuid NOT NULL,
    description uuid NOT NULL,
    start_time bigint NOT NULL,
    end_time bigint
);


ALTER TABLE bbq.sessions OWNER TO scott;

--
-- TOC entry 217 (class 1259 OID 16727)
-- Name: subject_states; Type: TABLE; Schema: bbq; Owner: scott
--

CREATE TABLE bbq.subject_states (
    id uuid NOT NULL,
    subject_id uuid NOT NULL,
    state character varying(64) NOT NULL,
    temperature integer NOT NULL
);


ALTER TABLE bbq.subject_states OWNER TO scott;

--
-- TOC entry 216 (class 1259 OID 16720)
-- Name: subjects; Type: TABLE; Schema: bbq; Owner: scott
--

CREATE TABLE bbq.subjects (
    id uuid NOT NULL,
    name character varying(128) NOT NULL,
    description text NOT NULL
);


ALTER TABLE bbq.subjects OWNER TO scott;

--
-- TOC entry 3482 (class 2606 OID 16719)
-- Name: devices pk_devices_id; Type: CONSTRAINT; Schema: bbq; Owner: scott
--

ALTER TABLE ONLY bbq.devices
    ADD CONSTRAINT pk_devices_id PRIMARY KEY (id);


--
-- TOC entry 3497 (class 2606 OID 16766)
-- Name: sensor_readings pk_readings_id; Type: CONSTRAINT; Schema: bbq; Owner: scott
--

ALTER TABLE ONLY bbq.sensor_readings
    ADD CONSTRAINT pk_readings_id PRIMARY KEY (id);


--
-- TOC entry 3493 (class 2606 OID 16761)
-- Name: sensors pk_sensors_id; Type: CONSTRAINT; Schema: bbq; Owner: scott
--

ALTER TABLE ONLY bbq.sensors
    ADD CONSTRAINT pk_sensors_id PRIMARY KEY (id);


--
-- TOC entry 3491 (class 2606 OID 16742)
-- Name: sessions pk_sessions_id; Type: CONSTRAINT; Schema: bbq; Owner: scott
--

ALTER TABLE ONLY bbq.sessions
    ADD CONSTRAINT pk_sessions_id PRIMARY KEY (id);


--
-- TOC entry 3487 (class 2606 OID 16731)
-- Name: subject_states pk_states_id; Type: CONSTRAINT; Schema: bbq; Owner: scott
--

ALTER TABLE ONLY bbq.subject_states
    ADD CONSTRAINT pk_states_id PRIMARY KEY (id);


--
-- TOC entry 3484 (class 2606 OID 16726)
-- Name: subjects pk_subjects_id; Type: CONSTRAINT; Schema: bbq; Owner: scott
--

ALTER TABLE ONLY bbq.subjects
    ADD CONSTRAINT pk_subjects_id PRIMARY KEY (id);


--
-- TOC entry 3494 (class 1259 OID 16778)
-- Name: ix_readings_sensors; Type: INDEX; Schema: bbq; Owner: scott
--

CREATE INDEX ix_readings_sensors ON bbq.sensor_readings USING btree (sensor_id);


--
-- TOC entry 3495 (class 1259 OID 16777)
-- Name: ix_readings_sessions; Type: INDEX; Schema: bbq; Owner: scott
--

CREATE INDEX ix_readings_sessions ON bbq.sensor_readings USING btree (session_id);


--
-- TOC entry 3488 (class 1259 OID 16753)
-- Name: ix_sessions_devices; Type: INDEX; Schema: bbq; Owner: scott
--

CREATE INDEX ix_sessions_devices ON bbq.sessions USING btree (device_id);


--
-- TOC entry 3489 (class 1259 OID 16754)
-- Name: ix_sessions_states; Type: INDEX; Schema: bbq; Owner: scott
--

CREATE INDEX ix_sessions_states ON bbq.sessions USING btree (desired_state);


--
-- TOC entry 3485 (class 1259 OID 16737)
-- Name: ix_states_subject; Type: INDEX; Schema: bbq; Owner: scott
--

CREATE INDEX ix_states_subject ON bbq.subject_states USING btree (subject_id);


--
-- TOC entry 3501 (class 2606 OID 16772)
-- Name: sensor_readings fk_readings_sensors; Type: FK CONSTRAINT; Schema: bbq; Owner: scott
--

ALTER TABLE ONLY bbq.sensor_readings
    ADD CONSTRAINT fk_readings_sensors FOREIGN KEY (sensor_id) REFERENCES bbq.sensors(id);


--
-- TOC entry 3502 (class 2606 OID 16767)
-- Name: sensor_readings fk_readings_sessions; Type: FK CONSTRAINT; Schema: bbq; Owner: scott
--

ALTER TABLE ONLY bbq.sensor_readings
    ADD CONSTRAINT fk_readings_sessions FOREIGN KEY (session_id) REFERENCES bbq.sessions(id);


--
-- TOC entry 3499 (class 2606 OID 16743)
-- Name: sessions fk_sessions_devices; Type: FK CONSTRAINT; Schema: bbq; Owner: scott
--

ALTER TABLE ONLY bbq.sessions
    ADD CONSTRAINT fk_sessions_devices FOREIGN KEY (device_id) REFERENCES bbq.devices(id);


--
-- TOC entry 3500 (class 2606 OID 16748)
-- Name: sessions fk_sessions_states; Type: FK CONSTRAINT; Schema: bbq; Owner: scott
--

ALTER TABLE ONLY bbq.sessions
    ADD CONSTRAINT fk_sessions_states FOREIGN KEY (desired_state) REFERENCES bbq.subject_states(id);


--
-- TOC entry 3498 (class 2606 OID 16732)
-- Name: subject_states fk_states_subjects; Type: FK CONSTRAINT; Schema: bbq; Owner: scott
--

ALTER TABLE ONLY bbq.subject_states
    ADD CONSTRAINT fk_states_subjects FOREIGN KEY (subject_id) REFERENCES bbq.subjects(id);


-- Completed on 2023-06-27 21:04:58 EDT

--
-- PostgreSQL database dump complete
--