import React from 'react';
import {shallow} from 'enzyme';
import toJson from 'enzyme-to-json';
import TransactionForm from './TransactionForm';

describe('Component TransactionForm', () => {
    it('renders as expected', () => {
        const amount = [{title:'ACCOUNT'}]
        const component = shallow(<TransactionForm amount={amount}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot()
    })
});