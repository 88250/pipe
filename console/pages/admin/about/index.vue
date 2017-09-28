<template>
  <div class="card">
    <div class="card__body">
      <h1>{{ $t('about', $store.state.locale)}}</h1>
    </div>
    {{ isLatest }}
  </div>
</template>

<script>
  export default {
    data () {
      return {
        isLatest: true
      }
    },
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('about', this.$store.state.locale)}`
      }
    },
    async mounted () {
      const responseData = await this.axios.get(`/console/version/latest`)
      if (responseData) {
        this.$set(this, 'isLatest', this.$store.state.version === responseData.version)
      }
    }
  }
</script>

<style lang="sass">
  .admin__index
    li
      line-height: 20px
      a
        display: block
</style>
