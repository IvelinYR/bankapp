export default function accounts(state = [], action) {
    switch (action.type) {
        case 'ADD_ACCOUNT_SUCCESS':
            return state
        case 'LOAD_ACCOUNTS_SUCCESS':
            let accountList = action.payload.data
            return accountList;
        default:
            return state
    }
}