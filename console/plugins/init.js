import Vue from 'vue'
import Vuetify from 'vuetify'
import MavonEditor from 'mavon-editor'
import Icon from '~/components/Icon'

export default ({ app, isClient, store }) => {
  Vue.component(Icon.name, Icon)
  Vue.use(Vuetify)
  Vue.use(MavonEditor)
  // for mavon-editor, it needs $el
  app.el = '#__nuxt'
}
