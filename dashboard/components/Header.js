import React from 'react';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';

import Container from './Container';
import './Header.css';

export default function Header() {
  return (
    <div className="lk-header">
      <Container>
        <div className="lk-header__wrapper">
          <ul>
            <li><Link to="/">LAIKA</Link></li>
            <li><Link to="/environments/new">New Environment</Link></li>
            <li><Link to="/features/new">New Feature</Link></li>
          </ul>
        </div>
      </Container>
    </div>
  );
}
