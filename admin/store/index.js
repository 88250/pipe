import axios from 'axios'

export const state = () => ({
  locale: 'zh',
  userName: 'solo',
  version: '1.0.0'
})

export const mutations = {
  setBaseInfo (state, data) {
    state.locale = data.locale
    state.userName = data.userName
    state.version = data.version
  }
}

export const actions = {
  async getBaseInfo ({ commit }) {
    try {
      const responseData = await axios.get('http://localhost:8888/base')
      commit('setBaseInfo', responseData)
    } catch (e) {
      throw e
    }
  }
}
