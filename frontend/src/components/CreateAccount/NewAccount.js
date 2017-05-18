import React, {Component} from 'react';
import {Link} from 'react-router';
import './NewAccount.css';

export default class NewAccount extends Component {
    constructor() {
        super();
        this.state = {
            type: 'VISA',
            currency: 'GBP'
        };

        this.handleTypeChange = this.handleTypeChange.bind(this);
        this.handleCurrencyChange = this.handleCurrencyChange.bind(this);
        this.handleAddAccount = this.handleAddAccount.bind(this);
    }

    handleTypeChange(e) {
        this.setState({type: e.target.value})
    }

    handleCurrencyChange(e) {
        this.setState({currency: e.target.value})
    }

    handleAddAccount() {
        this.props.onSubmitAccount({
            Type: this.state.type,
            Total: +0,
            Currency: this.state.currency
        })
    }

    render() {
        let error = this.props.accounts.error;
        return (
            <div >
                <Link to="/home" className="button-back">Back</Link>
                <Link to="/login" className="button-logout">Logout</Link>
                <div className="table">
                    <h1>New Account</h1>
                    <table>
                        <thead>
                        <tr>
                            <td>Account Type</td>
                            <td>Currency</td>
                        </tr>
                        </thead>
                        <tbody>
                        <tr>
                            <th>
                                <div>
                                    <select id="lang" onChange={this.handleTypeChange}>
                                        <option value="VISA">VISA</option>
                                        <option value="Mastercard">Mastercard</option>
                                        <option value="Maestro">Maestro</option>
                                    </select>
                                </div>
                            </th>
                            <th>
                                <div>
                                    <select id="lang2" onChange={this.handleCurrencyChange}>
                                        <option value="GBP">GBP</option>
                                        <option value="EUR">EUR</option>
                                        <option value="BGN">BGN</option>
                                        <option value="USD">USD</option>
                                        <option value="RUB">RUB</option>
                                    </select>
                                </div>
                            </th>
                        </tr>
                        </tbody>
                    </table>
                    <Link to={"/home"}>
                        <button className="btn" onClick={this.handleAddAccount}>Create Accounts</button>
                    </Link>
                    <h1>{error}</h1>
                </div>
            </div>
        )
    }
}
