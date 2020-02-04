//@flow

import React from "react";
import { transport } from "../../transport";
import { API_SERVER } from "../../config";
import CreateMonitor from "./createMonitor";
//type State = {
//    devices: Array<Object>
//}
class Monitors extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      monitors: []
    };
  }

  createMonitor = m => {
    console.log(m);
    transport
      .post(`${API_SERVER}v1/bbq/monitors`, m)
      .then(r => this.getMonitors())
      .catch(err => this.setState({ err }));
  };

  getMonitors = () =>
    transport
      .get(`${API_SERVER}v1/bbq/monitors`)
      .then(resp => this.setState({ monitors: resp.data }))
      .catch(err => console.log(err));

  componentDidMount() {
    this.getMonitors();
  }

  render() {
    return (
      <React.Fragment>
        {" "}
        <div className="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3">
          <h1 className="h2">Monitors</h1>
        </div>
        <div className="bbq-button-strip">
          <CreateMonitor
            buttonClassName="btn btn-outline-success btn-sm margin-bottom-10"
            buttonText="Create New Monitor"
            saveMonitor={m => {
              this.createMonitor(m);
            }}
          />
        </div>
        <table className="table table-hover table-striped">
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Address</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            {this.state.monitors.map(d => (
              <tr key={d.id}>
                <td>{d.id}</td>
                <td>{d.name}</td>
                <td>{d.address}</td>
                <td>{d.description}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </React.Fragment>
    );
  }
}

export default Monitors;
