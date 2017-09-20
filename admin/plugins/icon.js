import Vue from 'vue'
import Icon from '~/components/Icon'

export default ({ app, isClient, store }) => {
  Vue.component(Icon.name, Icon)
}
