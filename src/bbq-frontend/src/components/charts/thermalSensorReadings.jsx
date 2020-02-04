import React, { useState, useEffect } from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend
} from "recharts";

import { transport } from "../../transport";
import { API_SERVER } from "../../config";

export default function ThermalSensorReadings({ sessionid }) {
  const [data, setData] = useState([]);
  // const { sessionid } = useParams();
  useEffect(() => {
    console.log(sessionid);

    const fetchData = async () => {
      const result = await transport.get(
        `${API_SERVER}v1/data/sensors/${sessionid}/raw`
      );
      setData(result.data);
    };

    fetchData();
  }, [sessionid]);

  console.log(data);
  return (
    <div>
      <LineChart
        width={1200}
        height={600}
        data={data}
        margin={{
          top: 5,
          right: 30,
          left: 20,
          bottom: 5
        }}
      >
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="name" />
        <YAxis />
        <Tooltip />
        <Legend />
        <Line type="monotone" dot={false} dataKey="probe0" stroke="#0000FF" />
        <Line type="monotone" dot={false} dataKey="probe1" stroke="#008000" />
        <Line type="monotone" dot={false} dataKey="probe2" stroke="#FF0000" />
        <Line type="monotone" dot={false} dataKey="probe3" stroke="#800000" />
      </LineChart>
    </div>
  );
}
