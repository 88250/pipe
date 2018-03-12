<template>
  <div class="card">
    <div class="card__body">
      <div class="fn-flex admin__about">
        <div class="about__side">
          <img src="~/static/images/logo.png"/> <br>
          <a class="btn btn--info btn--margin-t30 btn--block btn--success"
             target="_blank"
             href="https://b3log.org/donate.html">{{ $t('becomeSponsor', $store.state.locale) }}</a>
        </div>
        <div class="fn-flex-1 pipe-content__reset">
          <h2 class="fn-clear" v-if="isLatest">
            <span class="fn-left">
              {{ $t('about1', $store.state.locale) }}
              <a :href="download" target="_blank">{{ version }}</a>
            </span>
            <iframe class="about__github fn-left"
                    src="https://ghbtns.com/github-btn.html?user=b3log&repo=pipe&type=star&count=true&size=large"
                    frameborder="0" scrolling="0" width="160px" height="30px"></iframe>
          </h2>
          <h2 class="fn-clear" v-else>
            <span class="fn-left">
               {{ $t('about2', $store.state.locale) }}
              <a class="ft-danger" :href="download" target="_blank">{{ version }}</a>
            </span>
            <iframe class="about__github fn-left"
                    src="https://ghbtns.com/github-btn.html?user=b3log&repo=pipe&type=star&count=true&size=large"
                    frameborder="0" scrolling="0" width="160px" height="30px"></iframe>
          </h2>
          <p v-html="$t('about4', $store.state.locale)"></p>
          <p v-html="$t('about3', $store.state.locale)"></p>
          <p>
            <a href="https://github.com/b3log/pipe-themes" target="_blank" class="about__link btn btn--info">
              <v-icon>github</v-icon>
              {{ $t('theme', $store.state.locale) }}
            </a>
            <a href="https://hacpai.com/article/1513761942333" target="_blank" class="about__link btn btn--info">
              <v-icon>file-text</v-icon>
              {{ $t('doc', $store.state.locale) }}
            </a>
            <a href="https://hacpai.com/tag/Pipe" target="_blank" class="about__link btn btn--info">
              <v-icon>hacpai-logo2</v-icon>
              {{ $t('community', $store.state.locale) }}
            </a>
            <a href="https://b3log.org" target="_blank" class="about__link btn btn--info">
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

  .admin__about
    .pipe-content__reset h2
      font-size: 1.4em
    .about
      &__side
        margin: 30px 50px 0 30px
        width: 112px
      &__github
        border: 0
        margin-left: 20px
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
      .about__github
        float: none
        margin: 10px 0 0 0
      .about__side
        text-align: center
        margin: 15px
        width: auto
      .pipe-content__reset h2
        font-size: 1.2em
</style>
