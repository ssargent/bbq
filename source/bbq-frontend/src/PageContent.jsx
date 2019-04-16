import React from "react";
import { BrowserRouter as Router, Link, Route } from "react-router-dom";
import Dashboard from "./components/dashboard";
import Devices from "./components/Devices";
import Monitors from "./components/Monitors";

export default function PageContent({}) {
    return (
    <Router> 

    <div className="container-fluid">
    <div className="row">
      <nav className="col-md-2 d-none d-md-block bg-light sidebar">
        <div className="sidebar-sticky">
          <ul className="nav flex-column">
                  <li className="nav-item"> 
              <a className="nav-link active" href="#">
                <span data-feather="home"></span>
                Dashboard <span className="sr-only">(current)</span>
              </a>
            </li>
            <li className="nav-item">
              <Link className="nav-link" to="/devices">Devices</Link> 
            </li>
            <li className="nav-item">
            <Link className="nav-link" to="/monitors">Monitors</Link> 
            </li>
            <li className="nav-item">
              <a className="nav-link" href="#">
                <span data-feather="users"></span>
                Sessions
              </a>
            </li>

          </ul>
        </div>
      </nav>
  
      <main role="main" className="col-md-9 ml-sm-auto col-lg-10 px-4">
       <Route 
        path="/"
        exact={true}
        component={Dashboard}
       />
              <Route 
        path="/devices" 
        component={Devices}
       />
                <Route 
        path="/monitors" 
        component={Monitors}
       />
  
        
              </main>
    </div>
  </div>
  </Router>
  )
}