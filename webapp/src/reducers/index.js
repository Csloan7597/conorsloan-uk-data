import {routerReducer} from 'react-router-redux'
import * as actionTypes from '../actions/action-types'
import {combineReducers} from 'redux';

import aboutMeReducer from './about';
import homeReducer from './home';

const init = {
  aboutMeData: {},
  homeData: {}
}

export default combineReducers({
  routing: routerReducer,
  aboutMeData: aboutMeReducer,
  homeData: homeReducer
})
