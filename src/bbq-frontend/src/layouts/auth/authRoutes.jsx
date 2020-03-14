import React from "react";
import { Route } from "react-router-dom";

import AuthComponent from "./authComponent";

const AuthRoutes = ({ component: Component, ...rest }) => {
  return (
    <Route
      {...rest}
      render={props => (
        <AuthComponent>
          <Component {...props} />
        </AuthComponent>
      )}
    />
  );
};

export default AuthRoutes;
