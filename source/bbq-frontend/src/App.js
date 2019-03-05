import React, { Component } from 'react';
import TopNav from './TopNav';
import PageContent from './PageContent';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      devices: []
    }
  }
   
  render() {
    return (
      <React.Fragment>
     <TopNav />
     <PageContent />
     </React.Fragment>
    );
  }
}

export default App;
