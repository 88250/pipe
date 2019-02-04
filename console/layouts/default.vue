<template>
  <v-app :class="$store.state.bodySide" id="pipe">
    <pipe-header from="default"/>
    <side v-if="$route.path.indexOf('/admin') > -1"/>
    <div class="main">
      <div class="content" v-if="$route.path.indexOf('/admin') > -1">
        <nuxt/>
      </div>
      <template v-else>
        <nuxt/>
      </template>
      <pipe-footer/>
    </div>
    <v-snackbar
      :top="true"
      v-model="snack"
      :timeout="$store.state.snackModify === 'error' ? 6000 : 3000"
      :color="$store.state.snackModify === 'success' ? 'snack--info' : 'snack--error'"
    >
      {{ $store.state.snackMsg }}
      <span @click="snack = false"><v-icon>cancel</v-icon></span>
    </v-snackbar>
  </v-app>
</template>

<script>
  import Side from '~/components/Side'
  import PipeHeader from '~/components/Header'
  import PipeFooter from '~/components/Footer'

  export default {
    data () {
      return {
        snack: false
      }
    },
    watch: {
      '$store.state.snackBar': function (val) {
        this.$set(this, 'snack', val)
      },
      snack: function (val) {
        if (val === false) {
          this.$store.commit('setSnackBar', {
            snackBar: false,
            snackMsg: ''
          })
        }
      }
    },
    components: {
      Side,
      PipeHeader,
      PipeFooter
    },
    mounted () {
      if (document.documentElement.clientWidth < 721 || this.$route.path.indexOf('/admin') === -1) {
        this.$store.commit('setBodySide', '')
      } else {
        this.$store.commit('setBodySide', 'body--side')
      }
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'

  #__nuxt, #pipe, .main, #__layout
    height: 100%

  .main
    padding-top: 60px
    display: flex
    flex-direction: column
    transition: $transition
    box-sizing: border-box

  .body--side .main
    margin-left: 240px

  .content
    background-color: $blue-lighter
    padding: 30px
    flex: 1

  @media (max-width: 768px)
    .body--side .main
      margin-left: 0
    .content
      padding: 15px
</style>
