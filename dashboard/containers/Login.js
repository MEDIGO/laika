import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';

import LoginForm from '../components/LoginForm';
import { logIn, logOut } from '../utils/auth';

class Login extends Component {
  constructor(props) {
    super(props);

    this.state = {
      error: null,
    };

    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit({ username, password }) {
    this.setState({ error: null });

    logIn(username, password)
      .then(() => this.props.history.push('/'))
      .catch(() => {
        logOut();
        this.setState({ error: 'Please enter the correct username and password.' });
      });
  }

  render() {
    return <LoginForm onSubmit={this.handleSubmit} errorText={this.state.error} />;
  }
}

Login.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func,
  }).isRequired,
};

export default withRouter(Login);
