import React from 'react'
import { string, node } from 'prop-types'

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
  title: string,
  children: node
}

Card.defaultProps = {
  title: null,
  children: null
}
