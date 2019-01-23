<template>
  <div>
    <div class="card fn__clear card__body">

      <v-form ref="form">

        <v-select
          :label="$t('feedOutputMode', $store.state.locale)"
          v-model="feedOutputMode"
          :items="feedOutputModeItems"
          append-icon=""
        ></v-select>

        <div class="alert alert--danger" v-show="error">
          <v-icon>danger</v-icon>
          <span>{{ errorMsg }}</span>
        </div>
      </v-form>

      <v-btn class="fn__right btn--margin-t30 btn--info btn--space" @click="update">
        {{ $t('confirm', $store.state.locale) }}
      </v-btn>
    </div>
  </div>
</template>

<script>
  import { numberOnly } from '~/plugins/validate'

  export default {
    data () {
      return {
        requiredRules: [
          (v) => numberOnly.call(this, v)
        ],
        feedOutputMode: 0,
        feedOutputModeItems: [{
          'text': `${this.$t('abstract', this.$store.state.locale)}`,
          'value': 0
        }, {
          'text': `${this.$t('fullArticle', this.$store.state.locale)}`,
          'value': 1
        }],
        error: false,
        errorMsg: ''
      }
    },
    head () {
      return {
        title: `${this.$t('feed', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      async update () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.put('/console/settings/feed', {
          feedOutputMode: this.feedOutputMode
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
      const responseData = await this.axios.get('/console/settings/feed')
      if (responseData) {
        this.$set(this, 'feedOutputMode', responseData.feedOutputMode)
      }
    }
  }
</script>
