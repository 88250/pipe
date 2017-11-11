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
          @change="setLocalstorage('title')"
        ></v-text-field>

        <v-editor :height="300" v-model="content" @input="parseMarkdown"></v-editor>

        <v-select
          v-model="tags"
          :label="$t('tags', $store.state.locale)"
          chips
          tags
          :items="$store.state.tagsItems"
          required
          :rules="tagsRules"
          @change="setLocalstorage"
        ></v-select>

        <v-text-field
          :label="$t('links', $store.state.locale)"
          v-model="url"
          @change="setLocalstorage('url')"
        ></v-text-field>

        <label class="checkbox">
          <input
            type="checkbox"
            :checked="commentable"
            @change="setLocalstorage('commentable')"
            @click="commentable = !commentable"/><span
          class="checkbox__icon"></span>
          {{ $t('allowComment', $store.state.locale) }}
        </label>

        <label class="checkbox btn--space">
          <input
            type="checkbox"
            :checked="useThumbs"
            @change="setLocalstorage('useThumbs')"
            @click="useThumbs = !useThumbs"/><span
          class="checkbox__icon"></span>
          {{ $t('useThumb', $store.state.locale) }}
        </label>
      </v-form>
      <div class="alert alert--danger" v-show="error">
        <v-icon>danger</v-icon>
        <span>{{ errorMsg }}</span>
      </div>

      <div class="article-post__carousel" v-show="useThumbs">
        <v-carousel
          :cycle="false"
          icon="circle"
          left-control-icon="angle-left"
          right-control-icon="angle-right">
          <v-carousel-item v-for="(item,i) in thumbs" v-bind:src="item" :key="i"></v-carousel-item>
        </v-carousel>
      </div>

      <div class="fn-right">
        <v-btn @click="remove" class="btn--danger btn--margin-t30" v-if="$route.query.id">
          {{ $t('delete', $store.state.locale) }}
        </v-btn>
        <v-btn @click="edit($route.query.id)" class="btn--info btn--space btn--margin-t30" v-if="$route.query.id">
          {{ $t('edit', $store.state.locale) }}
        </v-btn>
        <v-btn @click="edit()" class="btn--info btn--margin-t30" v-else>{{ $t('publish', $store.state.locale)
          }}
        </v-btn>
      </div>
    </div>
  </div>
</template>

