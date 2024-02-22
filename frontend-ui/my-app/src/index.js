import React, { Component } from 'react';
import { render } from 'react-dom';
import { BrowserRouter as Router, Route, Link, Switch  } from 'react-router-dom';
import { RouteBody } from '../src/routes/routes';
import 'bootstrap/dist/css/bootstrap.min.css';

import '../src/assets/styles.css';

class App extends Component {
  constructor() {
    super();
    this.state = {
      name: 'React'

      
    };
  }
  
  render() {
    return (
      <RouteBody />
    );
  }
}

render(<App />, document.getElementById('root'));
