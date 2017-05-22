import {connect} from 'react-redux';
import {loadAccounts} from '../action/account';
import {logout} from '../action/login'
import AccountList from '../components/AccountList/AccountList'

const mapStateToProps = (state) => {
    return {
        accounts: state.accounts,
    };
};

const mapDispatchToProps = {
    onSubmitLogout: logout,
    loadAccounts: loadAccounts,
};

export default connect(mapStateToProps, mapDispatchToProps)(AccountList);