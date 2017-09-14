import React, { Component } from 'react';
import { BrowserRouter as Router, Link, Route } from 'react-router-dom'

import routes from './routes'

const RouteWithSubRoutes = (route) => (
  <Route exact path={route.path} render={props => (
    // pass the sub-routes down to keep nesting
    <route.component {...props} routes={route.routes}/>
  )}/>
)

class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <ul>
            <li><Link to="/login">Login</Link></li>
            <li><Link to="/admin">Admin</Link></li>
            <li><Link to="/admin/comments">commets</Link></li>
            <li><Link to="/admin/settings">settings</Link></li>
            <li><Link to="/admin/settings/users">users</Link></li>
          </ul>

          {routes.map((route, i) => (
            <RouteWithSubRoutes key={i} {...route}/>
          ))}
        </div>
      </Router>
    );
  }
}


export default App;