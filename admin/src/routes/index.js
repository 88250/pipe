import Login from '../containers/login'
import Admin from '../containers/admin/index'
import Comments from '../containers/admin/comments'
import Settings from '../containers/admin/settings'
import Users from '../containers/admin/settings/users'

const routes = [
  { path: '/login',
    component: Login
  },
  { path: '/admin',
    component: Admin,
    routes: [
      { path: '/admin/comments',
        component: Comments
      },
      { path: '/admin/settings',
        component: Settings,
        routes: [
          { path: '/admin/settings/users',
            component: Users
          }
        ]
      }
    ]
  }
]

export default routes



