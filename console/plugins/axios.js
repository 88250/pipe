import axios from 'axios'
import VueAxios from 'vue-axios'
import Vue from 'vue'

export default (ctx) => {
  const customAxios = axios.create({
    baseURL: process.env.AxiosBaseURL
  })

  Vue.use(VueAxios, customAxios)

  customAxios.interceptors.request.use((config) => {
    // clear cache
    // if (config.method === 'get') {
    //   let char = '?'
    //   if (config.url.split('?').length > 1) {
    //     char = '&'
    //   }
    //   config.url += `${char}${(new Date()).getTime()}`
    // }
    return config
  })

  customAxios.interceptors.response.use((response) => {
    if (response.config.method === 'get' || response.config.method === 'delete') {
      // get and delete use snack tip
      if (response.data.code === 0) {
        return response.data.data
      } else if(response.data.code == -2) {
        window.location.href = `${ctx.env.Server}/start`
      } else {
        ctx.store.commit('setSnackBar', {
          snackBar: true,
          snackMsg: response.data.msg
        })
      }
    } else {
      // other, deal with yourself
      return response.data
    }
  }, (error) => {
    ctx.store.commit('setSnackBar', {
      snackBar: true,
      snackMsg: ctx.app.i18n.t('requestError', ctx.app.store.state.locale)
    })
    return Promise.reject(error)
  })

  return customAxios
}
