import React, { useState, useEffect } from "react";
import moment from "moment";

import { BrowserRouter as Router, Link, Route } from "react-router-dom";
import { transport } from "../transport";
import { API_SERVER } from "../config";

export default function Sessions() {
    const [data, setData] = useState([]);
    useEffect(() => {

        const fetchData = async () => {
            const result = await transport.get(`${API_SERVER}v1/bbq/sessions`);
            setData(result.data);
        };

        fetchData();
    },[]);

    console.log(data);
    return (
      <div>
        <table className="table">
          <thead className="thread-dark">
              <tr>
                  <th scope="col">When</th>
                  <th scope="col">Name</th>
                  <th scope="col">Description</th>
                  <th scope="col">Device</th>
                  <th scope="col">Monitor</th>
                  <th></th>
              </tr>
          </thead>
          <tbody>
          {data && data.map(s => (
              <tr key={s.uid}>
                  <td>{moment(s.starttime).fromNow()}</td>
                  <td>{s.name}</td>
                  <td>{s.description}</td>
                  <td>{s.device}</td>
                  <td>{s.monitor}</td>
                  <td><Link className="btn btn-link btn-xs" to={`/cookingsession/${s.uid}`}>Chart</Link> </td>
              </tr> 
          ))}
          </tbody>
        </table>
      </div>
    );
};