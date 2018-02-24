import React from 'react'
import { oneOf, node } from 'prop-types'

import './Tag.css'

export default function Tag({ children, type }) {
  return <span className={`lk-tag lk-tag--${type}`}>{children}</span>
}

Tag.propTypes = {
  children: node,
  type: oneOf(['success', 'default'])
}

Tag.defaultProps = {
  children: null,
  type: 'default'
}
