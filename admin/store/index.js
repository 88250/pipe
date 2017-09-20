import axios from 'axios'

export const state = () => ({
  locale: 'en',
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
  async nuxtClientInit ({ commit }, { app, req }) {
    try {
      const responseData = await axios.get('http://localhost:8888/base')
      commit('setBaseInfo', responseData.data)
      app.i18n.locale = responseData.data.lang
    } catch (e) {
      console.error(e)
    }
  }
}
