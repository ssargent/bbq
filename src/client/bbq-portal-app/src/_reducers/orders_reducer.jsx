import { ORDERS_LOAD_ORDER, ORDERS_ORDER_PLACED } from "../_actions/types";

export default function(state = { loading: true }, action) {
  switch (action.type) {
    case ORDERS_LOAD_ORDER:
      return {
        ...state,
        order: action.order
      };
    case ORDERS_ORDER_PLACED:
      return {
        ...state,
        order: action.order,
        placed: action.placed
      };
    default:
      return state;
  }
}
