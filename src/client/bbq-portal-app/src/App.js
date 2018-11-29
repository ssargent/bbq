import numeral from "numeral";
import React, { Component } from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import "./App.css";
import Menu from "./components/menu";
import MyOrder from "./components/myOrder";
import StoreNav from "./components/storeNav";
import * as orderActions from "./_actions/orderActions";

class App extends Component {
  componentDidMount() {
    this.props.orderActions.loadCart();
  }
  render() {
    const order = this.props.orders.order || { lines: [] };
    const orderTotal = order.lines.reduce(
      (accum, line) => accum + line.lineTotal,
      0
    );
    return (
      <div>
        <StoreNav />
        <div className="App row">
          <div className="col-md-9">
            <Menu />
            {orderTotal > 0 && (
              <button className="btn btn-success" style={{ marginTop: "50px" }}>
                I'm Hungry! Order {numeral(orderTotal).format("$0,0.00")} worth
                of yum!
              </button>
            )}
          </div>
          <div className="col-md-3">
            <MyOrder className="col-md-3" order={this.props.orders.order} />
          </div>
        </div>
      </div>
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
