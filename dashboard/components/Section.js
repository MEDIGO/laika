import React from 'react'
import PropTypes from 'prop-types'

import './Section.css'

export default function Section({ title, children }) {
  const items = children.map((c, i) => (
    <div key={i.toString()} className='lk-section__item'>
      {c}
    </div>
  ))

  return (
    <div className='lk-section'>
      <h3 className='lk-section__title'>{title}</h3>
      <div className='lk-section__content'>{items}</div>
    </div>
  )
}

Section.propTypes = {
  title: PropTypes.string.isRequired,
  children: PropTypes.node.isRequired
}
