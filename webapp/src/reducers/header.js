import * as actionTypes from '../actions/action-types'

const init = {
  projectNavList: null
};

export default (state = init, action) => {
  switch (action.type) {
    case actionTypes.SET_PROJECT_NAV_LIST:
      return {
        projectNavList: action.data
      };
    default:
      return state;
  }
}
