<template>
  <div class="console" id="particles">
    <div class="card" v-if="error.statusCode === 404">
      <h2 class="card__title">404</h2>
      <div class="card__body fn__clear">
        <div class="error__description">Page not found</div>
        <div class="fn__right">
          <a href="/">{{ $t('index', $store.state.locale)}}</a>
          |
          <a href="https://hacpai.com">{{ $t('hacpai', $store.state.locale)}}</a>
        </div>
      </div>
    </div>
    <div class="card" v-else>
      <h2 class="card__title">50X</h2>
      <div class="card__body fn__clear">
        <div class="error__description">{{error.statusCode}}</div>
        <div class="fn__right">
          <a href="https://github.com/88250/pipe/issues/new">{{ $t('reportIssue', $store.state.locale)}}</a> |
          <a href="https://hacpai.com">{{ $t('hacpai', $store.state.locale)}}</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import 'particles.js'
  import { initParticlesJS } from '~/plugins/utils'

  export default {
    head () {
      return {
        title: `${this.error.statusCode === 404 ? 404 : '50x'} - ${this.$store.state.blogTitle}`,
      }
    },
    props: {
      error: {
        type: Object,
        required: true,
      },
    },
    mounted () {
      initParticlesJS('particles')
    },
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'
  #__nuxt, #pipe, .main, #__layout
    height: 100%

  .main
    padding-top: 60px
    display: flex
    flex-direction: column
    transition: $transition
    box-sizing: border-box

  .error__description
    margin: 50px 0
    font-size: 1.25rem
    font-weight: 300
</style>
