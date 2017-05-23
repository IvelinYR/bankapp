import * as actions from './account';

const ADD_ACCOUNT = 'ADD_ACCOUNT';
const LOAD_ACCOUNT = 'LOAD_ACCOUNT';

describe('Account actions', () => {
    it('should dispatch an action addAccount ', () => {
        const expectedAction = {
            type: 'ADD_ACCOUNT',
            payload: {
                request: {
                    method: 'post',
                    url: '/v1/users/me/new-account'
                }
            }
        };
        expect(
            actions.addAccount()
        ).toEqual(expectedAction);
    });
    it('should dispatch an action loadAccount ', () => {
        const expectedAction = {
            type: 'LOAD_ACCOUNTS',
            payload: {
                request: {
                    method: 'get',
                    url: '/v1/users/me/accounts'
                }
            }
        };
        expect(
            actions.loadAccounts()
        ).toEqual(expectedAction);
    });

});