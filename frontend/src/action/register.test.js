import * as actions from './register';

const REGISTER = 'REGISTER';

describe('Register actions', () => {
    it('should dispatch an action register ', () => {
        const expectedAction = {
            type: 'REGISTER',
            payload: {
                request: {
                    method: 'post',
                    url: '/v1/users/signup',
                }
            }
        };
        expect(
            actions.register()
        ).toEqual(expectedAction);
    });
});