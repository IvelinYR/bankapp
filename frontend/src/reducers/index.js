import {combineReducers} from 'redux'
import {routeReducer} from 'redux-simple-router';
import {reducer as formReducer} from 'redux-form';

import accounts from './account';
import login from './login';
import transaction from './loadTransaction'
import amount from './transaction'

const rootReducer = combineReducers({
    accounts: accounts,
    login: login,
    transaction: transaction,
    amount: amount,
    routing: routeReducer,
    form: formReducer
});

export default rootReducer;