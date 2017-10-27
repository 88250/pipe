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
  role: 3, // 0-no login, 1-admin, 2-blog admin, 3-blog user, 4-visitor
  blogs: [{
    title: '',
    id: ''
  }],
  snackMsg: '',
  snackBar: false,
  snackModify: 'error',
  menu: [],
  tagsItems: []
})

export const mutations = {
  setMenu (state, data) {
    state.menu = data
  },
  setStatus (state, data) {
    state.locale = data.locale
    state.version = data.version
    state.isInit = data.inited
    state.role = data.role
    state.name = data.name
    state.nickname = data.nickname
    state.blogTitle = data.blogTitle
    state.blogURL = data.blogURL
    state.blogs = data.blogs
    state.avatarURL = data.avatarURL
  },
  setLocale (state, locale) {
    state.locale = locale
  },
  setIsInit (state, isInit) {
    state.isInit = isInit
  },
  setBlog (state, data) {
    state.blogTitle = data.title
    state.blogURL = data.path
    state.role = data.role
  },
  setSnackBar (state, data) {
    state.snackBar = data.snackBar
    state.snackMsg = data.snackMsg
    if (data.snackModify) {
      state.snackModify = data.snackModify
    } else {
      state.snackModify = 'error'
    }
  },
  setTagsItems (state, data) {
    state.tagsItems = data
  }
}

export const actions = {
  async nuxtClientInit ({commit}, {app}) {
    try {
      const responseData = await vueAxios().get('/status')
      if (responseData) {
        if (app.i18n.messages[responseData.locale]) {
          app.i18n.locale = responseData.locale
        } else {
          const message = require(`../../i18n/${responseData.locale}.json`)
          app.i18n.setLocaleMessage(responseData.locale, message)
        }
        commit('setStatus', responseData)
      }
    } catch (e) {
      console.error(e)
    }
  },
  setLocaleMessage ({commit}, locale) {
    if (this.app.i18n.messages[locale]) {
      this.app.i18n.locale = locale
    } else {
      const message = require(`../../i18n/${locale}.json`)
      this.app.i18n.setLocaleMessage(locale, message)
    }
    commit('setLocale', locale)
  },
  async getTags ({commit, state}) {
    if (state.tagsItems.length > 0) {
      return
    }
    const tagResponseData = await vueAxios().get('/console/tags/')
    if (tagResponseData) {
      let tagList = []
      tagResponseData.map((v) => {
        tagList.push(v.title)
      })
      commit('setTagsItems', tagList)
    }
  }
}
