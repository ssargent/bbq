import numeral from "numeral";
import React from "react";
import { Link } from "react-router-dom";
import Menu from "./menu";
import MyOrder from "./myOrder";

export default function StoreFront(props: { order: object }) {
  const order = props.order || { lines: [] };
  const orderTotal = order.lines.reduce(
    (accum, line) => accum + line.lineTotal,
    0
  );
  return (
    <div className="App row">
      <div className="col-md-9">
        <Menu />
        {orderTotal > 0 && (
          <Link
            className="btn btn-success"
            style={{ marginTop: "50px" }}
            to="/order-now"
          >
            I'm Hungry! Order {numeral(orderTotal).format("$0,0.00")} worth of
            yum!
          </Link>
        )}
      </div>
      <div className="col-md-3">
        <MyOrder className="col-md-3" order={props.order} />
      </div>
    </div>
  );
}