<script>
  import {required, maxSize} from '~/plugins/validate'
  import {asyncLoadScript} from '~/plugins/utils'

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
        tags: [],
        commentable: true,
        useThumbs: false,
        thumbs: ['', '', '', '', '', '']
      }
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t(this.$route.query.id ? 'editArticle' : 'postArticle', this.$store.state.locale)}`
      }
    },
    methods: {
      _paseMD (text, previewRef) {
        previewRef.innerHTML = `<div class="pipe-content__reset">${text}</div>`
        let hasMathJax = false
        let hasFlow = false
        if (text.indexOf('$\\') > -1 || text.indexOf('$$') > -1) {
          hasMathJax = true
        }

        if (text.indexOf('<code class="language-flow"') > -1) {
          hasFlow = true
        }

        if (hasMathJax) {
          if (typeof MathJax !== 'undefined') {
            window.MathJax.Hub.Typeset()
          } else {
            asyncLoadScript('https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.2/MathJax.js?config=TeX-MML-AM_CHTML',
              () => {
                window.MathJax.Hub.Config({
                  tex2jax: {
                    inlineMath: [['$', '$'], ['\\(', '\\)']],
                    displayMath: [['$$', '$$']],
                    processEscapes: true,
                    processEnvironments: true,
                    skipTags: ['pre', 'code', 'script']
                  }
                })
              })
          }
        }

        if (hasFlow) {
          const initFlow = function () {
            document.querySelectorAll('.pipe-content__reset .language-flow').forEach(function (it, index) {
              const id = 'pipeFlow' + (new Date()).getTime() + index
              it.style.display = 'none'
              const diagram = window.flowchart.parse(it.textContent)
              it.parentElement.outerHTML = '<div class="ft-center" id="' + id + '"></div>'
              diagram.drawSVG(id)
              document.getElementById(id).firstChild.style.height = 'auto'
              document.getElementById(id).firstChild.style.width = 'auto'
            })
          }

          if (typeof (flowchart) !== 'undefined') {
            initFlow()
          } else {
            asyncLoadScript('https://static.hacpai.com/js/lib/flowchart/flowchart.min.js', initFlow())
          }
        }
      },
      async parseMarkdown (value, previewRef) {
        this.setLocalstorage('content')
        if (previewRef) {
          const responseData = await this.axios.post('/console/markdown', {
            mdText: this.content
          })
          if (responseData.code === 0) {
            this._paseMD(responseData.data.html, previewRef)
            this.$set(this, 'error', false)
            this.$set(this, 'errorMsg', '')
          } else {
            this.$set(this, 'error', true)
            this.$set(this, 'errorMsg', responseData.msg)
          }
        }
      },
      setLocalstorage (type) {
        if (this.$route.query.id) {
          return
        }

        if (typeof arguments[0] === 'object') {
          localStorage.setItem('article-tags', arguments[0])
          return
        }
        switch (type) {
          case 'title':
            localStorage.setItem('article-title', this.title)
            break
          case 'url':
            localStorage.setItem('article-url', this.url)
            break
          case 'commentable':
            localStorage.setItem('article-commentable', this.commentable)
            break
          case 'useThumbs':
            localStorage.setItem('article-useThumbs', this.useThumbs)
            break
          case 'content':
            localStorage.setItem('article-content', this.content)
            break
          default:
            break
        }
      },
      async getThumbs () {
        const responseData = await this.axios.get(`console/thumbs?n=5&w=768&h=180`)
        if (responseData) {
          this.$set(this, 'thumbs', responseData)
        }
      },
      async edit (id) {
        if (!this.$refs.form.validate()) {
          return
        }

        let content = this.content
        if (this.useThumbs) {
          document.querySelectorAll('.carousel__item').forEach((item, index) => {
            if (item.style.display !== 'none') {
              content = `![](${this.thumbs[index]})\n\n` + content
            }
          })
        }
        const responseData = await this.axios[id ? 'put' : 'post'](`/console/articles${id ? '/' + id : ''}`, {
          title: this.title,
          content: content,
          path: this.url,
          tags: this.tags.toString(),
          commentable: this.commentable
        })
        if (responseData.code === 0) {
          if (!id) {
            localStorage.removeItem('article-title')
            localStorage.removeItem('article-content')
            localStorage.removeItem('article-tags')
            localStorage.removeItem('article-url')
            localStorage.removeItem('article-commentable')
            localStorage.removeItem('article-useThumbs')
          }
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$router.push('/admin/articles')
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      },
      async remove () {
        const responseData = await this.axios.delete(`/console/articles/${this.$route.query.id}`)
        if (responseData === null) {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('deleteSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
          this.$router.push('/admin/articles')
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
          this.$set(this, 'url', responseData.path)
          this.$set(this, 'tags', responseData.tags.split(','))
          this.$set(this, 'commentable', responseData.commentable)
        }
      } else {
        // set storage
        setTimeout(() => {
          if (localStorage.getItem('article-title')) {
            this.title = localStorage.getItem('article-title')
            this.$set(this, 'title', localStorage.getItem('article-title'))
          }
          if (localStorage.getItem('article-content')) {
            this.$set(this, 'content', localStorage.getItem('article-content'))
          }
          if (localStorage.getItem('article-tags')) {
            this.$set(this, 'tags', localStorage.getItem('article-tags').split(','))
          }
          if (localStorage.getItem('article-url')) {
            this.$set(this, 'url', localStorage.getItem('article-url'))
          }
          if (localStorage.getItem('article-commentable')) {
            this.$set(this, 'commentable', localStorage.getItem('article-commentable'))
          }
          if (localStorage.getItem('article-useThumbs')) {
            this.$set(this, 'useThumbs', localStorage.getItem('article-useThumbs'))
          }
        })
      }

      // get tags
      this.$store.dispatch('getTags')

      this.getThumbs()
    }
  }
</script>
<style lang="sass">
  .article-post__carousel
    margin: 0 auto
    width: 720px
</style>
