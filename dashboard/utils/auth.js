import React, { Component } from 'react';
import { getUser } from '../utils/api';

export function logIn(username, password) {
  localStorage.setItem('credentials.username', username);
  localStorage.setItem('credentials.password', password);
  return getUser(username);
}

export function isLoggedIn() {
  const username = localStorage.getItem('credentials.username');
  const password = localStorage.getItem('credentials.password');
  return Boolean(username && password);
}

export function logOut() {
  localStorage.removeItem('credentials.username');
  localStorage.removeItem('credentials.password');
}

export function withAuth(WrappedComponent) {
  return class extends Component {
    componentDidMount() {
      // eslint-disable-next-line
      if (!isLoggedIn()) this.props.history.push('/login');
    }

    render() {
      return isLoggedIn() ? <WrappedComponent {...this.props} /> : null;
    }
  };
}
