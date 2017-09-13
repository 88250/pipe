import React, { Component } from 'react';
import {
  Route,
  Link
} from 'react-router-dom'
import Navigation from './routes/Tools'

class App extends Component {
  render() {
    return (
      <div>
        <nav>
          <Link to="/tools">Tools</Link>
          <Link to="/">Index</Link>
        </nav>
        <div>
          <Route path="/tools" component={Navigation}/>
        </div>
      </div>
    );
  }
}

export default App;