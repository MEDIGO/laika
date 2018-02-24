import React from 'react'
import { string, bool, func } from 'prop-types'

import './Toggle.css'

export default function Toggle({ name, value, onChange }) {
  const handleChange = e => {
    onChange(name, e.target.checked)
  }

  return (
    <label htmlFor={name} className='lk-toggle'>
      <input
        id={name}
        name={name}
        type='checkbox'
        checked={value}
        onChange={handleChange}
      />
      <span />
    </label>
  )
}

Toggle.propTypes = {
  name: string.isRequired,
  value: bool.isRequired,
  onChange: func.isRequired
}
