//@flow

import React from "react";

import { Link } from "react-router-dom";

export default function LeftNav() {
  return (
    <nav className="col-md-2 d-none d-md-block bg-light sidebar">
      <div className="sidebar-sticky">
        <ul className="nav flex-column">
          <li className="nav-item">
            <Link className="nav-link" to="/">
              Dashboard
            </Link>
          </li>
          <li className="nav-item">
            <Link className="nav-link" to="/devices">
              Devices
            </Link>
          </li>
          <li className="nav-item">
            <Link className="nav-link" to="/monitors">
              Monitors
            </Link>
          </li>
          <li className="nav-item">
            <Link className="nav-link" to="/sessions">
              Sessions
            </Link>
          </li>
        </ul>
      </div>
    </nav>
  );
}
