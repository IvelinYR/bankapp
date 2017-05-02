import {browserHistory} from 'react-router';

export default function login(state = [], action) {
    switch (action.type) {
        case 'LOGIN_SUCCESS':
            return state;
        case 'LOGIN_FAIL':
            browserHistory.push('/login');
            return state;
        default:
            return state
    }

}
