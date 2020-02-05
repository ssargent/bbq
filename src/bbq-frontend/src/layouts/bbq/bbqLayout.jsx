//@flow
//import { findIconDefinition, library } from '@fortawesome/fontawesome-svg-core';
//import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from "react";
import { Switch } from "react-router-dom";
import PrivateRoute from "../privateRoute";
import Dashboard from "../../components/dashboard";
import Devices from "../../components/device/devices";
import Monitors from "../../components/monitor/monitors";
import Sessions from "../../components/session/sessions";
import Session from "../../components/session/session";
import LeftNav from "../../leftNav";
import TopNav from "../../components/TopNav";

export default function BbqLayout() {
  //console.log(match);
  return (
    <React.Fragment>
      <TopNav />
      <div className="container-fluid">
        <div className="row">
          <LeftNav />

          <main role="main" className="col-md-9 ml-sm-auto col-lg-10 px-4">
            <Switch>
              <PrivateRoute path="/devices" comp={Devices} />

              <PrivateRoute path="/" exact={true} comp={Dashboard} />
              <PrivateRoute path="/monitors" exact={true} comp={Monitors} />
              <PrivateRoute path="/sessions" exact={true} comp={Sessions} />
              <PrivateRoute
                path="/cookingsession/:sessionid"
                exact={true}
                comp={Session}
              />
            </Switch>
          </main>
        </div>
      </div>
    </React.Fragment>
  );
}
