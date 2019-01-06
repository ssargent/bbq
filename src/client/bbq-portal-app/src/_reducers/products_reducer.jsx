import { PRODUCTS_LOADED, PRODUCTS_LOADING } from "../_actions/types";

export default function(state = { loading: true }, action) {
  switch (action.type) {
    case PRODUCTS_LOADED:
      console.log(action);
      return {
        ...state,
        products: action.products,
        loading: action.loading
      };
    case PRODUCTS_LOADING:
      return {
        ...state,
        loading: action.loading
      };
    default:
      return state;
  }
}
