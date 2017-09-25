<template>
  <div class="module">
    Post
  </div>
</template>

<script>
  export default {
    layout: 'admin',
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('articleList', this.$store.state.locale)}`
      }
    }
  }
</script>

<style>
</style>
