<template>
  <header class="header fn-flex">
    <div class="header__logo" v-if="from === 'admin'">
      <nuxt-link :to="$store.state.blogPath">
        <img src="~static/images/logo.png"/>
        {{ $store.state.blogTitle }}
      </nuxt-link>
    </div>
    <div class="header__nav fn-flex-1" data-app="true">
      <div v-if="$store.state.name === ''">
        <nuxt-link :to="`/login?goto=${$route.path.indexOf('/login') > -1 ? '/' : $route.fullPath}`">
          {{ $t('login', $store.state.locale) }}
        </nuxt-link>
      </div>
      <div v-else>
        {{ $store.state.nickname }}
        <v-menu
          v-if="$route.path.indexOf('/admin') > -1"
          :nudge-bottom="38"
          :nudge-right="24"
          :nudge-width="100">
          <v-toolbar-title slot="activator">
            <button class="btn btn--success">
              {{ $store.state.blogTitle }}
              <icon v-if="$store.state.blogs.length > 1" icon="chevron-down"/>
            </button>
          </v-toolbar-title>
          <v-list>
            <v-list-tile>
              <v-list-tile-title v-for="item in $store.state.blogs" :key="item.id">
                <div @click="switchBlog(item)">{{ item.title }}</div>
              </v-list-tile-title>
            </v-list-tile>
          </v-list>
        </v-menu>

        <nuxt-link v-if="$route.path.indexOf('/admin') === -1" to="/admin">{{ $t('manage', $store.state.locale) }}
        </nuxt-link>
        <button class="btn btn--danger btn--space" @click="logout">{{ $t('logout', $store.state.locale) }}</button>
      </div>
    </div>
  </header>
</template>

<script>
  export default {
    props: ['from'],
    methods: {
      async switchBlog (item) {
        if (item.path === this.$store.state.blogPath) {
          return
        }
        const responseData = await this.axios.post('/console/blog/switch', item.id)
        if (responseData.code === 0) {
          this.$store.commit('setBlog', item)
          this.$router.go()
        } else {
          this.commit('setSnackBar', {
            snackBar: true,
            snackMsg: responseData.msg
          })
        }
      },
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
