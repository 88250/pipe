<template>
  <div class="init fn-flex-1">
    <v-stepper v-model="step">
      <v-stepper-header>
        <v-stepper-step step="1">{{ $t('guide', $store.state.locale) }}</v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step step="2">{{ $t('init', $store.state.locale) }}</v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step step="3">{{ $t('welcome', $store.state.locale) }}</v-stepper-step>
      </v-stepper-header>

      <v-stepper-content step="1" class="init__card">
        <div class="fn-clear init__center">
          <h1 class="init__card-title">{{ $t('pipeUseHacpaiAccount', $store.state.locale)}}</h1>
          <a class="card card--dark fn-left"
             href="https://hacpai.com/register"
             target="_blank">
            <div class="card__body fn-flex">
              <v-icon>hacpai-logo</v-icon>
              <div class="fn-flex-1">
                <h3>{{ $t('registerHacpaiAccount', $store.state.locale) }}</h3>
              </div>
            </div>
          </a>
          <a class="card card--danger fn-right"
             href="https://hacpai.com/login"
             target="_blank">
            <div class="card__body fn-flex">
              <v-icon>hacpai-logo</v-icon>
              <div class="fn-flex-1">
                <h3>
                  {{ $t('loginHacpai', $store.state.locale) }}
                </h3>
              </div>
            </div>
          </a>
          <v-btn
            v-if="$store.state.name !== ''"
            class="btn--info fn-right"
            @click="step = 2">{{ $t('nextStep', $store.state.locale) }}
          </v-btn>
        </div>
      </v-stepper-content>

      <v-stepper-content step="2">
        <div class="init__center fn-clear">
          <img class="avatar" :src="`${$store.state.avatarURL}?imageView2/1/w/128/h/128/interlace/1/q/100`"/>
          <h1 class="ft-center init__name">{{ $store.state.name }}</h1>
          <div class="alert alert--danger" v-show="postInitError">
            <v-icon>danger</v-icon>
            <span>{{ postInitErrorMsg }}</span>
          </div>

          <div class="fn-right">
            <v-btn class="btn--info" @click="step = 1">{{ $t('preStep', $store.state.locale) }}</v-btn>
            <v-btn class="btn--success btn--space" @click="init">{{ $t('init', $store.state.locale) }}</v-btn>
          </div>
        </div>
      </v-stepper-content>

      <v-stepper-content step="3" class="init__card fn-clear">
        <div class="fn-clear init__center">
          <a :href="`/blogs/${$store.state.name}`" class="card card--danger init__card--last">
            <div class="card__body fn-flex">
              <img src="~/static/images/logo.jpg"/>
              <div class="fn-flex-1">
                <h3>{{ $t('welcome', $store.state.locale) }} Pipe</h3>
              </div>
            </div>
          </a>
        </div>
      </v-stepper-content>
    </v-stepper>
  </div>
</template>

<script>
  export default {
    layout: 'console',
    data () {
      return {
        step: this.$store.state.name === '' ? 1 : 2,
        postInitError: false,
        postInitErrorMsg: ''
      }
    },
    head () {
      return {
        title: this.$t('welcome', this.$store.state.locale)
      }
    },
    methods: {
      async init () {
        const responseData = await this.axios.post('/init')
        if (responseData.code === 0) {
          this.$set(this, 'step', 3)
          this.$set(this, 'postInitError', false)
          this.$set(this, 'postInitErrorMsg', '')
          this.$store.commit('setIsInit', true)
        } else {
          this.$set(this, 'postInitError', true)
          this.$set(this, 'postInitErrorMsg', responseData.msg)
        }
      }
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'
  .init
    &__center
      max-width: 630px
      margin: 0 auto

      .avatar
        margin: 30px auto 30px
        display: block
        height: 89px
        width: 89px
    &__name
      margin-bottom: 99px
    &__card
      &-title
        margin: 30px 0 60px
      &--last
       margin: 107px auto 0
      .card
        display: block
        width: 280px
        margin-bottom: 60px
        &:hover
          text-decoration: none
          opacity: .9
          box-shadow: 0 14px 26px -12px rgba(23, 105, 255, .42), 0 4px 23px 0 rgba(0, 0, 0, .12), 0 8px 10px -5px rgba(23, 105, 255, .2)
        svg,
        img
          height: 63px
          width: 63px
          margin-right: 15px
          border-radius: 50%
          color: #fff
        h3
          line-height: 63px
          text-align: center
          color: #fff

  @media (max-width: 768px)
    .init__center
      height: auto
    .init__card .card
      margin-top: 0
      width: 100%
      margin-bottom: 15px
    .init__card-title
      margin: 0 0 15px 0
</style>
