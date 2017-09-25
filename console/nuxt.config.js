const env = require(`./config/env.json`)

module.exports = {
  /*
  ** Headers of the page
  */
  head: {
    title: 'Solo.go',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { name: 'robots', content: 'none' },
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
    ]
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#1976d2' },
  css: [
    '~/assets/scss/main.scss'
  ],
  plugins: [
    { src: '~/plugins/i18n.js', ssr: false },
    { src: '~/plugins/icon.js', ssr: false },
    { src: '~/plugins/vuetify.js', ssr: false },
    { src: '~/plugins/nuxt-client-init.js', ssr: false }
  ],
  mode: 'spa',
  /*
  ** Build configuration
  */
  build: {
    vendor: ['vue-i18n', '~/assets/symbol.js', 'axios', 'vuetify'],
    publicPath: env.publicPath,
    extractCSS: true,
    ssr: false,
    /*
    ** Run ESLint on save
    */
    extend (config, ctx) {
      if (ctx.dev && ctx.isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/
        })
      }
    }
  }
}
