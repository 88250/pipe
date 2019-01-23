<template>
  <div class="console" id="particles">
    <div class="card">
      <h2 class="card__title">404</h2>
      <div class="card__body fn__clear">
        <div class="error__description">Page not found</div>
        <div class="fn__right">
          <nuxt-link to="/">{{ $t('index', $store.state.locale)}}</nuxt-link>
          |
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
        title: '404 - Pipe'
      }
    },
    mounted () {
      initParticlesJS('particles')
    }
  }
</script>

<style lang="sass">
  .error__description
    margin: 50px 0
    font-size: 1.25rem
    font-weight: 300
</style>
