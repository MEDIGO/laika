import React from 'react';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';
import EnvironmentForm from '../components/EnvironmentForm';
import { createEnvironment } from '../utils/api';
import { withAuth } from '../utils/auth';

function EnvironmentCreate({ history }) {
  const handleSubmit = ({ name }) => {
    createEnvironment({ name })
      .then(() => history.push('/'));
  };

  return (
    <EnvironmentForm
      titleText="Create an Environment"
      submitText="Create Environment"
      onSubmit={handleSubmit}
    />)
  ;
}

EnvironmentCreate.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func,
  }).isRequired,
};

export default withAuth(withRouter(EnvironmentCreate));
