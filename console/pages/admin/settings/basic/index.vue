<template>
  <div>
    <div class="card fn-clear card__body">

      <v-form>
        <v-text-field
          :label="$t('basicBlogTitle', $store.state.locale)"
          v-model="basicBlogTitle"
        ></v-text-field>
        <v-text-field
          :label="$t('basicBlogSubtitle', $store.state.locale)"
          v-model="basicBlogSubtitle"
        ></v-text-field>
        <v-text-field
          :label="$t('basicFaviconURL', $store.state.locale)"
          v-model="basicFaviconURL"
        ></v-text-field>
        <v-text-field
          label="Meta Keywords"
          v-model="basicMetaKeywords"
        ></v-text-field>
        <v-text-field
          label="Meta Description"
          v-model="basicMetaDescription"
        ></v-text-field>
        <v-text-field
          label="HTML head"
          v-model="basicHeader"
        ></v-text-field>
        <v-text-field
          :label="$t('basicFooter', $store.state.locale)"
          v-model="basicFooter"
        ></v-text-field>
        <v-text-field
          :label="$t('basicNoticeBoard', $store.state.locale)"
          v-model="basicNoticeBoard"
        ></v-text-field>
        <label class="checkbox">
          <input type="checkbox" :checked="basicCommentable" @click="basicCommentable = !basicCommentable"/><span
          class="checkbox__icon"></span>
          {{ $t('allowComment', $store.state.locale) }}
        </label>

        <div class="alert alert--danger" v-show="error">
          <icon icon="danger"/>
          <span>{{ errorMsg }}</span>
        </div>
      </v-form>
      <v-btn class="fn-right btn btn--margin-t30 btn--info btn--space" @click="update">
        {{ $t('confirm', $store.state.locale) }}
      </v-btn>
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        basicBlogTitle: '',
        basicBlogSubtitle: '',
        basicFaviconURL: '',
        basicHeader: '',
        basicFooter: '',
        basicMetaKeywords: '',
        basicMetaDescription: '',
        basicNoticeBoard: '',
        basicCommentable: true,
        error: false,
        errorMsg: ''
      }
    },
    head () {
      return {
        title: `${this.$store.state.basicBlogTitle} - ${this.$t('baseInfo', this.$store.state.locale)}`
      }
    },
    methods: {
      async update () {
        const responseData = await this.axios.put('/console/settings/basic', {
          basicBlogTitle: this.basicBlogTitle,
          basicBlogSubtitle: this.basicBlogSubtitle,
          basicFaviconURL: this.basicFaviconURL,
          basicHeader: this.basicHeader,
          basicFooter: this.basicFooter,
          basicMetaKeywords: this.basicMetaKeywords,
          basicMetaDescription: this.basicMetaDescription,
          basicNoticeBoard: this.basicNoticeBoard,
          basicCommentable: this.basicCommentable
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
      const responseData = await this.axios.get('/console/settings/basic')
      if (responseData) {
        this.$set(this, 'basicBlogTitle', responseData.basicBlogTitle)
        this.$set(this, 'basicBlogTitle', responseData.basicBlogTitle)
        this.$set(this, 'basicFaviconURL', responseData.basicFaviconURL)
        this.$set(this, 'basicHeader', responseData.basicHeader)
        this.$set(this, 'basicFooter', responseData.basicFooter)
        this.$set(this, 'basicMetaKeywords', responseData.basicMetaKeywords)
        this.$set(this, 'basicMetaDescription', responseData.basicMetaDescription)
        this.$set(this, 'basicNoticeBoard', responseData.basicNoticeBoard)
        this.$set(this, 'basicCommentable', responseData.basicCommentable)
      }
    }
  }
</script>
