<template>
  <div class="console" id="particles">
    <div class="card login">
      <h1>{{$t('register', $store.state.locale)}}</h1>
      <div class="ft__center login__content" v-if="account===''">
        <div class="fn__pointer" @click="loginGitHub">
          {{ $t('useGitHub', $store.state.locale) }}{{ $t('login', $store.state.locale) }}
          <div class="login__github"></div>
          <img class="fn__none" src="~assets/images/github.gif"/>
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
            @keyup.13="register"
          ></v-text-field>
        </v-form>
        <div class="alert alert--danger" v-show="error">
          <v-icon>danger</v-icon>
          <span>{{ errorMsg }}</span>
        </div>
        <div class="fn__right">
          <v-btn class="btn--info" @click="account = ''">{{ $t('preStep', $store.state.locale) }}</v-btn>
          <v-btn
            class="btn--success btn--space"
            @click="register">{{ $t('register', $store.state.locale) }}
          </v-btn>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import 'particles.js'
  import sha512crypt from 'sha512crypt-node'
  import { initParticlesJS } from '~/plugins/utils'
  import { required, maxSize } from '~/plugins/validate'

  export default {
    head () {
      return {
        title: this.$t('register', this.$store.state.locale)
      }
    },
    data () {
      return {
        clickedGitHub: false,
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
      loginGitHub () {
        this.$store.commit('setSnackBar', {
          snackBar: true,
          snackMsg: 'Loading...',
          snackModify: 'success'
        })
        if (!this.clickedGitHub) {
          window.location.href = `${process.env.AxiosBaseURL}/oauth/github/redirect`
          this.$set(this, 'clickedGitHub', true)
        }
      },
      async register () {
        if (!this.$refs.accountForm.validate()) {
          return
        }
        const responseData = await this.axios.post('/register', {
          name: this.userName,
          password: sha512crypt.sha512crypt(this.userPassword, `$6$5000$${Math.random().toString(36)}`)
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
