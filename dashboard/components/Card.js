import React from 'react'
import PropTypes from 'prop-types'

import './Card.css'

export default function Card({ title, children }) {
  return (
    <div className='lk-card'>
      <h3 className='lk-card__title'>{title}</h3>
      <div className='lk-card__content'>{children}</div>
    </div>
  )
}

Card.propTypes = {
  title: PropTypes.string,
  children: PropTypes.node
}

Card.defaultProps = {
  title: null,
  children: null
}
