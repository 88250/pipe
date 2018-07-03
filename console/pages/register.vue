<template>
  <div class="console" id="particles">
    <div class="card login">
      <h1>{{ $t('register', $store.state.locale) }}</h1>
      <div class="ft-center login__content fn-flex" v-if="account===''">
        <div class="fn-flex-1">
          <a class="card card--dark login__image"
             href="https://hacpai.com/register"
             target="_blank">
            <v-icon>hacpai-logo</v-icon>
          </a>
          <a class="login__link" href="https://hacpai.com/register"
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
            @keyup.13="register"
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
