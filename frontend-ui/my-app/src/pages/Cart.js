import React, { Component } from 'react';
import { render } from 'react-dom';
import '../assets/styles.css';


export class Cart extends Component {
  constructor() {
    super();
    this.state = {
      name: 'React'
    };
  }

  render() {
    return (
      <div>
        <h2>Cart Page</h2>
      </div>
    );
  }
}
