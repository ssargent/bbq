import React, { Component } from 'react';
import axios from "axios";
import {
  Collapse,
  Navbar,
  NavbarToggler,
  NavbarBrand,
  Nav,
  NavItem,
  NavLink,
  UncontrolledDropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem } from 'reactstrap';
import logo from './logo.svg';
import './TopNav.css';

class TopNav extends Component {
  render() {
    return (
        <nav className="navbar navbar-dark fixed-top bg-dark flex-md-nowrap p-0 shadow">
        <a className="navbar-brand col-sm-3 col-md-2 mr-0" href="#">BBQ by Mythical Code Labs</a>
  
  <ul className="navbar-nav px-3">
    <li className="nav-item text-nowrap">
      <a className="nav-link" href="#">Sign out</a>
    </li>
  </ul> 
        </nav>

    );
  }
}

export default TopNav;
