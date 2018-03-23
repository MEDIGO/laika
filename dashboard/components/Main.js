import React from 'react'
import { node } from 'prop-types'

import './Main.css'

const Main = ({ children }) => <main className='lk-main'>{children}</main>

Main.propTypes = {
  children: node
}

Main.defaultProps = {
  children: null
}

export default Main
