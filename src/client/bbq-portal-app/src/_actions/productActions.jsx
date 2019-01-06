//@flow
import axios from "axios";
import { API_HOST } from "../config";
import { PRODUCTS_LOADED, PRODUCTS_LOADING } from "./types";

export function productsLoading(loading: boolean) {
  return {
    loading,
    type: PRODUCTS_LOADING
  };
}

export function productsLoaded(loaded, products) {
  return {
    loading: !loaded,
    products,
    type: PRODUCTS_LOADED
  };
}

export function loadProducts() {
  return dispatch => {
    dispatch(productsLoading(true));
    console.log("Starting to load products");
    axios.get(`${API_HOST}/api/products`).then(resp => {
      console.log(resp.data);
      dispatch(productsLoaded(true, resp.data));
    });
  };
}
