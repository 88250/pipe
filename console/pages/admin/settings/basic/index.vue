<template>
  <div>
    <div class="card fn-clear card__body">

      <v-form>
        <v-text-field
          :label="$t('blogTitle', $store.state.locale)"
          v-model="blogTitle"
        ></v-text-field>
        <v-text-field
          :label="$t('blogSubtitle', $store.state.locale)"
          v-model="blogSubtitle"
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
          label="HTML head"
          v-model="header"
        ></v-text-field>
        <v-text-field
          :label="$t('footer', $store.state.locale)"
          v-model="footer"
        ></v-text-field>
        <v-text-field
          :label="$t('noticeBoard', $store.state.locale)"
          v-model="noticeBoard"
        ></v-text-field>
        <label class="checkbox">
          <input type="checkbox" :checked="commentable" @click="commentable = !commentable"/><span
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
        blogTitle: '',
        blogSubtitle: '',
        faviconURL: '',
        header: '',
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
        title: `${this.$store.state.blogTitle} - ${this.$t('baseInfo', this.$store.state.locale)}`
      }
    },
    methods: {
      async update () {
        const responseData = await this.axios.put('/console/settings/basic', {
          blogTitle: this.blogTitle,
          blogSubtitle: this.blogSubtitle,
          faviconURL: this.faviconURL,
          header: this.header,
          footer: this.footer,
          metaKeywords: this.metaKeywords,
          metaDescription: this.metaDescription,
          noticeBoard: this.noticeBoard,
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
      const responseData = await this.axios.get('/console/settings/basic')
      if (responseData) {
        this.$set(this, 'blogTitle', responseData.blogTitle)
        this.$set(this, 'blogSubtitle', responseData.blogSubtitle)
        this.$set(this, 'faviconURL', responseData.faviconURL)
        this.$set(this, 'header', responseData.header)
        this.$set(this, 'footer', responseData.footer)
        this.$set(this, 'metaKeywords', responseData.metaKeywords)
        this.$set(this, 'metaDescription', responseData.metaDescription)
        this.$set(this, 'noticeBoard', responseData.noticeBoard)
        this.$set(this, 'commentable', responseData.commentable)
      }
    }
  }
</script>
