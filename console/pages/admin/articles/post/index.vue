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

        <v-editor
          :uploadURL="`/upload`"
          :uploadMax="10"
          :height="300"
          :value="content"
          :label="label"
          :fetchUpload="fetchUpload"
          @change="parseMarkdown"></v-editor>

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
          :rules="linkRules"
          :counter="255"
          @keyup="setLocalstorage('url')"
        ></v-text-field>

        <v-text-field
          :label="$t('createdTime', $store.state.locale) + '[YYYY-MM-DD HH:mm:ss]'"
          v-model="time"
          :rules="timeRules"
          @keyup="setLocalstorage('time')"
        ></v-text-field>

        <v-text-field
          :label="$t('abstract', $store.state.locale)"
          v-model="abstract"
          multi-line
          @keyup="setLocalstorage('abstract')"
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

        <label class="checkbox btn--space">
          <input
            type="checkbox"
            :checked="topped"
            @change="setLocalstorage('topped')"
            @click="topped = !topped"/><span
          class="checkbox__icon"></span>
          {{ $t('top', $store.state.locale) }}
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
        <span
          :aria-label="$t('refresh', $store.state.locale)"
          class="article-post__carousel-refresh pipe-tooltipped pipe-tooltipped--n"
          @click="getThumbs"><v-icon>refresh</v-icon></span>
      </div>

      <div class="fn-right">
        <v-btn @click="remove" class="btn--danger btn--margin-t30" v-if="$route.query.id">
          {{ $t('delete', $store.state.locale) }}
        </v-btn>
        <v-btn @click="edit($route.query.id)" class="btn--info btn--space btn--margin-t30" v-if="$route.query.id">
          {{ $t('submit', $store.state.locale) }}
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
  import {asyncLoadScript, LazyLoadImage} from '~/plugins/utils'

  export default {
    data () {
      return {
        label: {
          loading: this.$t('uploading', this.$store.state.locale),
          error: this.$t('uploadError', this.$store.state.locale),
          over: this.$t('uploadOver', this.$store.state.locale),
          emoji: this.$t('emoji', this.$store.state.locale) + ' <ctrl+/>',
          bold: this.$t('bold', this.$store.state.locale) + ' <ctrl+b>',
          italic: this.$t('italic', this.$store.state.locale) + ' <ctrl+i>',
          quote: this.$t('quote', this.$store.state.locale) + ' <ctrl+e>',
          link: this.$t('link', this.$store.state.locale) + ' <ctrl+k>',
          upload: this.$t('upload', this.$store.state.locale),
          unorderedList: this.$t('unorderedList', this.$store.state.locale) + ' <ctrl+l>',
          orderedList: this.$t('orderedList', this.$store.state.locale) + ' <ctrl+shift+k>',
          preview: this.$t('preview', this.$store.state.locale) + ' <ctrl+d>',
          fullscreen: this.$t('fullscreen', this.$store.state.locale) + ' <ctrl+shift+a>',
          help: this.$t('question', this.$store.state.locale)
        },
        error: false,
        errorMsg: '',
        content: '',
        title: '',
        titleRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 128)
        ],
        linkRules: [
          (v) => maxSize.call(this, v, 255)
        ],
        tagsRules: [
          (v) => this.tags.length > 0 || this.$t('required', this.$store.state.locale)
        ],
        timeRules: [
          (v) => (v.length === 0 || /^(((20[0-3][0-9]-(0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|(20[0-3][0-9]-(0[2469]|11)-(0[1-9]|[12][0-9]|30))) (20|21|22|23|[0-1][0-9]):[0-5][0-9]:[0-5][0-9])$/.test(v)) || this.$t('createdTime', this.$store.state.locale) + '[YYYY-MM-DD HH:mm:ss]'
        ],
        url: '',
        time: '',
        abstract: '',
        tags: [],
        commentable: true,
        useThumbs: false,
        topped: false,
        thumbs: ['', '', '', '', '', '']
      }
    },
    head () {
      return {
        title: `${this.$t(this.$route.query.id ? 'editArticle' : 'postArticle', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      async fetchUpload (url, succCB) {
        const responseData = await this.axios.post(`${this.$store.state.blogURL}/fetch-upload`, {
          url
        })
        if (responseData.code === 0) {
          succCB(responseData.data.originalURL, responseData.data.url)
        } else {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: responseData.msg
          })
        }
      },
      _paseMD (text, previewRef) {
        previewRef.innerHTML = `<div class="pipe-content__reset">${text}</div>`
        LazyLoadImage()
        let hasMathJax = false
        let hasFlow = false
        if (text.split('$').length > 2 ||
          (text.split('\\(').length > 1 &&
            text.split('\\)').length > 1)) {
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
              it.parentElement.outerHTML = `<div class="ft-center" id="${id}"></div>`
              diagram.drawSVG(id)
              document.getElementById(id).firstChild.style.height = 'auto'
              document.getElementById(id).firstChild.style.width = 'auto'
            })
          }

          if (typeof (flowchart) !== 'undefined') {
            initFlow()
          } else {
            asyncLoadScript((process.env.StaticServer || process.env.Server) + '/theme/js/lib/flowchart.min.js', initFlow)
          }
        }
      },
      async parseMarkdown (value, previewRef) {
        this.$set(this, 'content', value)
        this.setLocalstorage('content')
        if (previewRef) {
          if (value.replace(/(^\s*)|(\s*)$/g, '') === '') {
            this._paseMD('', previewRef)
            return
          }
          const responseData = await this.axios.post('/console/markdown', {
            mdText: value
          })
          if (responseData.code === 0) {
            this._paseMD(responseData.data.html, previewRef)
          } else {
            this.$store.commit('setSnackBar', {
              snackBar: true,
              snackMsg: responseData.msg
            })
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
          case 'time':
            localStorage.setItem('article-time', this.time)
            break
          case 'abstract':
            localStorage.setItem('article-abstract', this.abstract)
            break
          case 'commentable':
            localStorage.setItem('article-commentable', this.commentable)
            break
          case 'useThumbs':
            localStorage.setItem('article-useThumbs', this.useThumbs)
            break
          case 'topped':
            localStorage.setItem('article-topped', this.topped)
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
              content = `![](${this.thumbs[index].replace('imageView2/1/w/768/h/180/interlace/1/q/100',
                'imageView2/1/w/960/h/520/interlace/1/q/100')})\n\n` + content
            }
          })
        }
        const responseData = await this.axios[id ? 'put' : 'post'](`/console/articles${id ? '/' + id : ''}`, {
          title: this.title,
          content: content,
          path: this.url,
          tags: this.tags.toString(),
          commentable: this.commentable,
          topped: this.topped,
          abstract: this.abstract,
          time: this.time === '' ? '' : this.time.replace(' ', 'T') + '+08:00'
        })
        if (responseData.code === 0) {
          if (!id) {
            localStorage.removeItem('article-title')
            localStorage.removeItem('article-content')
            localStorage.removeItem('article-tags')
            localStorage.removeItem('article-url')
            localStorage.removeItem('article-time')
            localStorage.removeItem('article-abstract')
            localStorage.removeItem('article-commentable')
            localStorage.removeItem('article-useThumbs')
            localStorage.removeItem('article-topped')
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
        if (!confirm(this.$t('confirmDelete', this.$store.state.locale))) {
          return
        }
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
          this.$set(this, 'time', responseData.time.replace('T', ' ').substr(0, 19))
          this.$set(this, 'abstract', responseData.abstract)
          this.$set(this, 'tags', responseData.tags.split(','))
          this.$set(this, 'commentable', responseData.commentable)
          this.$set(this, 'topped', responseData.topped)
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
          if (localStorage.getItem('article-time')) {
            this.$set(this, 'time', localStorage.getItem('article-time'))
          }
          if (localStorage.getItem('article-abstract')) {
            this.$set(this, 'abstract', localStorage.getItem('article-abstract'))
          }
          if (localStorage.getItem('article-commentable')) {
            this.$set(this, 'commentable', localStorage.getItem('article-commentable') === 'true')
          }
          if (localStorage.getItem('article-useThumbs')) {
            this.$set(this, 'useThumbs', localStorage.getItem('article-useThumbs') === 'true')
          }
          if (localStorage.getItem('article-topped')) {
            this.$set(this, 'topped', localStorage.getItem('article-topped') === 'true')
          }

          this.parseMarkdown(this.content)
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
    position: relative
    &-refresh
      position: absolute
      right: 15px
      bottom: 15px
      cursor: pointer
      z-index: 10
      svg
        color: #fff
        height: 20px
        width: 20px

    .pipe-content__reset img
    cursor: auto
</style>
