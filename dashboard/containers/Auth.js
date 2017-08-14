import { Component } from 'react';
import { withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';
import { isLoggedIn } from '../utils/auth';

class Auth extends Component {
  componentDidMount() {
    if (!isLoggedIn()) this.props.history.push('/login');
  }

  render() {
    return isLoggedIn() ? this.props.children : null;
  }
}

Auth.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func,
  }).isRequired,
  children: PropTypes.node,
};

Auth.defaultProps = {
  children: null,
};

export default withRouter(Auth);
