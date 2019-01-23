<template>
  <div>
    <div class="card fn__clear card__body">

      <v-form>
        <v-text-field
          :label="$t('blogURL', $store.state.locale)"
          v-model="blogURL"
        ></v-text-field>
        <v-text-field
          :label="$t('blogTitle', $store.state.locale)"
          v-model="blogTitle"
        ></v-text-field>
        <v-text-field
          :label="$t('blogSubtitle', $store.state.locale)"
          v-model="blogSubtitle"
        ></v-text-field>
        <v-text-field
          :label="$t('logoURL', $store.state.locale)"
          v-model="logoURL"
        ></v-text-field>
        <v-text-field
          :label="$t('faviconURL', $store.state.locale)"
          v-model="faviconURL"
        ></v-text-field>
        <v-text-field
          label="Meta Keywords"
          v-model="metaKeywords"
        ></v-text-field>
        <v-text-field
          label="Meta Description"
          v-model="metaDescription"
        ></v-text-field>
        <v-text-field
          multi-line
          :label="$t('footer', $store.state.locale)"
          v-model="footer"
        ></v-text-field>
        <v-text-field
          multi-line
          :label="$t('noticeBoard', $store.state.locale)"
          v-model="noticeBoard"
        ></v-text-field>
        <label class="checkbox">
          <input type="checkbox" :checked="commentable" @click="commentable = !commentable"/><span
          class="checkbox__icon"></span>
          {{ $t('allowComment', $store.state.locale) }}
        </label>

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
  export default {
    data () {
      return {
        blogURL: '',
        blogTitle: '',
        blogSubtitle: '',
        logoURL: '',
        faviconURL: '',
        footer: '',
        metaKeywords: '',
        metaDescription: '',
        noticeBoard: '',
        commentable: true,
        error: false,
        errorMsg: ''
      }
    },
    head () {
      return {
        title: `${this.$t('baseInfo', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      async update () {
        const responseData = await this.axios.put('/console/settings/basic', {
          basicBlogURL: this.blogURL,
          basicBlogTitle: this.blogTitle,
          basicBlogSubtitle: this.blogSubtitle,
          basicLogoURL: this.logoURL,
          basicFaviconURL: this.faviconURL,
          basicFooter: this.footer,
          basicMetaKeywords: this.metaKeywords,
          basicMetaDescription: this.metaDescription,
          basicNoticeBoard: this.noticeBoard,
          basicCommentable: this.commentable
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
      const responseData = await this.axios.get('/console/settings/basic')
      if (responseData) {
        this.$set(this, 'blogURL', responseData.basicBlogURL)
        this.$set(this, 'blogTitle', responseData.basicBlogTitle)
        this.$set(this, 'blogSubtitle', responseData.basicBlogSubtitle)
        this.$set(this, 'logoURL', responseData.basicLogoURL)
        this.$set(this, 'faviconURL', responseData.basicFaviconURL)
        this.$set(this, 'footer', responseData.basicFooter)
        this.$set(this, 'metaKeywords', responseData.basicMetaKeywords)
        this.$set(this, 'metaDescription', responseData.basicMetaDescription)
        this.$set(this, 'noticeBoard', responseData.basicNoticeBoard)
        this.$set(this, 'commentable', responseData.basicCommentable)
      }
    }
  }
</script>
