import React from 'react';
import ReactDOM from 'react-dom';
import allReducers from './reducers'
import {IndexRoute, Router, Route, browserHistory} from 'react-router'
import {createStore, applyMiddleware} from 'redux';
import {composeWithDevTools} from 'redux-devtools-extension';
import {Provider} from 'react-redux';
import {axiosLoad}  from './axios.mock';
import axiosMiddleware from 'redux-axios-middleware';
import './index.css';

import App from './App';
import AccountListContainer from './container/AccountListContainer';
import CreateAccountContainer from './container/CreateAccountContainer';
import LoginContainer from './container/LoginContainer';
import TransactionContainer from './container/TransactionContainer';

let axios = require('axios');
let store = createStore(allReducers, {}, composeWithDevTools(applyMiddleware(axiosMiddleware(axios))));
axios.defaults.baseURL = '/';

if (process.env.NODE_ENV === 'development') {
    axiosLoad()
}

ReactDOM.render(
    <Provider store={store}>
        <Router history={browserHistory}>
            <Route path="/" component={App}>
                <IndexRoute component={LoginContainer}/>
                <Route path="home" component={AccountListContainer}/>
                <Route path="account" component={CreateAccountContainer}/>
                <Route path="/transaction/:id" component={TransactionContainer}/>
                <Route path="login" component={LoginContainer}/>
            </Route>
        </Router>
    </Provider>,
    document.getElementById('root')
);