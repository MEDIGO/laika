import React, { Component } from 'react';
import { withRouter } from 'react-router-dom';
import FeatureList from '../components/FeatureList';
import { listFeatures } from '../utils/api';

class Home extends Component {
  constructor(props) {
    super(props);

    this.state = {
      features: [],
    };
  }

  componentDidMount() {
    listFeatures().then(features => this.setState({ features }));
  }

  render() {
    return <FeatureList features={this.state.features} />;
  }
}

export default withRouter(Home);
