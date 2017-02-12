import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as homeActions from '../actions/home-actions';
import * as projectActions from '../actions/project-actions';

import PageHeading from './PageHeading';

class HomeContainer extends Component {

  componentWillMount() {
    if (this.props.homeData.tagline == null) {
      this.props.homeActions.getTagline()
    }
    if (this.props.homeData.glanceItems == null) {
      this.props.homeActions.getGlanceItems()
    }
    if (this.props.projectData.projects == null) {
      this.props.projectActions.getProjects()
    }
  }

  render() {
    const {homeData, homeActions, projects, projectActions} = this.props

    return (
      <div className="container">
        <PageHeading
            title="Conor Sloan | Software Developer"
            tagline={homeData.tagline == null ? "" : homeData.tagline}
        />
        <div>
            <p> Conor Sloan -  {homeData.tagline} </p>
            {
              homeData.glanceItems === null
               ? <p>Loading</p>
               : <ul>
                  {homeData.glanceItems.map( (glanceItem) => <li> {glanceItem.title} </li>)}
                 </ul>
            }

        </div>
      </div>
    )
  }

}

function mapStateToProps(state) {
  return {
    homeData: state.homeData,
    projectData: state.projectData
  }
}

function mapDispatchToProps(dispatch) {
  return {
    homeActions: bindActionCreators(homeActions, dispatch),
    projectActions: bindActionCreators(projectActions, dispatch)
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(HomeContainer)
