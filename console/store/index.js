import vueAxios from '~/plugins/axios'

export const state = () => ({
  locale: 'zh_CN',
  userName: 'solo.go',
  version: '1.0.0',
  isInit: true,
  snackMsg: '',
  snackBar: false,
  snackModify: 'error'
})

export const mutations = {
  setBaseInfo (state, data) {
    state.locale = data.locale
    state.version = data.version
    state.isInit = data.inited
  },
  setLocale (state, locale) {
    state.locale = locale
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
