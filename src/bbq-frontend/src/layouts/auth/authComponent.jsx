import React from "react";

const AuthComponent = props => {
  return (
    <div id="wrapper">
      <div id="content">{props.children}</div>
    </div>
  );
};

export default AuthComponent;
