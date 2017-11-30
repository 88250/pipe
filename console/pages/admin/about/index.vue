<template>
  <div class="card">
    <div class="card__body">
      <div class="fn-flex admin__about">
        <div class="about__side">
          <img src="~/static/images/logo.png"/> <br>
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
          <p class="fn-clear">
            <span class="fn-left">
              <a href='https://hacpai.com/tag/Pipe' target='_blank'>Pipe</a> 是一款开源（<a
              href='http://www.gnu.org/licenses/gpl-3.0.html' target='_blank'>GPLv3</a>）的博客平台，由
              <a href="https://github.com/b3log" target="_blank">B3log 开源</a>组织维护。
            </span>
            <iframe class="about__github fn-left"
                    src="https://ghbtns.com/github-btn.html?user=b3log&repo=pipe&type=star&count=true&size=large"
                    frameborder="0" scrolling="0" width="160px" height="30px"></iframe>
          </p>
          <p>
            如果你对开源感兴趣，想贡献自己的一份力量，欢迎<a
              href='https://hacpai.com/article/1463025124998' target='_blank'>加入我们</a>！</p>
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
        title: `${this.$t('about', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
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
    &__github
      border: 0
      margin: -4px 0 0 10px
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
