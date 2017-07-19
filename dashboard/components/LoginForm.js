import React, { Component } from 'react';
import PropTypes from 'prop-types';

import Input from './Input';
import Form from './Form';
import Card from './Card';

function check(condition, message) {
  return condition ? null : message;
}

export default class LoginForm extends Component {
  constructor(props) {
    super(props);

    this.state = {
      username: '',
      password: '',
      errors: {
        username: null,
        password: null,
      },
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(name, value) {
    this.setState({
      [name]: value,
    });
  }

  handleSubmit(e) {
    e.preventDefault();

    const { username, password } = this.state;

    const errors = {
      username: check(username, 'Username is required'),
      password: check(password, 'Password is required'),
    };

    this.setState({ errors });

    if (errors.username || errors.password) return;

    this.props.onSubmit({
      username: this.state.username,
      password: this.state.password,
    });
  }

  render() {
    return (
      <Card title="Log In">
        <Form submitText={'Log In'} onSubmit={this.handleSubmit} errorText={this.props.errorText} >
          <Input
            label="Username"
            name="username"
            value={this.state.username}
            required
            onChange={this.handleChange}
            error={this.state.errors.username}
            placeholder="Username"
          />
          <Input
            label="Password"
            name="password"
            type="password"
            value={this.state.password}
            required
            onChange={this.handleChange}
            error={this.state.errors.password}
            placeholder="Password"
          />
        </Form>
      </Card>
    );
  }
}

LoginForm.propTypes = {
  onSubmit: PropTypes.func.isRequired,
  errorText: PropTypes.string,
};

LoginForm.defaultProps = {
  errorText: null,
};
