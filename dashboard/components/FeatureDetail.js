import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';

import Toggle from './Toggle';
import Section from './Section';
import './FeatureDetail.css';

import { capitalize } from '../utils/string';

export default function FeatureDetail({ environments, feature, onToggle }) {
  const envStatus = environments.map(env =>
    <div key={env.name}>
      <span className="lk-feature-details__environment-name">
        {capitalize(env.name)} <span>({env.name})</span>
      </span>
      <span className="lk-feature-details__environment-control">
        <Toggle name={env.name} value={feature.status[env.name]} onChange={onToggle} />
      </span>
    </div>,
  );

  return (
    <div className="lk-feature-detail">
      <div className="lk-feature-detail__header">
        <h2>{feature.name}</h2>
        <div>Created {moment(feature.created_at).fromNow()}</div>
      </div>
      <Section title="Environments">
        {envStatus}
      </Section>
    </div>
  );
}

FeatureDetail.propTypes = {
  environments: PropTypes.arrayOf(PropTypes.shape({
    name: PropTypes.string,
  })).isRequired,
  feature: PropTypes.shape({
    name: PropTypes.string,
  }).isRequired,
  onToggle: PropTypes.func.isRequired,
};
