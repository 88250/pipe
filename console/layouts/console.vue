<template>
  <div class="console__panel">
    <pipe-header from="console"/>
    <div class="console__content">
      <nuxt/>
      <pipe-footer/>
    </div>
    <v-snackbar
      :top="true"
      v-model="snack"
      :timeout="$store.state.snackModify === 'error' ? 6000 : 3000"
      :color="$store.state.snackModify === 'success' ? 'snack--info' : 'snack--error'"
    >
      {{ $store.state.snackMsg }}
      <span @click="snack = false"><v-icon>close</v-icon></span>
    </v-snackbar>
  </div>
</template>

<script>
  import PipeFooter from '~/components/Footer'
  import PipeHeader from '~/components/Header'

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
      PipeFooter,
      PipeHeader
    }
  }
</script>

<style lang="sass">
  .console
    &__panel
      height: 100%
    &__content
      height: 100%
      box-sizing: border-box
      display: flex
      padding-top: 60px
      flex-direction: column
</style>
