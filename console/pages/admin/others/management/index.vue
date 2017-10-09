<template>
  <div>
    <div class="card fn-clear card__body">
export, import, b3log key, rm unused tag
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        locale: '',
        timeZone: '',
        articleListStyle: '',
        mostUseTagListSize: '',
        recentCommentListSize: '',
        mostCommentArticleListSize: '',
        mostViewArticleListSize: '',
        articleListPageSize: '',
        articleListWindowSize: '',
        randomArticleListSize: '',
        relevantArticleListSize: '',
        externalRelevantArticleListSize: '',
        enableArticleUpdateHint: '',
        feedOutputMode: '',
        feedOutputCnt: '',
        error: false,
        errorMsg: ''
      }
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t('preference', this.$store.state.locale)}`
      }
    },
    methods: {
      update (id) {
        this.$set(this, 'showForm', true)
        this.$set(this, 'editId', id)
      }
    },
    mounted () {
      // get
    }
  }
</script>
