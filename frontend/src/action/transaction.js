export function deposit(data) {

    return {
        type: 'DEPOSIT',
        payload: {
            request: {
                method: 'post',
                data: data,
                url: '/v1/users/me/deposit'
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
                url: '/v1/users/me/withdraw'
            }
        }
    }
}

export function loadTransactions(data) {
    console.log(data);
    return {
        type: 'LOAD_TRANSACTIONS',
        payload: {
            request: {
                method:'post',
                data:data,
                url: '/v1/users/me/account-history'
            }
        }
    }
}