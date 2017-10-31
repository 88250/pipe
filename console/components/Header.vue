<template>
  <header class="header fn-flex">
    <div class="header__logo" v-if="from === 'admin'">
      <a :href="$store.state.blogURL">
        <img src="~static/images/logo.png"/>
        {{ $store.state.blogTitle }}
      </a>
    </div>
    <div class="header__nav fn-flex-1 fn-flex">
      <div class="header__bar--theme" v-if="$store.state.role === 0">
        <a class="btn--space btn--radius btn" href="https://hacpai.com/register">
          {{ $t('register', $store.state.locale) }}
        </a>
        <a href="https://hacpai.com/login">
          {{ $t('login', $store.state.locale) }}
        </a>
      </div>
      <template v-else>
        <span class="header__bar--icon fn-flex-1" v-if="$route.path.indexOf('/admin') > -1">
          <div class="side__icon" @click="toggleSide">
            <span class="side__icon-line"></span>
            <span class="side__icon-line side__icon-line--middle"></span>
            <span class="side__icon-line"></span>
          </div>
        </span>
        <div :class="$route.path.indexOf('/admin') == -1 ? 'header__bar--theme' : 'header__bar--admin'">
          <v-btn class="btn--small btn--danger btn--space" @click="logout">{{ $t('logout', $store.state.locale) }}
          </v-btn>
          <nuxt-link
            class="btn--space"
            v-if="$route.path.indexOf('/admin') === -1 && $store.state.role !== 0 && $store.state.role !== 4"
            to="/admin">
            {{ $t('manage', $store.state.locale) }}
          </nuxt-link>
          <v-menu
            class="btn--space"
            z-index="100"
            v-if="$route.path.indexOf('/admin') > -1 && $store.state.blogs.length > 1"
            :nudge-bottom="28"
            :open-on-hover="true">
            <v-toolbar-title slot="activator">
              <v-btn class="btn--small btn--success">
                {{ $store.state.blogTitle }}
                <v-icon>arrow_drop_down</v-icon>
              </v-btn>
            </v-toolbar-title>
            <v-list>
              <v-list-tile @click="switchBlog(item)"
                           v-for="item in $store.state.blogs"
                           :key="item.id" class="list__tile--link">
                {{ item.title }}
              </v-list-tile>
            </v-list>
          </v-menu>
          <div class="avatar avatar--small tooltipped tooltipped--w"
               :style="`background-image: url(${$store.state.avatarURL})`"
               :aria-label="$store.state.nickname || $store.state.name"></div>
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
    z-index: 10
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
        top: 6px
        background-color: $theme-accent
      &-line--middle
        opacity: 0
      &-line:last-child
        width: 100%
        transform: rotateZ(45deg)
        top: -10px
        background-color: $theme-accent

    &__logo
      display: flex
      background-color: $white
      width: 0
      overflow: hidden
      align-items: center
      transition: $transition
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
    &__bar--admin
      padding-right: 30px
      display: flex
      align-items: center
      flex-direction: row-reverse
    &__bar--theme
      width: 100%
      padding-right: 30px
      display: flex
      align-items: center
      flex-direction: row-reverse

  @media (max-width: 768px)
    .header__bar--theme,
    .header__bar--admin
      padding-right: 15px
    .header .side__icon
      padding: 0 15px
</style>
