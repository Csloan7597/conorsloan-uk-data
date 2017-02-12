import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as aboutMeActions from '../actions/about-actions';

import PageHeading from './PageHeading';

class AboutContainer extends Component {

  componentWillMount() {
    if (this.props.aboutMeData == null) {
        this.props.aboutMeActions.getAboutMeData();
    }
  }

  render() {
    const {aboutMeData, aboutMeActions} = this.props;

    return (
      <div className="container">
        <PageHeading title="About Me" tagline="Because I can tell you're interested..." />
        <p> ABOUT ME {aboutMeData === null ? "null" : aboutMeData.content[0]}</p>
      </div>
    );
  }

}

function mapStateToProps(state) {
  return {
    aboutMeData: state.aboutMeData
  }
}

function mapDispatchToProps(dispatch) {
  return {
    aboutMeActions: bindActionCreators(aboutMeActions, dispatch)
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(AboutContainer)
