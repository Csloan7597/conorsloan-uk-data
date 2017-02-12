import {routerReducer} from 'react-router-redux'
import * as actionTypes from '../actions/action-types'
import {combineReducers} from 'redux';

import aboutMeReducer from './about';
import homeReducer from './home';
import headerReducer from './header'

const init = {
  aboutMeData: {},
  homeData: {},
  navData: {}
}

export default combineReducers({
  routing: routerReducer,
  aboutMeData: aboutMeReducer,
  homeData: homeReducer,
  navData: headerReducer
})
