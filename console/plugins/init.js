import Vue from 'vue'
import VueI18n from 'vue-i18n'
import Vuetify from 'vuetify'

export default ({app}) => {
  Vue.use(Vuetify)
  Vue.use(VueI18n)
  app.i18n = new VueI18n()
}
