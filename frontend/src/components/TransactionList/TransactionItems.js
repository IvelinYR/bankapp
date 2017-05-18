import React from 'react';
import './Transaction.css';
import PropTypes from 'prop-types';

export default class TransactionItems extends React.Component {
    render() {
        let prefix = '';
        let style = {color: 'black'};
        let transaction = this.props.transaction;

        if (this.props.transaction.type === "Withdrawal") {
            prefix = "-";
        }
        return (
            <tr>
                <th >{transaction.type}</th>
                <th >{transaction.date}</th>
                <th style={style}>{prefix} {transaction.Amount}</th>
            </tr>
        );
    }
}

TransactionItems.propTypes = {
    transaction:PropTypes.shape({
        type: PropTypes.string,
        date: PropTypes.string,
        Amount: PropTypes.number
    }).isRequired

};