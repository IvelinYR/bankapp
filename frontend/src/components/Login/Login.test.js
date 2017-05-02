import React from 'react';
import {shallow} from 'enzyme';
import toJson from 'enzyme-to-json';
import Login from './LoginForm';

describe('Component Login', () => {
    it('renders as expected', () => {
        const component = shallow(<Login/>);
        const tree = toJson(component);
        expect(tree).toMatchSnapshot()
    })
});