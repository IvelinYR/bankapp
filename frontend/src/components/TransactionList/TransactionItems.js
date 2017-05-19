import React from 'react';
import './Transaction.css';
import PropTypes from 'prop-types';

export default class TransactionItems extends React.Component {
    render() {
        let prefix = '';
        let style = {color: 'black'};
        let transaction = this.props.transaction;

        if (transaction.Type === "Withdrawal") {
            prefix = "-";
        }
        return (
            <tr>
                <th >{transaction.Type}</th>
                <th >{transaction.Date}</th>
                <th style={style}>{prefix} {transaction.Amount}</th>
            </tr>
        );
    }
}

TransactionItems.propTypes = {
    transaction:PropTypes.shape({
        Type: PropTypes.string,
        Date: PropTypes.string,
        Amount: PropTypes.number
    }).isRequired

};