import axios from 'axios'
import VueAxios from 'vue-axios'
import Vue from 'vue'

export default () => {
  const customAxios = axios.create({
    baseURL: process.env.clientBaseURL
  })

  Vue.use(VueAxios, customAxios)

  customAxios.interceptors.response.use(function (response) {
    return response.data
  }, function (error) {
    return Promise.reject(error)
  })

  return customAxios
}
