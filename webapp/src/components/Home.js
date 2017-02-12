import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as homeActions from '../actions/home-actions';

class HomeContainer extends Component {

  componentWillMount() {
    if (this.props.homeData.tagline == null) {
      this.props.homeActions.getTagline()
    }
    if (this.props.homeData.glanceItems == null) {
      this.props.homeActions.getGlanceItems()
    }
    //this.props.homeActions.getImages()
  }

  render() {
    const {homeData, homeActions} = this.props

    return (
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
    )
  }

}

function mapStateToProps(state) {
  return {
    homeData: state.homeData
  }
}

function mapDispatchToProps(dispatch) {
  return {
    homeActions: bindActionCreators(homeActions, dispatch)
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(HomeContainer)
