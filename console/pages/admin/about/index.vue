<template>
  <div class="card">
    <div class="card__body">
      <div class="fn-flex admin__about">
        <div class="about__side">
          <img src="~/static/images/logo.jpg"/> <br>
          <a class="btn btn--info btn--margin-t30 btn--block btn--success"
             target="_blank"
             href="http://b3log.org/donate.html">{{ $t('becomeSponsor', $store.state.locale) }}</a>
        </div>
        <div class="fn-flex-1 pipe-content__reset">
          <h2 v-if="isLatest">
            {{ $t('about1', $store.state.locale) }}
            <a :href="download" target="_blank">{{ version }}</a>
          </h2>
          <h2 v-else>
            {{ $t('about2', $store.state.locale) }}
            <a class="ft-danger" :href="download" target="_blank">{{ version }}</a>
          </h2>
          <p><a href='https://hacpai.com/tag/Pipe' target='_blank'>Pipe</a> 是一款开源（<a
            href='http://www.gnu.org/licenses/gpl-3.0.html' target='_blank'>GPLv3</a>）的博客平台，双击即可运行。</p>
          <p><a href='http://b3log.org' target='_blank'>B3log</a> 提倡平等、自由、奔放，并正在尝试构建个人博客 +
            <a href='https://hacpai.com' target='_blank'>论坛</a>的互动体验。如果您有兴趣，<a
              href='https://hacpai.com/article/1463025124998' target='_blank'>加入我们</a>吧！</p>
          <p>
            <a href="http://b3log.org/services/index.html#pipe" target="_blank" class="about__link btn btn--info">
              <v-icon>user-circle</v-icon>
              服务
            </a>
            <a href="https://github.com/b3log/pipe-themes" target="_blank" class="about__link btn btn--info">
              <v-icon>github</v-icon>
              主题
            </a>
            <a href="https://hacpai.com/article/1492881378588" target="_blank" class="about__link btn btn--info">
              <v-icon>file-text</v-icon>
              文档
            </a>
            <a href="https://hacpai.com" target="_blank" class="about__link btn btn--info">
              <v-icon>hacpai-logo2</v-icon>
              社区
            </a>
            <a href="http://b3log.org" target="_blank" class="about__link btn btn--info">
              <v-icon>b3log-logo2</v-icon>
              B3log
            </a>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        isLatest: true,
        download: '',
        version: ''
      }
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t('about', this.$store.state.locale)}`
      }
    },
    async mounted () {
      const responseData = await this.axios.get(`check-version`)
      if (responseData) {
        this.$set(this, 'isLatest', this.$store.state.version === responseData.version)
        this.$set(this, 'download', responseData.download)
        this.$set(this, 'version', responseData.version)
      }
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'

  .admin__about .about
    &__side
      margin: 30px 50px 0 30px
      width: 112px
    &__link
      margin: 10px 20px 10px 20px
      .icon
        border-right: 1px dotted #fff
        padding-right: 10px
        margin: 0 10px 0 0
        height: 18px
        width: 18px

  @media (max-width: 768px)
    .admin__about
      display: block
      .about__side
        text-align: center
        margin: 15px
        width: auto
</style>
