<template>
  <div class="console" id="particles">
    <div class="card">
      <v-stepper v-model="step">
        <v-stepper-content step="1" class="fn-clear">
          <h1>{{ $t('pipeUseHacpaiAccount', $store.state.locale) }}</h1>
          <div class="ft-center init__content fn-clear">
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
          <h1>
            {{ $t('useHacpaiInit', $store.state.locale) }}
            {{ $store.state.name }}
            {{ $t('init', $store.state.locale) }}
          </h1>
          <div class="ft-center init__content fn-clear">
            <img class="avatar init__image card init__image--step2"
                 :src="`${$store.state.avatarURL}?imageView2/1/w/128/h/128/interlace/1/q/100`"/>
            <v-form ref="form" class="fn-flex" @submit.prevent="init">
              <v-text-field
                class="fn-flex-1"
                label="B3log Key"
                v-model="b3key"
                :counter="20"
                :rules="b3keyRules"
                required
              ></v-text-field>
              <a
                class="init__help"
                href="https://hacpai.com/settings/b3"
                target="_blank">{{ $t('check', $store.state.locale) }} B3log key</a>
            </v-form>
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
          <div class="ft-center init__content fn-clear">
            <a href="https://github.com/b3log/pipe"
               target="_blank"
               :aria-label="$t('openPipeTravel', $store.state.locale)"
               class="card init__image init__image--step3 pipe-tooltipped pipe-tooltipped--n">
                <img src="~/static/images/logo.png"/>
            </a>
            <div class="init__text">{{$t('starIsMotivation', $store.state.locale)}}</div>
            <iframe src="https://ghbtns.com/github-btn.html?user=b3log&repo=pipe&type=star&count=true&size=large"
                    frameborder="0" scrolling="0" width="160px" height="30px"></iframe>
          </div>
        </v-stepper-content>
      </v-stepper>
    </div>
  </div>
</template>

<script>
  import Vue from 'vue'
  import 'particles.js'
  import {initParticlesJS} from '~/plugins/utils'
  import { required, maxSize } from '~/plugins/validate'

  export default {
    data () {
      return {
        step: 4,
        postInitError: false,
        postInitErrorMsg: '',
        b3key: '',
        b3keyRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 20)
        ]
      }
    },
    head () {
      return {
        title: this.$t('welcome', this.$store.state.locale) + ' - Pipe'
      }
    },
    methods: {
      async init () {
        if (!this.$refs.form.validate()) {
          return
        }

        const responseData = await this.axios.post('/init', {
          b3key: this.b3key
        })
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
      Vue.nextTick(() => {
        this.$set(this, 'step', this.$store.state.name === '' ? 1 : 2)
      })
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'
  .init
    &__content
      height: 300px
    &__link
      margin: 0 25px
    &__help
      padding: 34px 0 0 20px
    &__image
      &.card
        display: block
        height: 120px
        width: 120px
        margin: 54px auto
        svg
          color: #fff
          height: 100px
          width: 100px
          padding: 10px
      &--step2.card
        margin: 20px auto 0
      &--step3.card
        margin: 30px auto 36px
    &__text
      margin-bottom: 30px
</style>
