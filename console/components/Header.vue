<template>
  <header class="header fn-flex">
    <div class="header__logo" v-if="from === 'admin'">
      <nuxt-link :to="$store.state.blogPath">
        <img src="~static/images/logo.png"/>
        {{ $store.state.blogTitle }}
      </nuxt-link>
    </div>
    <div class="header__nav fn-flex-1">
      <div v-if="$store.state.name === ''" class="header__bar--unlogin">
        <a :href="`/login?goto=${$route.path.indexOf('/login') > -1 ? '/' : $route.fullPath}`">{{ $t('login', $store.state.locale) }}</a>
        &nbsp;
        <nuxt-link to="/init">
          {{ $t('register', $store.state.locale) }}
        </nuxt-link>
      </div>
      <div v-else>
        <span class="header__bar" v-if="$route.path.indexOf('/admin') > -1" @click="toggleSide">
          <icon icon="bars"/>
        </span>
        {{ $store.state.nickname || $store.state.name }} &nbsp;
        <v-menu
          z-index="100"
          v-if="$route.path.indexOf('/admin') > -1 && $store.state.blogs.length > 1"
          :nudge-bottom="38">
          <v-toolbar-title slot="activator">
            <v-btn class="btn btn--success">
              {{ $store.state.blogTitle }}
              <icon icon="chevron-down"/>
            </v-btn>
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
        <v-btn class="btn btn--danger btn--space" @click="logout">{{ $t('logout', $store.state.locale) }}</v-btn>
      </div>
    </div>
  </header>
</template>

<script>
  export default {
    props: {
      from: {
        type: String,
        required: true
      }
    },
    methods: {
      toggleSide () {
        const className = document.querySelector('#sologo').className
        if (className === '') {
          document.querySelector('#sologo').className = 'body--side'
        } else {
          document.querySelector('#sologo').className = ''
        }
      },
      async switchBlog (item) {
        if (item.path === this.$store.state.blogPath) {
          return
        }
        const responseData = await this.axios.post(`/console/blog/switch/${item.id}`)
        if (responseData.code === 0) {
          item.role = responseData.data
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

  .body--side .header__logo
    display: flex

  .header
    position: fixed
    height: 60px
    width: 100%
    z-index: 10
    top: 0
    color: #fff

    &__logo
      display: none
      background-color: $white
      width: 240px
      align-items: center
      a
        color: $text-title
        margin: 0 auto
        font-size: 18px
        &:hover
          text-decoration: none

    &__bar
      float: left
      margin: 9px 15px
      cursor: pointer
      .icon
        height: 20px
        width: 20px
      &--unlogin
        margin-top: 8px
    &__nav
      background-color: $blue
      padding: 11px 15px 0 0
      text-align: right
      a
        color: #fff
</style>
