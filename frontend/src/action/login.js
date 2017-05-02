export function login(data) {
    return {
        type: 'LOGIN',
        payload: {
            request: {
                method: 'get',
                url: '/login',
                data: data
            }
        }
    }
}