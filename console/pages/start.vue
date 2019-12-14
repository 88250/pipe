<template>
  <div class="console" id="particles">
    <div class="card login__content" ref="content">
      <div @click="loginGitHub" class="login__icon">
        <v-icon>hacpai-logo</v-icon>
      </div>
      <v-btn class="btn--small btn--info" @click="loginGitHub">{{ $t('index2', $store.state.locale) }}</v-btn>
      <span>&nbsp;&nbsp;</span>
      <a href="https://hacpai.com/article/1576294445994" target="_blank"><v-icon>question</v-icon></a>
      <div class="start__space"></div>
      <div class="start__space"></div>
      <div class="start__space"></div>
      <div class="start__space"></div>
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
        showIntro: false,
      }
    },
    head () {
      return {
        title: this.$t('welcome', this.$store.state.locale) + ' - Pipe',
      }
    },
    methods: {
      toggleIntro () {
        this.$set(this, 'showIntro', !this.showIntro)
      },
      loginGitHub () {
        this.$store.commit('setSnackBar', {
          snackBar: true,
          snackMsg: this.$t('processing', this.$store.state.locale),
          snackModify: 'success',
        })
        if (!this.clickedGitHub) {
          window.location.href = `${process.env.AxiosBaseURL}/login/redirect?referer=${document.referrer}`
          this.$set(this, 'clickedGitHub', true)
        }
      },
    },
    mounted () {
      initParticlesJS('particles')
    },
  }
</script>

<style lang="sass">
  .ft__12
    font-size: 12px !important
  .start
    &__intro
      text-align: left
      width: 300px
      margin: 0 auto
      ul
        margin-bottom: 0 !important
    &__space
      height: 10px
    &__checkbox
      margin: 0 20px
      color: #999
      a
        text-decoration: underline
        color: #67757c

  @media (max-width: 470px)
    .start__checkbox
      line-height: 18px !important
</style>
