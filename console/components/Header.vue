<template>
  <header class="header fn-flex">
    <div class="header__logo">
      <img src="~static/images/logo.png"/>
      Solo
    </div>
    <div class="header__nav fn-flex-1">
      user logout
    </div>
  </header>
</template>

<style lang="sass">
  @import "~assets/scss/_variables.scss"

  .header
    position: fixed
    height: 60px
    width: 100%
    z-index: 10

    &__logo
      background-color: $white
      width: 240px

    &__nav
      background-color: $blue
</style>
