import {browserHistory} from 'react-router';

export default function register(state = [], action) {
    switch (action.type) {
        case 'REGISTER_SUCCESS':
            return Object.assign({}, state, {
                name: action.payload.data.username,
            });

        case 'REGISTER_FAIL':
            browserHistory.push('/register');
            return Object.assign({}, state, {
                error: action.error.response.data
            });
        default:
            return state
    }

}