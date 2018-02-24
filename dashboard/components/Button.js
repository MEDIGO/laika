import React from 'react'
import PropTypes from 'prop-types'

import './Button.css'

export default function Button({ label, type, onClick }) {
  const state = onClick ? type : 'disabled'

  return (
    <button className={`lk-button lk-button--${state}`} onClick={onClick}>
      {label}
    </button>
  )
}

Button.propTypes = {
  label: PropTypes.string.isRequired,
  type: PropTypes.oneOf(['primary', 'default']),
  onClick: PropTypes.func
}

Button.defaultProps = {
  type: 'default',
  onClick: null
}
