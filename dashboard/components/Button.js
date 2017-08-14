import React from 'react';
import PropTypes from 'prop-types';

import './Button.css';

export default function Button({ label, type, onClick }) {
  return (
    <button
      className={`lk-button lk-button--${type}`}
      onClick={onClick}
    >
      {label}
    </button>
  );
}

Button.propTypes = {
  label: PropTypes.string.isRequired,
  type: PropTypes.oneOf(['primary', 'default']),
  onClick: PropTypes.func.isRequired,
};

Button.defaultProps = {
  type: 'default',
};
