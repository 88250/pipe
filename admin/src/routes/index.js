import Login from 'login'

const routes = [
  { path: '/login',
    component: Login
  },
  { path: '/admin',
    component: Index,
    routes: [
      { path: '/admin/comments',
        component: Bus
      },
      { path: '/admin/settings',
        component: Cart
      }
    ]
  }
]

export default routes



