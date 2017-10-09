<template>
  <aside class="side">
    <nav>
      <v-list>
        <v-list-group
          v-for="item in items"
          :value="item.active"
          :key="item.title"
          v-if="$store.state.role <= item.role">
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
          <v-list-tile
            ripple
            v-for="subItem in item.items"
            :key="subItem.title"
            v-if="$store.state.role <= subItem.role">
            <nuxt-link :to="subItem.link">{{ subItem.title }}</nuxt-link>
          </v-list-tile>
        </v-list-group>
      </v-list>
    </nav>
  </aside>
</template>

<script>
  const genNavData = (app, locale) => [
    {
      title: app.$t('home', locale),
      icon: 'home',
      link: '/admin/',
      role: 2
    },
    {
      title: app.$t('postArticle', locale),
      icon: 'add',
      link: '/admin/articles/post',
      role: 2
    },
    {
      title: app.$t('themeList', locale),
      icon: 'theme',
      link: '/admin/themes',
      role: 1
    },
    {
      title: app.$t('manage', locale),
      icon: 'manage',
      active: app.$route.path.indexOf('management') > -1,
      role: 2,
      items: [
        {
          title: app.$t('articleList', locale),
          link: '/admin/articles/management',
          role: 2
        },
        {
          title: app.$t('commentList', locale),
          link: '/admin/comments/management',
          role: 2
        },
        {
          title: app.$t('categoryList', locale),
          link: '/admin/categories/management',
          role: 1
        },
        {
          title: app.$t('navigationList', locale),
          link: '/admin/navigations/management',
          role: 1
        },
        {
          title: app.$t('userList', locale),
          link: '/admin/users/management',
          role: 1
        },
        {
          title: app.$t('blogManage', locale),
          link: '/admin/blogs/management',
          role: 0
        }
      ]
    },
    {
      title: app.$t('setting', locale),
      icon: 'setting',
      active: app.$route.path.indexOf('settings') > -1,
      role: 1,
      items: [
        {
          title: app.$t('baseInfo', locale),
          link: '/admin/settings/basic',
          role: 1
        },
        {
          title: app.$t('preference', locale),
          link: '/admin/settings/preference',
          role: 1
        },
        {
          title: app.$t('signs', locale),
          link: '/admin/settings/sign',
          role: 1
        },
        {
          title: app.$t('internationalization', locale),
          link: '/admin/settings/internationalization',
          role: 1
        },
        {
          title: app.$t('feed', locale),
          link: '/admin/settings/feed',
          role: 1
        }
      ]
    },
    {
      title: app.$t('others', locale),
      icon: 'inbox',
      link: '/admin/others/management',
      role: 2
    },
    {
      title: app.$t('about', locale),
      icon: 'info',
      link: '/admin/about',
      role: 2
    }
  ]

  export default {
    watch: {
      '$store.state.locale': function (val) {
        this.$set(this, 'items', genNavData(this, val))
      }
    },
    data () {
      return {
        items: genNavData(this, this.$store.state.locale)
      }
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'

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
