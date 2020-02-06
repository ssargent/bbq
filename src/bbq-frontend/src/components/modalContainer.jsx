//@flow
import React, { useState } from "react";
import { Modal, ModalHeader, ModalBody, ModalFooter } from "reactstrap";

export default function ModalContainer(props) {
  const [showModal, setShowModal] = useState(false);
  const toggle = () => {
    setShowModal(!showModal);
  };

  return (
    <React.Fragment>
      <button className={props.buttonClassName} onClick={toggle}>
        {props.buttonText}
      </button>
      <Modal>
        <ModalHeader>{props.title}</ModalHeader>
        <ModalBody>{props.children}</ModalBody>
      </Modal>
    </React.Fragment>
  );
}
