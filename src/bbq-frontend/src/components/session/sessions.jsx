//@flow
import React, { useState, useEffect } from "react";
import moment from "moment";

import { Link } from "react-router-dom";
import { transport } from "../../transport";
import { API_SERVER } from "../../config";
import CreateSession from "./createSession";

export default function Sessions() {
  const [data, setData] = useState([]);

  const fetchData = async () => {
    const result = await transport.get(`${API_SERVER}v1/bbq/sessions`);
    setData(result.data);
  };
  useEffect(() => {
    fetchData();
  }, []);

  const createSession = async s => {
    const resp = await transport.post("v1/bbq/sessions", s);
    await fetchData();
    return true;
  };

  return (
    <div>
      <div className="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3">
        <h1 className="h2">Cooking Sessions</h1>
      </div>
      <div className="bbq-button-strip">
        <CreateSession
          buttonClassName="btn btn-outline-success btn-sm margin-bottom-10"
          buttonText="Start Cooking!"
          title="Let's Cook Something"
          saveSession={s => {
            return createSession(s);
          }}
        />
      </div>
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
          {data &&
            data.map(s => (
              <tr key={s.uid}>
                <td>{moment(s.starttime).fromNow()}</td>
                <td>{s.name}</td>
                <td>{s.description}</td>
                <td>{s.device}</td>
                <td>{s.monitor}</td>
                <td>
                  <Link
                    className="btn btn-link btn-xs"
                    to={`/cookingsession/${s.uid}`}
                  >
                    Chart
                  </Link>
                </td>
              </tr>
            ))}
        </tbody>
      </table>
    </div>
  );
}
