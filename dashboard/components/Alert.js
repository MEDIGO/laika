import React from 'react';
import PropTypes from 'prop-types';

import './Alert.css';

export default function Alert({ type, children }) {
  return <div className={`lk-alert lk-alert--${type}`}>{children}</div>;
}

Alert.propTypes = {
  type: PropTypes.oneOf(['danger']),
  children: PropTypes.node,
};

Alert.defaultProps = {
  type: 'danger',
  children: null,
};
