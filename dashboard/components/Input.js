import React from 'react'
import PropTypes from 'prop-types'

import './Input.css'

export default function Input({
  label,
  name,
  type,
  required,
  value,
  error,
  onChange,
  placeholder,
  autoFocus
}) {
  return (
    <div className='lk-input'>
      <label className='lk-input__label' htmlFor={name}>
        {label}
        {required ? '*' : null}
      </label>
      <input
        id={name}
        className='lk-input__input'
        value={value}
        onChange={e => onChange(name, e.target.value)}
        required={required}
        type={type}
        placeholder={placeholder}
        autoFocus={autoFocus}
      />
      {error ? <div>{error}</div> : null}
    </div>
  )
}

Input.propTypes = {
  label: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  value: PropTypes.string.isRequired,
  required: PropTypes.bool,
  error: PropTypes.string,
  onChange: PropTypes.func.isRequired,
  type: PropTypes.string,
  placeholder: PropTypes.string,
  autoFocus: PropTypes.bool
}

Input.defaultProps = {
  required: false,
  error: '',
  value: '',
  type: '',
  placeholder: '',
  autoFocus: false
}
