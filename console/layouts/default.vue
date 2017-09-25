<template>
  <div>
    <solo-header/>
    <nuxt class="content"/>
    <solo-footer/>
  </div>
</template>

<script>
  import SoloFooter from '~/components/Footer'
  import SoloHeader from '~/components/Header'

  export default {
    components: {
      SoloFooter,
      SoloHeader
    }
  }
</script>

<style lang="sass">
  .content
    margin-top: 60px
</style>
