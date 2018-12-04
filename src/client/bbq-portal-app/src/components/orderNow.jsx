import React from "react";
import { connect } from "react-redux";
import { Button, Form, FormGroup, Input, Label } from "reactstrap";
import { bindActionCreators } from "redux";
import * as orderActions from "../_actions/orderActions";
import * as productActions from "../_actions/productActions";
import MyOrder from "./myOrder";

class OrderNow extends React.Component<*, *> {
  constructor(props) {
    super(props);

    this.state = {
      customerName: ""
    };
  }

  componentDidMount() {
    this.props.orderActions.loadCart();
  }
  render() {
    return (
      <div className="container order-now">
        <MyOrder order={this.props.orders.order} />
        {this.props.orders.placed !== true && (
          <Form>
            <FormGroup>
              <Label for="customerName">Customer Name</Label>
              <Input
                type="text"
                name="customerName"
                id="customerName"
                placeholder="Your Name"
                onChange={e => this.setState({ customerName: e.target.value })}
                value={this.state.customerName}
              />
            </FormGroup>

            <Button
              onClick={() =>
                this.props.orderActions.placeOrder(
                  this.state.customerName,
                  this.props.orders.order
                )
              }
            >
              Place Order
            </Button>
          </Form>
        )}
        <div>
          {this.props.orders.placed && <div>Your Order has been placed!</div>}
        </div>
      </div>
    );
  }
}

OrderNow = connect(
  (state, ownProps) => {
    return {
      orders: state.orders
    };
  },
  (dispatch: Function) => {
    return {
      productActions: bindActionCreators(productActions, dispatch),
      orderActions: bindActionCreators(orderActions, dispatch)
    };
  }
)(OrderNow);

export default OrderNow;
