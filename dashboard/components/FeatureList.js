import React from 'react'
import { arrayOf, shape, string } from 'prop-types'
import { Link } from 'react-router-dom'
import moment from 'moment'

import Tag from './Tag'
import Section from './Section'

import './FeatureList.css'

const sort = (features) => features.sort((a, b) => {
  if (a.created_at < b.created_at) return 1
  if (a.created_at > b.created_at) return -1
  return 0
})

const parseStatus = (environments, status) =>
  environments.map(env => ({
    name: env.name,
    enabled: status[env.name]
  }))

const FeatureList = ({ environments, features }) => {
  const items = sort(features).map(feature => (
    <div key={feature.name} className='lk-feature-list__item'>
      <Link to={`/features/${window.encodeURIComponent(feature.name)}`}>
        <div className='lk-feature-list__name'>
          <span>{feature.name}</span>
          <span className='lk-feature-list__status-list'>
            {parseStatus(environments, feature.status).map(status => (
              <Tag key={status.name} type={status.enabled ? 'success' : null}>
                {status.name}
              </Tag>
            ))}
          </span>
        </div>
        <div className='lk-feature-list__time'>
          Created {moment(feature.created_at).fromNow()}
        </div>
      </Link>
    </div>
  ))

  return (
    <div className='lk-feature-list'>
      <Section title={`Features (${items.length})`}>{items}</Section>
    </div>
  )
}

FeatureList.propTypes = {
  features: arrayOf(
    shape({
      name: string
    })
  ).isRequired,
  environments: arrayOf(
    shape({
      name: string
    })
  ).isRequired
}

export default FeatureList
