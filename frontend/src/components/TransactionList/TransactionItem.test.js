import React from 'react';
import {shallow} from 'enzyme';
import toJson from 'enzyme-to-json';
import TransactionItems from './TransactionItems';

describe('Component TransactionItems', () => {
    it('renders as expected', () => {
        const transaction = [{'type':'VISA', 'date':'10-10-1011',  'operation':'123'}];
        const component = shallow(<TransactionItems transaction={transaction}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot()
    })
});