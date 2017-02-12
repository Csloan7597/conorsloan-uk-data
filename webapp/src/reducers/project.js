import * as actionTypes from '../actions/action-types'

const init = {};

export default (state = init, action) => {
  switch (action.type) {
    case actionTypes.SET_PROJECTS:
      return { projects: action.data };
    default:
      return state;
  }
}
