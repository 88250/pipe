<template>
  <div class="error">
    <pipe-header from="error"/>
    <div class="error__content fn-clear">
      <div v-if="error.statusCode === 404">
        <h1>404</h1>
        <div class="error__description">Page not found</div>
        <div class="fn-right">
          <nuxt-link to="/">{{ $t('index', $store.state.locale)}}</nuxt-link>
          |
          <a href="https://hacpai.com">{{ $t('hacpai', $store.state.locale)}}</a>
        </div>
      </div>
      <div v-else>
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

    &__content
      width: 600px
      margin: 60px auto 0
      padding: 80px 0

      h1
        font-size: 4.5rem
        font-weight: 300
        line-height: 1.1

    &__description
      margin: 50px 0
      font-size: 1.25rem
      font-weight: 300
</style>
