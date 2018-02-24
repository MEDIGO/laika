import React, { Component } from 'react'
import { withRouter } from 'react-router-dom'
import FeatureList from '../components/FeatureList'
import { listFeatures, listEnvironments } from '../utils/api'

class Home extends Component {
  constructor(props) {
    super(props)

    this.state = {
      environments: [],
      features: []
    }
  }

  componentDidMount() {
    Promise.all([listEnvironments(), listFeatures()]).then(
      ([environments, features]) => this.setState({ environments, features })
    )
  }

  render() {
    return (
      <FeatureList
        environments={this.state.environments}
        features={this.state.features}
      />
    )
  }
}

export default withRouter(Home)
