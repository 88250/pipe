export const genMenuData = (app, locale) => [
  {
    title: app.$t('home', locale),
    icon: 'home',
    link: '/admin/',
    role: 3
  },
  {
    title: app.$t('postArticle', locale),
    icon: 'add',
    link: '/admin/articles/post',
    role: 3
  },
  {
    title: app.$t('themeList', locale),
    icon: 'theme',
    link: '/admin/themes',
    role: 2
  },
  {
    title: app.$t('manage', locale),
    icon: 'manage',
    active: true,
    role: 3,
    items: [
      {
        title: app.$t('articleList', locale),
        link: '/admin/articles',
        role: 3
      },
      {
        title: app.$t('commentList', locale),
        link: '/admin/comments',
        role: 3
      },
      {
        title: app.$t('categoryList', locale),
        link: '/admin/categories',
        role: 2
      },
      {
        title: app.$t('navigationList', locale),
        link: '/admin/navigations',
        role: 2
      },
      {
        title: app.$t('userList', locale),
        link: '/admin/users',
        role: 3
      } /*,
      {
        title: app.$t('blogManage', locale),
        link: '/admin/blogs',
        role: 3
      } */
    ]
  },
  {
    title: app.$t('setting', locale),
    icon: 'setting',
    active: app.$route.path.indexOf('settings') > -1,
    role: 2,
    items: [
      {
        title: app.$t('baseInfo', locale),
        link: '/admin/settings/basic',
        role: 2
      },
      {
        title: app.$t('preference', locale),
        link: '/admin/settings/preference',
        role: 2
      },
      {
        title: app.$t('signs', locale),
        link: '/admin/settings/sign',
        role: 2
      },
      {
        title: app.$t('internationalization', locale),
        link: '/admin/settings/i18n',
        role: 2
      },
      {
        title: app.$t('feed', locale),
        link: '/admin/settings/feed',
        role: 2
      }
    ]
  },
  {
    title: app.$t('others', locale),
    icon: 'inbox',
    link: '/admin/others',
    role: 3
  },
  {
    title: app.$t('about', locale),
    icon: 'info',
    link: '/admin/about',
    role: 3
  }
]

export const initParticlesJS = (id) => {
  window.particlesJS(id, {
    'particles': {
      'number': {
        'value': 6,
        'density': {
          'enable': true,
          'value_area': 200
        }
      },
      'color': {
        'value': '#bbb'
      },
      'opacity': {
        'value': 0.5,
        'anim': {
          'speed': 1,
          'opacity_min': 0.1
        }
      },
      'size': {
        'value': 10,
        'random': true,
        'anim': {
          'enable': false,
          'speed': 80,
          'size_min': 0.1,
          'sync': false
        }
      },
      'line_linked': {
        'enable': true,
        'distance': 300,
        'color': '#bbb',
        'opacity': 0.4,
        'width': 1
      },
      'move': {
        'enable': true,
        'speed': 2,
        'direction': 'none',
        'straight': false,
        'out_mode': 'out',
        'bounce': false,
        'attract': {
          'enable': false,
          'rotateX': 300,
          'rotateY': 600
        }
      }
    },
    'interactivity': {
      'detect_on': 'canvas',
      'events': {
        'onclick': {
          'enable': false
        },
        'resize': true
      },
      'modes': {
        'grab': {
          'distance': 400,
          'line_linked': {
            'opacity': 0.7
          }
        },
        'bubble': {
          'distance': 800,
          'size': 80,
          'duration': 2,
          'opacity': 0.8,
          'speed': 3
        },
        'repulse': {
          'distance': 400,
          'duration': 0.4
        },
        'push': {
          'particles_nb': 4
        },
        'remove': {
          'particles_nb': 2
        }
      }
    },
    'retina_detect': true
  })
}
