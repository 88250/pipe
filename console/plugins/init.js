import Vue from 'vue'
import VueI18n from 'vue-i18n'
import Vuetify from 'vuetify'
import MavonEditor from 'mavon-editor'
import Icon from '~/components/Icon'

export default ({ app, isClient, store }) => {
  Vue.component(Icon.name, Icon)
  Vue.use(Vuetify)
  Vue.use(MavonEditor)
  // for mavon-editor, it needs $el
  app.el = '#__nuxt'
  Vue.use(VueI18n)
  app.i18n = new VueI18n()
}
