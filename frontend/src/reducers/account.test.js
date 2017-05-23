import accounts from './account';

describe('Account reducer', () => {
    it('should return the initial state', () => {
        expect(
            accounts(undefined, {})
        ).toEqual([])
    });
    it('should handle LOAD_ACCOUNTS', () => {
        expect(
            accounts([{'id': '10', 'title': 'ACCOUNT-10', 'type': 'VISA', "amount": '1000', 'currency': 'BGN'},
            ], {
                type: "LOAD_ACCOUNTS",
            })
        ).toEqual(
            [
                {'id': '10', 'title': 'ACCOUNT-10', 'type': 'VISA', "amount": '1000', 'currency': 'BGN'},

            ]
        )
    });
    it('should handle LOAD_ACCOUNTS', () => {
        expect(
            accounts([
                {'id': '10', 'title': 'ACCOUNT-10', 'type': 'VISA', "amount": '1000', 'currency': 'BGN'},
                {'id': '10', 'title': 'ACCOUNT-10', 'type': 'VISA', "amount": '1000', 'currency': 'BGN'},
                {'id': '10', 'title': 'ACCOUNT-10', 'type': 'VISA', "amount": '1000', 'currency': 'BGN'},
            ], {
                type: "LOAD_ACCOUNTS",
            })
        ).toEqual(
            [
                {'id': '10', 'title': 'ACCOUNT-10', 'type': 'VISA', "amount": '1000', 'currency': 'BGN'},
                {'id': '10', 'title': 'ACCOUNT-10', 'type': 'VISA', "amount": '1000', 'currency': 'BGN'},
                {'id': '10', 'title': 'ACCOUNT-10', 'type': 'VISA', "amount": '1000', 'currency': 'BGN'},
            ]
        )
    });
});