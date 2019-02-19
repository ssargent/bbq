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
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      devices: []
    }
  }
  componentDidMount() {
    axios.get("https://bbq.k8s.ssargent.net/v1/development/bbq/devices")
      .then(resp => this.setState({ devices: resp.data}))
      .catch(err => console.log(err));
  }
  render() {
    return (
       <Navbar color="dark" dark expand="md">
          <NavbarBrand href="/">Go BBQ</NavbarBrand>
          <NavbarToggler onClick={this.toggle} />
          <Collapse isOpen={this.state.isOpen} navbar>
            <Nav className="ml-auto" navbar>
            <UncontrolledDropdown nav inNavbar>
                <DropdownToggle nav caret>
                  Monitors
                </DropdownToggle>
                <DropdownMenu right>
                {this.state.devices.map(d => (<DropdownItem key={d.id}>
          {d.name}
         </DropdownItem>))}
                  <DropdownItem divider />
                  <DropdownItem>
                    Add New
                  </DropdownItem>
                </DropdownMenu>
              </UncontrolledDropdown>
              <UncontrolledDropdown nav inNavbar>
                <DropdownToggle nav caret>
                  Devices
                </DropdownToggle>
                <DropdownMenu right>
                {this.state.devices.map(d => (<DropdownItem key={d.id}>
          {d.name}
         </DropdownItem>))}
                  <DropdownItem divider />
                  <DropdownItem>
                    Add New
                  </DropdownItem>
                </DropdownMenu>
              </UncontrolledDropdown>
            </Nav>
          </Collapse>
        </Navbar>

    );
  }
}

export default App;
