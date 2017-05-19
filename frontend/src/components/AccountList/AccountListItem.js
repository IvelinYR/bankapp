import React from 'react';
import {Link} from 'react-router';
import PropTypes from 'prop-types';

export default class AccountListItem extends React.Component {
    render() {
        const {account} = this.props;

        return (
            <tr>
                <th><Link to={"/transaction/" + account.AccountID} className='iban'>{account.title}</Link></th>
                <th className="type">{account.Type}</th>
                <th className="amount">{account.Currency + ' ' + account.Total}</th>
            </tr>
        );
    };
}

AccountListItem.propTypes = {
    account: PropTypes.shape({
        AccountID: PropTypes.string,
        title: PropTypes.string,
        Type: PropTypes.string,
        Total: PropTypes.number,
        Currency: PropTypes.string
    }).isRequired
};
