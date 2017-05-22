import {connect} from 'react-redux';
import {addAccount} from '../action/account';
import {loadAccounts} from '../action/account';
import {logout} from '../action/login'

import NewAccount from '../components/CreateAccount/NewAccount'

const mapStateToProps = (state) => {
    return {
        accounts: state.accounts,
    };
};
const mapDispatchToProps = {
    onSubmitAccount: addAccount,
    onSubmitLogout: logout,
    loadAccounts: loadAccounts
};

export default connect(mapStateToProps, mapDispatchToProps)(NewAccount);