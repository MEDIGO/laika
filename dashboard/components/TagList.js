import React from 'react'
import { arrayOf, shape, string, bool } from 'prop-types'

import Tag from './Tag'
import './TagList.css'

const TagList = ({ tags }) =>
  <span className='lk-tag-list__status-list'>
    {tags.map(tag => (
      <Tag key={tag.name} type={tag.enabled ? 'success' : null}>
        {tag.name}
      </Tag>
    ))}
  </span>

TagList.propTypes = {
  tags: arrayOf(
    shape({
      name: string,
      enabled: bool
    })
  )
}

export default TagList
