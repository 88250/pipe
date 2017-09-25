import axios from 'axios'

export const state = () => ({
  locale: 'zh_CN',
  userName: 'solo',
  version: '1.0.0'
})

export const mutations = {
  setBaseInfo (state, data) {
    state.locale = data.lang
    state.userName = data.name
    state.version = data.version
  },
  setLocale (state, locale) {
    state.locale = locale
  }
}

export const actions = {
  async nuxtClientInit ({ commit }, { app }) {
    try {
      const responseData = await axios.get('http://localhost:8888/base')
      if (app.i18n.messages[responseData.data.lang]) {
        app.i18n.locale = responseData.data.lang
      } else {
        const message = require(`../../i18n/${responseData.data.lang}.json`)
        app.i18n.setLocaleMessage(responseData.data.lang, message)
      }
      commit('setBaseInfo', responseData.data)
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
