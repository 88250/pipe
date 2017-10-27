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
