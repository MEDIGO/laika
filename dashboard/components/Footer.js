import React from 'react'
import Container from './Container'

import './Footer.css'

export default function Footer() {
  return (
    <footer className='lk-footer'>
      <Container>
        <span>MEDIGO GmbH © 2017</span>
        <a href='https://github.com/MEDIGO/laika'>Contribute</a>
      </Container>
    </footer>
  )
}
