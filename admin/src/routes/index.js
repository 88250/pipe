import Login from './login'
import Admin from './admin/index'
import Comments from './admin/comments'
import Settings from './admin/settings'
import Users from './admin/settings/users'

const routes = [
  { path: '/login',
    component: Login
  },
  { path: '/admin',
    component: Admin,
  },
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

export default routes



