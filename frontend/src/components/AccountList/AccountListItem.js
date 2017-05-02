import React from 'react';
import {Link} from 'react-router';
import PropTypes from 'prop-types';

export default class AccountListItem extends React.Component {
    render() {
        const {account} = this.props;
        return (
            <tr>
                <th><Link to={"/transaction/" + account.id} className='iban'>{account.title}</Link></th>
                <th className="type">{account.type}</th>
                <th className="amount">{account.currency + ' ' + account.amount}</th>
            </tr>
        );
    };
}

AccountListItem.propTypes = {
    account: PropTypes.shape({
        id: PropTypes.number,
        title: PropTypes.string,
        type: PropTypes.string,
        amount: PropTypes.number,
        currency: PropTypes.string
    }).isRequired
};
