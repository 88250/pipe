import Vue from 'vue'
import Vuetify from 'vuetify'
import Icon from '~/components/Icon'

export default ({ app, isClient, store }) => {
  Vue.component(Icon.name, Icon)
  Vue.use(Vuetify)
}
