import { Component } from 'react';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';

import { logOut } from '../utils/auth';

class Logout extends Component {
  componentDidMount() {
    logOut();
    this.props.history.push('/login');
  }

  render() {
    return null;
  }
}

Logout.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func,
  }).isRequired,
};

export default withRouter(Logout);
