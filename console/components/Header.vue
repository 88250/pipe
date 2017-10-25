<template>
  <header class="header fn-flex">
    <div class="header__logo" v-if="from === 'admin'">
      <nuxt-link :to="$store.state.blogURL">
        <img src="~static/images/logo.png"/>
        {{ $store.state.blogTitle }}
      </nuxt-link>
    </div>
    <div class="header__nav fn-flex-1 fn-flex">
      <div v-if="$store.state.name === ''" class="header__bar--theme">
        <a :href="`/login?goto=${$route.path.indexOf('/login') > -1 ? '/' : $route.fullPath}`"
          v-if="$route.path !== '/login'">{{ $t('login', $store.state.locale) }}</a>
        <nuxt-link class="btn--space" to="/init" v-if="$route.path !== '/init'">{{ $t('register', $store.state.locale) }}</nuxt-link>
      </div>
      <template v-else>
        <span class="header__bar--icon fn-flex-1" v-if="$route.path.indexOf('/admin') > -1" >
          <v-icon @click="toggleSide">bars</v-icon>
        </span>
        <div :class="$route.path.indexOf('/admin') == -1 ? 'header__bar--theme' : 'header__bar--admin'">
          {{ $store.state.nickname || $store.state.name }} &nbsp;
          <v-menu
            z-index="100"
            v-if="$route.path.indexOf('/admin') > -1 && $store.state.blogs.length > 1"
            :nudge-bottom="38">
            <v-toolbar-title slot="activator">
              <v-btn class="btn btn--success">
                {{ $store.state.blogTitle }}
                <v-icon>arrow_drop_down</v-icon>
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
          <v-btn class="btn--small btn--danger btn--space" @click="logout">{{ $t('logout', $store.state.locale) }}</v-btn>
        </div>
      </template>
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
        const className = document.querySelector('#pipe').className
        if (className === '') {
          document.querySelector('#pipe').className = 'body--side'
        } else {
          document.querySelector('#pipe').className = ''
        }
      },
      async switchBlog (item) {
        if (item.URL === this.$store.state.blogURL) {
          return
        }
        const responseData = await this.axios.post(`/console/blogs/switch/${item.id}`)
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
    &__nav
      background-color: $blue
      align-items: center
      a
        color: #fff
    &__bar--icon .icon
      cursor: pointer
      margin: 0 15px
      height: 20px
      width: 20px

    &__bar--admin
      padding-right: 30px
    &__bar--theme
      text-align: right
      width: 100%
      padding-right: 30px
  @media (max-width: 768px)
    .header__bar--theme,
    .header__bar--admin
      padding-right: 15px
</style>
