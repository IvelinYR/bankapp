export default function accounts(state = [], action) {
    switch (action.type) {
        case 'ADD_ACCOUNT_SUCCESS':
            let obj = (action.payload.data);
            const newAccount = {
                id: obj.id,
                title: obj.title,
                type: obj.type,
                amount: obj.amount,
                currency: obj.currency
            };
            return [...state, newAccount];
        case 'LOAD_ACCOUNTS_SUCCESS':
            return action.payload.data;
        default:
            return state
    }
}