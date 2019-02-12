<template>
  <div class="console" id="particles">
    <div class="card" ref="content">
      <div class="login__content">
        <div class="fn__pointer" @click="loginGitHub">
          <br/><br/>
          <v-btn class="btn--small btn--info">{{ $t('index2', $store.state.locale) }}</v-btn>
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
          snackMsg: this.$t('processing', this.$store.state.locale),
          snackModify: 'success',
        })
        if (!this.clickedGitHub) {
          window.location.href = `${process.env.AxiosBaseURL}/oauth/github/redirect?referer=${document.referrer}`
          this.$set(this, 'clickedGitHub', true)
        }
      },
    },
    mounted () {
      initParticlesJS('particles')
    },
  }
</script>
