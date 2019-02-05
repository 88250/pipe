<template>
  <div class="console" id="particles">
    <div class="card" ref="content">
      <div class="login__content">
        <div class="fn__pointer" @click="loginGitHub">
          <div v-html="$t('index2', $store.state.locale)"></div>
          <div class="login__github"></div>
          <img class="fn__none" src="~assets/images/github.gif"/>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import 'particles.js'
  import { initParticlesJS } from '~/plugins/utils'

  export default {
    data () {
      return {
        clickedGitHub: false,
      }
    },
    head () {
      return {
        title: this.$t('welcome', this.$store.state.locale) + ' - Pipe',
      }
    },
    methods: {
      loginGitHub () {
        this.$store.commit('setSnackBar', {
          snackBar: true,
          snackMsg: 'Loading...',
          snackModify: 'success',
        })
        if (!this.clickedGitHub) {
          window.location.href = `${process.env.AxiosBaseURL}/oauth/github/redirect`
          this.$set(this, 'clickedGitHub', true)
        }
      },
    },
    mounted () {
      initParticlesJS('particles')
    },
  }
</script>
