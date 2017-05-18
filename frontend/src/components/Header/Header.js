import React, {Component} from 'react';
import './Header.css';
import logo from './ING-logo-v2.jpg';

export default class Header extends Component {
    render() {
        return (
            <div className="header">
                <div><img src={logo} alt="" className="img-logo"/></div>
            </div>
        )
    }
}