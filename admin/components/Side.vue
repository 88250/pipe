<template>
  <aside class="side">
    <nav>
      <v-list>
        <v-list-group v-for="item in items" :value="item.active">
          <v-list-tile ripple slot="item" @click=goLink(item.link)>
            <v-list-tile-content>
              <v-list-tile-title>
                <icon :icon="item.icon"></icon>
                {{ item.title }}
              </v-list-tile-title>
            </v-list-tile-content>
            <v-list-tile-action v-if="item.items">
              <icon icon="angle-down"></icon>
            </v-list-tile-action>
          </v-list-tile>
          <v-list-tile ripple v-for="subItem in item.items" @click=goLink(subItem.link)>
              {{ subItem.title }}
          </v-list-tile>
        </v-list-group>
      </v-list>
    </nav>
  </aside>
</template>

<script>
  export default {
    data () {
      return {
        items: [
          {
            title: this.$t('home'),
            icon: 'home',
            link: '/'
          },
          {
            title: this.$t('postArticle'),
            icon: 'add',
            link: '/articles/post'
          },
          {
            title: this.$t('skinList'),
            icon: 'skin',
            link: '/skins'
          },
          {
            title: this.$t('manage'),
            icon: 'manage',
            active: true,
            items: [
              {
                title: this.$t('articleList'),
                link: '/articles/management'
              },
              {
                title: this.$t('commentList'),
                link: '/comments/management'
              },
              {
                title: this.$t('categoryList'),
                link: '/categories/management'
              },
              {
                title: this.$t('navigationList'),
                link: '/navigation/management'
              },
              {
                title: this.$t('linkList'),
                link: '/links/management'
              },
              {
                title: this.$t('userList'),
                link: '/users/management'
              },
              {
                title: this.$t('otherList'),
                link: '/others/management'
              }
            ]
          },
          {
            title: this.$t('setting'),
            icon: 'setting',
            items: [
              {
                title: this.$t('configuration'),
                link: '/configurations/setting'
              },
              {
                title: this.$t('signs'),
                link: '/signs/setting'
              },
              {
                title: this.$t('parameters'),
                link: '/parameters/setting'
              },
              {
                title: this.$t('configuration'),
                link: '/configurations/setting'
              },
              {
                title: this.$t('B3log'),
                link: '/b3log/setting'
              }
            ]
          },
          {
            title: this.$t('about'),
            icon: 'info',
            link: '/about'
          }
        ]
      }
    },
    methods: {
      goLink (link) {
        if (!link) {
          return
        }
        this.$router.push(link)
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
    top: 60px
    box-shadow: 1px 0px 20px rgba(0, 0, 0, 0.08)
    overflow: auto
    padding-bottom: 60px

  .side::-webkit-scrollbar
    display: none
</style>
