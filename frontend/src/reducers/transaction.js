export default function amount(state = [], action) {
    switch (action.type) {
        case 'DEPOSIT_SUCCESS':
            let after_deposit = action.payload.data;
            return Object.assign({}, state, {
                amount: after_deposit.Total,
            });
        case 'WITHDRAWAL_SUCCESS':
            let after_withdrawal = action.payload.data;
            return Object.assign({}, state, {
                amount: after_withdrawal.Total,
            });
        default:
            return state
    }
}