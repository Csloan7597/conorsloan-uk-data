import * as api from '../api';
import * as actionTypes from './action-types'

export const getAboutMeData = () => dispatch => {
  // TODO: Error handling, clear errors, etc
  return api.fetchAboutMeData().then(json => {
    dispatch({
      type: actionTypes.SET_ABOUTME_DATA,
      data: json
    });
  });
}
