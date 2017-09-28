<template>
  <div class="card">
    <div class="card__body fn-clear">
      <v-form ref="form">
        <v-text-field
          :label="$t('title', $store.state.locale)"
          v-model="title"
          :rules="titleRules"
          :counter="10"
          required
        ></v-text-field>

        <mavon-editor v-model="content"/>

        <v-text-field
          :label="$t('tags', $store.state.locale)"
          v-model="tags"
          :rules="titleRules"
          :counter="10"
          required
        ></v-text-field>

        <mavon-editor v-model="abstract"/>

        <v-text-field
          :label="$t('links', $store.state.locale)"
          v-model="permalink"
        ></v-text-field>

        <v-text-field
          :label="$t('visitPassword', $store.state.locale)"
          v-model="password"
        ></v-text-field>

        <label class="checkbox">
          <input type="checkbox" :checked="commentable" @click="commentable = !commentable"/><span
          class="checkbox__icon"></span>
          {{ $t('allowComment', $store.state.locale) }}
        </label>
      </v-form>
      <div class="alert alert--danger" v-show="error">
        <icon icon="danger"/>
        <span>{{ errorMsg }}</span>
      </div>
      <div class="fn-right">
        <button @click="edit" class="btn btn--info btn--margin-t30" v-if="$route.query.id">
          {{ $t('edit', $store.state.locale) }}
        </button>
        <button @click="publish" class="btn btn--info btn--margin-t30" v-else>{{ $t('publish', $store.state.locale)
          }}
        </button>
      </div>
    </div>
  </div>
</template>

<script>
  import 'mavon-editor/dist/css/index.css'

  export default {
    data () {
      return {
        error: false,
        errorMsg: '',
        content: '',
        abstract: '',
        title: '',
        titleRules: [
          (v) => !!v || this.$t('required', this.$store.state.locale),
          (v) => v.length <= 32 || this.$t('validateRule2', this.$store.state.locale)
        ],
        permalink: '',
        password: '',
        tags: '',
        commentable: false
      }
    },
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t(this.$route.query.id ? 'editArticle' : 'postArticle', this.$store.state.locale)}`
      }
    },
    methods: {
      async edit () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.put(`/console/articles/${this.$route.query.id}`, {
          title: this.title,
          abstract: this.abstract,
          content: this.content,
          permalink: this.permalink,
          password: this.password,
          tags: this.tags,
          commentable: this.commentable
        })
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$router.push('/admin/articles/management')
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      },
      async publish () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.post(`/console/articles/`, {
          title: this.title,
          abstract: this.abstract,
          content: this.content,
          permalink: this.permalink,
          password: this.password,
          tags: this.tags,
          commentable: this.commentable
        })
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$router.push('/admin/articles/management')
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      }
    },
    async mounted () {
      const id = this.$route.query.id
      if (id) {
        const responseData = await this.axios.get(`/console/articles/${id}`)
        if (responseData) {
          this.$set(this, 'title', responseData.title)
          this.$set(this, 'abstract', responseData.abstract)
          this.$set(this, 'content', responseData.content)
          this.$set(this, 'permalink', responseData.permalink)
          this.$set(this, 'password', responseData.password)
          this.$set(this, 'tags', responseData.tags)
          this.$set(this, 'commentable', responseData.commentable)
        }
      }
    }
  }
</script>
