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

        <div class="alert alert--danger" v-show="error">
          <icon icon="danger"/>
          <span>{{ errorMsg }}</span>
        </div>
      </v-form>
      <button class="fn-right btn btn--margin-t30 btn--info btn--space" @click="update">
        {{ $t('confirm', $store.state.locale) }}
      </button>
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        requiredRules: [
          (v) => /^\d+$/.test(v) || this.$t('validateRule3', this.$store.state.locale)
        ],
        blogTitle: '',
        blogSubtitle: '',
        header: '',
        footer: '',
        metaKeywords: '',
        metaDescription: '',
        noticeBoard: '',
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
        const responseData = await this.axios.post('/console/configurations', {
          blogTitle: this.blogTitle,
          blogSubtitle: this.blogSubtitle,
          header: this.header,
          footer: this.footer,
          metaKeywords: this.metaKeywords,
          metaDescription: this.metaDescription,
          noticeBoard: this.noticeBoard
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
      const responseData = await this.axios.get('/console/configurations')
      if (responseData) {
        this.$set(this, 'blogTitle', responseData.blogTitle)
        this.$set(this, 'blogSubtitle', responseData.blogSubtitle)
        this.$set(this, 'header', responseData.header)
        this.$set(this, 'footer', responseData.footer)
        this.$set(this, 'metaKeywords', responseData.metaKeywords)
        this.$set(this, 'metaDescription', responseData.metaDescription)
        this.$set(this, 'noticeBoard', responseData.noticeBoard)
      }
    }
  }
</script>
