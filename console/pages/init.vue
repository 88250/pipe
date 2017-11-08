<template>
  <div class="console" id="particles">
    <div class="card">
      <v-stepper v-model="step">
        <v-stepper-content step="1" class="fn-clear">
          <h1>{{ $t('pipeUseHacpaiAccount', $store.state.locale) }}</h1>
          <div class="ft-center init__content">
            <a class="card card--dark init__image"
               href="https://hacpai.com"
               target="_blank">
              <v-icon>hacpai-logo</v-icon>
            </a>
            <a class="init__link" href="https://hacpai.com/register"
               target="_blank">
              {{ $t('registerHacpaiAccount', $store.state.locale) }}
            </a>
            <a class="init__link" href="https://hacpai.com/login"
               target="_blank">
              {{ $t('loginHacpai', $store.state.locale) }}
            </a>
          </div>
          <v-btn
            v-if="$store.state.name !== ''"
            class="btn--info fn-right"
            @click="step = 2">{{ $t('nextStep', $store.state.locale) }}
          </v-btn>
        </v-stepper-content>

        <v-stepper-content step="2" class="fn-clear">
          <h1>{{ $t('useHacpaiInit', $store.state.locale) }}</h1>
          <div class="ft-center init__content">
            <img class="avatar init__image card"
                 :src="`${$store.state.avatarURL}?imageView2/1/w/128/h/128/interlace/1/q/100`"/>
            <div>{{ $store.state.name }}</div>
            <div class="alert alert--danger" v-show="postInitError">
              <v-icon>danger</v-icon>
              <span>{{ postInitErrorMsg }}</span>
            </div>
          </div>
          <div class="fn-right">
            <v-btn class="btn--info" @click="step = 1">{{ $t('preStep', $store.state.locale) }}</v-btn>
            <v-btn class="btn--success btn--space" @click="init">{{ $t('init', $store.state.locale) }}</v-btn>
          </div>
        </v-stepper-content>

        <v-stepper-content step="3" class="fn-clear">
          <h1>{{ $t('welcome', $store.state.locale) }} Pipe</h1>
          <div class="ft-center init__content">
            <a :href="`/blogs/${$store.state.name}`" class="card init__image">
                <img src="~/static/images/logo.jpg"/>
            </a>
            <div>
              <nuxt-link to="/admin">{{ $t('openPipeTravel', $store.state.locale) }}</nuxt-link>
            </div>
          </div>
        </v-stepper-content>
      </v-stepper>
    </div>
  </div>
</template>

<script>
  import 'particles.js'
  import {initParticlesJS} from '~/plugins/utils'

  export default {
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
    },
    mounted () {
      initParticlesJS('particles')
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'
  .init
    &__content
      height: 310px
    &__link
      margin: 0 25px
    &__image.card
      display: block
      height: 120px
      width: 120px
      margin: 50px auto
      svg
        color: #fff
        height: 100px
        width: 100px
        padding: 10px

</style>
