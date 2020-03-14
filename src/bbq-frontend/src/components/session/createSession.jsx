//@flow
import React, { useEffect, useState } from "react";
import { Modal, ModalHeader, ModalBody, ModalFooter, Button } from "reactstrap";
import { transport } from "../../transport";

export default function CreateSession(props) {
  const [showModal, setShowModal] = useState(false);
  const [devices, setDevices] = useState([]);
  const [error, setError] = useState(undefined);
  const [subjects, setSubjects] = useState([]);
  const [monitors, setMonitors] = useState([]);

  const [sessionName, setSessionName] = useState("");
  const [sessionDescription, setSessionDescription] = useState("");
  const [sessionSubject, setSessionSubject] = useState("");
  const [sessionWeight, setSessionWeight] = useState(0);
  const [sessionDevice, setSessionDevice] = useState("");
  const [sessionMonitor, setSessionMonitor] = useState("");

  const toggle = () => {
    setShowModal(!showModal);
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const devResp = await transport.get(`v1/bbq/devices`);
        const subResp = await transport.get(`v1/bbq/subjects`);
        const monResp = await transport.get(`v1/bbq/monitors`);

        setDevices(devResp.data);
        setSubjects(subResp.data);
        setMonitors(monResp.data);
      } catch (apiError) {
        setError(apiError);
      }
    };

    fetchData();
  }, []);

  return (
    <div className="bbq-thermometer-cooking-session">
      <React.Fragment>
        <button className={props.buttonClassName} onClick={toggle}>
          {props.buttonText}
        </button>
        <Modal isOpen={showModal} toggle={toggle}>
          <ModalHeader>{props.title}</ModalHeader>
          <ModalBody>
            {error != undefined && (
              <div className="alert alert-danger">
                {error.request.responseURL} - {error.response.status} -{" "}
                {error.response.statusText}
              </div>
            )}
            <form>
              <div className="form-group">
                <label htmlFor="inputName" className="col-form-label">
                  Name
                </label>

                <input
                  type="text"
                  className="form-control"
                  id="inputName"
                  value={sessionName}
                  onChange={e => setSessionName(e.target.value)}
                />
              </div>
              <div className="form-group">
                <label htmlFor="inputDscription" className="col-form-label">
                  Description
                </label>

                <input
                  type="text"
                  className="form-control"
                  id="inputDescription"
                  value={sessionDescription}
                  onChange={e => setSessionDescription(e.target.value)}
                />
              </div>
              <div className="form-group">
                <label htmlFor="inputPassword3" className=" col-form-label">
                  Subject
                </label>
                <select
                  className="form-control"
                  id="inputSubject"
                  value={sessionSubject}
                  onChange={e => setSessionSubject(e.target.value)}
                >
                  <option key="">Select a Subject</option>
                  {subjects.map(s => (
                    <option key={s.uid}>{s.name}</option>
                  ))}
                </select>
              </div>
              <div className="form-group">
                <label htmlFor="inputWeight" className=" col-form-label">
                  Weight
                </label>

                <input
                  type="text"
                  className="form-control"
                  id="inputWeight"
                  value={sessionWeight}
                  onChange={e => setSessionWeight(e.target.value)}
                />
              </div>
              <div className="form-group">
                <label htmlFor="inputDevice" className="col-form-label">
                  Device
                </label>
                <select
                  className="form-control"
                  id="inputDevice"
                  value={sessionDevice}
                  onChange={e => setSessionDevice(e.target.value)}
                >
                  <option key="">Select a Device</option>
                  {devices.map(d => (
                    <option key={d.uid} value={d.name}>
                      {d.description}
                    </option>
                  ))}
                </select>
              </div>
              <div className="form-group">
                <label htmlFor="inputMonitor" className=" col-form-label">
                  Monitor
                </label>

                <select
                  className="form-control"
                  id="inputMonitor"
                  value={sessionMonitor}
                  onChange={e => setSessionMonitor(e.target.value)}
                >
                  <option key="">Select a Monitor</option>
                  {monitors.map(m => (
                    <option key={m.uid} value={m.name}>
                      {m.description}
                    </option>
                  ))}
                </select>
              </div>
            </form>
          </ModalBody>
          <ModalFooter>
            <Button
              color="primary"
              onClick={() => {
                //toggle();

                const session = {
                  name: sessionName,
                  description: sessionDescription,
                  subject: sessionSubject,
                  type: "Uncategorized",
                  weight: parseFloat(sessionWeight),
                  device: sessionDevice,
                  monitor: sessionMonitor
                };

                console.log(session);
                if (props.saveSession(session)) toggle();
                //    setMonitorName("");
                //    setMonitorDesc("");
                //    setMonitorAddress("");
              }}
            >
              Save
            </Button>
            <Button
              color="secondary"
              onClick={() => {
                //    setMonitorName("");
                //    setMonitorDesc("");
                //    setMonitorAddress("");
                toggle();
              }}
            >
              Cancel
            </Button>
          </ModalFooter>
        </Modal>
      </React.Fragment>
    </div>
  );
}
