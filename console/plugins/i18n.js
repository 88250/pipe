import Vue from 'vue'
import VueI18n from 'vue-i18n'

Vue.use(VueI18n)

export default ({ app, isClient, store }) => {
  app.i18n = new VueI18n({
    locale: store.state.locale,
    messages: {
      'en_US': require('../../i18n/en_US.json'),
      'zh_CN': require('../../i18n/zh_CN.json')
    }
  })
}
