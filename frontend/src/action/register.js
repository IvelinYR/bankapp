export function register(data) {
    return {
        type: 'REGISTER',
        payload: {
            request: {
                method: 'post',
                url: '/v1/users/signup',
                data: data
            }
        }
    }
}