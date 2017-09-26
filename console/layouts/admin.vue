<template>
  <div>
    <solo-header/>
    <side/>
    <div class="main">
      <div class="content">
        <nuxt/>
      </div>
      <solo-footer/>
    </div>
  </div>
</template>

<script>
  import Side from '~/components/Side'
  import SoloHeader from '~/components/Header'
  import SoloFooter from '~/components/Footer'

  export default {
    middleware: 'authenticated',
    components: {
      Side,
      SoloHeader,
      SoloFooter
    }
  }
</script>

<style lang="sass">
  @import "~assets/scss/_variables.scss"
  .main
    padding-top: 60px
    margin-left: 240px

  .content
    background-color: $blue-lighter
    min-height: 360px
    padding: 30px
</style>
