<template>
  <div class="default">
    <solo-header from="default"/>
    <nuxt class="default__content"/>
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
  .default
    &__content
      margin-top: 60px
</style>
