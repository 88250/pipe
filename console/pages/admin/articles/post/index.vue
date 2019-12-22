<template>
  <div class="card">
    <div class="card__body fn__clear">
      <v-form ref="form">
        <v-text-field
          :label="$t('title', $store.state.locale)"
          v-model="title"
          :rules="titleRules"
          :counter="128"
          required
          @change="setLocalstorage('title')"
        ></v-text-field>

        <div id="contentEditor" style="height: 480px;background-color: #f6f8fa"></div>

        <v-select
          v-model="tags"
          :label="$t('tags', $store.state.locale)"
          chips
          tags
          :items="$store.state.tagsItems"
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

        <div id="abstractEditor" style="height: 160px;background-color: #f6f8fa"></div>

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
        <label class="checkbox btn--space">
          <input
            type="checkbox"
            :checked="syncToCommunity"
            @change="setLocalstorage('syncToCommunity')"
            @click="syncToCommunity = !syncToCommunity"/><span
          class="checkbox__icon"></span>
          {{ $t('syncToCommunity', $store.state.locale) }}
          (<a href="https://hacpai.com/article/1546941897596" target="_blank">?</a>)
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

      <div class="fn__right">
        <v-btn @click="remove" class="btn--danger btn--margin-t30" v-if="$route.query.id">
          {{ $t('delete', $store.state.locale) }}
        </v-btn>
        <v-btn @click="edit($route.query.id)" class="btn--info btn--space btn--margin-t30" v-if="$route.query.id">
          {{ $t('submit', $store.state.locale) }}
        </v-btn>
        <v-btn @click="edit()" class="btn--info btn--margin-t30" v-else>{{ $t('publish', $store.state.locale) }}
        </v-btn>
      </div>
    </div>
  </div>
</template>

