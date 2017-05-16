import {connect} from 'react-redux';
import {deposit} from '../action/transaction';
import {withdrawal} from '../action/transaction';
import {loadTransactions} from '../action/transaction';
import {loadAmount} from '../action/transaction';

import TransactionForm from '../components/Transaction/TransactionForm'

const mapStateToProps = (state) => {
    return {
        amount: state.amount,
        transaction: state.transaction
    };
};

const mapDispatchToProps = {
    deposit: deposit,
    withdrawal: withdrawal,
    onSubmitTransaction: loadTransactions,
    onSubmitAmount: loadAmount

};

export default connect(mapStateToProps, mapDispatchToProps)(TransactionForm);