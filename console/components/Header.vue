<template>
  <header class="header fn-flex">
    <div class="header__logo" v-if="from === 'admin'">
      <img src="~static/images/logo.png"/>
      Solo
    </div>
    <div class="header__nav fn-flex-1">
      <div v-if="$store.state.name === ''">
        <nuxt-link :to="`/login?goto=${$route.path.indexOf('/login') > -1 ? '/' : $route.fullPath}`">
          {{ $t('login', $store.state.locale) }}
        </nuxt-link>
      </div>
      <div v-else>
        {{ $store.state.nickname }}
        <nuxt-link to="/admin">{{ $t('manage', $store.state.locale) }}</nuxt-link>
        <button class="btn btn--danger btn--space" @click="logout">{{ $t('logout', $store.state.locale) }}</button>
      </div>
    </div>
  </header>
</template>

<script>
  export default {
    props: ['from'],
    methods: {
      async logout () {
        const responseData = await this.axios.post('/logout')
        if (responseData.code === 0) {
          this.$store.commit('setUserInfo', null)
          this.$router.push('/')
        } else {
          this.commit('setSnackBar', {
            snackBar: true,
            snackMsg: responseData.msg
          })
        }
      }
    }
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
