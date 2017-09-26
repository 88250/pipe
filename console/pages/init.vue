<template>
  <div class="init">
    <v-stepper v-model="e1">
      <v-stepper-header>
        <v-stepper-step step="1">Name of step 1</v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step step="2">Name of step 2</v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step step="3">Name of step 3</v-stepper-step>
      </v-stepper-header>
      <v-stepper-content step="1">
        <v-form v-model="valid" class="init__form">
          <v-text-field
            label="Name"
            v-model="name"
            :rules="nameRules"
            :counter="10"
            required
          ></v-text-field>
          <v-text-field
            label="E-mail"
            v-model="email"
            :rules="emailRules"
            required
          ></v-text-field>
        </v-form>
        <v-btn primary @click.native="checkHP">Continue</v-btn>
        <v-btn flat>Cancel</v-btn>
      </v-stepper-content>
      <v-stepper-content step="2">
        <v-card class="grey lighten-1 mb-5" height="200px">2</v-card>
        <v-btn primary @click.native="e1 = 3">Continue</v-btn>
        <v-btn flat>Cancel</v-btn>
      </v-stepper-content>
      <v-stepper-content step="3">
        <v-card class="grey lighten-1 mb-5" height="200px">3</v-card>
        <v-btn primary @click.native="e1 = 1">Continue</v-btn>
        <v-btn flat>Cancel</v-btn>
      </v-stepper-content>
    </v-stepper>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        e1: 0,
        valid: false,
        name: '',
        nameRules: [
          (v) => !!v || 'Name is required',
          (v) => v.length <= 10 || 'Name must be less than 10 characters'
        ],
        email: '',
        emailRules: [
          (v) => !!v || 'E-mail is required'
        ]
      }
    },
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('init', this.$store.state.locale)}`
      }
    },
    methods: {
      async checkHP () {
        const resultData = await this.axios.post('/hp/apis/check-account', {
          userName: 'solo',
          userEmail: 'b3log.solo@gmail.com',
          userB3Key: 'sss'
        })

        console.log(resultData)
      }
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'
  .init
    &__form
      width: 600px
      margin: 0 auto
</style>
