<template>
  <div class="card__body fn-clear">
    <v-form ref="form">
      <v-text-field
        :label="$t('account', $store.state.locale)"
        v-model="name"
        :counter="32"
        :rules="titleRules"
        required
      ></v-text-field>
      <v-text-field
        :label="$t('nickname', $store.state.locale)"
        v-model="nickname"
        :rules="titleRules"
        :counter="32"
        required
      ></v-text-field>
      <v-text-field
        :label="$t('hacpaiEmail', $store.state.locale)"
        v-model="email"
        :rules="titleRules"
        :counter="32"
        required
      ></v-text-field>
      <v-text-field
        :label="$t('password', $store.state.locale)"
        v-model="password"
        :counter="32"
        type="password"
        required
      ></v-text-field>
      <v-text-field
        :label="$t('avatarURL', $store.state.locale)"
        v-model="avatarURL"
      ></v-text-field>
      <div class="alert alert--danger" v-show="error">
        <icon icon="danger"/>
        <span>{{ errorMsg }}</span>
      </div>
    </v-form>
    <button class="fn-right btn btn--margin-t30 btn--info btn--space" @click="created">
      {{ $t('confirm', $store.state.locale) }}
    </button>
    <button class="fn-right btn btn--margin-t30 btn--danger btn--space" @click="$emit('update:show', false)">
      {{ $t('cancel', $store.state.locale) }}
    </button>
  </div>
</template>

<script>
  import md5 from 'blueimp-md5'

  export default {
    props: ['id'],
    data () {
      return {
        errorMsg: '',
        error: false,
        name: '',
        nickname: '',
        avatarURL: '',
        password: '',
        email: '',
        titleRules: [
          (v) => !!v || this.$t('required', this.$store.state.locale)
        ]
      }
    },
    watch: {
      id: function () {
        this.init()
      }
    },
    methods: {
      async created () {
        if (!this.$refs.form.validate()) {
          return
        }
        let responseData = {}
        if (this.id === '') {
          responseData = await this.axios.post('/console/users', {
            name: this.name,
            nickname: this.nickname,
            email: this.email,
            password: md5(this.password),
            avatarURL: this.avatarURL
          })
        } else {
          responseData = await this.axios.put(`/console/users/${this.id}`, {
            name: this.name,
            nickname: this.nickname,
            email: this.email,
            password: md5(this.password),
            avatarURL: this.avatarURL
          })
        }

        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$emit('addSuccess')
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      },
      async init () {
        if (this.id === '') {
          return
        }
        const responseData = await this.axios.get(`/console/users/${this.id}`)
        if (responseData) {
          this.$set(this, 'name', responseData.name)
          this.$set(this, 'nickname', responseData.nickname)
          this.$set(this, 'email', responseData.email)
          this.$set(this, 'avatarURL', responseData.avatarURL)
          this.$set(this, 'password', responseData.password)
        }
      }
    },
    mounted () {
      this.init()
    }
  }
</script>
