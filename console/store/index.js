import axios from 'axios'

export const state = () => ({
  locale: 'en_US',
  userName: 'solo',
  version: '1.0.0'
})

export const mutations = {
  setBaseInfo (state, data) {
    state.locale = data.lang
    state.userName = data.name
    state.version = data.version
  }
}

export const actions = {
  async nuxtClientInit ({ commit }, { app }) {
    try {
      const responseData = await axios.get('http://localhost:8888/base')
      commit('setBaseInfo', responseData.data)
      // app.i18n.messages['zh_CN'] = { home: '1231 ' }
      // app.i18n.locale = responseData.data.lang
      // app.i18n.setLocaleMessage('zh_CN', { home: '111' })
      console.log(222)
    } catch (e) {
      console.error(e)
    }
  },
  changeLocale ({ commit }, { $i18n }) {
    $i18n.locale = 'en_US'
  }
}
