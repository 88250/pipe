import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
import routes from './routes'

ReactDOM.render(
  <App routes={routes} />,
  document.getElementById('root')
)

registerServiceWorker();
