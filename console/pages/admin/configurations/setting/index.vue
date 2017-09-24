<template>
  <div>
    <button @click="changeLang">lang</button>
  </div>
</template>

<script>
  export default {
    layout: 'admin',
    methods: {
      changeLang () {
        const locale = this.$i18n.locale === 'en_US' ? 'zh_CN' : 'en_US'
        this.$store.commit('setBaseInfo', { lang: locale })
        this.$i18n.locale = locale
      }
    }
  }
</script>

<style>
</style>
