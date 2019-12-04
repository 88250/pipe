<template>
  <footer class="footer">
    Powered by <a href="https://b3log.org/" target="_blank">B3log 开源</a> •
    <a href="https://github.com/88250/pipe" target="_blank">Pipe</a> {{$store.state.version}}
  </footer>
</template>

<style lang="sass">
  @import '~assets/scss/_variables'

  .footer
    padding: 17px 15px
    border-top: 1px solid $border-color
    font-size: 16px
    background-color: $white
</style>
