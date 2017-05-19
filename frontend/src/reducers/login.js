import {browserHistory} from 'react-router';
import * as Cookies from "js-cookie";

export default function login(state = [], action) {
    switch (action.type) {
        case 'LOGIN_SUCCESS':
            const name = "SID";
            const value = action.payload.data;
            Cookies.set(name, value);
            return Object.assign({}, state, {
                name: action.payload.data
            });
        case 'LOGIN_FAIL':
            browserHistory.push('/login');
            return state;
        default:
            return state
    }

}
