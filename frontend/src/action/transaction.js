export function deposit(data) {
    return {
        type: 'DEPOSIT',
        payload: {
            request: {
                method: 'post',
                data: data,
                url: '/transaction/deposits'
            }
        }
    }
}

export function withdrawal(data) {
    return {
        type: 'WITHDRAWAL',
        payload: {
            request: {
                method: 'post',
                data: data,
                url: '/transaction/withdrawals'
            }
        }
    }
}

export function loadTransactions(data) {
    return {
        type: 'LOAD_TRANSACTIONS',
        payload: {
            request: {
                method:'post',
                data:data,
                url: '/history'
            }
        }
    }
}

export function loadAmount(data) {
    return {
        type: 'LOAD_AMOUNT',
        payload: {
            request: {
                method:'post',
                data:data,
                url: '/transaction'
            }
        }
    }
}


