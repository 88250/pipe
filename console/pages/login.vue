<template>
  <div class="login" id="login">
    <div class="card">
      <h2 class="card__title">{{ $t('login', $store.state.locale) }}</h2>
      <div class="card__body fn-clear">
        <v-form class="init__center" ref="form">
          <v-text-field
            :label="$t('accountOrEmail', $store.state.locale)"
            v-model="accountOrEmail"
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
            @keyup.enter="login"
          ></v-text-field>
          <div class="alert alert--danger" v-show="error">
            <v-icon>danger</v-icon>
            <span>{{ errorMsg }}</span>
          </div>
        </v-form>
        <v-btn class="fn-right btn--margin-t30 btn--info" @click="login">
          {{ $t('confirm', $store.state.locale) }}
        </v-btn>
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
        accountOrEmail: '',
        password: '',
        userNameRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 16)
        ],
        error: false,
        errorMsg: ''
      }
    },
    methods: {
      async login () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.post('/login', {
          name: this.accountOrEmail,
          password: this.password
        })
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$store.commit('setStatus', responseData.data)
          this.$router.push(this.$route.query.goto || '/admin')
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

<style lang="sass">
  @import '~assets/scss/_variables'
  .login
    background-color: $blue-lighter
    position: relative
    flex: 1
    display: flex
    align-items: center
    overflow: hidden
    .particles-js-canvas-el
      position: absolute
      top: 0
    .card
      width: 650px
      margin: 0 auto
      position: relative
      z-index: 1
  @media (max-width: 768px)
    .login
      padding: 0 15px
</style>
