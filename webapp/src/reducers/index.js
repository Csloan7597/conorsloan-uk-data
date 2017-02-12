import {routerReducer} from 'react-router-redux'
import * as actionTypes from '../actions/action-types'
import {combineReducers} from 'redux';

import aboutMeReducer from './about';
import homeReducer from './home';
import headerReducer from './header'
import projectReducer from './project'

const init = {
  aboutMeData: {},
  homeData: {},
  navData: {},
  projectData: {}
}

export default combineReducers({
  routing: routerReducer,
  aboutMeData: aboutMeReducer,
  homeData: homeReducer,
  navData: headerReducer,
  projectData: projectReducer
})
