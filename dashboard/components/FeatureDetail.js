import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';

import Toggle from './Toggle';
import Section from './Section';
import './FeatureDetail.css';

import { capitalize } from '../utils/string';

export default function FeatureDetail({ feature, onToggle }) {
  const environments = Object.keys(feature.status).map(name =>
    <div key={name}>
      <span className="lk-feature-details__environment-name">
        {capitalize(name)} <span>({name})</span>
      </span>
      <span className="lk-feature-details__environment-control">
        <Toggle name={name} value={feature.status[name]} onChange={onToggle} />
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
        {environments}
      </Section>
    </div>
  );
}

FeatureDetail.propTypes = {
  feature: PropTypes.shape({
    name: PropTypes.string,
  }).isRequired,
  onToggle: PropTypes.func.isRequired,
};
