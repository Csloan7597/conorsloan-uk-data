import * as api from '../api';
import * as actionTypes from './action-types'

export const getProjectNavList = () => dispatch => {
  console.log(api);
  // TODO: Error handling, clear errors, etc
  return api.fetchProjectNavList().then(json => {
    dispatch({
      type: actionTypes.SET_PROJECT_NAV_LIST,
      data: json
    });
  });
}
