export default function transactions(state = [], action) {
    switch (action.type) {
        case 'LOAD_TRANSACTIONS_SUCCESS':
            let transactionList = action.payload.data;
            return transactionList;
        default:
            return state
    }
}