import React, {Component} from 'react';
import {shallow} from 'enzyme';
import toJson from 'enzyme-to-json';
import AccountList from './AccountList';
import AccountListItem from './AccountListItem';

describe('Component AccountList', () => {
    it('should show text "NO ACCOUNTS" when accounts array is empty', () => {
        const account = [{}];
        const component = shallow(<AccountList accounts={account}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot();
        expect(component.contains("NO ACCOUNTS")).toBe(true)
    });
    it('renders single account', () => {
        const account= [{'id': '10', 'title': 'ACCOUNT-10', 'type': 'VISA', "amount": 1000, 'currency': 'BGN'}];
        const component = shallow(<AccountList accounts={account}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot();
        expect(component.find(AccountListItem).length).toBe(1)
    });
    it('renders multiple accounts', () => {
        const accounts = [
            {'id': '10', 'title': 'ACCOUNT-10', 'type': 'VISA', "amount": '1000', 'currency': 'BGN'},
            {'id': '11', 'title': 'ACCOUNT-11', 'type': 'VISA', "amount": '105', 'currency': 'BGN'},
            {'id': '12', 'title': 'ACCOUNT-12', 'type': 'VISA', "amount": '102', 'currency': 'BGN'},
            {'id': '13', 'title': 'ACCOUNT-13', 'type': 'VISA', "amount": '101', 'currency': 'BGN'}
        ];
        const component = shallow(<AccountList accounts={accounts}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot();
        expect(component.find(AccountListItem).length).toBe(4)
    });
    it("should be changed the style of the attribute ", () => {
        const account = [{'id': '11', 'title': 'ACCOUNT-11', 'type': 'VISA', "amount": '10', 'currency': 'BGN'}];
        const style = {display: "none"};
        const component = shallow(<AccountList accounts={account}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot();
        expect(component.containsMatchingElement(<div style={style}>NO ACCOUNTS</div>)).toEqual(true)
    });
});