import React from 'react'
import { oneOf, node } from 'prop-types'

import './Tag.css'

const Tag = ({ children, type }) =>
  <span className={`lk-tag lk-tag--${type}`}>{children}</span>

Tag.propTypes = {
  children: node,
  type: oneOf(['success', 'default'])
}

Tag.defaultProps = {
  children: null,
  type: 'default'
}

export default Tag
