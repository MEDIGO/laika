import React, { Component } from 'react';
import PropTypes from 'prop-types';

import Form from './Form';
import Input from './Input';
import Card from './Card';

export default class FeatureForm extends Component {
  constructor(props) {
    super(props);

    this.state = {};

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(name, value) {
    this.setState({ [name]: value });
  }

  handleSubmit() {
    this.props.onSubmit(this.state);
  }

  render() {
    return (
      <Card title={this.props.titleText}>
        <Form
          onSubmit={this.handleSubmit}
          submitText={this.props.submitText}
        >
          <Input
            label="Name"
            name="name"
            value={this.state.name}
            required
            onChange={this.handleChange}
            placeholder="e.g. development"
          />
        </Form>
      </Card>
    );
  }
}

FeatureForm.propTypes = {
  onSubmit: PropTypes.func.isRequired,
  submitText: PropTypes.string,
  titleText: PropTypes.string.isRequired,
};

FeatureForm.defaultProps = {
  submitText: null,
};
