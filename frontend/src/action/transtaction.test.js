import * as actions from './transaction';

const DEPOSIT = 'DEPOSIT';
const WITHDRAWAL = 'WITHDRAWAL';
const LOAD_TRANSACTIONS = 'LOAD_TRANSACTIONS';


describe('Transaction actions', () => {
    it('should dispatch an action deposit ', () => {
        const expectedAction = {
            type: 'DEPOSIT',
            payload: {
                request: {
                    method: 'post',
                    url: '/v1/users/me/deposit'
                }
            }
        };
        expect(
            actions.deposit()
        ).toEqual(expectedAction);
    });
    it('should dispatch an action withdrawal ', () => {
        const expectedAction = {
            type: 'WITHDRAWAL',
            payload: {
                request: {
                    method: 'post',
                    url: '/v1/users/me/withdraw'
                }
            }
        };
        expect(
            actions.withdrawal()
        ).toEqual(expectedAction);
    });
    it('should dispatch an action loadTransactions ', () => {
        const expectedAction = {
            type: 'LOAD_TRANSACTIONS',
            payload: {
                request: {
                    method:'post',
                    url: '/v1/users/me/account-history'
                }
            }
        };
        expect(
            actions.loadTransactions()
        ).toEqual(expectedAction);
    });
});