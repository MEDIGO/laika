import React from 'react';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';
import FeatureForm from '../components/FeatureForm';
import { createFeature } from '../utils/api';

function FeatureCreate({ history }) {
  const handleSubmit = ({ name }) => {
    createFeature({ name })
      .then(() => history.push('/'));
  };

  return (
    <FeatureForm
      titleText="Create a Feature"
      submitText="Create Feature"
      onSubmit={handleSubmit}
    />
  );
}

FeatureCreate.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func,
  }).isRequired,
};

export default withRouter(FeatureCreate);
