import React, {Component} from 'react';
import {Link} from 'react-router';
import './Register.css';

export default class LoginForm extends Component {
    constructor(props) {
        super(props);
        this.state = {username: '', password: '',name:'', email:'', age:''};
        this.bindEventHandlers();
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
            case 'name':
                this.setState({name: event.target.value});
                break;
            case 'email':
                this.setState({email: event.target.value});
                break;
            case 'age':
                this.setState({age: event.target.value});
                break;
            default:
                break;
        }
    }

    onSubmitHandler() {
        this.props.onSubmitRegister({
            Username: this.state.username,
            Password: this.state.password,
            Name: this.state.name,
            Email: this.state.email,
            Age: +this.state.age
        })
    }

    render() {
        let error = this.props.register.error;
        return (
            <div>
                <Link to="/login" className="button-login">Login</Link>
                <br/>
                <br/>
                <h1>Register Page</h1>
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
                <div className="form-group">
                    <label>Name</label>
                    <br/>
                    <input
                        className="form-control"
                        type="text"
                        name="name"
                        value={this.props.email}
                        onChange={this.onChangeHandler}
                    />
                </div>

                <div className="form-group">
                    <label>Email</label>
                    <br/>
                    <input
                        className="form-control"
                        type="text"
                        name="email"
                        value={this.props.password}
                        onChange={this.onChangeHandler}
                    />
                </div>
                <div className="form-group">
                    <label>Age</label>
                    <br/>
                    <input
                        className="form-control"
                        type="number"
                        name="age"
                        value={this.props.password}
                        onChange={this.onChangeHandler}
                    />
                </div>
                <Link to={"/login"}>
                    <button className="login-center" onClick={this.onSubmitHandler}>Register</button>
                </Link>
                <h2 className="error-info">{error}</h2>
            </div>
        );
    }
}