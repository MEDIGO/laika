import React from 'react';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';
import EnvironmentForm from '../components/EnvironmentForm';
import { createEnvironment } from '../utils/api';

function EnvironmentCreate({ history }) {
  const handleSubmit = ({ name }) => {
    createEnvironment({ name })
      .then(() => history.push('/'));
  };

  return (
    <EnvironmentForm
      titleText="Create an environment"
      submitText="Create environment"
      onSubmit={handleSubmit}
    />)
  ;
}

EnvironmentCreate.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func,
  }).isRequired,
};

export default withRouter(EnvironmentCreate);
