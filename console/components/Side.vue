<template>
  <aside class="side">
    <nav>{{ $t('configuration') }}
      <v-list>
        <v-list-group v-for="item in items" :value="item.active" :key="item.title">
          <v-list-tile ripple slot="item">
            <nuxt-link :to="item.link" v-if="!item.items">
              <icon :icon="item.icon"></icon>
              {{ item.title }}
            </nuxt-link>
            <v-list-tile-content ripple v-else>
              <icon :icon="item.icon"></icon>
              {{ item.title }}
            </v-list-tile-content>
            <v-list-tile-action v-if="item.items">
              <icon icon="angle-down"></icon>
            </v-list-tile-action>
          </v-list-tile>
          <v-list-tile ripple v-for="subItem in item.items" :key="subItem.title">
            <nuxt-link :to="subItem.link">{{ subItem.title }}</nuxt-link>
          </v-list-tile>
        </v-list-group>
      </v-list>
    </nav>
  </aside>
</template>

<script>
  export default {
    data () {
      console.log(this.$store.state.locale)
      return {
        items: [
          {
            title: this.$t('home', 'en_US'),
            icon: 'home',
            link: '/admin/'
          },
          {
            title: this.$t('postArticle', 'zh_CN'),
            icon: 'add',
            link: '/admin/articles/post'
          },
          {
            title: this.$t('skinList', this.$store.state.locale),
            icon: 'skin',
            link: '/admin/skins'
          },
          {
            title: this.$t('manage', this.$store.state.locale),
            icon: 'manage',
            active: true,
            items: [
              {
                title: this.$t('articleList', this.$store.state.locale),
                link: '/admin/articles/management'
              },
              {
                title: this.$t('commentList', this.$store.state.locale),
                link: '/admin/comments/management'
              },
              {
                title: this.$t('categoryList', this.$store.state.locale),
                link: '/admin/categories/management'
              },
              {
                title: this.$t('navigationList', this.$store.state.locale),
                link: '/admin/navigation/management'
              },
              {
                title: this.$t('linkList', this.$store.state.locale),
                link: '/admin/links/management'
              },
              {
                title: this.$t('userList', this.$store.state.locale),
                link: '/admin/users/management'
              },
              {
                title: this.$t('others', this.$store.state.locale),
                link: '/admin/others/management'
              }
            ]
          },
          {
            title: this.$t('setting', this.$store.state.locale),
            icon: 'setting',
            items: [
              {
                title: this.$t('configuration', this.$store.state.locale),
                link: '/admin/configurations/setting'
              },
              {
                title: this.$t('signs', this.$store.state.locale),
                link: '/admin/signs/setting'
              },
              {
                title: this.$t('parameters', this.$store.state.locale),
                link: '/admin/parameters/setting'
              },
              {
                title: this.$t('upload', this.$store.state.locale),
                link: '/admin/upload/setting'
              },
              {
                title: 'B3log',
                link: '/admin/b3log/setting'
              }
            ]
          },
          {
            title: this.$t('about', this.$store.state.locale),
            icon: 'info',
            link: '/admin/about'
          }
        ]
      }
    }
  }
</script>

<style lang="sass">
  @import "~assets/scss/_variables.scss"

  .side
    width: 240px
    background-color: $white
    position: fixed
    height: 100%
    top: 0
    box-shadow: 1px 0px 20px rgba(0, 0, 0, 0.08)
    overflow: auto
    padding-top: 60px
    box-sizing: border-box

    .list
      transition: height 0.4s cubic-bezier(0.4, 0, 0.2, 1)
      list-style-type: none

      a
        display: flex
        height: 44px
        text-decoration: none
        align-items: center
        padding: 0 15px
        margin: 0
        transition: 0.3s cubic-bezier(0.25, 0.8, 0.5, 1)
        position: relative
        user-select: none
        color: #607d8b
        border-left: 3px solid transparent
        width: 100%
        box-sizing: border-box

        &:hover
          color: $blue

        &.nuxt-link-exact-active
          border-left-color: $blue
          color: $blue

      .list--group a
        padding-left: 43px

        &.nuxt-link-exact-active
          border-left-color: transparent

      .icon
        margin-right: 10px

      .list--group__header
        height: 44px
        line-height: 44px
        align-items: center
        display: flex
        cursor: pointer

        li
          width: 100%

          .list__tile
            display: flex

        .list__tile__content
          padding-left: 15px
          border-left: 3px solid transparent
          flex: 1

        .list__tile__action

      .list--group__header--active .list__tile__content
        border-left-color: $blue
        color: $blue

  .side::-webkit-scrollbar
    display: none
</style>
