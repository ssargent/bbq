import React, { Component } from "react";
import { connect } from "react-redux";
import { BrowserRouter as Router, Route } from "react-router-dom";
import { bindActionCreators } from "redux";
import "./App.css";
import OrderNow from "./components/orderNow";
import StoreFront from "./components/storeFront";
import * as orderActions from "./_actions/orderActions";

class App extends Component {
  componentDidMount() {
    this.props.orderActions.loadCart();
  }
  render() {
    const order = this.props.orders.order || { lines: [] };

    return (
      <Router>
        <div>
          <Route path="/order-now" component={OrderNow} />
          <Route
            path="/"
            exact
            render={props => <StoreFront order={order} />}
          />
        </div>
      </Router>
    );
  }
}

App = connect(
  (state, ownProps) => {
    return {
      orders: state.orders
    };
  },
  (dispatch: Function) => {
    return {
      orderActions: bindActionCreators(orderActions, dispatch)
    };
  }
)(App);

export default App;
