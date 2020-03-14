//@flow

import React, { useState } from "react";
import axios from "axios";
import { useHistory, useLocation } from "react-router-dom";
import { Button } from "reactstrap";
import { API_SERVER } from "../config";
import "./pageLogin.css";

export default function PageLogin({}) {
  let history = useHistory();
  let location = useLocation();
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");

  let { from } = location.state || { from: { pathname: "/" } };
  let onLogin = () => {
    const loginModel = {
      loginname: login,
      password: password
    };

    let url = "";

    url = `${API_SERVER}v1/system/accounts/signin`;

    axios
      .post(url, loginModel)
      .then(resp => {
        if (resp.data.success === true) {
          localStorage.setItem("bbq-authenticated", JSON.stringify(resp.data));
          history.replace(from);
        }
      })
      .catch(err => {
        console.log(err);
      });

    /*  fakeAuth.authenticate(() => {
      history.replace(from);
    });*/
  };

  return (
    <div className="container h-100" style={{ marginTop: "300px" }}>
      <div className="row h-100 justify-content-center align-items-center">
        <div className="card text-white bg-secondary  mb-3  ">
          <div className="card-header">Please Login</div>
          <div class="card-body">
            <div>
              <h4>The World's Most Over Engineered BBQ Thermometer</h4>
            </div>
            <form className="form-login" style={{ paddingTop: "15px" }}>
              <div className="form-group">
                <input
                  type="text"
                  className="form-control"
                  placeholder="Your Login"
                  onChange={e => setLogin(e.target.value)}
                  value={login}
                />
              </div>
              <div className="form-group">
                <input
                  type="password"
                  className="form-control"
                  placeholder="Your Password"
                  onChange={e => setPassword(e.target.value)}
                  value={password}
                />
              </div>
            </form>

            <div className="text-right">
              <Button color="primary" onClick={onLogin}>
                Login
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
