import React, { Component } from 'react'
import moment from 'moment'
import { shape, object, func } from 'prop-types'
import { withRouter } from 'react-router-dom'
import FeatureDetailComponent from '../components/FeatureDetail'
import { getFeature, toggleFeature, deleteFeature } from '../utils/api'

class FeatureDetail extends Component {
  constructor(props) {
    super(props)

    this.state = {
      loading: true,
      environments: [],
      feature: null
    }

    this.handleToggle = this.handleToggle.bind(this)
    this.handleDelete = this.handleDelete.bind(this)
  }

  componentDidMount() {
    getFeature(window.decodeURIComponent(this.props.match.params.name)).then(
      feature =>
        this.setState({
          loading: false,
          environments: feature.feature_status,
          feature
        })
    )
  }

  handleToggle(name, value) {
    const envs = this.state.environments.map(e => {
      if (e.name === name) {
        return Object.assign({}, e, {
          status: value,
          toggled_at: moment()
        })
      }
      return e
    })
    toggleFeature(name, this.state.feature.name, value).then(() => {
      this.setState({
        environments: envs
      })
    })
  }

  handleDelete(name) {
    deleteFeature(name).then(() => this.props.history.push('/'))
  }

  render() {
    if (this.state.loading) return null
    return (
      <FeatureDetailComponent
        environments={this.state.environments}
        feature={this.state.feature}
        onToggle={this.handleToggle}
        onDelete={this.handleDelete}
      />
    )
  }
}

FeatureDetail.propTypes = {
  match: shape({
    params: object
  }).isRequired,
  history: shape({
    push: func
  }).isRequired
}

export default withRouter(FeatureDetail)
