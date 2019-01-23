<template>
  <aside class="side fn__flex">
    <nav>
      <v-list v-for="item in $store.state.menu" :key="item.title">
        <v-list-group
          :value="item.active"
          v-if="$store.state.role <= item.role">
          <v-list-tile
            @click=""
            ripple
            slot="item">
            <nuxt-link :to="item.link" v-if="!item.items">
              <v-icon>{{item.icon}}</v-icon>
              {{ item.title }}
            </nuxt-link>
            <template v-else>
              <v-list-tile-content>
                <v-icon>{{item.icon}}</v-icon>
                {{ item.title }}
              </v-list-tile-content>
              <v-list-tile-action>
                <v-icon>angle-down</v-icon>
              </v-list-tile-action>
            </template>
          </v-list-tile>
          <v-list-tile
            ripple
            v-for="subItem in item.items"
            :key="subItem.title"
            @click=""
            v-if="$store.state.role <= subItem.role">
            <nuxt-link :to="subItem.link">{{ subItem.title }}</nuxt-link>
          </v-list-tile>
        </v-list-group>
      </v-list>
    </nav>
    <div class="side__mobile fn__flex-1" @click="closeSide"></div>
  </aside>
</template>

<script>
  import {genMenuData} from '~/plugins/utils'

  export default {
    methods: {
      closeSide () {
        document.querySelector('#pipe').className = ''
      }
    },
    mounted () {
      this.$store.commit('setMenu', genMenuData(this, this.$store.state.locale))
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'

  .body--side .side
    left: 0

  .side
    left: -250px
    position: fixed
    height: 100%
    top: 0
    padding-top: 60px
    box-sizing: border-box
    z-index: 9
    transition: $transition

    nav
      width: 240px
      background-color: $white
      box-shadow: 1px 0px 20px rgba(0, 0, 0, 0.08)
      overflow: auto
      &::-webkit-scrollbar
        display: none

    .list
      padding: 0

      .list__tile a.nuxt-link-exact-active
        border-left-color: $blue

      .list--group .list__tile a
        padding-left: 30px

      .list__tile
        padding: 0
        line-height: 50px
        a,
        .list__tile__content
          border-left: 3px solid transparent
          color: #607d8b
          padding: 0 15px
          width: 100%
          box-sizing: border-box
          display: flex
          align-items: center
          .icon
            margin-right: 15px
          &:hover,
          &.nuxt-link-exact-active
            color: $blue
            text-decoration: none
            .icon
              color: $blue

  .side::-webkit-scrollbar
    display: none

  @media (max-width: 768px)
    .side
      width: 100%
      left: -100%
    .side__mobile
      background-color: $fade
      display: none
    .body--side .side__mobile
      display: block
</style>
