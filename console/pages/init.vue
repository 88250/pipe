<template>
  <div class="init">
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
              <icon icon="hacpai-logo"/>
              <div class="fn-flex-1">
                <h2>{{ $t('hacpai', $store.state.locale) }}</h2>
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
                <h2>solo.go</h2>
                <div>{{ $t('onlineDemo', $store.state.locale) }}</div>
              </div>
            </div>
          </a>
        </div>
        <button class="btn btn--info fn-right" @click="step = 2">{{ $t('nextStep', $store.state.locale) }}</button>
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
          ></v-text-field>
          <div class="alert alert--danger" v-show="postError">
            <icon icon="danger"/>
            <span v-html="$t('hacpaiRule', $store.state.locale)"></span>
          </div>
        </v-form>
        <div class="fn-right">
          <button class="btn btn--info" @click="step = 1">{{ $t('preStep', $store.state.locale) }}</button>
          <button class="btn btn--success btn--space" @click="checkHP">{{ $t('confirm', $store.state.locale) }}</button>
        </div>
      </v-stepper-content>

      <v-stepper-content step="3" class="fn-clear">
        <v-form class="init__center" ref="initForm">
          <v-text-field
            :label="$t('nickName', $store.state.locale)"
            v-model="name"
            :counter="16"
            :rules="userNameRules"
            required
          ></v-text-field>
          <v-text-field
            :label="$t('password', $store.state.locale)"
            v-model="password"
            :rules="userNameRules"
            :counter="16"
            required
            type="password"
          ></v-text-field>
          <div class="alert alert--danger" v-show="postInitError">
            <icon icon="danger"/>
            <span>{{ postInitErrorMsg }}</span>
          </div>
        </v-form>
        <div class="fn-right">
          <button class="btn btn--info" @click="step = 2">{{ $t('preStep', $store.state.locale) }}</button>
          <button class="btn btn--success btn--space" @click="init">{{ $t('confirm', $store.state.locale) }}</button>
        </div>
      </v-stepper-content>

      <v-stepper-content step="4" class="init__card fn-clear">
        <div class="fn-clear init__center">
          <nuxt-link to="/" class="card card--danger init__card-welcome">
            <div class="card__body fn-flex">
              <img src="~/static/images/logo.jpg"/>
              <div class="fn-flex-1">
                <h2>solo.go</h2>
                <div>{{ $t('welcome', $store.state.locale) }}</div>
              </div>
            </div>
          </nuxt-link>
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
        step: 1,
        userName: '',
        userNameRules: [
          (v) => !!v || this.$t('required', this.$store.state.locale),
          (v) => v.length <= 16 || this.$t('validateRule', this.$store.state.locale)
        ],
        userEmail: '',
        userEmailRules: [
          (v) => !!v || this.$t('required', this.$store.state.locale),
          (v) => /^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$/.test(v) ||
            this.$t('emailRule', this.$store.state.locale)
        ],
        userB3Key: '',
        userB3KeyRules: [
          (v) => v.length <= 32 || this.$t('validateRule2', this.$store.state.locale)
        ],
        postError: false,
        postInitError: false,
        postInitErrorMsg: '',
        name: '',
        password: ''
      }
    },
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('init', this.$store.state.locale)}`
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
        if (!this.$refs.initForm.validate()) {
          return
        }
        const responseData = await this.axios.post('/init', {
          name: this.name,
          email: this.userEmail,
          password: this.password,
          b3key: this.userB3Key
        })
        if (responseData.code === 0) {
          this.$set(this, 'step', 4)
          this.$set(this, 'postInitError', false)
          this.$set(this, 'postInitErrorMsg', '')
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
      width: 630px
      margin: 50px auto
      height: 300px

    &__card
      &-welcome.card.card--danger
        width: auto
      .card
        display: block
        width: 280px
        color: #fff
        margin-top: 98px
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
        img
          background-color: #fff
        h2
          line-height: 32px
          margin-bottom: 10px
</style>
