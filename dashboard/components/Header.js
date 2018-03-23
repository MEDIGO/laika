import React from 'react'
import { Link } from 'react-router-dom'

import Container from './Container'
import Logo from './Logo'
import './Header.css'

const Header = () => (
  <header className='lk-header'>
    <Container>
      <nav className='lk-header__nav'>
        <ul className='lk-header__nav-items'>
          <li>
            <Link to='/'>
              <Logo />
            </Link>
          </li>
          <li>
            <Link to='/new/environment'>New environment</Link>
          </li>
          <li>
            <Link to='/new/feature'>New feature</Link>
          </li>
        </ul>
      </nav>
    </Container>
  </header>
)

export default Header
