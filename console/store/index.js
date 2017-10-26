import vueAxios from '~/plugins/axios'

export const state = () => ({
  locale: 'zh_CN',
  version: '1.0.0',
  isInit: true,
  name: '',
  nickname: '',
  blogTitle: '',
  avatarURL: '',
  blogURL: '/',
  role: 2,
  blogs: [{
    title: '',
    id: ''
  }],
  snackMsg: '',
  snackBar: false,
  snackModify: 'error',
  menu: []
})

export const mutations = {
  setMenu (state, data) {
    state.menu = data
  },
  setBaseInfo (state, data) {
    state.locale = data.locale
    state.version = data.version
    state.isInit = data.inited
  },
  setLocale (state, locale) {
    state.locale = locale
  },
  setIsInit (state, isInit) {
    state.isInit = isInit
  },
  setUserInfo (state, data) {
    if (data) {
      state.name = data.name
      state.nickname = data.nickname
      state.blogTitle = data.blogTitle
      state.blogURL = data.blogURL
      state.role = data.role
      state.blogs = data.blogs
      state.avatarURL = data.avatarURL
      localStorage.setItem('userInfo', JSON.stringify(data))
    } else {
      state.name = ''
      state.nickname = ''
      state.blogTitle = ''
      state.blogURL = '/'
      state.role = 2
      state.avatarURL = ''
      state.blogs = [{
        title: '',
        id: ''
      }]
      localStorage.removeItem('userInfo')
    }
  },
  getUserInfo (state) {
    const userInfo = localStorage.getItem('userInfo')
    if (userInfo) {
      const userInfoJSON = JSON.parse(userInfo)
      state.name = userInfoJSON.name
      state.nickname = userInfoJSON.nickname
      state.blogTitle = userInfoJSON.blogTitle
      state.blogURL = userInfoJSON.blogURL
      state.role = userInfoJSON.role
      state.blogs = userInfoJSON.blogs
      state.avatarURL = userInfoJSON.avatarURL
      return userInfoJSON
    } else {
      state.name = ''
      state.nickname = ''
      state.blogTitle = ''
      state.blogURL = '/'
      state.avatarURL = ''
      state.role = 2
      state.blogs = [{
        title: '',
        id: ''
      }]
    }
  },
  setBlog (state, data) {
    const userInfo = localStorage.getItem('userInfo')
    if (!userInfo) {
      return
    }
    const userInfoJSON = JSON.parse(userInfo)
    if (data) {
      state.blogTitle = data.title
      state.blogURL = data.path
      state.role = data.role
      userInfoJSON.blogURL = data.path
      userInfoJSON.blogTitle = data.title
      userInfoJSON.role = data.role
    } else {
      state.blogTitle = ''
      state.blogURL = '/'
      state.role = 2
      userInfoJSON.blogURL = '/'
      userInfoJSON.blogTitle = ''
      userInfoJSON.role = 2
    }
    localStorage.setItem('userInfo', JSON.stringify(userInfoJSON))
  },
  setSnackBar (state, data) {
    state.snackBar = data.snackBar
    state.snackMsg = data.snackMsg
    if (data.snackModify) {
      state.snackModify = data.snackModify
    } else {
      state.snackModify = 'error'
    }
  }
}

export const actions = {
  async nuxtClientInit ({ commit }, { app }) {
    try {
      const responseData = await vueAxios().get('/status')
      if (responseData) {
        if (app.i18n.messages[responseData.locale]) {
          app.i18n.locale = responseData.locale
        } else {
          const message = require(`../../i18n/${responseData.locale}.json`)
          app.i18n.setLocaleMessage(responseData.locale, message)
        }
        commit('setBaseInfo', responseData)
      }
    } catch (e) {
      console.error(e)
    }
  },
  setLocaleMessage ({ commit }, locale) {
    if (this.app.i18n.messages[locale]) {
      this.app.i18n.locale = locale
    } else {
      const message = require(`../../i18n/${locale}.json`)
      this.app.i18n.setLocaleMessage(locale, message)
    }
    commit('setLocale', locale)
  }
}
