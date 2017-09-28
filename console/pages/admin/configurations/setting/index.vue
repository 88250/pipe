<template>
  <div>
    <button @click="changeLang">lang</button>
  </div>
</template>

<script>
  export default {
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('configuration', this.$store.state.locale)}`
      }
    },
    methods: {
      changeLang () {
        const locale = this.$store.state.locale === 'en_US' ? 'zh_CN' : 'en_US'
        this.$store.dispatch('setLocaleMessage', locale)
      }
    }
  }
</script>
