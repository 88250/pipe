<template>
  <div class="default-content">
    welcome
  </div>
</template>

<script>
  export default {
    head () {
      return {
        title: `${this.$store.state.userName}`
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
