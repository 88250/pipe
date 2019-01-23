<template>
  <div class="card fn__clear card__body">
    <v-form>
      <v-text-field
        :label="$t('signs', $store.state.locale)"
        v-model="sign"
        multi-line
        @keyup.ctrl.13="update"
        @keyup.meta.13="update"
      ></v-text-field>

      <div class="alert alert--danger" v-show="error">
        <v-icon>danger</v-icon>
        <span>{{ errorMsg }}</span>
      </div>
    </v-form>

    <v-btn class="fn__right btn--margin-t30 btn--info btn--space" @click="update">
      {{ $t('confirm', $store.state.locale) }}
    </v-btn>
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
        title: `${this.$t('signs', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      async update () {
        const responseData = await this.axios.put('/console/settings/sign', {
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
      const responseData = await this.axios.get('/console/settings/sign')
      if (responseData) {
        this.$set(this, 'sign', responseData)
      }
    }
  }
</script>
