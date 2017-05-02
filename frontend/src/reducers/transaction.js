export default function amount(state = [], action) {
    switch (action.type) {
        case 'DEPOSIT_SUCCESS':
            let after_deposit = action.payload.data;
            return Object.assign({}, state, {
                amount: after_deposit.amount,
                currency: after_deposit.currency,
                id: after_deposit.id,
                title: after_deposit.title,
                type: after_deposit.type
            });
        case 'WITHDRAWAL_SUCCESS':
            let after_withdrawal = action.payload.data;
            return Object.assign({}, state, {
                amount: after_withdrawal.amount,
                currency: after_withdrawal.currency,
                id: after_withdrawal.id,
                title: after_withdrawal.title,
                type: after_withdrawal.type
            });
        case 'LOAD_AMOUNT_SUCCESS':
            let currentAccount = action.payload.data;
            return Object.assign({}, state, {
                amount: currentAccount.amount,
                currency: currentAccount.currency,
                id: currentAccount.id,
                title: currentAccount.title,
                type: currentAccount.type
            });
        default:
            return state
    }
}