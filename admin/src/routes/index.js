import DefaultLayout from '../layouts'
import admin from './Admin'

/*  Note: Instead of using JSX, we recommend using react-router
    PlainRoute objects to build route definitions.   */

export const createRoutes = (store) => ({
  path        : '/',
  component   : DefaultLayout,
  indexRoute  : admin,
  childRoutes : [
   // CounterRoute(store)
  ]
})

export default createRoutes
