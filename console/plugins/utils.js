export const genMenuData = (app, locale) => [
  {
    title: app.$t('home', locale),
    icon: 'home',
    link: '/admin/',
    role: 2
  },
  {
    title: app.$t('postArticle', locale),
    icon: 'add',
    link: '/admin/articles/post',
    role: 2
  },
  {
    title: app.$t('themeList', locale),
    icon: 'theme',
    link: '/admin/themes',
    role: 1
  },
  {
    title: app.$t('manage', locale),
    icon: 'manage',
    active: true,
    role: 2,
    items: [
      {
        title: app.$t('articleList', locale),
        link: '/admin/articles',
        role: 2
      },
      {
        title: app.$t('commentList', locale),
        link: '/admin/comments',
        role: 2
      },
      {
        title: app.$t('categoryList', locale),
        link: '/admin/categories',
        role: 1
      },
      {
        title: app.$t('navigationList', locale),
        link: '/admin/navigations',
        role: 1
      },
      {
        title: app.$t('userList', locale),
        link: '/admin/users',
        role: 2
      } /*,
      {
        title: app.$t('blogManage', locale),
        link: '/admin/blogs',
        role: 2
      } */
    ]
  },
  {
    title: app.$t('setting', locale),
    icon: 'setting',
    active: app.$route.path.indexOf('settings') > -1,
    role: 1,
    items: [
      {
        title: app.$t('baseInfo', locale),
        link: '/admin/settings/basic',
        role: 1
      },
      {
        title: app.$t('preference', locale),
        link: '/admin/settings/preference',
        role: 1
      },
      {
        title: app.$t('signs', locale),
        link: '/admin/settings/sign',
        role: 1
      },
      {
        title: app.$t('internationalization', locale),
        link: '/admin/settings/i18n',
        role: 1
      },
      {
        title: app.$t('feed', locale),
        link: '/admin/settings/feed',
        role: 1
      }
    ]
  },
  {
    title: app.$t('others', locale),
    icon: 'inbox',
    link: '/admin/others',
    role: 2
  },
  {
    title: app.$t('about', locale),
    icon: 'info',
    link: '/admin/about',
    role: 2
  }
]
