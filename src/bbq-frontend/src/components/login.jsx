//@flow

import React from "react"
import axios from "axios"
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from 'reactstrap';
import { API_SERVER } from "../config"

class Login extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            showLoginModal: false,
            isAuthenticated: false,
            loginName: '',
            password: ''
        };
    }

    toggle = () => {
      this.setState(prevState => ({
          showLoginModal: !prevState.showLoginModal
      }));
    };


    login = () => {
      const loginModel = {
        loginname: this.state.loginName,
        password: this.state.password
      };

      let url = "";
    
      if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
          url = `${API_SERVER}v1/system/accounts/login`
      } else {
          url = `${API_SERVER}v1/system/accounts/signin`
      }

      axios
        .post(url, loginModel)
        .then(resp => {
          if(resp.data.success === true) {
            localStorage.setItem("bbq-authenticated", JSON.stringify(resp.data))
            console.log("Successfully Authenticated");
            this.setState({ showLoginModal: false });
          }
        })
        .catch(err => {
          console.log(err);
        });
    };

    logout = () => {

    };

    componentDidMount() {
      if(localStorage.getItem("bbq-authenticated") === null) {
        this.isAuthenticated = false;
        this.setState({ showLoginModal: true });
      }
      else
        this.isAuthenticated = true;
    }

    render() {

      if(this.isAuthenticated) {
        return <button onClick={() => localStorage.removeItem("bbq-authenticated")}>Logout</button>
      }

      return (
        <div>
          <a onClick={this.toggle}>Login</a>
          <Modal isOpen={this.state.showLoginModal} toggle={this.toggle} className={this.props.className}>
            <ModalHeader toggle={this.toggle}>Login</ModalHeader>
            <ModalBody>
              <form>
                <div className="form-group">
                  <input type="text" className="form-control" placeholder="Your Email *" 
                  onChange={e => this.setState({ loginName: e.target.value})}
                  value={this.state.loginName} />
                </div>
                <div className="form-group">
                  <input type="password" className="form-control" placeholder="Your Password *" 
                   onChange={e => this.setState({ password: e.target.value})}
                   value={this.state.password} 
                  />
                </div>
                <div className="form-group">
                  <a href="#" className="ForgetPwd">Forget Password?</a>
                </div>
              </form>            
            </ModalBody>
            <ModalFooter>
              <Button color="primary" onClick={this.login}>Login</Button>
              <Button color="secondary" onClick={this.toggle}>Cancel</Button>
            </ModalFooter>
          </Modal>
        </div>
    )
    }
}

export default Login;