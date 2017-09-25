<template>
  <div>
    login
  </div>
</template>

<script>
  export default {
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('login', this.$store.state.locale)}`
      }
    },
    mounted () {
    }
  }
</script>
