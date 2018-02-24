import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';
import EnvironmentForm from '../components/EnvironmentForm';
import { createEnvironment } from '../utils/api';

class EnvironmentCreate extends Component {
  constructor(props) {
    super(props);

    this.state = {
      errorText: null,
    };

    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit(env) {
    createEnvironment(env)
      .then(() => this.props.history.push('/'))
      .catch(err => this.setState({ errorText: err.message }));
  }

  render() {
    return (
      <EnvironmentForm
        errorText={this.state.errorText}
        titleText="Create an environment"
        submitText="Create environment"
        onSubmit={this.handleSubmit}
      />)
    ;
  }
}

EnvironmentCreate.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func,
  }).isRequired,
};

export default withRouter(EnvironmentCreate);
