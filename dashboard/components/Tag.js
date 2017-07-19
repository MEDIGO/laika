import React from 'react';
import PropTypes from 'prop-types';

import './Tag.css';

export default function Tag({ children, type }) {
  return <span className={`lk-tag lk-tag--${type}`}>{children}</span>;
}

Tag.propTypes = {
  children: PropTypes.node,
  type: PropTypes.oneOf(['success', 'default']),
};

Tag.defaultProps = {
  children: null,
  type: 'default',
};
