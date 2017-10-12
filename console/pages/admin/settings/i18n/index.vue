<template>
  <div>
    <div class="card fn-clear card__body">

      <v-form>
        <v-select
          :label="$t('i18nLocale', $store.state.locale)"
          :items="localeItems"
          v-model="i18nLocale"
          append-icon=""
        ></v-select>
        <v-text-field
          :label="$t('i18nTimezone', $store.state.locale)"
          v-model="i18nTimezone"
          readonly
        ></v-text-field>
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
        i18nLocale: this.$store.state.locale,
        localeItems: [{
          'text': '简体中文',
          'value': 'zh_CN'
        }, {
          'text': 'English(US)',
          'value': 'en_US'
        }],
        i18nTimezone: 'Asia/Shanghai',
        error: false,
        errorMsg: ''
      }
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t('internationalization', this.$store.state.locale)}`
      }
    },
    methods: {
      async update () {
        const responseData = await this.axios.put('/console/settings/i18n', {
          i18nLocale: this.i18nLocale,
          i18nTimezone: this.i18nTimezone
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
      const responseData = await this.axios.get('/console/settings/i18n')
      if (responseData) {
        this.$set(this, 'i18nLocale', responseData.i18nLocale)
        this.$set(this, 'i18nTimezone', responseData.i18nTimezone)
      }
    }
  }
</script>
