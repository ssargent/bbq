import { combineReducers } from "redux";
//import { reducer as form } from "redux-form";
import ordersReducer from "./orders_reducer";
import productsReducer from "./products_reducer";

const bbqStoreApp = combineReducers({
  // form,
  products: productsReducer,
  orders: ordersReducer
});

export default bbqStoreApp;
