import {connect} from 'react-redux';
import {loadAccounts} from '../action/account';
import AccountList from '../components/AccountList/AccountList'

const mapStateToProps = (state) => {
    return {
        accounts: state.accounts,
    };
};

const mapDispatchToProps = {
    loadAccounts: loadAccounts
};

export default connect(mapStateToProps, mapDispatchToProps)(AccountList);