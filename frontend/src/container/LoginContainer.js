import {connect} from 'react-redux';
import {login} from '../action/login';
import LoginForm from '../components/Login/LoginForm';

const mapStateToProps = (state) => {
    return {
        login: state.login,
    };
};

const mapDispatchToProps = {
    onSubmitLogin: login
};


export default connect(mapStateToProps, mapDispatchToProps)(LoginForm);