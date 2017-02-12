import * as api from '../api';
import * as actionTypes from './action-types';

export const getProjects = () => dispatch => {
    return api.fetchProjects().then(json =>
      dispatch({
        type: actionTypes.SET_PROJECTS,
        data: json
      })
    )
}
