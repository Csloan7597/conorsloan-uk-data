import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as homeActions from '../actions/home-actions';

class HomeContainer extends Component {

  componentWillMount() {
    this.props.homeActions.getTagline()
    this.props.homeActions.getGlanceItems()
    //this.props.homeActions.getImages()
  }

  render() {
    const {homeData, homeActions} = this.props
    console.log("PRINTING PROPS");
    console.log(this.props);

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
