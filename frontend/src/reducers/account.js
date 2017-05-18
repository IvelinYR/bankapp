import {browserHistory} from 'react-router';

export default function accounts(state = [], action) {
    switch (action.type) {
        case 'ADD_ACCOUNT_SUCCESS':
            let obj = (action.payload.data);
            const newAccount = {
                id: obj.id,
                title: obj.title,
                Type: obj.Type,
                Total: obj.Total,
                Currency: obj.Currency
            };
            return [...state, newAccount];
        case 'ADD_ACCOUNT_FAIL':
            browserHistory.push('/account');
            return Object.assign({}, state, {
                error: action.error.response.data
            });
        case 'LOAD_ACCOUNTS_SUCCESS':
            return action.payload.data;
        default:
            return state
    }
}