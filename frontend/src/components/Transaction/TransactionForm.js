import React, {Component} from 'react';
import './TransactionForm.css';
import TransactionsList from '../TransactionList/TransactionsList';

export default class TransactionForm extends Component {
    constructor(props) {
        super(props);
        this.state = {
            account: [],
            transaction: [],
            value: '',
            amount: 0
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleAddDeposit = this.handleAddDeposit.bind(this);
        this.handleAddWithdrawals = this.handleAddWithdrawals.bind(this);
    }

    handleAddDeposit() {
        this.props.deposit({
            id: Number(this.props.params.id),
            operation: Number(this.state.value),
            type: "Deposit",
            date: ""
        });
        this.props.onSubmitTransaction({
            id: this.props.params.id,
            operation: this.state.value
        });
        this.setState({value: ''});
    }

    handleAddWithdrawals() {
        this.props.withdrawal({
            id: Number(this.props.params.id),
            operation: Number(this.state.value),
            type: "Withdrawal",
            date: ""
        });
        this.props.onSubmitTransaction({
            id: this.props.params.id,
            operation: this.state.value
        });
        this.setState({value: ''})
    }

    componentDidMount() {
        this.props.onSubmitTransaction({
            id: this.props.params.id,
            operation: this.state.value
        });
        this.props.onSubmitAmount({
            id: Number(this.props.params.id)
        })
    }

    handleChange(event) {
        this.setState({value: event.target.value});
    }

    render() {
        let account = this.props.amount;
        return (
            <div>
                <br/>
                <h1>{account.title}</h1>
                <h2>{account.currency} {account.amount}</h2>
                <input id="transaction" type="text" value={this.state.value} onChange={this.handleChange}/><br/>
                <button className="deposits" onClick={this.handleAddDeposit}>Deposit</button>
                <button className="withdrawals" onClick={this.handleAddWithdrawals}>Withdrawals</button>
                <TransactionsList transaction={this.props.transaction}/>
            </div>
        )
    }
}
