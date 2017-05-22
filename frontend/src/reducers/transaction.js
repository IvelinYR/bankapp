export default function amount(state = [], action) {
    switch (action.type) {
        case 'DEPOSIT_SUCCESS':
            let after_deposit = action.payload.data;
            return Object.assign({}, state, {
                amount: after_deposit.Amount,
                error: ''
            });
        case 'WITHDRAWAL_SUCCESS':
            let after_withdrawal = action.payload.data;
            return Object.assign({}, state, {
                amount: after_withdrawal.Amount,
                error: ''
            });
        case 'WITHDRAWAL_FAIL':
            return Object.assign({}, state, {
                error: action.error.response.data.Message,
            });
        default:
            return state
    }
}