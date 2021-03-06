import React, {Component} from 'react'
import AccountListItem from './AccountListItem'
import {Link} from 'react-router'
import './AccountsList.css'

export default class Accounts extends Component {
    constructor() {
        super();
        this.handleLogout = this.handleLogout.bind(this);
    }

    handleLogout() {
        this.props.onSubmitLogout({})
    }

    componentDidMount() {
        this.props.loadAccounts({})
    }

    render() {
        const {accounts} = this.props;

        let style = {display: "inline-block"};
        if (accounts.length > 0) {
            style = {display: "none"};
        }

        return (
            <div >
                <Link to="/login">
                    <button className="button-logout" onClick={this.handleLogout}>Logout</button>
                </Link>
                <br/>
                <div className="table">
                    <Link to="/account" className="active">NewAccount</Link>
                    <table>
                        <thead className="table-head">
                        <tr>
                            <td>Account Number</td>
                            <td>Account Type</td>
                            <td>Balance</td>
                        </tr>
                        </thead>
                        <tbody>
                        {accounts.map((account, i) => {
                                return <AccountListItem key={i} account={account}/>;
                            }
                        )}
                        </tbody>
                    </table>
                    <div className="table-accounts" style={style}>NO ACCOUNTS</div>
                </div>
            </div>
        );
    }
};