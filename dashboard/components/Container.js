import React from 'react'
import { node } from 'prop-types'

import './Container.css'

export default function Container({ children }) {
  return <div className='lk-container'>{children}</div>
}

Container.propTypes = {
  children: node
}

Container.defaultProps = {
  children: null
}
