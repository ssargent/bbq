//@flow
import React from "react";
import "./TopNav.css";
import Login from "./components/login";

export default function TopNav() {
  return (
    <nav className="navbar navbar-dark fixed-top bg-dark flex-md-nowrap p-0 shadow">
      <a
        className="navbar-brand col-sm-3 col-md-2 mr-0"
        href="https://mythicalcodelabs.com/bbq"
      >
        BBQ <small>by Mythical Code Labs</small>
      </a>

      <ul className="navbar-nav px-3">
        <li className="nav-item text-nowrap">
          <Login />
        </li>
      </ul>
    </nav>
  );
}
