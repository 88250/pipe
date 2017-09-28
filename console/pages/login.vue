<template>
  <div class="login">
    <div class="card">
      <div class="card__body fn-clear">
        <h2>{{ $t('login', $store.state.locale) }}</h2>
        <v-form class="init__center" ref="form">
          <v-text-field
            :label="$t('accountOrEmail', $store.state.locale)"
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
            @keyup.enter="login"
          ></v-text-field>
          <div class="alert alert--danger" v-show="error">
            <icon icon="danger"/>
            <span>{{ errorMsg }}</span>
          </div>
        </v-form>
        <button class="fn-right btn btn--margin-t30 btn--info btn--space" @click="login">
          {{ $t('confirm', $store.state.locale) }}
        </button>
      </div>
    </div>
  </div>
</template>

<script>
  import md5 from 'blueimp-md5'
  export default {
    layout: 'console',
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('login', this.$store.state.locale)}`
      }
    },
    data () {
      return {
        name: '',
        password: '',
        userNameRules: [
          (v) => !!v || this.$t('required', this.$store.state.locale),
          (v) => v.length <= 16 || this.$t('validateRule', this.$store.state.locale)
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
          nameOrEmail: this.name,
          passwordHashed: md5(this.password)
        })
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$router.push(this.$route.query.goto)
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      }
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'
  .login
    background-color: $blue-lighter
    padding: 50px 0
    .card
      width: 650px
      margin: 0 auto
</style>
