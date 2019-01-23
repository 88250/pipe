<template>
  <div>
    <div class="card fn__clear card__body">

      <v-form>
        <v-select
          :label="$t('language', $store.state.locale)"
          :items="localeItems"
          v-model="locale"
          append-icon=""
        ></v-select>
        <v-text-field
          :label="$t('timezone', $store.state.locale)"
          v-model="timezone"
          readonly
        ></v-text-field>
      </v-form>

      <v-btn class="fn__right btn--margin-t30 btn--info btn--space" @click="update">
        {{ $t('confirm', $store.state.locale) }}
      </v-btn>
    </div>
  </div>
</template>

<script>
  import { genMenuData } from '~/plugins/utils'

  export default {
    data () {
      return {
        locale: this.$store.state.locale,
        localeItems: [{
          'text': '简体中文',
          'value': 'zh_CN'
        }, {
          'text': 'English(US)',
          'value': 'en_US'
        }],
        timezone: 'Asia/Shanghai',
        error: false,
        errorMsg: ''
      }
    },
    head () {
      return {
        title: `${this.$t('internationalization', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      async update () {
        const responseData = await this.axios.put('/console/settings/i18n', {
          i18nLocale: this.locale,
          i18nTimezone: this.timezone
        })

        if (responseData.code === 0) {
          this.$store.dispatch('setLocaleMessage', this.locale)
          this.$store.commit('setMenu', genMenuData(this, this.$store.state.locale))
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
        this.$set(this, 'locale', responseData.i18nLocale)
        this.$set(this, 'timezone', responseData.i18nTimezone)
      }
    }
  }
</script>
