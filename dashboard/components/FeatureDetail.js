import React, { Component } from 'react'
import PropTypes from 'prop-types'
import moment from 'moment'

import Button from './Button'
import Toggle from './Toggle'
import Section from './Section'
import './FeatureDetail.css'

import { capitalize } from '../utils/string'

export default class FeatureDetail extends Component {
  constructor(props) {
    super(props)

    this.state = {
      deleteUnlocked: false
    }

    this.lockDelete = this.lockDelete.bind(this)
    this.unlockDelete = this.unlockDelete.bind(this)
  }

  lockDelete() {
    this.setState({ deleteUnlocked: false })
  }

  unlockDelete() {
    this.setState({ deleteUnlocked: true })
  }

  render() {
    const { environments, feature, onToggle, onDelete } = this.props
    const { deleteUnlocked } = this.state

    const envStatus = environments.map(env => (
      <div key={env.name}>
        <span className='lk-feature-details__environment-name'>
          {capitalize(env.name)} <span>({env.name})</span>
        </span>
        <span className='lk-feature-details__environment-control'>
          <Toggle
            name={env.name}
            value={feature.status[env.name]}
            onChange={onToggle}
          />
        </span>
      </div>
    ))

    const cancel = deleteUnlocked ? (
      <Button onClick={this.lockDelete} label='Cancel' />
    ) : null
    const del = deleteUnlocked ? (
      <Button
        onClick={() => onDelete(feature.name)}
        label='Confirm deletion'
        type='primary'
      />
    ) : null

    return (
      <div className='lk-feature-detail'>
        <div className='lk-feature-detail__header'>
          <h2>{feature.name}</h2>
          <div>Created {moment(feature.created_at).fromNow()}</div>
          <div>
            <Button
              onClick={deleteUnlocked ? null : this.unlockDelete}
              label='Delete feature'
              type='primary'
            />
            {cancel}
            {del}
          </div>
        </div>
        <Section title='Environments'>{envStatus}</Section>
      </div>
    )
  }
}

FeatureDetail.propTypes = {
  environments: PropTypes.arrayOf(
    PropTypes.shape({
      name: PropTypes.string
    })
  ).isRequired,
  feature: PropTypes.shape({
    name: PropTypes.string
  }).isRequired,
  onToggle: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired
}
