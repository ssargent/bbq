import React, { useState, useEffect } from "react";
import moment from "moment";
import { useParams } from "react-router-dom";
import { BrowserRouter as Router, Link, Route } from "react-router-dom";
import { transport } from "../transport";
import { API_SERVER } from "../config";
import ThermalSensorReadings from "./charts/thermalSensorReadings";

export default function Sessions() {
    const [data, setData] = useState([]);
    const { sessionid } = useParams();
    useEffect(() => {

        const fetchData = async () => {
            const result = await transport.get(`${API_SERVER}v1/bbq/sessions/${sessionid}`);
            setData(result.data);
        };

        fetchData();
    },[sessionid]);

    console.log(data);
/*
{
    "id": 1,
    "name": "Pulled Pork",
    "description": "Pulled Pork",
    "subject": "",
    "type": "",
    "weight": 5,
    "device": "LBGE",
    "monitor": "InkBird 4 Probe",
    "starttime": "2019-04-06T10:25:59.977852Z",
    "endtime": {
        "Time": "2019-06-06T11:32:28.811119Z",
        "Valid": true
    },
    "tenantid": "c28f15a7-24f3-5ead-8403-b4d08312801e",
    "uid": "4a78179d-5cf6-432d-b5a8-d87de7b136e5"
}
*/
    const endTime =  data.endtime ? moment(data.endtime.Time) : moment();
    const startTime = moment(data.starttime);

    var duration = moment.duration(endTime.diff(startTime))

    return (
      <div style={{ marginTop: "50px"}}>
        <table className="table">
            <tr>
                <th>Name</th>
                <td>{data.name}</td>
            </tr>
            <tr>
                <th>Description</th>
                <td>{data.description}</td>
            </tr>
            <tr>
                <th>Time</th>
                <td>
                    {`${duration.get('hours')} hours and ${duration.get('minutes')} minutes`}
                    </td>
            </tr>   
            <tr>
                <th>Date</th>
                <td>{moment(data.starttime).format('MMMM Do YYYY')}</td>
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
};