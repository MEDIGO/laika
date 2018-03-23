import React from 'react'
import { string, oneOf, func } from 'prop-types'

import './Button.css'

const Button = ({ label, type, onClick }) => {
  const state = onClick ? type : 'hidden'

  return (
    <button className={`lk-button lk-button--${state}`} onClick={onClick}>
      {label}
    </button>
  )
}

Button.propTypes = {
  label: string.isRequired,
  type: oneOf(['primary', 'default']),
  onClick: func
}

Button.defaultProps = {
  type: 'default',
  onClick: null
}

export default Button
