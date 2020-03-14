/*import React, { Component } from "react";
import { Route, Redirect } from "react-router-dom";

// A wrapper for <Route> that redirects to the login
// screen if you're not yet authenticated.
function PrivateRoute({ comp, ...rest }) {
  // enhance this with api call of some kind.
  const isAuthenticated = () => {
    console.log(
      "is authentcated",
      localStorage.getItem("bbq-authenticated") != null
    );
    return localStorage.getItem("bbq-authenticated") != null;
  };

  return (
    <Route
      render={props =>
        !isAuthenticated() ? <Redirect to="/login" /> : <Component {...props} />
      }
    />
  );
}

export default PrivateRoute;
*/

import React from "react";
import { Route, Redirect } from "react-router-dom";
//import { isLogin } from "../utils";

const PrivateRoute = ({ comp: Component, ...rest }) => {
  const isAuthenticated = () => {
    console.log(
      "is authentcated",
      localStorage.getItem("bbq-authenticated") != null
    );
    return localStorage.getItem("bbq-authenticated") != null;
  };
  return (
    // Show the component only when the user is logged in
    // Otherwise, redirect the user to /signin page
    <Route
      {...rest}
      render={props =>
        isAuthenticated() ? <Component {...props} /> : <Redirect to="/login" />
      }
    />
  );
};

export default PrivateRoute;
