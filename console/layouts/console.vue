<template>
  <div class="default">
    <pipe-header from="default"/>
    <div class="default__content">
      <nuxt/>
    </div>
    <pipe-footer/>
    <v-snackbar
      :top="true"
      v-model="snack"
      :timeout="$store.state.snackModify === 'error' ? 6000 : 3000"
      :success="$store.state.snackModify === 'success'"
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
  .default
    &__content
      margin-top: 60px
</style>
