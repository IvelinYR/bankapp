import transactions from './loadTransaction';

describe('LoadTransaction reducer', () => {
    it('should return the initial state', () => {
        expect(
            transactions(undefined, {})
        ).toEqual([])
    });
    it('should handle LOAD_TRANSACTIONS', () => {
        expect(
            transactions([ {'type': 'VISA', 'date': '20-10-1011', "Amount": 244}
            ], {
                type: "LOAD_TRANSACTIONS",
            })
        ).toEqual(
            [
                {'type': 'VISA', 'date': '20-10-1011', "Amount": 244}
            ]
        )
    });
    it('should handle LOAD_TRANSACTIONS', () => {
        expect(
            transactions([
                {'type': 'VISA', 'date': '10-10-1011', "Amount": 123},
                {'type': 'VISA', 'date': '20-10-1011', "Amount": 244},
                {'type': 'VISA', 'date': '15-10-1011', "Amount": 423}
            ], {
                type: "LOAD_TRANSACTIONS",
            })
        ).toEqual(
            [
                {'type': 'VISA', 'date': '10-10-1011', "Amount": 123},
                {'type': 'VISA', 'date': '20-10-1011', "Amount": 244},
                {'type': 'VISA', 'date': '15-10-1011', "Amount": 423}
            ]
        )
    });
});