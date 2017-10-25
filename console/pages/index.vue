<template>
  <h1 class="fn-flex-1">
    welcome
  </h1>
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
