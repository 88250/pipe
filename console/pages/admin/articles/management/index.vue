<template>
  <div class="card fn-clear">
    <ul class="list">
      <li v-for="item in list" class="fn-flex">
        <div class="fn-flex-1">
          <div class="list__title">
            <nuxt-link to="/">React Native 系列【五】View 滑动时禁止 WebView 滚动</nuxt-link>
          </div>
          <div class="list__meta">
            <nuxt-link to="/">Webview</nuxt-link>&nbsp;
            <nuxt-link to="/">gesture</nuxt-link>&nbsp;
            <nuxt-link to="/">React</nuxt-link>&nbsp;
            <a>18 评论</a> •
            900 浏览 •
            <time>2017-10-23</time>
          </div>
        </div>
        <div>
          <button class="btn btn--info btn--small">{{ $t('edit', $store.state.locale)}}</button>
          <button class="btn btn--danger btn--space btn--small">{{ $t('delete', $store.state.locale)}}</button>
          <button class="btn btn--success btn--space btn--small">{{ $t('top', $store.state.locale)}}</button>
        </div>
      </li>
    </ul>
    <v-pagination
      class="fn-right"
      :length="49"
      v-model="page"
      circle
      total-visible="10"
      next-icon=">"
      prev-icon="<"
    ></v-pagination>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        page: 1,
        list: []
      }
    },
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('articleList', this.$store.state.locale)}`
      }
    },
    async mounted () {
      const responseData = await this.axios('/console/articles')
      if (responseData) {
        this.$set(this, 'list', responseData.articles)
      }
    }
  }
</script>

<style>
</style>
