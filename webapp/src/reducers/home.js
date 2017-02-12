import * as actionTypes from '../actions/action-types'

const init = {
  tagline: null,
  glanceItems: null
};

export default (state = init, action) => {
  switch (action.type) {
    case actionTypes.SET_TAGLINE:
      return {...state, tagline: action.data};
    case actionTypes.SET_GLANCE_ITEMS:
      return {...state, glanceItems: action.data};
    default:
      return state;
  }
}
