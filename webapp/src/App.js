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
        {this.props.children}
        <Footer />
      </div>
    );
  }
}

export default App;
