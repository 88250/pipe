import React from 'react'

import SubRoutes from '../../../components/SubRoutes'

export const Index = ({routes}) => (
  <div>
    <h4>Admin Index</h4>
    {routes.map((route, i) => (
      <SubRoutes key={i} {...route}/>
    ))}
  </div>
)

export default Index