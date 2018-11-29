import numeral from "numeral";
import React from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import * as orderActions from "../_actions/orderActions";
import * as productActions from "../_actions/productActions";

type Props = {
  actions: object,
  products: Array<object>
};

class Menu extends React.Component<Props, *> {
  componentDidMount() {
    this.props.productActions.loadProducts();
  }

  render() {
    if (
      this.props.products === undefined ||
      this.props.products.loading === true
    )
      return <div>Loading...</div>;
    const products = this.props.products.products || [];
    return (
      <div>
        {products.map(p => {
          return (
            <div className="row" key={p.id}>
              <div className="col-md-10">
                <span className="product-name">{p.name}</span>{" "}
                <span>
                  {numeral(p.price).format("$0,0.00")}/{p.unit}
                </span>
                <div>
                  <em property="italic">{p.description}</em>
                </div>
              </div>
              <div className="col-md-2">
                <button
                  className="btn btn-primary"
                  onClick={() =>
                    this.props.orderActions.addToCart({
                      productId: p.id,
                      quantity: 1
                    })
                  }
                >
                  Add to Cart
                </button>
              </div>
            </div>
          );
        })}
      </div>
    );
  }
}

Menu = connect(
  (state, ownProps) => {
    return {
      products: state.products
    };
  },
  (dispatch: Function) => {
    return {
      productActions: bindActionCreators(productActions, dispatch),
      orderActions: bindActionCreators(orderActions, dispatch)
    };
  }
)(Menu);

export default Menu;
