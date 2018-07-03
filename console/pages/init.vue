<template>
  <div class="console" id="particles">
    <div class="card">
      <v-stepper v-model="step">
        <v-stepper-content step="1" class="fn-clear">
          <h1>Pipe {{ $t(account === 'pipe' ? 'register' : 'guide', $store.state.locale) }}</h1>
          <div class="ft-center login__content fn-flex" v-if="account===''">
            <div class="fn-flex-1">
              <a class="card card--dark login__image"
                 href="https://hacpai.com/login"
                 target="_blank">
                <v-icon>hacpai-logo</v-icon>
              </a>
              <a class="login__link" href="https://hacpai.com/login"
                 target="_blank">
                {{ $t('useHacpaiInit', $store.state.locale) }}
              </a>
            </div>
            <div class="fn-flex-1">
              <div class="card login__image fn-pointer"
                 @click="account='pipe'">
                <img src="~/static/images/logo.png"/>
              </div>
              <div class="login__link fn-pointer" @click="account='pipe'">
                {{ $t('usePipeInit', $store.state.locale) }}
              </div>
            </div>
          </div>
          <div v-if="account==='pipe'">
            <br>
            <v-form ref="accountForm" @submit.prevent="init">
              <v-text-field
                :label="$t('userName', $store.state.locale)"
                v-model="userName"
                :counter="16"
                :rules="requiredRules"
                required
              ></v-text-field>
              <v-text-field
                :label="$t('password', $store.state.locale)"
                v-model="userPassword"
                :counter="16"
                :rules="requiredRules"
                required
                type="password"
                @keyup.13="goStep2"
              ></v-text-field>
            </v-form>
            <div class="alert alert--danger" v-show="registerError">
              <v-icon>danger</v-icon>
              <span>{{ registerErrorMsg }}</span>
            </div>
          </div>
          <div class="fn-right">
            <v-btn  v-if="account === 'pipe'"
              class="btn--info" @click="account = ''">{{ $t('preStep', $store.state.locale) }}</v-btn>
            <v-btn
              v-if="$store.state.name !== '' || account === 'pipe'"
              class="btn--success btn--space"
              @click="goStep2">{{ $t(account === 'pipe' ? 'init' : 'nextStep', $store.state.locale) }}
            </v-btn>
          </div>
        </v-stepper-content>

        <v-stepper-content step="2" class="fn-clear">
          <h1>
            {{ $t('useHacpaiInit', $store.state.locale) }}
            {{ $store.state.name }}
            {{ $t('init', $store.state.locale) }}
          </h1>
          <div class="ft-center login__content fn-clear">
            <img class="avatar login__image card init__image--step2"
                 :src="`${$store.state.avatarURL}?imageView2/1/w/128/h/128/interlace/1/q/100`"/>
            <v-form ref="form" class="fn-flex" @submit.prevent="init">
              <v-text-field
                class="fn-flex-1"
                label="B3log Key"
                v-model="b3key"
                :counter="20"
                :rules="b3keyRules"
                required
                @keyup.13="init"
              ></v-text-field>
              <a
                class="init__help"
                href="https://hacpai.com/settings/b3"
                target="_blank">{{ $t('check', $store.state.locale) }}/{{ $t('setting', $store.state.locale) }} B3log Key</a>
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
          <div class="ft-center login__content fn-clear">
            <a href="https://github.com/b3log/pipe"
               target="_blank"
               :aria-label="$t('openPipeTravel', $store.state.locale)"
               class="card login__image init__image--step3 pipe-tooltipped pipe-tooltipped--n">
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
  import sha512crypt from 'sha512crypt-node'
  import {initParticlesJS} from '~/plugins/utils'
  import { required, maxSize } from '~/plugins/validate'

  export default {
    data () {
      return {
        userName: '',
        userPassword: '',
        account: '',
        step: 4,
        postInitError: false,
        postInitErrorMsg: '',
        registerError: false,
        registerErrorMsg: '',
        b3key: '',
        b3keyRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 20)
        ],
        requiredRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 16)
        ]
      }
    },
    head () {
      return {
        title: this.$t('welcome', this.$store.state.locale) + ' - Pipe'
      }
    },
    methods: {
      async goStep2 () {
        if (this.account === '') {
          this.$set(this, 'step', 2)
        }

        if (this.account === 'pipe') {
          if (!this.$refs.accountForm.validate()) {
            return
          }
          const responseData = await this.axios.post('/init/local', {
            name: this.userName,
            password: sha512crypt.sha512crypt(this.userPassword, `$6$5000$${Math.random().toString(36)}`)
          })
          if (responseData.code === 0) {
            this.$set(this, 'step', 3)
            this.$set(this, 'registerError', false)
            this.$set(this, 'registerErrorMsg', '')
            this.$store.commit('setIsInit', true)
            const stateResponseData = await this.axios.get('/status')
            if (stateResponseData) {
              this.$store.commit('setStatus', stateResponseData)
            }
          } else {
            this.$set(this, 'registerError', true)
            this.$set(this, 'registerErrorMsg', responseData.msg)
          }
        }
      },
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
    &__help
      padding: 34px 0 0 20px
    &__image
      &--step2.card
        margin: 20px auto 0
      &--step3.card
        margin: 30px auto 36px
    &__text
      margin-bottom: 30px
</style>
