<template>
  <div data-app="true">
    <solo-header from="admin"/>
    <side/>
    <div class="main">
      <div class="content">
        <nuxt/>
      </div>
      <solo-footer/>
    </div>
    <v-snackbar
      :top="true"
      v-model="snack"
      :timeout="$store.state.snackModify === 'error' ? 6000 : 3000"
      :success="$store.state.snackModify === 'success'"
    >
      {{ $store.state.snackMsg }}
      <span @click="snack = false"><icon icon="close"/></span>
    </v-snackbar>
  </div>
</template>

<script>
  import Side from '~/components/Side'
  import SoloHeader from '~/components/Header'
  import SoloFooter from '~/components/Footer'

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
    middleware: 'authenticated',
    components: {
      Side,
      SoloHeader,
      SoloFooter
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'
  .main
    padding-top: 60px
    margin-left: 240px

  .content
    background-color: $blue-lighter
    padding: 30px
    min-height: 650px
</style>
