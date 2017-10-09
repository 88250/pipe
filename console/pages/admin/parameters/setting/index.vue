<template>
  <div>
    <div class="card fn-clear card__body">

      <v-form ref="form">
        <v-text-field
          :label="$t('language', $store.state.locale)"
          v-model="locale"
          :counter="32"
          required
        ></v-text-field>
        <v-text-field
          :label="$t('timeZone', $store.state.locale)"
          v-model="timeZone"
          :counter="32"
          required
        ></v-text-field>
        <v-text-field
          :label="$t('articleListStyle', $store.state.locale)"
          v-model="articleListStyle"
          :counter="32"
        ></v-text-field>
        <v-text-field
          :label="$t('mostUseTagListSize', $store.state.locale)"
          v-model="mostUseTagListSize"
          :counter="32"
        ></v-text-field>
        <v-text-field
          :label="$t('recentCommentListSize', $store.state.locale)"
          v-model="recentCommentListSize"
          :counter="32"
        ></v-text-field>
        <v-text-field
          :label="$t('mostCommentArticleListSize', $store.state.locale)"
          v-model="mostCommentArticleListSize"
          :counter="32"
        ></v-text-field>
        <v-text-field
          :label="$t('mostViewArticleListSize', $store.state.locale)"
          v-model="mostViewArticleListSize"
          :counter="32"
        ></v-text-field>
        <v-text-field
          :label="$t('articleListPageSize', $store.state.locale)"
          v-model="articleListPageSize"
          :counter="32"
        ></v-text-field>
        <v-text-field
          :label="$t('articleListWindowSize', $store.state.locale)"
          v-model="articleListWindowSize"
          :counter="32"
        ></v-text-field>
        <v-text-field
          :label="$t('randomArticleListSize', $store.state.locale)"
          v-model="randomArticleListSize"
          :counter="32"
        ></v-text-field>
        <v-text-field
          :label="$t('relevantArticleListSize', $store.state.locale)"
          v-model="relevantArticleListSize"
          :counter="32"
        ></v-text-field>
        <v-text-field
          :label="$t('externalRelevantArticleListSize', $store.state.locale)"
          v-model="externalRelevantArticleListSize"
          :counter="32"
        ></v-text-field>
        <v-text-field
          :label="$t('enableArticleUpdateHint', $store.state.locale)"
          v-model="enableArticleUpdateHint"
          :counter="32"
        ></v-text-field>
        <v-text-field
          :label="$t('feedOutputMode', $store.state.locale)"
          v-model="feedOutputMode"
          :counter="32"
        ></v-text-field>
        <v-text-field
          :label="$t('feedOutputCnt', $store.state.locale)"
          v-model="feedOutputCnt"
          :counter="32"
        ></v-text-field>

        <div class="alert alert--danger" v-show="error">
          <icon icon="danger"/>
          <span>{{ errorMsg }}</span>
        </div>
      </v-form>
      <button class="fn-right btn btn--margin-t30 btn--info btn--space" @click="update">
        {{ $t('confirm', $store.state.locale) }}
      </button>
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
        title: `${this.$store.state.blogTitle} - ${this.$t('parameters', this.$store.state.locale)}`
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
