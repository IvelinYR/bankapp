import React, {Component} from 'react';
import './Login.css'
import {Link} from 'react-router';
import AlertContainer from 'react-alert';
import logo from './cros.png'

export default class LoginForm extends Component {
    constructor(props) {
        super(props);
        this.state = {username: '', password: ''};
        this.alertOptions = {
            offset: 14,
            position: 'bottom right',
            theme: 'dark',
            time: 5000,
            transition: 'scale'
        };
        this.bindEventHandlers();
    }

    componentDidMount() {
        this.msg.show('Username or Password was WRONG!!', {
            time: 4000,
            type: 'error',
            icon: <img src={logo} alt="" height="35"/>
        });
    }

    bindEventHandlers() {
        this.onChangeHandler = this.onChangeHandler.bind(this);
        this.onSubmitHandler = this.onSubmitHandler.bind(this);
    }

    onChangeHandler(event) {
        switch (event.target.name) {
            case 'username':
                this.setState({username: event.target.value});
                break;
            case 'password':
                this.setState({password: event.target.value});
                break;
            default:
                break;
        }
    }

    onSubmitHandler() {
        this.props.onSubmitLogin({
            Username: this.state.username,
            Password: this.state.password
        })
    }

    render() {
        return (
            <div>
                <Link to="/register" className="button-login">Sign in</Link>
                <br/>
                <br/>
                <h1>Login Page</h1>
                <div className="form-group">
                    <label>Username</label>
                    <br/>
                    <input
                        className="form-control"
                        type="text"
                        name="username"
                        value={this.props.username}
                        onChange={this.onChangeHandler}
                    />
                </div>
                <div className="form-group">
                    <label>Password</label>
                    <br/>
                    <input
                        className="form-control"
                        type="password"
                        name="password"
                        value={this.props.password}
                        onChange={this.onChangeHandler}
                    />
                </div>
                <Link to={"/home"}>
                    <button className="login-center" onClick={this.onSubmitHandler}>Login</button>
                </Link>
                <AlertContainer ref={a => this.msg = a} {...this.alertOptions} />
            </div>
        );
    }
}