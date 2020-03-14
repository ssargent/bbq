//@flow
import React, { useEffect, useState } from "react";
import { Modal, ModalHeader, ModalBody, ModalFooter, Button } from "reactstrap";
import { Dropdown, Menu, Icon } from "antd";
import { transport } from "../../transport";
import { API_SERVER } from "../../config";

export default function AdvancedCreateSession(props) {
  const [showModal, setShowModal] = useState(false);
  const [devices, setDevices] = useState([]);
  const [error, setError] = useState({});
  const [subjects, setSubjects] = useState([]);

  const toggle = () => {
    setShowModal(!showModal);
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const result = await transport.get(`${API_SERVER}v1/bbq/devices`);
        setDevices(result.data);

        const subjectsResult = await transport.get(
          `${API_SERVER}v1/bbq/subjects`
        );
        setSubjects(subjectsResult.data);
      } catch (apiError) {
        setError(apiError);
      }
    };

    fetchData();
  }, []);

  const deviceMenu = () => {
    return (
      <Menu>
        {devices.map(d => (
          <Menu.Item key={d.id}>{d.name}</Menu.Item>
        ))}
      </Menu>
    );
  };

  const menu = deviceMenu();
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
              <div>
                I am making Pulled Pork on the
                <Dropdown overlay={menu}>
                  <a className="ant-dropdown-link" href="#">
                    Hover me <Icon type="down" />
                  </a>
                </Dropdown>{" "}
                using the Inkbird 4 Monitor
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
