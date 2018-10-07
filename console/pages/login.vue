<template>
  <div class="console" id="particles">
    <div class="card login">
      <h1>{{ $t('login', $store.state.locale) }}</h1>
      <div class="ft-center login__content" v-if="account===''">
        <a :href="`${baseURL}/oauth/github/redirect`">
          {{ $t('useGitHub', $store.state.locale) }}{{ $t('login', $store.state.locale) }}
          <div class="login__github"></div>
          <img class="fn-none" src="~/static/images/github.gif"/>
        </a>

        <a class="login__link fn-flex-center" href="https://hacpai.com/login"
           target="_blank">
          <v-icon>hacpai-logo</v-icon>
          <div>&nbsp;{{ $t('useHacpaiInit', $store.state.locale) }}</div>
        </a>
        <div class="login__link fn-pointer fn-flex-center" @click="account='pipe'">
          <img width="16" src="~/static/images/logo.png"/>
          <span>&nbsp;{{ $t('usePipeInit', $store.state.locale) }}</span>
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
            @keyup.13="login"
          ></v-text-field>
        </v-form>
        <div class="alert alert--danger" v-show="error">
          <v-icon>danger</v-icon>
          <span>{{ errorMsg }}</span>
        </div>
        <div class="fn-right">
          <v-btn class="btn--info" @click="account = ''">{{ $t('preStep', $store.state.locale) }}</v-btn>
          <v-btn
            class="btn--success btn--space"
            @click="login">{{ $t('login', $store.state.locale) }}
          </v-btn>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import 'particles.js'
  import { initParticlesJS } from '~/plugins/utils'
  import { required, maxSize } from '~/plugins/validate'

  export default {
    head () {
      return {
        title: this.$t('login', this.$store.state.locale)
      }
    },
    data () {
      return {
        baseURL: process.env.AxiosBaseURL,
        account: '',
        userName: '',
        userPassword: '',
        requiredRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 16)
        ],
        error: false,
        errorMsg: ''
      }
    },
    methods: {
      async login () {
        if (!this.$refs.accountForm.validate()) {
          return
        }
        const responseData = await this.axios.post('/login', {
          name: this.userName,
          password: this.userPassword
        })
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          const stateResponseData = await this.axios.get('/status')
          if (stateResponseData) {
            this.$store.commit('setStatus', stateResponseData)
          }
          this.$router.push('/admin')
          if (document.documentElement.clientWidth >= 768) {
            this.$store.commit('setBodySide', 'body--side')
          }
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      }
    },
    mounted () {
      initParticlesJS('particles')
    }
  }
</script>
