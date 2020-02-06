import React, { Component } from "react";
import { BrowserRouter as Router, Switch } from "react-router-dom";
import AuthRoutes from "./layouts/auth/authRoutes";
import BbqRoutes from "./layouts/bbq/bbqRoutes";
import Dashboard from "./components/dashboard";
import PageLogin from "./components/pageLogin";
import "./App.css";

class DebugRouter extends Router {
  constructor(props) {
    super(props);
    console.log("initial history is: ", JSON.stringify(this.history, null, 2));
    this.history.listen((location, action) => {
      console.log(
        `The current URL is ${location.pathname}${location.search}${location.hash}`
      );
      console.log(
        `The last navigation action was ${action}`,
        JSON.stringify(this.history, null, 2)
      );
    });
  }
}

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      devices: []
    };
  }

  render() {
    return (
      <Router>
        <Switch>
          <AuthRoutes exact path="/login" component={PageLogin} />
          <BbqRoutes path="/" component={Dashboard} />
        </Switch>
      </Router>
    );
  }
}

export default App;
