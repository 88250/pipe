import Vue from 'vue'
import VueI18n from 'vue-i18n'

Vue.use(VueI18n)

export default ({ app, store }) => {
  let messages = {}
  messages[store.state.locale] = require(`../../i18n/${store.state.locale}.json`)
  messages['zh_CN'] = require('../../i18n/zh_CN.json')
  app.i18n = new VueI18n({
    locale: store.state.locale,
    fallbackLocale: 'en_US',
    messages
  })
}
