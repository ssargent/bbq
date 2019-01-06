import "bootstrap/dist/css/bootstrap.css";
import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { applyMiddleware, createStore } from "redux";
import reduxthunk from "redux-thunk";
import App from "./App";
import "./index.css";
import * as serviceWorker from "./serviceWorker";
import bbqStoreApp from "./_reducers";

const createStoreWithMiddleware = applyMiddleware(reduxthunk)(createStore);

let store = createStoreWithMiddleware(
  bbqStoreApp,
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
);

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();
