<template>
  <div class="card__body fn-clear">
    <v-form ref="form">
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
      <v-text-field
        :label="$t('blogURL', $store.state.locale)"
        v-model="blogURL"
      ></v-text-field>
      <v-text-field
        :label="$t('commonUser', $store.state.locale)"
        v-model="blogMembers"
      ></v-text-field>
      <v-text-field
        :label="$t('blogAdmin', $store.state.locale)"
        v-model="blogAdmin"
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
    <v-btn class="fn-right btn btn--margin-t30 btn--info btn--space" @click="created">
      {{ $t('confirm', $store.state.locale) }}
    </v-btn>
    <v-btn class="fn-right btn btn--margin-t30 btn--danger btn--space" @click="$emit('update:show', false)">
      {{ $t('cancel', $store.state.locale) }}
    </v-btn>
  </div>
</template>

<script>
  export default {
    props: {
      id: {
        type: String,
        required: true
      }
    },
    data () {
      return {
        errorMsg: '',
        error: false,
        blogTitle: '',
        blogSubtitle: '',
        faviconURL: '',
        header: '',
        footer: '',
        metaKeywords: '',
        metaDescription: '',
        noticeBoard: '',
        blogURL: '',
        blogMembers: '',
        blogAdmin: '',
        commentable: true
      }
    },
    watch: {
      id: function () {
        this.init()
      }
    },
    methods: {
      async created () {
        if (!this.$refs.form.validate()) {
          return
        }
        let responseData = {}
        const requestData = {
          blogTitle: this.blogTitle,
          blogSubtitle: this.blogSubtitle,
          faviconURL: this.faviconURL,
          header: this.header,
          footer: this.footer,
          metaKeywords: this.metaKeywords,
          metaDescription: this.metaDescription,
          noticeBoard: this.noticeBoard,
          blogURL: this.blogURL,
          blogMembers: this.blogMembers,
          blogAdmin: this.blogAdmin,
          commentable: this.commentable
        }
        if (this.id === '') {
          responseData = await this.axios.post('/console/blogs', requestData)
        } else {
          responseData = await this.axios.put(`/console/blogs/${this.id}`, requestData)
        }

        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$emit('addSuccess')
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      },
      async init () {
        if (this.id === '') {
          return
        }
        const responseData = await this.axios.get(`/console/blogs/${this.id}`)
        if (responseData) {
          this.$set(this, 'blogTitle', responseData.blogTitle)
          this.$set(this, 'blogSubtitle', responseData.blogSubtitle)
          this.$set(this, 'faviconURL', responseData.faviconURL)
          this.$set(this, 'header', responseData.header)
          this.$set(this, 'footer', responseData.footer)
          this.$set(this, 'metaKeywords', responseData.metaKeywords)
          this.$set(this, 'metaDescription', responseData.metaDescription)
          this.$set(this, 'noticeBoard', responseData.noticeBoard)
          this.$set(this, 'blogURL', responseData.blogURL)
          this.$set(this, 'blogMembers', responseData.blogMembers)
          this.$set(this, 'blogAdmin', responseData.blogAdmin)
          this.$set(this, 'commentable', responseData.commentable)
        }
      }
    },
    mounted () {
      this.init()
    }
  }
</script>
