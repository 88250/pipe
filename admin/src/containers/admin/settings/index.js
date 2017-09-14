import React from 'react'

import SubRoutes from '../../../components/SubRoutes'

export const Settings = ({routes}) => (
  <div>
    <h4>Admin Settings</h4>
    {routes.map((route, i) => (
      <SubRoutes key={i} {...route}/>
    ))}
  </div>
)

export default Settings