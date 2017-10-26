<template>
  <div class="init fn-flex-1">
    <v-stepper v-model="step">
      <v-stepper-header>
        <v-stepper-step step="1">{{ $t('guide', $store.state.locale) }}</v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step step="2">{{ $t('verify', $store.state.locale) }}</v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step step="3">{{ $t('init', $store.state.locale) }}</v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step step="4">{{ $t('welcome', $store.state.locale) }}</v-stepper-step>
      </v-stepper-header>

      <v-stepper-content step="1" class="init__card fn-clear">
        <div class="fn-clear init__center">
          <a class="card card--dark fn-left"
             href="https://hacpai.com/register?r=Vanessa"
             target="_blank">
            <div class="card__body fn-flex">
              <v-icon>hacpai-logo</v-icon>
              <div class="fn-flex-1">
                <h3>{{ $t('hacpai', $store.state.locale) }}</h3>
                <div>{{ $t('registerAccount', $store.state.locale) }}</div>
              </div>
            </div>
          </a>
          <a class="card card--danger fn-right"
             href="http://demo.b3log.org"
             target="_blank">
            <div class="card__body fn-flex">
              <img src="~/static/images/logo.jpg"/>
              <div class="fn-flex-1">
                <h3>Pipe</h3>
                <div>{{ $t('onlineDemo', $store.state.locale) }}</div>
              </div>
            </div>
          </a>
        </div>
        <v-btn class="btn--info fn-right" @click="step = 2">{{ $t('nextStep', $store.state.locale) }}</v-btn>
      </v-stepper-content>

      <v-stepper-content step="2" class="fn-clear">
        <v-form class="init__center" ref="form">
          <v-text-field
            :label="$t('hacpaiEmail', $store.state.locale)"
            v-model="userEmail"
            :rules="userEmailRules"
            required
          ></v-text-field>
          <v-text-field
            :label="$t('hacpaiAccount', $store.state.locale)"
            v-model="userName"
            :rules="userNameRules"
            :counter="16"
            required
          ></v-text-field>
          <v-text-field
            label="B3log Key"
            v-model="userB3Key"
            :rules="userB3KeyRules"
            :counter="32"
            @keyup.enter="checkHP"
          ></v-text-field>
          <div class="alert alert--danger" v-show="postError">
            <v-icon>danger</v-icon>
            <span v-html="$t('hacpaiRule', $store.state.locale)"></span>
          </div>
        </v-form>
        <div class="fn-right">
          <v-btn class="btn--info" @click="step = 1">{{ $t('preStep', $store.state.locale) }}</v-btn>
          <v-btn class="btn--success btn--space" @click="checkHP">{{ $t('confirm', $store.state.locale) }}</v-btn>
        </div>
      </v-stepper-content>

      <v-stepper-content step="3" class="fn-clear">
        <v-form class="init__center" ref="initForm">
          <v-text-field
            :label="$t('hacpaiEmail', $store.state.locale)"
            v-model="userEmail"
            :readonly="true"
          ></v-text-field>
          <v-text-field
            :label="$t('hacpaiAccount', $store.state.locale)"
            v-model="userName"
            :readonly="true"
          ></v-text-field>
          <v-text-field
            label="B3log Key"
            v-model="userB3Key"
            :readonly="true"
          ></v-text-field>
          <v-text-field
            :label="$t('password', $store.state.locale)"
            v-model="password"
            :rules="userNameRules"
            :counter="16"
            required
            type="password"
            @keyup.enter="init"
          ></v-text-field>
          <div class="alert alert--danger" v-show="postInitError">
            <v-icon>danger</v-icon>
            <span>{{ postInitErrorMsg }}</span>
          </div>
        </v-form>
        <div class="fn-right">
          <v-btn class="btn--info" @click="step = 2">{{ $t('preStep', $store.state.locale) }}</v-btn>
          <v-btn class="btn--success btn--space" @click="init">{{ $t('confirm', $store.state.locale) }}</v-btn>
        </div>
      </v-stepper-content>

      <v-stepper-content step="4" class="init__card fn-clear">
        <div class="fn-clear init__center">
          <a href="/" class="card card--danger init__card-welcome">
            <div class="card__body fn-flex">
              <img src="~/static/images/logo.jpg"/>
              <div class="fn-flex-1">
                <h3>Pipe</h3>
                <div>{{ $t('welcome', $store.state.locale) }}</div>
              </div>
            </div>
          </a>
        </div>
      </v-stepper-content>
    </v-stepper>
  </div>
</template>

<script>
  import md5 from 'blueimp-md5'
  import { required, maxSize, email } from '~/plugins/validate'

  export default {
    layout: 'console',
    data () {
      return {
        step: 1,
        userName: '',
        userNameRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 16)
        ],
        userEmail: '',
        userEmailRules: [
          (v) => required.call(this, v),
          (v) => email.call(this, v)
        ],
        userB3Key: '',
        userB3KeyRules: [
          (v) => maxSize.call(this, v, 32)
        ],
        postError: false,
        postInitError: false,
        postInitErrorMsg: '',
        password: ''
      }
    },
    head () {
      return {
        title: this.$t('welcome', this.$store.state.locale)
      }
    },
    methods: {
      async checkHP () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.post('/hp/apis/check-account', {
          userName: this.userName,
          userEmail: this.userEmail,
          userB3Key: this.userB3Key
        })
        if (responseData.code === 0) {
          this.$set(this, 'step', 3)
          this.$set(this, 'postError', false)
        } else {
          this.$set(this, 'postError', true)
        }
      },
      async init () {
        this.$set(this, 'step', 4)
        if (!this.$refs.initForm.validate()) {
          return
        }
        const responseData = await this.axios.post('/init', {
          name: this.userName,
          email: this.userEmail,
          passwordHashed: md5(this.password),
          b3key: this.userB3Key
        })
        if (responseData.code === 0) {
          this.$set(this, 'step', 4)
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
    .stepper__wrapper
      height: 460px
    &__center
      max-width: 630px
      height: 380px
      margin: 0 auto
    &__card
      &-welcome.card.card--danger
        width: auto
      .card
        display: block
        width: 280px
        color: #fff
        margin-top: 138px
        &:hover
          text-decoration: none
          opacity: .9
          box-shadow: 0 14px 26px -12px rgba(23, 105, 255, .42), 0 4px 23px 0 rgba(0, 0, 0, .12), 0 8px 10px -5px rgba(23, 105, 255, .2)
        img,
        svg
          height: 63px
          width: 63px
          margin-right: 15px
          border-radius: 50%
          color: #fff
        img
          background-color: #fff
        h3
          line-height: 32px
          margin-bottom: 10px
          color: #fff

  @media (max-width: 768px)
    .init__center
      height: auto
    .init__card .card
      margin-top: 0
      width: 100%
      margin-bottom: 15px
</style>
