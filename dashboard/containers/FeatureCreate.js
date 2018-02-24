import React, { Component } from 'react'
import { shape, func } from 'prop-types'
import { withRouter } from 'react-router-dom'
import FeatureForm from '../components/FeatureForm'
import { createFeature } from '../utils/api'

class FeatureCreate extends Component {
  constructor(props) {
    super(props)

    this.state = {
      errorText: null
    }

    this.handleSubmit = this.handleSubmit.bind(this)
  }

  handleSubmit(feature) {
    createFeature(feature)
      .then(() => this.props.history.push('/'))
      .catch(err => this.setState({ errorText: err.message }))
  }

  render() {
    return (
      <FeatureForm
        titleText='Create a feature'
        submitText='Create feature'
        errorText={this.state.errorText}
        onSubmit={this.handleSubmit}
      />
    )
  }
}

FeatureCreate.propTypes = {
  history: shape({
    push: func
  }).isRequired
}

export default withRouter(FeatureCreate)
