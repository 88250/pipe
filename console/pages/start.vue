<template>
  <div class="console" id="particles">
    <div class="card login__content" ref="content">
      <div class="login__github" @click="loginGitHub"></div>
      <img class="fn__none" src="~assets/images/github.gif"/>
      <v-btn class="btn--small btn--info" @click="loginGitHub">{{ $t('index2', $store.state.locale) }}</v-btn>
      <div class="start__space"></div>
      <a class="ft__12 fn__pointer" @click="toggleIntro">查看 GitHub 数据使用说明</a>
      <div class="vditor-reset ft__12 start__intro" v-show="showIntro">
        <ul>
          <li>获取用户名、头像等用于初始化</li>
          <li>获取公开仓库信息用于展示</li>
          <li>不会对你的已有数据进行写入</li>
        </ul>
      </div>
      <div class="start__space"></div>
      <label class="start__checkbox checkbox ft__12">
        <input
          type="checkbox"
          :checked="isAgreen"
          @click="isAgreen = !isAgreen"/><span
        class="checkbox__icon"></span>
        是否愿意在 GitHub 上收藏该<a href="https://github.com/b3log/pipe" target="_blank">项目</a>、关注<a
        href="https://github.com/88250" target="_blank">开发者</a>并加入 <a href="https://github.com/b3log" target="_blank">B3log
        开源组织</a>
      </label>
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
        isAgreen: true,
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
          window.location.href = `${process.env.AxiosBaseURL}/oauth/github/redirect?referer=${document.referrer}__${this.isAgreen
            ? '0'
            : '1'}`
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
