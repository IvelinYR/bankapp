import {combineReducers} from 'redux'
import {routeReducer} from 'redux-simple-router';
import {reducer as formReducer} from 'redux-form';

import accounts from './account';
import transaction from './loadTransaction';
import amount from './transaction';
import register from './register';

const rootReducer = combineReducers({
    accounts: accounts,
    transaction: transaction,
    amount: amount,
    register: register,
    routing: routeReducer,
    form: formReducer
});

export default rootReducer;
