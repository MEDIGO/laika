import React from 'react';
import PropType from 'prop-types';

import './Toggle.css';

export default function Toggle({ name, value, onChange }) {
  const handleChange = (e) => {
    onChange(name, e.target.checked);
  };

  return (
    <label htmlFor={name} className="lk-toggle">
      <input
        id={name}
        name={name}
        type="checkbox"
        checked={value}
        onChange={handleChange}
      />
      <span />
    </label>
  );
}

Toggle.propTypes = {
  name: PropType.string.isRequired,
  value: PropType.bool.isRequired,
  onChange: PropType.func.isRequired,
};
