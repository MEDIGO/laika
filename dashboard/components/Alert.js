import React from 'react'
import { oneOf, node } from 'prop-types'

import './Alert.css'

const Alert = ({ type, children }) => {
  return <div className={`lk-alert lk-alert--${type}`}>{children}</div>
}

Alert.propTypes = {
  type: oneOf(['danger']),
  children: node
}

Alert.defaultProps = {
  type: 'danger',
  children: null
}

export default Alert
