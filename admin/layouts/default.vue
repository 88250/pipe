<template>
  <div>
    <solo-header/>
    <side/>
    <nuxt/>
  </div>
</template>

<script>
  import Side from '~/components/Side'
  import SoloHeader from '~/components/Header'

  export default {
    components: {
      Side,
      SoloHeader
    }
  }
</script>

<style>
</style>
