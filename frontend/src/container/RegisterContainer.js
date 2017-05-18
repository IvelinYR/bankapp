import {connect} from 'react-redux';
import {register} from '../action/register'
import Register from '../components/Register/Register';

const mapStateToProps = (state) => {
    return {
        register: state.register
    };
};
const mapDispatchToProps = {
    onSubmitRegister: register
};

export default connect(mapStateToProps, mapDispatchToProps)(Register)