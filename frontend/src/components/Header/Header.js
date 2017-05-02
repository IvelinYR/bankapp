import React, {Component} from 'react';
import {Link} from 'react-router';
import './Header.css'
import logo from './ING-logo-v2.jpg'

export default class Header extends Component {
    render() {
        return (
            <div className="header">
                <Link to="/home" className="button-back">back</Link>
                <Link to="/login" className="button-login">login</Link>
                <div><img src={logo} alt="" className="img-logo"/></div>
                <p>{this.props.username}</p>
            </div>
        )
    }
}