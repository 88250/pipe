<template>
  <div class="card">
    <div class="card__body fn-clear">
      <v-form ref="form">
        <v-text-field
          :label="$t('title', $store.state.locale)"
          v-model="title"
          :rules="titleRules"
          :counter="128"
          required
        ></v-text-field>

        <mavon-editor v-model="content"/>

        <v-select
          v-model="tags"
          :label="$t('tags', $store.state.locale)"
          chips
          tags
          :items="tagsItems"
          required
          append-icon=""
          :rules="tagsRules"
        ></v-select>

        <v-text-field
          :label="$t('links', $store.state.locale)"
          v-model="url"
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
        <v-btn @click="edit" class="btn btn--info btn--margin-t30" v-if="$route.query.id">
          {{ $t('edit', $store.state.locale) }}
        </v-btn>
        <v-btn @click="publish" class="btn btn--info btn--margin-t30" v-else>{{ $t('publish', $store.state.locale)
          }}
        </v-btn>
      </div>
    </div>
  </div>
</template>

<script>
  import 'mavon-editor/dist/css/index.css'
  import { required, maxSize } from '~/plugins/validate'

  export default {
    data () {
      return {
        error: false,
        errorMsg: '',
        content: '',
        title: '',
        titleRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 128)
        ],
        tagsRules: [
          (v) => this.tags.length > 0 || this.$t('required', this.$store.state.locale)
        ],
        url: '',
        password: '',
        tags: [],
        tagsItems: ['sologo'],
        commentable: false
      }
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t(this.$route.query.id ? 'editArticle' : 'postArticle', this.$store.state.locale)}`
      }
    },
    methods: {
      async edit () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.put(`/console/articles/${this.$route.query.id}`, {
          title: this.title,
          content: this.content,
          url: this.url,
          password: this.password,
          tags: this.tags.toString(),
          commentable: this.commentable
        })
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$router.push('/admin/articles')
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
          content: this.content,
          url: this.url,
          password: this.password,
          tags: this.tags.toString(),
          commentable: this.commentable
        })
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$router.push('/admin/articles')
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
          this.$set(this, 'content', responseData.content)
          this.$set(this, 'url', responseData.url)
          this.$set(this, 'password', responseData.password)
          this.$set(this, 'tags', responseData.tags.split(','))
          this.$set(this, 'commentable', responseData.commentable)
        }
      }

      // get tags
      const tagResponseData = await this.axios.get('/console/tags/')
      if (tagResponseData) {
        let tagList = []
        tagResponseData.map((v) => {
          tagList.push(v.title)
        })
        this.$set(this, 'tagsItems', tagList)
      }
    }
  }
</script>
