<template>
  <div>
    <div class="card fn-clear card__body">

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
        title: `${this.$store.state.blogTitle} - ${this.$t('internationalization', this.$store.state.locale)}`
      }
    },
    methods: {
      async update () {
        const responseData = await this.axios.put('/console/settings/i18n', {
          locale: this.locale,
          timezone: this.timezone
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
        this.$set(this, 'locale', responseData.locale)
        this.$set(this, 'timezone', responseData.timezone)
      }
    }
  }
</script>
