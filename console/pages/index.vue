<template>
  <div class="default-content">
    welcome
  </div>
</template>

<script>
  export default {
    layout: 'console',
    head () {
      return {
        title: this.$t('index', this.$store.state.locale)
      }
    },
    mounted () {
    }
  }
</script>

<style lang="sass">
  .default-content
    margin-top: 60px
</style>
