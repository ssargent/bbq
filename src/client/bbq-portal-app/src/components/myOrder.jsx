import numeral from "numeral";
import React from "react";
import { Table } from "reactstrap";

export default function MyOrder(props: { order: object }) {
  var order = props.order || { lines: [] };
  return (
    <div>
      <h2>My Order</h2>
      <Table striped size="sm" responsive>
        <thead>
          <tr>
            <th>Quantity</th>
            <th>Item</th>
            <th>Price</th>
          </tr>
        </thead>
        <tfoot>
          <tr>
            <td colSpan="3" style={{ textAlign: "right" }}>
              {numeral(
                order.lines.reduce((accum, line) => accum + line.lineTotal, 0)
              ).format("$0,0.00")}
            </td>
          </tr>
        </tfoot>
        <tbody>
          {order.lines.map(dl => (
            <tr key={dl.id}>
              <td>{dl.quantity}</td>
              <td>{dl.productName}</td>
              <td style={{ textAlign: "right" }}>
                {numeral(dl.lineTotal).format("$0,0.00")}
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
}
