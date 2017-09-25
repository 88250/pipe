<template>
  <div class="module">
    社区推荐
  </div>
</template>

<script>
  export default {
    layout: 'admin',
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('home', this.$store.state.locale)}`
      }
    },
    mounted () {
    }
  }
</script>

<style>
</style>
