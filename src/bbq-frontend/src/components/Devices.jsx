//@flow

import React from "react";
import axios from "axios";

//type State = {
//    devices: Array<Object>
//}
class Devices extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            devices: []
        }
    }

    componentDidMount() {
        axios.get("https://bbq.k8s.ssargent.net/v1/bbq/devices",
        { withCredentials: true })
            .then(resp => this.setState({ devices: resp.data }))
            .catch(err => console.log(err));
    }
    render() {
        return (  <React.Fragment>   <div className="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
         
     
        <h1 className="h2">Devices</h1>
        <div className="btn-toolbar mb-2 mb-md-0">
          <div className="btn-group mr-2">
            <button type="button" className="btn btn-sm btn-outline-secondary">Share</button>
            <button type="button" className="btn btn-sm btn-outline-secondary">Export</button>
          </div>
          <button type="button" className="btn btn-sm btn-outline-secondary dropdown-toggle">
            <span data-feather="calendar"></span>
            This week
          </button>
        </div></div>
        <table className="table table-hover table-striped">
            <thead>
                <tr>
                    <th>ID</th><th>Name</th><th>Description</th>
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
        </table></React.Fragment> 
        )
    }
}

export default Devices;