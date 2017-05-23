import * as actions from './login';

const LOGIN = 'LOGIN';
const LOGOUT = 'LOGOUT';

describe('Login actions', () => {
    it('should dispatch an action login ', () => {
        const expectedAction = {
            type: 'LOGIN',
            payload: {
                request: {
                    method: 'post',
                    url: '/v1/users/login',
                }
            }
        };
        expect(
            actions.login()
        ).toEqual(expectedAction);
    });
    it('should dispatch an action logout ', () => {
        const expectedAction = {
            type: 'LOGOUT',
            payload: {
                request: {
                    method: 'post',
                    url: '/v1/users/logout',
                }
            }
        };
        expect(
            actions.logout()
        ).toEqual(expectedAction);
    });

});