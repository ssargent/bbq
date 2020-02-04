//@flow
//import { findIconDefinition, library } from '@fortawesome/fontawesome-svg-core';
//import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from "react";
import { BrowserRouter as Router, Route } from "react-router-dom";
import Dashboard from "./components/dashboard";
import Devices from "./components/device/devices";
import Monitors from "./components/monitor/monitors";
import Sessions from "./components/session/sessions";
import Session from "./components/session/session";
import LeftNav from "./leftNav";

export default function PageContent() {
  return (
    <Router>
      <div className="container-fluid">
        <div className="row">
          <LeftNav />

          <main role="main" className="col-md-9 ml-sm-auto col-lg-10 px-4">
            <Route path="/" exact={true} component={Dashboard} />
            <Route path="/devices" component={Devices} />
            <Route path="/monitors" component={Monitors} />
            <Route path="/sessions" component={Sessions} />
            <Route path="/cookingsession/:sessionid" component={Session} />
          </main>
        </div>
      </div>
    </Router>
  );
}
