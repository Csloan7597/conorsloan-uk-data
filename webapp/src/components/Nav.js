import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as navActions from '../actions/nav-actions';
import { Router, Route, Link, browserHistory } from 'react-router';
import {Navbar, NavItem, NavDropdown, MenuItem, Nav} from 'react-bootstrap';
import { IndexLinkContainer, LinkContainer } from 'react-router-bootstrap';

class NavigationBar extends Component {

  componentWillMount() {
    console.log("TERYAKI FRIES");
    console.log(this.props)
    if (this.props.navData.projectNavList === null) {
        console.log("Gonna get stuff")
        this.props.navActions.getProjectNavList();
    }
  }

  render() {

    const {navData, routing, navActions} = this.props;

    let pathname = routing.locationBeforeTransitions.pathname !== null ?
                   routing.locationBeforeTransitions.pathname : "";

    return (
      <Navbar collapseOnSelect>
        <Navbar.Header>
          <Navbar.Brand>
            <Link to="/">LOGO</Link>
          </Navbar.Brand>
          <Navbar.Toggle />
        </Navbar.Header>
        <Navbar.Collapse>
          <Nav>
            <IndexLinkContainer to="/">
              <NavItem eventKey={1}>
                Home
              </NavItem>
            </IndexLinkContainer>
            <LinkContainer to="/about">
              <NavItem eventKey={2}>
                About Me
              </NavItem>
            </LinkContainer>
            <NavDropdown eventKey={3} title="Projects" id="basic-nav-dropdown">
              <LinkContainer to="/projects">
                <MenuItem eventKey={3.1}>All Projects</MenuItem>
              </LinkContainer>
              <MenuItem divider />
              {
                navData.projectNavList === null ? ""
                : navData.projectNavList.map( (project, index) => {
                  let eventKey = 3 + ((index + 2) / 10);
                  let path = "/project/" + project.id;
                  return <LinkContainer to={path}>
                    <MenuItem eventKey={eventKey}>{project.name}</MenuItem>
                  </LinkContainer>
                })
              }
            </NavDropdown>
          </Nav>
        </Navbar.Collapse>
      </Navbar>
    )
  }
}

function mapStateToProps(state) {
  return {
    navData: state.navData,
    routing: state.routing
  }
}

function mapDispatchToProps(dispatch) {
  return {
    navActions: bindActionCreators(navActions, dispatch)
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(NavigationBar)


// <Navbar inverse collapseOnSelect>
//   <Navbar.Header>
//     <Navbar.Brand>
//       <a href="#">React-Bootstrap</a>
//     </Navbar.Brand>
//     <Navbar.Toggle />
//   </Navbar.Header>
//   <Navbar.Collapse>
//     <Nav>
//       <NavItem eventKey={1} href="#">Link</NavItem>
//       <NavItem eventKey={2} href="#">Link</NavItem>
//       <NavDropdown eventKey={3} title="Dropdown" id="basic-nav-dropdown">
//         <MenuItem eventKey={3.1}>Action</MenuItem>
//         <MenuItem eventKey={3.2}>Another action</MenuItem>
//         <MenuItem eventKey={3.3}>Something else here</MenuItem>
//         <MenuItem divider />
//         <MenuItem eventKey={3.3}>Separated link</MenuItem>
//       </NavDropdown>
//     </Nav>
//     <Nav pullRight>
//       <NavItem eventKey={1} href="#">Link Right</NavItem>
//       <NavItem eventKey={2} href="#">Link Right</NavItem>
//     </Nav>
//   </Navbar.Collapse>
// </Navbar>
