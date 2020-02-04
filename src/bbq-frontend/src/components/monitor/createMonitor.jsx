import React, { useState, useEffect } from "react";
import { Modal, ModalHeader, ModalBody, ModalFooter, Button } from "reactstrap";

export default function CreateMonitor({
  buttonClassName,
  buttonText,
  modalClassName,
  saveMonitor
}) {
  const [monitorName, setMonitorName] = useState("");
  const [monitorDesc, setMonitorDesc] = useState("");
  const [monitorAddress, setMonitorAddress] = useState("");

  const [showModal, setShowModal] = useState(false);
  const toggle = () => setShowModal(!showModal);

  return (
    <React.Fragment>
      <button className={buttonClassName} onClick={toggle}>
        {buttonText}
      </button>
      <Modal isOpen={showModal} toggle={toggle} className={modalClassName}>
        <ModalHeader toggle={toggle}>Create Monitor</ModalHeader>
        <ModalBody>
          <form>
            <div className="form-group">
              <label htmlFor="device-name">Name</label>
              <input
                type="text"
                className="form-control"
                id="monitor-name"
                value={monitorName}
                onChange={e => setMonitorName(e.target.value)}
                placeholder="Your Awesome Monitor"
              ></input>
            </div>
            <div className="form-group">
              <label htmlFor="monitor-description">Description</label>
              <input
                type="text"
                className="form-control"
                id="monitor-description"
                value={monitorDesc}
                onChange={e => setMonitorDesc(e.target.value)}
                placeholder="Your Awesome Monitor"
              ></input>
            </div>
            <div className="form-group">
              <label htmlFor="monitor-address">Address</label>
              <input
                type="text"
                className="form-control"
                id="monitor-address"
                value={monitorAddress}
                onChange={e => setMonitorAddress(e.target.value)}
                placeholder="Your Awesome Monitor"
              ></input>
            </div>
          </form>
        </ModalBody>
        <ModalFooter>
          <Button
            color="primary"
            onClick={() => {
              saveMonitor({
                name: monitorName,
                description: monitorDesc,
                address: monitorAddress
              });
              toggle();
              setMonitorName("");
              setMonitorDesc("");
              setMonitorAddress("");
            }}
          >
            Save
          </Button>
          <Button
            color="secondary"
            onClick={() => {
              setMonitorName("");
              setMonitorDesc("");
              setMonitorAddress("");
              toggle();
            }}
          >
            Cancel
          </Button>
        </ModalFooter>
      </Modal>
    </React.Fragment>
  );
}
