<template>
  <header class="header fn-flex">
    <div :class="$route.path.indexOf('/admin') > -1 ? 'header__logo' : 'header__logo header__logo--theme'">
      <a :href="$store.state.blogURL || '/'">
        <img class="header__logo-img" src="~static/images/logo.png"/>
        {{ $store.state.blogTitle || 'Pipe' }}
      </a>
    </div>
    <div class="header__nav fn-flex-1 fn-flex">
      <template v-if="$store.state.role === 0">
        <span class="fn-flex-1"> &nbsp;</span>
        <span>
          <nuxt-link to="/login">{{ $t('login', $store.state.locale)}}</nuxt-link>
          <nuxt-link to="/register" class="btn--space btn--success btn btn--small">
            {{ $t('register', $store.state.locale)}}
          </nuxt-link>
        </span>
      </template>
      <template v-else>
        <span class="header__bar--icon fn-flex-1">
          <span v-if="$route.path.indexOf('/admin') > -1 && from !== 'error'">
            <div class="side__icon fn-left" @click="toggleSide">
              <span class="side__icon-line"></span>
              <span class="side__icon-line side__icon-line--middle"></span>
              <span class="side__icon-line"></span>
            </div>
            <a :href="$store.state.blogURL">
              <img class="header__logo-img fn-none" src="~static/images/logo.png"/>
            </a>
          </span>
          <template v-else>&nbsp;</template>
        </span>
        <v-menu
          z-index="100"
          :min-width="120"
          :open-on-hover="true"
          :nudge-bottom="30"
          :nudge-right="30">
          <v-toolbar-title slot="activator">
            <div class="avatar avatar--small pipe-tooltipped pipe-tooltipped--w"
                       :style="`background-image: url(${$store.state.avatarURL}?imageView2/1/w/64/h/64/interlace/1/q/100)`"
                       :aria-label="$store.state.nickname || $store.state.name"></div>
          </v-toolbar-title>
          <v-list>
            <v-list-tile @click="switchBlog(item)"
                         v-if="$store.state.blogs.length > 1"
                         v-for="item in $store.state.blogs"
                         :key="item.id" class="list__tile--link">
              {{ item.title }}
            </v-list-tile>
            <v-list-tile
              @click="goHome"
              v-if="$route.path.indexOf('/admin') > -1">
              {{ $t('index', $store.state.locale) }}
            </v-list-tile>
            <v-list-tile
              @click="goAdmin"
              v-if="$route.path.indexOf('/admin') === -1 && $store.state.role !== 0 && $store.state.isInit">
                {{ $t('manage', $store.state.locale) }}
            </v-list-tile>
            <v-list-tile @click="logout">
              {{ $t('logout', $store.state.locale) }}
            </v-list-tile>
          </v-list>
        </v-menu>
      </template>
    </div>
  </header>
</template>

<script>
  import { initParticlesJS } from '~/plugins/utils'

  export default {
    props: {
      from: {
        type: String,
        required: true
      }
    },
    methods: {
      goAdmin () {
        this.$router.push('/admin')
        if (document.documentElement.clientWidth >= 768) {
          this.$store.commit('setBodySide', 'body--side')
        }
      },
      goHome () {
        this.$router.push('/')
        this.$store.commit('setBodySide', '')
        setTimeout(() => {
          initParticlesJS('particles')
        }, 200)
      },
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
        // TODO
        const responseData = await this.axios.post('/logout')
        if (responseData.code === 0) {
          window.location.href = 'https://hacpai.com/logout'
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

  .body--side
    .header
      .side__icon
        &-line:first-child
          margin-top: 0
          transform: none
          top: auto
          background-color: #fff
        &-line:last-child
          transform: none
          top: auto
          background-color: #fff
        &-line--middle
          opacity: 1

        &:hover
          .side__icon-line:first-child
            width: 50%
            transform: rotateZ(-45deg)
            top: 4px
          .side__icon-line--middle
            width: 90%
          .side__icon-line:last-child
            width: 50%
            transform: rotateZ(45deg)
            top: -4px
    .header__logo
      width: 240px

  .header
    position: fixed
    height: 60px
    width: 100%
    z-index: 20
    top: 0
    color: #fff
    .side__icon
      width: 20px
      height: 20px
      padding: 0 30px
      cursor: pointer

      &-line
        display: block
        position: relative
        vertical-align: top
        height: 3px
        width: 100%
        background: #fff
        margin-top: 5px
        transition-duration: .2s
        transition-timing-function: ease-in-out
        transition-delay: 0s
        border-radius: 1px
        opacity: 1
      &-line:first-child
        width: 100%
        transform: rotateZ(-45deg)
        top: 4px
        background-color: $theme-accent
      &-line--middle
        opacity: 0
      &-line:last-child
        width: 100%
        transform: rotateZ(45deg)
        top: -12px
        background-color: $theme-accent

    &__logo
      display: flex
      background-color: $white
      width: 0
      overflow: hidden
      align-items: center
      transition: $transition

      &-img
        height: 30px
        width: 30px
      a
        color: $text-title
        margin: 0 auto
        font-size: 18px
        &:hover
          text-decoration: none
      &--theme
        width: auto
        padding-left: 30px
        background-color: $blue
        a
          color: #fff
    &__nav
      background-color: $blue
      align-items: center
      padding-right: 30px
      a
        color: #fff

  @media (max-width: 768px)
    .header__nav
      padding-right: 15px
    .header__logo--theme
      padding-left: 15px
    .header__bar--icon .fn-none
      display: block
      float: left
    .header .side__icon
      padding: 0 15px
      margin-top: 6px
      &-line:first-child
        margin-top: 0
        transform: none
        top: auto
        background-color: #fff
      &-line:last-child
        transform: none
        top: auto
        background-color: #fff
      &-line--middle
        opacity: 1
    .body--side .header .side__icon,
    .body--side .header .side__icon:hover
      .side__icon-line:first-child
        width: 100%
        transform: rotateZ(-45deg)
        top: 6px
        background-color: $theme-accent
      .side__icon-line--middle
        opacity: 0
      .side__icon-line:last-child
        width: 100%
        transform: rotateZ(45deg)
        top: -10px
        background-color: $theme-accent
    .body--side .header__logo
      width: 0
</style>
