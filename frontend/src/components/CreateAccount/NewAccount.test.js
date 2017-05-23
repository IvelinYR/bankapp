import React from 'react';
import {shallow} from 'enzyme';
import toJson from 'enzyme-to-json';
import NewAccount from './NewAccount';

describe('Component NewAccountu', () => {
    it('renders as expected', () => {
        const accounts = {error:"Wrong"}
        const component = shallow(<NewAccount accounts={accounts}/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot()
    })
});