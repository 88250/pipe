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
        title: `${this.$store.state.blogTitle}`
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
