import React from 'react';
import {shallow} from 'enzyme';
import toJson from 'enzyme-to-json';
import AccountListItem from './AccountListItem';

describe('Component AccountListItem', () => {
    it('renders as expected', () => {
        const account = [{'id':'10', 'title':'ACCOUNT-10', 'type':'VISA', "amount":'1000', 'currency':'BGN'}];
        const component = shallow(<AccountListItem account={account}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot()
    });
    it('renders single account', () => {
        const account = [{'id': '10', 'title': 'ACCOUNT-10', 'type': 'VISA', "amount": '1000', 'currency': 'BGN'}];
        const component = shallow(<AccountListItem account={account}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot();
    });
    it('renders multiple accounts', () => {
        const account = [
            {'id': '10', 'title': 'ACCOUNT-10', 'type': 'VISA', "amount": '1000', 'currency': 'BGN'},
            {'id': '11', 'title': 'ACCOUNT-11', 'type': 'VISA', "amount": '105', 'currency': 'BGN'},
            {'id': '12', 'title': 'ACCOUNT-12', 'type': 'VISA', "amount": '102', 'currency': 'BGN'},
            {'id': '13', 'title': 'ACCOUNT-13', 'type': 'VISA', "amount": '101', 'currency': 'BGN'}
        ];
        const component = shallow(<AccountListItem account={account}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot();
    });
});