<template>
  <div class="error">
    <pipe-header from="error"/>
    <div class="error__content">
      <div class="fn-clear fn-flex-1" v-if="error.statusCode === 404">
        <div class="fn-flex-1">
          <h1>404</h1>
          <div class="error__description">Page not found</div>
          <div class="fn-right">
            <nuxt-link to="/">{{ $t('index', $store.state.locale)}}</nuxt-link>
            |
            <a href="https://hacpai.com">{{ $t('hacpai', x$store.state.locale)}}</a>
          </div>
        </div>
      </div>
      <div class="fn-clear fn-flex-1" v-else>
        <div class="fn-flex-1">
          <h1>50X</h1>
          <div class="error__description">{{error.statusCode}}</div>
          <div class="fn-right">
            <a href="https://github.com/b3log/pipe/issues/new">{{ $t('reportIssue', $store.state.locale)}}</a> |
            <a href="https://hacpai.com">{{ $t('hacpai', $store.state.locale)}}</a>
          </div>
        </div>
      </div>
      <pipe-footer/>
    </div>
  </div>
</template>

<script>
  import PipeFooter from '~/components/Footer'
  import PipeHeader from '~/components/Header'

  export default {
    components: {
      PipeFooter,
      PipeHeader
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle ? this.$store.state.blogTitle + ' - ' : ''}${this.error.statusCode === 404 ? 404 : '50x'}`
      }
    },
    props: {
      error: {
        type: Object,
        required: true
      }
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'
  .error
    background-color: $blue-lighter
    height: 100%

    &__content
      height: 100%
      padding-top: 60px
      display: flex
      flex-direction: column
      box-sizing: border-box

      & > .fn-flex-1
        max-width: 630px
        width: 630px
        margin: 0 auto
        align-items: center
        display: flex

      h1
        font-size: 4.5rem
        font-weight: 300
        line-height: 1.1

    &__description
      margin: 50px 0
      font-size: 1.25rem
      font-weight: 300
  @media (max-width: 768px)
    .error__content > .fn-flex-1
      width: 100%
      padding: 0 15px
      box-sizing: border-box
</style>
