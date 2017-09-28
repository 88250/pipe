<template>
  <header class="header fn-flex">
    <div class="header__logo" v-if="from === 'admin'">
      <img src="~static/images/logo.png"/>
      Solo
    </div>
    <div class="header__nav fn-flex-1">
      <nuxt-link :to="`/login?goto=${$route.path.indexOf('/login') > -1 ? '/' : $route.fullPath}`">{{ $t('login', $store.state.locale) }}</nuxt-link>
    </div>
  </header>
</template>

<script>
  export default {
    props: ['from']
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'

  .header
    position: fixed
    height: 60px
    width: 100%
    z-index: 10
    top: 0

    &__logo
      background-color: $white
      width: 240px

    &__nav
      background-color: $blue
      a
        color: #fff
</style>
