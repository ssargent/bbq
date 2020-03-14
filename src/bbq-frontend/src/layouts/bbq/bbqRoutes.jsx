import React from "react";
import { Route } from "react-router-dom";

import BbqLayout from "./bbqLayout";

const BbqRoutes = ({ component: Component, ...rest }) => {
  return (
    <Route
      {...rest}
      render={props => (
        <BbqLayout>
          <Component {...props} />
        </BbqLayout>
      )}
    />
  );
};

export default BbqRoutes;
