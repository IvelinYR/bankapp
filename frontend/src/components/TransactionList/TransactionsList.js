import React, {Component} from 'react';
import TransactionItems from './TransactionItems'
import './Transaction.css'

export default class TransactionsList extends Component {
    render() {
        const {transaction} = this.props;
        let style = {display: "inline-block"};
        if (transaction.length > 0) {
            style = {display: "none"};
        }
        return (
            <div className="container">
                <div className="table">
                    <br/>
                    <table>
                        <thead>
                        <tr>
                            <td>Type</td>
                            <td>Date</td>
                            <td>Transaction</td>
                        </tr>
                        </thead>
                        <tbody>
                        {
                            transaction.map((transaction, i) => {
                                    return <TransactionItems key={i} transaction={transaction}/>;
                                }
                            )
                        }
                        </tbody>
                    </table>
                    <div className="table-block" style={style}>NO TRANSACTIONS</div>
                </div>
            </div>
        )
    }
}