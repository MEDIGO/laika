import React from 'react'
import PropTypes from 'prop-types'

import './Container.css'

export default function Container({ children }) {
  return <div className='lk-container'>{children}</div>
}

Container.propTypes = {
  children: PropTypes.node
}

Container.defaultProps = {
  children: null
}
