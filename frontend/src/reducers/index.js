import {combineReducers} from 'redux'
import {routeReducer} from 'redux-simple-router';
import {reducer as formReducer} from 'redux-form';

import accounts from './account';
import login from './login';
import name from './login';
import transaction from './loadTransaction';
import amount from './transaction';
import register from './register';

const rootReducer = combineReducers({
    accounts: accounts,
    login: login,
    transaction: transaction,
    amount: amount,
    register: register,
    name: name,
    routing: routeReducer,
    form: formReducer
});

export default rootReducer;
