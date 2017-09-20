import Vue from 'vue'
import VueI18n from 'vue-i18n'

Vue.use(VueI18n)

export default ({ app, isClient, store }) => {
  app.i18n = new VueI18n({
    locale: store.locale,
    messages: {
      'en': require('~/locales/en.json'),
      'zh': require('~/locales/zh.json')
    }
  })
}
