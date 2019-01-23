<template>
  <div>
    <div class="card fn__clear card__body">
      <v-form>
        <v-text-field
          multi-line
          label="Google AdSense ins"
          :rules="linkRules"
          :counter="255"
          v-model="adGoogleAdSenseArticleEmbed"
        ></v-text-field>
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
  import {maxSize} from '~/plugins/validate'

  export default {
    data () {
      return {
        adGoogleAdSenseArticleEmbed: '',
        linkRules: [
          (v) => maxSize.call(this, v, 255)
        ],
        error: false,
        errorMsg: ''
      }
    },
    head () {
      return {
        title: `${this.$t('3rdStatistic', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      async update () {
        const responseData = await this.axios.put('/console/settings/ad', {
          adGoogleAdSenseArticleEmbed: this.adGoogleAdSenseArticleEmbed
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
      const responseData = await this.axios.get('/console/settings/ad')
      if (responseData) {
        this.$set(this, 'adGoogleAdSenseArticleEmbed', responseData.adGoogleAdSenseArticleEmbed)
      }
    }
  }
</script>
