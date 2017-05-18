export function loadAccounts(data) {
    return {
        type: 'LOAD_ACCOUNTS',
        payload: {
            request: {
                method: 'get',
                data: data,
                url: '/v1/users/me/accounts'
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
                // headers: {'Set-Cookie': 'SID='+ getCookie("SID")},
                data: data,
                url: '/v1/users/me/new-account'
            }
        }
    }
}

