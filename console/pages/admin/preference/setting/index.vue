<template>
  <div>
    <div class="card fn-clear card__body">

      <v-form ref="form">
        <v-select
          :label="$t('language', $store.state.locale)"
          :items="localeItems"
          v-model="locale"
          append-icon=""
        ></v-select>
        <v-text-field
          :label="$t('timeZone', $store.state.locale)"
          v-model="timeZone"
          readonly
        ></v-text-field>
        <v-select
          :label="$t('articleListStyle', $store.state.locale)"
          v-model="articleListStyle"
          :items="articleListStyleItems"
          append-icon=""
        ></v-select>
        <v-text-field
          :label="$t('mostUseTagListSize', $store.state.locale)"
          v-model="mostUseTagListSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('recentCommentListSize', $store.state.locale)"
          v-model="recentCommentListSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('mostCommentArticleListSize', $store.state.locale)"
          v-model="mostCommentArticleListSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('mostViewArticleListSize', $store.state.locale)"
          v-model="mostViewArticleListSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('articleListPageSize', $store.state.locale)"
          v-model="articleListPageSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('articleListWindowSize', $store.state.locale)"
          v-model="articleListWindowSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('randomArticleListSize', $store.state.locale)"
          v-model="randomArticleListSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('relevantArticleListSize', $store.state.locale)"
          v-model="relevantArticleListSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('externalRelevantArticleListSize', $store.state.locale)"
          v-model="externalRelevantArticleListSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <label class="checkbox">
          <input type="checkbox"
                 :checked="enableArticleUpdateHint"
                 @click="enableArticleUpdateHint = !enableArticleUpdateHint"/><span
          class="checkbox__icon"></span>
          {{ $t('enableArticleUpdateHint', $store.state.locale) }}
        </label> <br/>
        <label class="checkbox">
          <input type="checkbox" :checked="commentable" @click="commentable = !commentable"/><span
          class="checkbox__icon"></span>
          {{ $t('allowComment', $store.state.locale) }}
        </label>
        <v-select
          :label="$t('feedOutputMode', $store.state.locale)"
          v-model="feedOutputMode"
          :items="feedOutputModeItems"
          append-icon=""
        ></v-select>
        <v-text-field
          :label="$t('feedOutputCnt', $store.state.locale)"
          v-model="feedOutputCnt"
          required
          :rules="requiredRules"
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
        requiredRules: [
          (v) => /^\d+$/.test(v) || this.$t('validateRule3', this.$store.state.locale)
        ],
        locale: this.$store.state.locale,
        localeItems: [{
          'text': '简体中文',
          'value': 'zh_CN'
        }, {
          'text': 'English(US)',
          'value': 'en_US'
        }],
        timeZone: 'Asia/Shanghai',
        articleListStyle: 'title',
        articleListStyleItems: [{
          'text': this.$t('title', this.$store.state.locale),
          'value': 'title'
        }, {
          'text': `${this.$t('title', this.$store.state.locale)}+${this.$t('abstract', this.$store.state.locale)}`,
          'value': 'titleAbstract'
        }, {
          'text': `${this.$t('title', this.$store.state.locale)}+${this.$t('content', this.$store.state.locale)}`,
          'value': 'titleContent'
        }],
        mostUseTagListSize: 10,
        recentCommentListSize: 10,
        mostCommentArticleListSize: 10,
        mostViewArticleListSize: 10,
        articleListPageSize: 15,
        articleListWindowSize: 20,
        randomArticleListSize: 10,
        relevantArticleListSize: 10,
        externalRelevantArticleListSize: 10,
        enableArticleUpdateHint: true,
        commentable: true,
        feedOutputMode: 'abstract',
        feedOutputModeItems: [{
          'text': `${this.$t('abstract', this.$store.state.locale)}`,
          'value': 'abstract'
        }, {
          'text': `${this.$t('fullArticle', this.$store.state.locale)}`,
          'value': 'full'
        }],
        feedOutputCnt: 10,
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
      async update () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.post('/console/categories', {
          locale: this.locale,
          timeZone: this.timeZone,
          articleListStyle: this.articleListStyle,
          mostUseTagListSize: this.mostUseTagListSize,
          recentCommentListSize: this.recentCommentListSize,
          mostCommentArticleListSize: this.mostCommentArticleListSize,
          mostViewArticleListSize: this.mostViewArticleListSize,
          articleListPageSize: this.articleListPageSize,
          articleListWindowSize: this.articleListWindowSize,
          randomArticleListSize: this.randomArticleListSize,
          relevantArticleListSize: this.relevantArticleListSize,
          externalRelevantArticleListSize: this.externalRelevantArticleListSize,
          enableArticleUpdateHint: this.enableArticleUpdateHint,
          feedOutputMode: this.feedOutputMode,
          feedOutputCnt: this.feedOutputCnt,
          commentable: this.commentable
        })

        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('setupSuccess', this.$store.state.locale),
            snackModify: 'success'
          })

          this.$store.dispatch('setLocaleMessage', this.locale)
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      }
    },
    async mounted () {
      const responseData = await this.axios.get('/console/preference')
      if (responseData) {
        this.$set(this, 'locale', responseData.locale)
        this.$set(this, 'timeZone', responseData.timeZone)
        this.$set(this, 'articleListStyle', responseData.articleListStyle)
        this.$set(this, 'mostUseTagListSize', responseData.mostUseTagListSize)
        this.$set(this, 'recentCommentListSize', responseData.recentCommentListSize)
        this.$set(this, 'mostCommentArticleListSize', responseData.mostCommentArticleListSize)
        this.$set(this, 'mostViewArticleListSize', responseData.mostViewArticleListSize)
        this.$set(this, 'articleListPageSize', responseData.articleListPageSize)
        this.$set(this, 'articleListWindowSize', responseData.articleListWindowSize)
        this.$set(this, 'randomArticleListSize', responseData.randomArticleListSize)
        this.$set(this, 'relevantArticleListSize', responseData.relevantArticleListSize)
        this.$set(this, 'externalRelevantArticleListSize', responseData.externalRelevantArticleListSize)
        this.$set(this, 'enableArticleUpdateHint', responseData.enableArticleUpdateHint)
        this.$set(this, 'feedOutputMode', responseData.feedOutputMode)
        this.$set(this, 'feedOutputCnt', responseData.feedOutputCnt)
        this.$set(this, 'commentable', responseData.commentable)
      }
    }
  }
</script>
