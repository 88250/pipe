import vueAxios from '~/plugins/axios'

export const state = () => ({
  locale: 'zh_CN',
  userName: 'solo',
  version: '1.0.0',
  isInit: true,
  snackMsg: '',
  snackBar: false,
  snackModify: 'error'
})

export const mutations = {
  setBaseInfo (state, data) {
    state.locale = data.lang
    state.name = data.name
    state.version = data.version
    state.isInit = data.isInit
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
      const responseData = await vueAxios().get('/base')
      if (app.i18n.messages[responseData.lang]) {
        app.i18n.locale = responseData.lang
      } else {
        const message = require(`../../i18n/${responseData.lang}.json`)
        app.i18n.setLocaleMessage(responseData.lang, message)
      }
      commit('setBaseInfo', responseData)
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
