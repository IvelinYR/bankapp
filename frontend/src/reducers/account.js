import * as Cookies from "js-cookie";
import {browserHistory} from 'react-router';

export default function accounts(state = [], action) {
    switch (action.type) {
        case 'ADD_ACCOUNT_SUCCESS':
            browserHistory.push('/home');
            return Object.assign({}, state, {
                error: ''
            });
        case 'ADD_ACCOUNT_FAIL':
            return Object.assign({}, state, {
                error: action.error.response.data.Message
            });
        case 'LOAD_ACCOUNTS_SUCCESS':
            let accountsList = action.payload.data;
            return accountsList;
        case 'LOGOUT_SUCCESS':
            Cookies.remove('SID');
            return state;
        case 'LOGIN_SUCCESS':
            const value = action.payload.data.SessionID;
            Cookies.set('SID', value);
            return state;
        case 'LOGIN_FAIL':
            browserHistory.push('/login');
            return state;
        default:
            return state
    }
}