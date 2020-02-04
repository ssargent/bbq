//@flow

import React from "react";
import { transport } from "../../transport";
import { API_SERVER } from "../../config";
import CreateDevice from "./createDevice";

//type State = {
//    devices: Array<Object>
//}
class Devices extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      devices: []
    };
  }

  getDevices = () => {
    transport
      .get(`${API_SERVER}v1/bbq/devices`)
      .then(resp => this.setState({ devices: resp.data }))
      .catch(err => console.log(err));
  };

  createDevice = d => {
    console.log(d);
    transport
      .post(`${API_SERVER}v1/bbq/devices`, d)
      .then(r => this.getDevices())
      .catch(err => this.setState({ err }));
  };

  componentDidMount() {
    this.getDevices();
  }

  render() {
    return (
      <React.Fragment>
        <div className="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3">
          <h1 className="h2">Devices</h1>
        </div>

        <div className="bbq-button-strip">
          <CreateDevice
            buttonClassName="btn btn-outline-success btn-sm margin-bottom-10"
            buttonText="Create New Device"
            saveDevice={d => {
              this.createDevice(d);
            }}
          />
        </div>
        <table className="table table-hover table-striped">
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            {this.state.devices.map(d => (
              <tr key={d.id}>
                <td>{d.id}</td>
                <td>{d.name}</td>
                <td>{d.description}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </React.Fragment>
    );
  }
}

export default Devices;
