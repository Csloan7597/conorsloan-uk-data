import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as aboutMeActions from '../actions/about-actions';

class AboutContainer extends Component {

  componentWillMount() {
    if (this.props.aboutMeData == null) {
        this.props.aboutMeActions.getAboutMeData();
    }
  }

  render() {
    const {aboutMeData, aboutMeActions} = this.props;

    return (
      <p> ABOUT ME {aboutMeData === null ? "null" : aboutMeData.content[0]}</p>
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
