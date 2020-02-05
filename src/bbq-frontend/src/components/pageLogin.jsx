//@flow

import React, { useState } from "react";
import axios from "axios";
import { useHistory, useLocation } from "react-router-dom";
import { Button } from "reactstrap";
import { API_SERVER } from "../config";

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

    if (!process.env.NODE_ENV || process.env.NODE_ENV === "development") {
      url = `${API_SERVER}v1/system/accounts/login`;
    } else {
      url = `${API_SERVER}v1/system/accounts/signin`;
    }

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
    <div className="row">
      <div className="col-md-4 offset-md-4" style={{ paddingTop: "300px" }}>
        <form className="form-login">
          <div className="form-group">
            <input
              type="text"
              className="form-control"
              placeholder="Your Email"
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

        <div className="pull-right">
          <Button color="primary" onClick={onLogin}>
            Login
          </Button>
        </div>
      </div>
    </div>
  );
}
