import React from 'react';
import {shallow} from 'enzyme';
import toJson from 'enzyme-to-json';
import TransactionsList from './TransactionsList';
import TransactionItems from './TransactionItems';

describe('Component TransactionsList', () => {
    it('renders as expected', () => {
        const transaction = [{'type': 'VISA', 'date': '10-10-1011', "Amount": 123}];
        const component = shallow(<TransactionsList transaction={transaction}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot()
    });
    it('should show text "NO ACCOUNTS" when transactions array is empty', () => {
        const transaction = [{}];
        const component = shallow(<TransactionsList transaction={transaction}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot();
        expect(component.contains("NO TRANSACTIONS")).toBe(true);
    });
    it('renders single transaction', () => {
        const transaction = [{'type': 'VISA', 'date': '10-10-1011', "Amount": 123}];
        const component = shallow(<TransactionsList transaction={transaction}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot();
        expect(component.find(TransactionItems).length).toBe(1);
    });
    it('renders multiple transactions', () => {
        const transaction = [
            {'type': 'VISA', 'date': '10-10-1011', "Amount": 123},
            {'type': 'VISA', 'date': '20-10-1011', "Amount": 244},
            {'type': 'VISA', 'date': '15-10-1011', "Amount": 423}
        ];
        const component = shallow(<TransactionsList transaction={transaction}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot();
        expect(component.find(TransactionItems).length).toBe(3)
    });
    it("should be changed the CSS style of the attribute ", () => {
        const transaction = [{'type': 'VISA', 'date': '10-10-1011', "Amount": 123}];
        const style = {display: "none"};
        const component = shallow(<TransactionsList transaction={transaction}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot();
        expect(component.containsMatchingElement(<div style={style}>NO TRANSACTIONS</div>)).toEqual(true)
    });
});