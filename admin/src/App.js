import React, { Component } from 'react';
import { BrowserRouter as Router, Link } from 'react-router-dom'

import SubRoutes from './components/SubRoutes'
import routes from './routes'

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

          {routes.map((route, i) => (
            <SubRoutes key={i} {...route}/>
          ))}
        </div>
      </Router>
    );
  }
}


export default App;