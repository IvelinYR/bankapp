export function loadAccounts() {
    return {
        type: 'LOAD_ACCOUNTS',
        payload: {
            request: {
                method: 'get',
                url: '/accounts'
            }
        }
    }
}

export function addAccount(data) {
    return {
        type: 'ADD_ACCOUNT',
        payload: {
            request: {
                method: 'post',
                data: data,
                url: '/accounts'
            }
        }
    }
}