<script>
  import { required, maxSize } from '~/plugins/validate'
  import { LazyLoadImage } from '~/plugins/utils'
  import Vditor from 'vditor'

  export default {
    data () {
      return {
        tokenURL: {
          URL: '',
          token: '',
        },
        error: false,
        errorMsg: '',
        content: '',
        originalContent: '',
        title: '',
        titleRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 128),
        ],
        linkRules: [
          (v) => maxSize.call(this, v, 255),
        ],
        timeRules: [
          (v) => (v.length === 0 ||
            /^(((20[0-3][0-9]-(0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|(20[0-3][0-9]-(0[2469]|11)-(0[1-9]|[12][0-9]|30))) (20|21|22|23|[0-1][0-9]):[0-5][0-9]:[0-5][0-9])$/.test(
              v)) || this.$t('createdTime', this.$store.state.locale) + '[YYYY-MM-DD HH:mm:ss]',
        ],
        url: '',
        time: '',
        abstractEditor: '',
        contentEditor: '',
        tags: [],
        commentable: true,
        useThumbs: false,
        topped: false,
        syncToCommunity: true,
        thumbs: ['', '', '', '', '', ''],
        edited: false,
      }
    },
    head () {
      return {
        title: `${this.$t(this.$route.query.id ? 'editArticle' : 'postArticle',
          this.$store.state.locale)} - ${this.$store.state.blogTitle}`,
      }
    },
    watch: {
      '$route.query.id': function (newValue, oldValue) {
        if (typeof newValue === 'undefined' && oldValue) {
          this._setDefaultLocalStorage()
          this.abstractEditor.enableCache()
          this.contentEditor.enableCache()
        } else {
          this.abstractEditor.disabledCache()
          this.contentEditor.disabledCache()
        }
      },
    },
    beforeRouteUpdate (to, from, next) {
      if (from.query.id && from.path === '/admin/articles/post' &&
        (this.edited || this.originalContent !== this.contentEditor.getValue())) {
        if (confirm(this.$t('isGoTo', this.$store.state.locale))) {
          next()
        }
      } else {
        next()
      }
    },
    beforeRouteLeave (to, from, next) {
      if (from.query.id && from.path === '/admin/articles/post' &&
        (this.edited || this.originalContent !== this.contentEditor.getValue())) {
        if (confirm(this.$t('isGoTo', this.$store.state.locale))) {
          next()
        }
      } else {
        next()
      }
    },
    methods: {
      _initEditor (data) {
        return new Vditor(data.id, {
          typewriterMode: true,
          tab: '\t',
          cache: this.$route.query.id ? false : true,
          hint: {
            emojiPath: 'https://cdn.jsdelivr.net/npm/vditor/dist/images/emoji',
          },
          preview: {
            delay: 500,
            mode: data.mode,
            url: `${process.env.Server}/api/console/markdown`,
            parse: (element) => {
              if (element.style.display === 'none') {
                return
              }
              LazyLoadImage()
              Vditor.highlightRender({
                style: 'github',
                enable: false,
              }, document)
            },
          },
          upload: {
            max: 10 * 1024 * 1024,
            url: this.tokenURL.URL,
            token: this.tokenURL.token,
            filename: name => name.replace(/[^(a-zA-Z0-9\u4e00-\u9fa5\.)]/g, '').
              replace(/[\?\\/:|<>\*\[\]\(\)\$%\{\}@~]/g, '').
              replace('/\\s/g', ''),
          },
          height: data.height,
          counter: 102400,
          resize: {
            enable: data.resize,
          },
          lang: this.$store.state.locale,
          placeholder: data.placeholder,
        })
      },
      setLocalstorage (type) {
        if (type !== 'content') {
          this.$set(this, 'edited', true)
        }

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
          case 'commentable':
            localStorage.setItem('article-commentable', this.commentable)
            break
          case 'useThumbs':
            localStorage.setItem('article-useThumbs', this.useThumbs)
            break
          case 'syncToCommunity':
            localStorage.setItem('article-syncToCommunity', this.syncToCommunity)
            break
          case 'topped':
            localStorage.setItem('article-topped', this.topped)
            break
          default:
            break
        }
      },
      async getThumbs () {
        const responseData = await this.axios.get('console/thumbs?n=5&w=768&h=432')
        if (responseData) {
          this.$set(this, 'thumbs', responseData)
        }
      },
      async edit (id) {
        if (!this.$refs.form.validate()) {
          return
        }
        if (this.contentEditor.getValue().length > 102400) {
          this.contentEditor.focus()
          return
        }
        if (this.abstractEditor.getValue().length > 102400) {
          this.abstractEditor.focus()
          return
        }

        let content = this.contentEditor.getValue()
        if (this.useThumbs) {
          document.querySelectorAll('.carousel__item').forEach((item, index) => {
            if (item.style.display !== 'none') {
              content = `![](${this.thumbs[index].replace('imageView2/1/w/768/h/180/interlace/1/q/100',
                'imageView2/1/w/960/h/540/interlace/1/q/100')})\n\n` + content
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
          abstract: this.abstractEditor.getValue(),
          syncToCommunity: this.syncToCommunity,
          time: this.time === '' ? '' : this.time.replace(' ', 'T') + '+08:00',
        })
        if (responseData.code === 0) {
          if (!id) {
            localStorage.removeItem('article-title')
            localStorage.removeItem('article-tags')
            localStorage.removeItem('article-url')
            localStorage.removeItem('article-time')
            localStorage.removeItem('article-commentable')
            localStorage.removeItem('article-useThumbs')
            localStorage.removeItem('article-syncToCommunity')
            localStorage.removeItem('article-topped')
            this.contentEditor.setValue('')
            this.abstractEditor.setValue('')
          }
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          if (id) {
            this.$set(this, 'edited', false)
            this.$set(this, 'originalContent', this.contentEditor.getValue())
          }
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
            snackModify: 'success',
          })
          this.$router.push('/admin/articles')
          this.contentEditor.setValue('')
          this.abstractEditor.setValue('')
        }
      },
      _setDefaultLocalStorage () {
        if (localStorage.getItem('article-title')) {
          this.title = localStorage.getItem('article-title')
          this.$set(this, 'title', localStorage.getItem('article-title'))
        } else {
          this.$set(this, 'title', '')
        }
        if (localStorage.getItem('vditorcontentEditor')) {
          this.$set(this, 'originalContent', localStorage.getItem('vditorcontentEditor'))
          this.contentEditor.setValue(localStorage.getItem('vditorcontentEditor'))
        } else {
          this.$set(this, 'originalContent', '')
          this.contentEditor.setValue('')
        }
        if (localStorage.getItem('vditorabstractEditor')) {
          this.abstractEditor.setValue(localStorage.getItem('vditorabstractEditor'))
        } else {
          this.abstractEditor.setValue('')
        }
        if (localStorage.getItem('article-tags')) {
          this.$set(this, 'tags', localStorage.getItem('article-tags').split(','))
        } else {
          this.$set(this, 'tags', [])
        }
        if (localStorage.getItem('article-url')) {
          this.$set(this, 'url', localStorage.getItem('article-url'))
        } else {
          this.$set(this, 'url', '')
        }
        if (localStorage.getItem('article-time')) {
          this.$set(this, 'time', localStorage.getItem('article-time'))
        } else {
          this.$set(this, 'time', '')
        }
        if (localStorage.getItem('article-commentable')) {
          this.$set(this, 'commentable', localStorage.getItem('article-commentable') === 'true')
        } else {
          this.$set(this, 'commentable', true)
        }
        if (localStorage.getItem('article-useThumbs')) {
          this.$set(this, 'useThumbs', localStorage.getItem('article-useThumbs') === 'true')
        } else {
          this.$set(this, 'useThumbs', false)
        }
        if (localStorage.getItem('article-syncToCommunity')) {
          this.$set(this, 'syncToCommunity', localStorage.getItem('article-syncToCommunity') === 'true')
        } else {
          this.$set(this, 'syncToCommunity', true)
        }
        if (localStorage.getItem('article-topped')) {
          this.$set(this, 'topped', localStorage.getItem('article-topped') === 'true')
        } else {
          this.$set(this, 'topped', false)
        }
        setTimeout(() => {
          document.querySelector('.input-group__input input').focus()
        }, 100)
      },
    },
    async mounted () {
      const responseData = await this.axios.get('console/upload/token')
      if (responseData) {
        this.$set(this, 'tokenURL', {
          token: responseData.uploadToken || '',
          URL: responseData.uploadURL || '',
        })
      }

      this.contentEditor = this._initEditor({
        id: 'contentEditor',
        mode: 'both',
        height: 480,
        placeholder: this.$t('inputContent', this.$store.state.locale),
        resize: false,
      })

      this.abstractEditor = this._initEditor({
        id: 'abstractEditor',
        height: 160,
        mode: 'editor',
        placeholder: this.$t('inputAbstract', this.$store.state.locale),
        resize: true,
      })

      const id = this.$route.query.id

      window.onbeforeunload = (event) => {
        if ((this.edited || this.originalContent !== this.contentEditor.getValue()) && id) {
          if (event) {
            event.returnValue = this.$t('isGoTo', this.$store.state.locale)
          }
          return this.$t('isGoTo', this.$store.state.locale)
        }
      }

      if (id) {
        const responseData = await this.axios.get(`/console/articles/${id}`)
        if (responseData) {
          this.$set(this, 'title', responseData.title)
          this.$set(this, 'originalContent', responseData.content)
          this.$set(this, 'url', responseData.path)
          this.$set(this, 'time', responseData.time.replace('T', ' ').substr(0, 19))
          this.$set(this, 'tags', responseData.tags.split(','))
          this.$set(this, 'commentable', responseData.commentable)
          this.$set(this, 'topped', responseData.topped)
          this.abstractEditor.setValue(responseData.abstract)
          this.contentEditor.setValue(responseData.content)
        }
      } else {
        // set storage
        this._setDefaultLocalStorage()
        setTimeout(() => {
          document.querySelector('.input-group__input input').focus()
        }, 100)
      }
      // get tags
      this.$store.dispatch('getTags')

      this.getThumbs()
    },
  }
</script>
<style lang="sass">
  .article-post__carousel
    margin: 0 auto
    width: 768px
    position: relative
    .carousel
      height: 432px
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

    .vditor-reset img
    cursor: auto
</style>
