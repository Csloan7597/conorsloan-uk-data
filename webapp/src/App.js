import React, { Component } from 'react';
import { Button } from 'react-bootstrap';
import { Router, Route, Link, browserHistory } from 'react-router'

import logo from './logo.svg';
import './App.css';

import Header from './components/Header';
import Footer from './components/Footer';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Header {...this.props.navData}/>
        <div className="container">
          {this.props.children}
        </div>
        <Footer />
      </div>
    );
  }
}

export default App;
