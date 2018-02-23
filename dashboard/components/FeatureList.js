import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import moment from 'moment';

import Tag from './Tag';
import Section from './Section';

import './FeatureList.css';

function sort(features) {
  return features.sort((a, b) => {
    if (a.created_at < b.created_at) return 1;
    if (a.created_at > b.created_at) return -1;
    return 0;
  });
}

function parseStatus(status) {
  return Object.keys(status).map(key => ({ name: key, enabled: status[key] }));
}

export default function FeatureList({ features }) {
  const items = sort(features).map(feature =>
    <div key={feature.name} className="lk-feature-list__item">
      <Link to={`/features/${window.encodeURIComponent(feature.name)}`}>
        <div className="lk-feature-list__name">
          <span>{feature.name}</span>
          <span className="lk-feature-list__status-list">
            {parseStatus(feature.status).map(status => <Tag key={status.name} type={status.enabled ? 'success' : null} >{status.name}</Tag>)}
          </span>
        </div>
        <div className="lk-feature-list__time">Created {moment(feature.created_at).fromNow()}</div>
      </Link>
    </div>,
  );

  return (
    <div className="lk-feature-list">
      <Section title="Features">
        {items}
      </Section>
    </div>
  );
}

FeatureList.propTypes = {
  features: PropTypes.arrayOf(PropTypes.shape({
    name: PropTypes.string,
  })).isRequired,
};
