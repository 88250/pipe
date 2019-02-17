<template>
  <div>
    <div class="card fn__clear card__body">

      <v-form ref="form">
        <v-select
          :label="$t('articleListStyle', $store.state.locale)"
          v-model="preferenceArticleListStyle"
          :items="preferenceArticleListStyleItems"
          append-icon=""
        ></v-select>
        <v-text-field
          :label="$t('mostUseTagListSize', $store.state.locale)"
          v-model="preferenceMostUseTagListSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('recentCommentListSize', $store.state.locale)"
          v-model="preferenceRecentCommentListSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('mostCommentArticleListSize', $store.state.locale)"
          v-model="preferenceMostCommentArticleListSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('mostViewArticleListSize', $store.state.locale)"
          v-model="preferenceMostViewArticleListSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('articleListPageSize', $store.state.locale)"
          v-model="preferenceArticleListPageSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('articleListWindowSize', $store.state.locale)"
          v-model="preferenceArticleListWindowSize"
          required
          :rules="requiredRules"
        ></v-text-field>
        <v-text-field
          :label="$t('recommendArticleListSize', $store.state.locale)"
          v-model="preferenceRecommendArticleListSize"
          required
          :rules="requiredRules"
        ></v-text-field>

        <div class="alert alert--danger" v-show="error">
          <v-icon>danger</v-icon>
          <span>{{ errorMsg }}</span>
        </div>
      </v-form>
      <v-btn class="fn__right btn--margin-t30 btn--info btn--space" @click="update">
        {{ $t('confirm', $store.state.locale) }}
      </v-btn>
    </div>
  </div>
</template>

<script>
  import { numberOnly } from '~/plugins/validate'

  export default {
    data () {
      return {
        requiredRules: [
          (v) => numberOnly.call(this, v)
        ],
        preferenceArticleListStyle: '1',
        preferenceArticleListStyleItems: [{
          'text': this.$t('title', this.$store.state.locale),
          'value': '0'
        }, {
          'text': `${this.$t('title', this.$store.state.locale)}+${this.$t('abstract', this.$store.state.locale)}`,
          'value': '1'
        }, {
          'text': `${this.$t('title', this.$store.state.locale)}+${this.$t('content', this.$store.state.locale)}`,
          'value': '2'
        }],
        preferenceMostUseTagListSize: 10,
        preferenceRecentCommentListSize: 10,
        preferenceMostCommentArticleListSize: 10,
        preferenceMostViewArticleListSize: 10,
        preferenceArticleListPageSize: 15,
        preferenceArticleListWindowSize: 20,
        preferenceRecommendArticleListSize: 0,
        error: false,
        errorMsg: ''
      }
    },
    head () {
      return {
        title: `${this.$t('preference', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      async update () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.put('/console/settings/preference', {
          preferenceArticleListStyle: this.preferenceArticleListStyle,
          preferenceMostUseTagListSize: this.preferenceMostUseTagListSize,
          preferenceRecentCommentListSize: this.preferenceRecentCommentListSize,
          preferenceMostCommentArticleListSize: this.preferenceMostCommentArticleListSize,
          preferenceMostViewArticleListSize: this.preferenceMostViewArticleListSize,
          preferenceArticleListPageSize: this.preferenceArticleListPageSize,
          preferenceArticleListWindowSize: this.preferenceArticleListWindowSize,
          preferenceRecommendArticleListSize: this.preferenceRecommendArticleListSize
        })

        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('setupSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      }
    },
    async mounted () {
      const responseData = await this.axios.get('/console/settings/preference')
      if (responseData) {
        this.$set(this, 'preferenceArticleListStyle', responseData.preferenceArticleListStyle)
        this.$set(this, 'preferenceMostUseTagListSize', responseData.preferenceMostUseTagListSize)
        this.$set(this, 'preferenceRecentCommentListSize', responseData.preferenceRecentCommentListSize)
        this.$set(this, 'preferenceMostCommentArticleListSize', responseData.preferenceMostCommentArticleListSize)
        this.$set(this, 'preferenceMostViewArticleListSize', responseData.preferenceMostViewArticleListSize)
        this.$set(this, 'preferenceArticleListPageSize', responseData.preferenceArticleListPageSize)
        this.$set(this, 'preferenceArticleListWindowSize', responseData.preferenceArticleListWindowSize)
        this.$set(this, 'preferenceRecommendArticleListSize', responseData.preferenceRecommendArticleListSize)
      }
    }
  }
</script>
