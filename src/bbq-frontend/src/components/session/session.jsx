//@flow

import React, { useState, useEffect } from "react";
import moment from "moment";
import { useParams } from "react-router-dom";
import { transport } from "../../transport";
import { API_SERVER } from "../../config";
import ThermalSensorReadings from "../charts/thermalSensorReadings";

export default function Sessions() {
  const [data, setData] = useState([]);
  const { sessionid } = useParams();
  useEffect(() => {
    const fetchData = async () => {
      const result = await transport.get(
        `${API_SERVER}/v1/bbq/sessions/${sessionid}`
      );
      setData(result.data);
    };

    fetchData();
  }, [sessionid]);

  let duration = {};

  if (data.endtime && data.endtime.Valid) {
    const endTime = data.endtime ? moment(data.endtime.Time) : moment();
    const startTime = moment(data.starttime);
    duration = moment.duration(endTime.diff(startTime));
  } else {
    const endTime = moment();
    const startTime = moment(data.starttime);
    duration = moment.duration(endTime.diff(startTime));
  }

  return (
    <div style={{ marginTop: "50px" }}>
      <h1>{data.name}</h1>
      {(!data.endtime || data.endtime.Valid == false) && (
        <div className="alert alert-info">
          This cook has no endtime and is likely still in progress.
        </div>
      )}
      <table className="table">
        <tr>
          <th>Description</th>
          <td>{data.description}</td>
        </tr>
        <tr>
          <th>Time</th>
          <td>
            {`${duration.get("hours")} hours and ${duration.get(
              "minutes"
            )} minutes`}
          </td>
        </tr>
        <tr>
          <th>Date</th>
          <td>{moment(data.starttime).format("MMMM Do YYYY")}</td>
        </tr>
        <tr>
          <th>Device</th>
          <td>{data.device}</td>
        </tr>
        <tr>
          <th>Monitor</th>
          <td>{data.monitor}</td>
        </tr>
      </table>
      <ThermalSensorReadings sessionid={sessionid} />
    </div>
  );
}
