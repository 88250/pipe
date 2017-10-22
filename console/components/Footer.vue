<template>
  <footer class="footer">
    Powered by <a href="http://b3log.org" target="_blank">B3log</a> •
    <a href="https://hacpai.com/tag/Pipe" target="_blank">Pipe</a>
    {{ $store.state.version }}
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
