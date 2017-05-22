import React, {Component} from 'react';
import './TransactionForm.css';
import TransactionsList from '../TransactionList/TransactionsList';
import {Link} from 'react-router';

export default class TransactionForm extends Component {
    constructor(props) {
        super(props);
        this.state = {
            account: [],
            transaction: [],
            value: '',
            validate:''
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleAddDeposit = this.handleAddDeposit.bind(this);
        this.handleAddWithdrawals = this.handleAddWithdrawals.bind(this);
    }

    handleAddDeposit() {
        this.props.deposit({
            Amount: Number(this.state.value),
            Currency: this.props.amount.currency,
            AccountID: this.props.params.id,
            Type: "Deposit",
            date: ""
        });
        this.props.onSubmitTransaction({
            AccountID: this.props.params.id,
        });
        this.setState({value: ''});
    }

    handleAddWithdrawals() {
        this.props.withdrawal({
            AccountID:this.props.params.id,
            Amount: Number(this.state.value),
            Currency: this.props.amount.currency,
            Type: "Withdrawal",
            date: ""
        });
        this.props.onSubmitTransaction({
            AccountID: this.props.params.id,
        });
        this.setState({value: ''})
    }

    componentDidMount() {
        this.props.onSubmitTransaction({
            AccountID: this.props.params.id,
        });
    }

    handleChange(event) {
        if(event.target.value <= 0){
            this.setState({validate:"Must be positive number"})
        } else {
            this.setState({value: event.target.value});
            this.setState({validate:""})
        }
    }

    render() {
        let AccountID = this.props.params.id;
        let accountList = this.props.accounts;
        let account = accountList.filter(function (obj) {
            return obj.AccountID === AccountID;
        });
        let currencyAccount = account[0];
        let error = this.props.amount.error;
        return (
            <div>
                <Link to="/home" className="button-back">Back</Link>
                <br/>
                <br/>
                <h2>{currencyAccount.Currency} {currencyAccount.Amount}</h2>
                <input id="transaction" type="number" value={this.state.value} onChange={this.handleChange}/><br/>
                <button className="deposits" onClick={this.handleAddDeposit}>Deposit</button>
                <button className="withdrawals" onClick={this.handleAddWithdrawals}>Withdrawals</button>
                <TransactionsList transaction={this.props.transaction}/>
                <h2 className="error-info">{error}{this.state.validate}</h2>
            </div>
        )
    }
}
