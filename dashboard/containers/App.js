import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import Home from './Home';
import Container from '../components/Container';
import FeatureCreate from './FeatureCreate';
import EnvironmentCreate from './EnvironmentCreate';
import FeatureDetail from './FeatureDetail';
import Header from '../components/Header';
import Footer from '../components/Footer';

export default function App() {
  return (
    <Router history={history}>
      <div>
        <Header />

        <Container>
          <Switch>
            <Route exact path="/" component={Home} />
            <Route exact path="/features/new" component={FeatureCreate} />
            <Route exact path="/features/:name" component={FeatureDetail} />
            <Route exact path="/environments/new" component={EnvironmentCreate} />
          </Switch>
        </Container>

        <Footer />
      </div>
    </Router>
  );
}
