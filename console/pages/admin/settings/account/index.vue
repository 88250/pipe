<template>
  <div>
    <div class="card fn-clear card__body">
      <v-form ref="passwordForm">
        <v-text-field
          :label="$t('password', $store.state.locale)"
          v-model="password"
          type="password"
          :rules="requiredRules"
          required
          @keyup.ctrl.13="passwordUpdate"
          @keyup.meta.13="passwordUpdate"
        ></v-text-field>
        <div class="alert alert--danger" v-show="errorPassword">
          <v-icon>danger</v-icon>
          <span>{{ errorPasswordMsg }}</span>
        </div>
      </v-form>

      <v-btn class="fn-right btn--margin-t30 btn--info btn--space" @click="passwordUpdate">
        {{ $t('confirm', $store.state.locale) }}
      </v-btn>
    </div>
    <br>
    <div class="card fn-clear card__body">
      <v-form ref="form">
        <v-text-field
          :label="$t('avatarURL', $store.state.locale)"
          v-model="avatarURL"
          :rules="avatarRules"
          :counter="255"
          required
          @keyup.ctrl.13="accountUpdate"
          @keyup.meta.13="accountUpdate"
        ></v-text-field>

        <v-text-field
          label="B3log Key"
          v-model="b3key"
          :counter="20"
          @keyup.ctrl.13="accountUpdate"
          @keyup.meta.13="accountUpdate"
        ></v-text-field>

        <div class="fn-clear">
          <a
            class="fn-right"
            href="https://hacpai.com/settings/b3"
            target="_blank">
            {{ $t('check', $store.state.locale) }}/{{ $t('setting', $store.state.locale) }} B3log Key
          </a>
        </div>
        <div class="alert alert--danger" v-show="error">
          <v-icon>danger</v-icon>
          <span>{{ errorMsg }}</span>
        </div>
      </v-form>

      <v-btn class="fn-right btn--margin-t30 btn--info btn--space" @click="accountUpdate">
        {{ $t('confirm', $store.state.locale) }}
      </v-btn>
    </div>
  </div>
</template>

<script>
  import sha512crypt from 'sha512crypt-node'
  import { maxSize, required } from '~/plugins/validate'

  export default {
    data () {
      return {
        requiredRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 20)
        ],
        avatarRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 255)
        ],
        error: false,
        errorMsg: '',
        errorPassword: false,
        errorPasswordMsg: '',
        password: '',
        b3key: '',
        avatarURL: ''
      }
    },
    head () {
      return {
        title: `${this.$t('account', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      async passwordUpdate () {
        if (!this.$refs.passwordForm.validate()) {
          return
        }
        const responseData = await this.axios.put('/console/settings/account/password', {
          password: sha512crypt.sha512crypt(this.password, `$6$5000$${Math.random().toString(36)}`)
        })

        if (responseData.code === 0) {
          this.$set(this, 'errorPassword', false)
          this.$set(this, 'errorPasswordMsg', '')
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('setupSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
        } else {
          this.$set(this, 'errorPassword', true)
          this.$set(this, 'errorPasswordMsg', responseData.msg)
        }
      },
      async accountUpdate () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.put('/console/settings/account', {
          b3key: this.b3key,
          avatarURL: this.avatarURL
        })

        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('setupSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      }
    },
    async mounted () {
      const responseData = await this.axios.get('/console/settings/account')
      if (responseData) {
        this.$set(this, 'b3key', responseData.b3Key)
        this.$set(this, 'avatarURL', responseData.avatarURL)
      }
    }
  }
</script>
