import React from 'react'
import { bool, func, string } from 'prop-types'

import './Input.css'

const Input = ({
  label,
  name,
  type,
  required,
  value,
  error,
  onChange,
  placeholder,
  autoFocus
}) =>
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

Input.propTypes = {
  label: string.isRequired,
  name: string.isRequired,
  value: string.isRequired,
  required: bool,
  error: string,
  onChange: func.isRequired,
  type: string,
  placeholder: string,
  autoFocus: bool
}

Input.defaultProps = {
  required: false,
  error: '',
  value: '',
  type: '',
  placeholder: '',
  autoFocus: false
}

export default Input
