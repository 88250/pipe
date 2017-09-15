import React, { Component } from 'react';
import { BrowserRouter as Router, Link, Switch } from 'react-router-dom'

import routes from './routes'
import SubRoutes from './components/SubRoutes'

class App extends Component {
  render() {
    return (
      <Router>
        <div>
        <ul>
          <li><Link to="/login">Login</Link></li>
          <li><Link to="/admin">Admin</Link></li>
          <li><Link to="/admin/comments">comments</Link></li>
          <li><Link to="/admin/settings">settings</Link></li>
          <li><Link to="/admin/settings/users">users</Link></li>
        </ul>
        <Switch>
          {routes.map((route, i) => (
            <SubRoutes key={i} {...route}/>
          ))}
        </Switch>
        </div>
      </Router>
    );
  }
}

export default App;