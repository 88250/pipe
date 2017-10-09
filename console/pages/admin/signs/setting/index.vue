<template>
  <div>
    <div class="card fn-clear card__body">

      <v-form ref="form">
        <v-text-field
          :label="$t('signs', $store.state.locale)"
          v-model="sign"
          multi-line
          @keyup.ctrl.enter="update"
        ></v-text-field>

        <div class="alert alert--danger" v-show="error">
          <icon icon="danger"/>
          <span>{{ errorMsg }}</span>
        </div>
      </v-form>

      <button class="fn-right btn btn--margin-t30 btn--info btn--space" @click="update">
        {{ $t('confirm', $store.state.locale) }}
      </button>
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        sign: '',
        error: false,
        errorMsg: ''
      }
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t('signs', this.$store.state.locale)}`
      }
    },
    methods: {
      async update () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.post('/console/signs', {
          sign: this.sign
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
      const responseData = await this.axios.get('/console/signs')
      if (responseData) {
        this.$set(this, 'sign', responseData)
      }
    }
  }
</script>
