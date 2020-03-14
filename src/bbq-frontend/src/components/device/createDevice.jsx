//@flow
import React, { useState } from "react";
import { Modal, ModalHeader, ModalBody, ModalFooter, Button } from "reactstrap";

export default function CreateDevice({
  buttonClassName,
  buttonText,
  modalClassName,
  saveDevice
}) {
  const [showModal, setShowModal] = useState(false);
  const [deviceName, setDeviceName] = useState("");
  const [deviceDesc, setDeviceDesc] = useState("");

  const toggle = () => setShowModal(!showModal);
  return (
    <React.Fragment>
      <button className={buttonClassName} onClick={toggle}>
        {buttonText}
      </button>
      <Modal isOpen={showModal} toggle={toggle} className={modalClassName}>
        <ModalHeader toggle={toggle}>Create Device</ModalHeader>
        <ModalBody>
          <form>
            <div className="form-group">
              <label htmlFor="device-name">Name</label>
              <input
                type="text"
                className="form-control"
                id="device-name"
                value={deviceName}
                onChange={e => setDeviceName(e.target.value)}
                placeholder="Your Awesome BBQ"
              ></input>
            </div>
            <div className="form-group">
              <label htmlFor="device-description">Description</label>
              <input
                type="text"
                className="form-control"
                id="device-description"
                value={deviceDesc}
                onChange={e => setDeviceDesc(e.target.value)}
                placeholder="Your Awesome BBQ"
              ></input>
            </div>
          </form>
        </ModalBody>
        <ModalFooter>
          <Button
            color="primary"
            onClick={() => {
              saveDevice({ name: deviceName, description: deviceDesc });
              toggle();
            }}
          >
            Save
          </Button>
          <Button
            color="secondary"
            onClick={() => {
              setDeviceDesc("");
              setDeviceName("");
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
