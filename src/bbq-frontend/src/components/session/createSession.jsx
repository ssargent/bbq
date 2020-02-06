//@flow
import React, { useEffect, useState } from "react";
import { Modal, ModalHeader, ModalBody, ModalFooter, Button } from "reactstrap";

export default function CreateSession(props) {
  const [showModal, setShowModal] = useState(false);
  const toggle = () => {
    setShowModal(!showModal);
  };
  /* 
       {
   	    "name": "Pulled Pork",
        "description": "Pulled Pork",
        "subject": "Not Specified",
        "type": "Uncategorized",
        "weight": 10,
        "device": "LBGE",
        "monitor": "InkBird 4 Probe",
    }
    */
  return (
    <div className="bbq-thermometer-cooking-session">
      <React.Fragment>
        <button className={props.buttonClassName} onClick={toggle}>
          {props.buttonText}
        </button>
        <Modal isOpen={showModal} toggle={toggle}>
          <ModalHeader>{props.title}</ModalHeader>
          <ModalBody>
            <form>
              <div className="form-group">
                <label htmlFor="inputEmail3" className="col-form-label">
                  Name
                </label>

                <input type="text" className="form-control" id="inputName" />
              </div>
              <div className="form-group">
                <label htmlFor="inputPassword3" className="col-form-label">
                  Description
                </label>

                <input
                  type="text"
                  className="form-control"
                  id="inputDescription"
                />
              </div>
              <div className="form-group">
                <label htmlFor="inputPassword3" className=" col-form-label">
                  Subject
                </label>

                <input type="text" className="form-control" id="inputSubject" />
              </div>
              <div className="form-group">
                <label htmlFor="inputPassword3" className=" col-form-label">
                  Weight
                </label>

                <input type="text" className="form-control" id="inputWeight" />
              </div>
              <div className="form-group">
                <label htmlFor="inputPassword3" className="col-form-label">
                  Device
                </label>

                <input type="text" className="form-control" id="inputDevice" />
              </div>
              <div className="form-group">
                <label htmlFor="inputPassword3" className=" col-form-label">
                  Monitor
                </label>

                <input type="text" className="form-control" id="inputMonitor" />
              </div>
            </form>
          </ModalBody>
          <ModalFooter>
            <Button
              color="primary"
              onClick={() => {
                toggle();
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
