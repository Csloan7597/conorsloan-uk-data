import * as actionTypes from '../actions/action-types'

const init = null;

export default (state = init, action) => {
  switch (action.type) {
    case actionTypes.SET_ABOUTME_DATA:
      return action.data;
    default:
      return state;
  }
}
