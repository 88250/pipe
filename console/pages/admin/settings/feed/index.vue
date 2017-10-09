<template>
  <div>
    <div class="card fn-clear card__body">

      <v-form ref="form">

        <v-select
          :label="$t('feedOutputMode', $store.state.locale)"
          v-model="feedOutputMode"
          :items="feedOutputModeItems"
          append-icon=""
        ></v-select>
        <v-text-field
          :label="$t('feedOutputCnt', $store.state.locale)"
          v-model="feedOutputCnt"
          required
          :rules="requiredRules"
          @keyup.enter="update"
        ></v-text-field>

        <div class="alert alert--danger" v-show="error">
          <icon icon="danger"/>
          <span>{{ errorMsg }}</span>
        </div>
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
        requiredRules: [
          (v) => /^\d+$/.test(v) || this.$t('validateRule3', this.$store.state.locale)
        ],
        feedOutputMode: 'abstract',
        feedOutputModeItems: [{
          'text': `${this.$t('abstract', this.$store.state.locale)}`,
          'value': 'abstract'
        }, {
          'text': `${this.$t('fullArticle', this.$store.state.locale)}`,
          'value': 'full'
        }],
        feedOutputCnt: 10,
        error: false,
        errorMsg: ''
      }
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t('feed', this.$store.state.locale)}`
      }
    },
    methods: {
      async update () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.put('/console/settings/feed', {
          outputMode: this.feedOutputMode,
          outputCnt: this.feedOutputCnt
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
        this.$set(this, 'feedOutputMode', responseData.outputMode)
        this.$set(this, 'feedOutputCnt', responseData.outputCnt)
      }
    }
  }
</script>
