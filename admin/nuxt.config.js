const env = require(`./config/env.json`)

module.exports = {
  /*
  ** Headers of the page
  */
  head: {
    title: 'solo-go-admin',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'The admin of solo.go' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#3B8070' },
  css: [
    '@/assets/scss/main.scss'
  ],
  plugins: ['~/plugins/i18n.js'],
  /*
  ** Build configuration
  */
  build: {
    vendor: ['vue-i18n'],
    publicPath: env.publicPath,
    extractCSS: true,
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
