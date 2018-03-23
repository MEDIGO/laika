import React from 'react'
import {
  BrowserRouter as Router,
  Redirect,
  Route,
  Switch
} from 'react-router-dom'

import Container from '../components/Container'
import Footer from '../components/Footer'
import Header from '../components/Header'
import Main from '../components/Main'

import EnvironmentCreate from './EnvironmentCreate'
import FeatureCreate from './FeatureCreate'
import FeatureDetail from './FeatureDetail'
import Home from './Home'

const App = () => {
  return (
    <Router>
      <div className='app'>
        <Header />

        <Main>
          <Container>
            <Switch>
              <Route exact path='/' component={Home} />
              <Route exact path='/features/:name' component={FeatureDetail} />
              <Route exact path='/new/feature' component={FeatureCreate} />
              <Route
                exact
                path='/new/environment'
                component={EnvironmentCreate}
              />
              <Redirect to='/' />
            </Switch>
          </Container>
        </Main>

        <Footer />
      </div>
    </Router>
  )
}

export default App
