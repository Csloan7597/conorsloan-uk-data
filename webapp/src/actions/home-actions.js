import * as api from '../api';
import * as actionTypes from './action-types'

export const getGlanceItems = () => dispatch => {
  console.log(api);
  // TODO: Error handling, clear errors, etc
  return api.fetchGlanceItems().then(json => {
    dispatch({
      type: actionTypes.SET_GLANCE_ITEMS,
      data: json
    });
  });
}

export const getTagline = () => dispatch => {
  console.log(api);
  // TODO: Error handling, clear errors, etc
  return api.fetchTagLine().then(json => {
    dispatch({
      type: actionTypes.SET_TAGLINE,
      data: json
    });
  });
}
