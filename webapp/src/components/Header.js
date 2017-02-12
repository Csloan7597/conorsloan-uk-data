import React, { Component } from 'react';
import { Router, Route, Link, browserHistory } from 'react-router';
import NavBar from './Nav';

export default (state) => {
  return (
      <div className="header">
        <NavBar {...state}/>
      </div>
  )
}
// TODO: This might not be the correct way to pass props...
