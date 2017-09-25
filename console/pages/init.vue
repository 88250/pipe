<template>
  <div class="application application--light">
    <v-stepper v-model="e1">
      <v-stepper-header>
        <v-stepper-step step="1">Name of step 1</v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step step="2">Name of step 2</v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step step="3">Name of step 3</v-stepper-step>
      </v-stepper-header>
      <v-stepper-content step="1">
        <v-card class="grey lighten-1 mb-5" height="200px">1</v-card>
        <v-btn primary @click.native="e1 = 2">Continue</v-btn>
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
        e1: 0
      }
    },
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('init', 'zh_CN')}`
      }
    },
    mounted () {
    }
  }
</script>
