<template>
  <div>
    <button @click="changeLang">lang</button>
  </div>
</template>

<script>
  export default {
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t('baseInfo', this.$store.state.locale)}`
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
