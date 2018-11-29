import { ORDERS_LOAD_ORDER } from "../_actions/types";

export default function(state = { loading: true }, action) {
  switch (action.type) {
    case ORDERS_LOAD_ORDER:
      return {
        ...state,
        order: action.order
      };
    default:
      return state;
  }
}
