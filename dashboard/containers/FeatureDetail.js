import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';
import FeatureDetailComponent from '../components/FeatureDetail';
import { getFeature, updateFeature } from '../utils/api';
import { withAuth } from '../utils/auth';

class FeatureDetail extends Component {
  constructor(props) {
    super(props);

    this.state = {
      loading: true,
      feature: null,
    };

    this.handleToggle = this.handleToggle.bind(this);
  }

  componentDidMount() {
    getFeature(this.props.match.params.name).then(feature => this.setState({
      loading: false,
      feature,
    }));
  }

  handleToggle(name, value) {
    updateFeature(this.props.match.params.name, { status: { [name]: value } }).then(() => {
      const feature = this.state.feature;
      feature.status[name] = value;
      this.setState({ feature });
    });
  }

  render() {
    if (this.state.loading) return null;
    return (
      <FeatureDetailComponent
        feature={this.state.feature}
        onToggle={this.handleToggle}
      />
    );
  }
}

FeatureDetail.propTypes = {
  match: PropTypes.shape({
    params: PropTypes.object,
  }).isRequired,
};

export default withAuth(withRouter(FeatureDetail));