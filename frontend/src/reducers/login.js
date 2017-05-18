import {browserHistory} from 'react-router';

export default function login(state = [], action) {
    switch (action.type) {
        case 'LOGIN_SUCCESS':
            const cname = "SID";
            const cvalue = action.payload.data;

        function setCookie(cname, cvalue, days) {
            let d = new Date();
            d.setTime(d.getTime() + (days * 1000));
            let expires = "expires=" + d.toUTCString();
            document.cookie = cname + "=" + cvalue + ";" + expires + "; path=/";
        }

            setCookie(cname, cvalue, 300);

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
