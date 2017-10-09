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
        :label="$t('blogPath', $store.state.locale)"
        v-model="blogPath"
      ></v-text-field>
      <v-text-field
        :label="$t('blogMember', $store.state.locale)"
        v-model="blogMembers"
      ></v-text-field>
      <v-text-field
        :label="$t('blogAdmin', $store.state.locale)"
        v-model="blogAdmin"
      ></v-text-field>
      <div class="alert alert--danger" v-show="error">
        <icon icon="danger"/>
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
    props: ['id'],
    data () {
      return {
        errorMsg: '',
        error: false,
        blogTitle: '',
        blogSubtitle: '',
        header: '',
        footer: '',
        metaKeywords: '',
        metaDescription: '',
        noticeBoard: '',
        blogPath: '',
        blogMembers: '',
        blogAdmin: ''
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
        if (this.id === '') {
          responseData = await this.axios.post('/console/blogs', {
            blogTitle: this.blogTitle,
            blogSubtitle: this.blogSubtitle,
            header: this.header,
            footer: this.footer,
            metaKeywords: this.metaKeywords,
            metaDescription: this.metaDescription,
            noticeBoard: this.noticeBoard,
            blogPath: this.blogPath,
            blogMembers: this.blogMembers,
            blogAdmin: this.blogAdmin
          })
        } else {
          responseData = await this.axios.put(`/console/blogs/${this.id}`, {
            blogTitle: this.blogTitle,
            blogSubtitle: this.blogSubtitle,
            header: this.header,
            footer: this.footer,
            metaKeywords: this.metaKeywords,
            metaDescription: this.metaDescription,
            noticeBoard: this.noticeBoard,
            blogPath: this.blogPath,
            blogMembers: this.blogMembers,
            blogAdmin: this.blogAdmin
          })
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
          this.$set(this, 'header', responseData.header)
          this.$set(this, 'footer', responseData.footer)
          this.$set(this, 'metaKeywords', responseData.metaKeywords)
          this.$set(this, 'metaDescription', responseData.metaDescription)
          this.$set(this, 'noticeBoard', responseData.noticeBoard)
          this.$set(this, 'blogPath', responseData.blogPath)
          this.$set(this, 'blogMembers', responseData.blogMembers)
          this.$set(this, 'blogAdmin', responseData.blogAdmin)
        }
      }
    },
    mounted () {
      this.init()
    }
  }
</script>
