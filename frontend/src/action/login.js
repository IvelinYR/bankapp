export function login(data) {
    return {
        type: 'LOGIN',
        payload: {
            request: {
                method: 'post',
                url: '/v1/users/login',
                data: data
            }
        }
    }
}

export function logout() {
    return {
        type: 'LOGOUT',
        payload: {
            request: {
                method: 'post',
                url: '/v1/users/logout',
            }
        }
    }
}