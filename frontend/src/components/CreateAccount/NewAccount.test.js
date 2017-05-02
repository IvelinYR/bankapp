import React from 'react';
import {shallow} from 'enzyme';
import toJson from 'enzyme-to-json';
import NewAccount from './NewAccount';

describe('Component NewAccountu', () => {
    it('renders as expected', () => {
        const component = shallow(<NewAccount/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot()
    })
});