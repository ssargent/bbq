//@flow
import axios from "axios";
import { API_HOST } from "../config";
import { ORDERS_LOAD_ORDER, ORDERS_ORDER_PLACED } from "./types";

export function createOrder(order: object) {
  return dispatch => {
    axios.post(`${API_HOST}/api/orders`, order).then(resp => {
      localStorage.setItem("current-order-id", resp.data.id);
      dispatch(loadOrder(resp.data));
    });
  };
}

export function loadCart() {
  return dispatch => {
    let orderId = localStorage.getItem("current-order-id");

    if (orderId !== undefined && orderId !== null) {
      axios.get(`${API_HOST}/api/orders/${orderId}`).then(resp => {
        dispatch(loadOrder(resp.data));
      });
    }
  };
}

export function addToCart(newItem: object) {
  return dispatch => {
    let orderId = localStorage.getItem("current-order-id");

    if (orderId === undefined || orderId === null)
      orderId = "00000000-0000-0000-0000-000000000000";

    axios
      .post(`${API_HOST}/api/orders/${orderId}/items`, newItem)
      .then(resp => {
        localStorage.setItem("current-order-id", resp.data.id);
        dispatch(loadOrder(resp.data));
      });
  };
}

export function placeOrder(customerName: string, order: object) {
  return dispatch => {
    order.customerName = customerName;
    const orderToPlace = { ...order, customerName: customerName };
    axios
      .post(`${API_HOST}/api/orders/${order.id}/checkout`, orderToPlace)
      .then(resp => {
        dispatch(orderPlaced(resp.data, true));
        localStorage.removeItem("current-order-id");
      });
  };
}

export function orderPlaced(order: object, placed: boolean) {
  return {
    order,
    placed,
    type: ORDERS_ORDER_PLACED
  };
}

export function loadOrder(order: object) {
  return {
    order,
    type: ORDERS_LOAD_ORDER
  };
}
