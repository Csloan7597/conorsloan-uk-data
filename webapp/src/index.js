import React from 'react';
import ReactDOM from 'react-dom';

import thunk from 'redux-thunk';
import { createStore, applyMiddleware } from 'redux';

import { Provider } from 'react-redux';
import { Router, Route, browserHistory, IndexRoute } from 'react-router';
import { syncHistoryWithStore, routerMiddleware } from 'react-router-redux';
import createLogger from 'redux-logger';

import App from './App';
import './index.css';

import HomeContainer from './components/Home'
import AboutContainer from './components/About'
import NotFound from './components/NotFound'

import rootReducer from './reducers';

const configureStore = (history) => {
  return createStore(
    rootReducer,
    applyMiddleware(
      thunk,
      createLogger(),
      routerMiddleware(history)
    )
  )
};

const store = configureStore(browserHistory);
const history = syncHistoryWithStore(browserHistory, store); // TODO: idk if need?

ReactDOM.render(
  (
    <Provider store={store}>
      <Router history={history}>
        <Route path="/" component={App}>
          <IndexRoute component={HomeContainer} />
          <Route path="about" component={AboutContainer} />
          <Route path="*" component={NotFound} />
        </Route>
      </Router>
    </Provider>
),
  document.getElementById('root')
);

// <Route path="projects" component={ProjectsContainer} />
// <Route path="cv" component={CvContainer} />
// <Route path="contact" component={ContactContainer} />
// <Route path="project/:projectID" component={ProjectContainer} />
