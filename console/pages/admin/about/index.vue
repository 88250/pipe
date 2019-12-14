<template>
  <div>
    <div class="card">
      <div class="card__body">
        <div class="fn__flex admin__about">
          <div class="about__side">
            <img src="~assets/images/logo.png"/>
          </div>
          <div class="fn__flex-1 vditor-reset">
            <h2 class="fn__clear" v-if="isLatest">
            <span class="fn__left">
              {{ $t('about1', $store.state.locale) }}
              <a :href="download" target="_blank">{{ version }}</a>
            </span>
              <iframe class="about__github fn__left"
                      src="https://ghbtns.com/github-btn.html?user=88250&repo=pipe&type=star&count=true&size=large"
                      frameborder="0" scrolling="0" width="160px" height="30px"></iframe>
            </h2>
            <h2 class="fn__clear" v-else>
            <span class="fn__left">
               {{ $t('about2', $store.state.locale) }}
              <a class="ft__danger" :href="download" target="_blank">{{ version }}</a>
            </span>
              <iframe class="about__github fn__left"
                      src="https://ghbtns.com/github-btn.html?user=88250&repo=pipe&type=star&count=true&size=large"
                      frameborder="0" scrolling="0" width="160px" height="30px"></iframe>
            </h2>
            <p v-html="$t('about4', $store.state.locale)"></p>
            <p v-html="$t('about3', $store.state.locale)"></p>
            <p>
              <a href="https://hacpai.com/article/1513761942333" target="_blank" class="about__link btn btn--success">
                <v-icon>file-text</v-icon>
                {{ $t('doc', $store.state.locale) }}
              </a>
              <a href="https://hacpai.com/tag/pipe" target="_blank" class="about__link btn btn--success">
                <v-icon>hacpai-logo</v-icon>
                {{ $t('community', $store.state.locale) }}
              </a>
              <a href="https://b3log.org" target="_blank" class="about__link btn btn--success">
                <v-icon>b3log-logo2</v-icon>
                B3log
              </a>
            </p>
            <br>
          </div>
        </div>
      </div>
    </div>
    <br>
    <div class="card">
      <div class="card__title">
        <h3>❤️ 欢迎成为我们的赞助者</h3>
      </div>
      <div class="card__body">
        <a href="https://b3log.org">B3log 开源组织</a>旗下包含
        <a href="https://sym.b3log.org/">Symphony</a>、
        <a href="https://solo.b3log.org/">Solo</a>、
        <a href="https://github.com/88250/pipe">Pipe</a>、
        <a href="https://github.com/88250/wide">Wide</a>、
        <a href="https://github.com/88250/latke">Latke</a>、
        <a href="https://github.com/vanessa219/vditor">Vditor</a>、
        <a href="https://github.com/88250/gulu">Gulu</a>&nbsp;等一系列开源项目。随着项目规模的增长，我们需要有相应的资金支持才能持续项目的维护和开发。
        <br/> <br/>
        如果你觉得 Pipe 还算好用，可通过支付宝对我们进行赞助，谢谢 🙏
        <br/> <br/>
        <div class="ft__center">
          <a class="btn btn--info" href="https://hacpai.com/sponsor">
            <svg viewBox="0 0 32 32" width="100%" height="100%"
                 className={classes.svg}>
              <path
                d="M32 21.906v-15.753c0-3.396-2.757-6.152-6.155-6.152h-19.692c-3.396 0-6.152 2.756-6.152 6.152v19.694c0 3.396 2.754 6.152 6.152 6.152h19.694c3.027 0 5.545-2.189 6.058-5.066-1.632-0.707-8.703-3.76-12.388-5.519-2.804 3.397-5.74 5.434-10.166 5.434s-7.38-2.726-7.025-6.062c0.234-2.19 1.736-5.771 8.26-5.157 3.438 0.323 5.012 0.965 7.815 1.89 0.726-1.329 1.329-2.794 1.785-4.35h-12.433v-1.233h6.151v-2.212h-7.503v-1.357h7.504v-3.195c0 0 0.068-0.499 0.62-0.499h3.077v3.692h7.999v1.357h-7.999v2.212h6.526c-0.6 2.442-1.51 4.686-2.651 6.645 1.895 0.686 10.523 3.324 10.523 3.324v0 0 0zM8.859 24.736c-4.677 0-5.417-2.953-5.168-4.187 0.246-1.227 1.6-2.831 4.201-2.831 2.987 0 5.664 0.767 8.876 2.328-2.256 2.94-5.029 4.69-7.908 4.69v0 0z"></path>
            </svg>
            &nbsp;
            使用支付宝进行赞助
          </a>
        </div>
      </div>
      <ul class="list" v-if="list.length > 0">
        <li
          v-for="item in list"
          :key="item.oId"
          class="fn__flex">
          <a :href="`https://hacpai.com/member/${item.paymentUserName}`"
             :aria-label="item.paymentUserName"
             v-if="item.paymentUserName"
             target="_blank"
             class="avatar avatar--mid avatar--space pipe-tooltipped pipe-tooltipped--n"
             :style="`background-image: url(${item.paymentUserAvatar48})`"></a>
          <img v-else
               class="avatar avatar--mid avatar--space"
               src="~assets/images/logo.png"/>
          <div class="fn__flex-1">
            <div class="fn__flex">
              <div class="fn__flex-1">
                <a v-if="item.paymentUserName" target="_blank"
                   class="list__title" :href="`https://hacpai.com/member/${item.paymentUserName}`">
                  {{ item.paymentUserName }}
                </a>
                <span v-else class="list__title">匿名好心人</span>
                <span class="ft__green">{{ item.paymentAmount }} RMB</span>
              </div>
              <time class="list__meta">{{ item.paymentTimeStr }}</time>
            </div>
            <div class="vditor__reset" v-html="item.paymentMemo"></div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        isLatest: true,
        download: '',
        version: '',
        list: [],
      }
    },
    head () {
      return {
        title: `${this.$t('about', this.$store.state.locale)} - ${this.$store.state.blogTitle}`,
      }
    },
    async mounted () {
      const responseData = await this.axios.get('check-version')
      if (responseData) {
        this.$set(this, 'isLatest', this.$store.state.version === responseData.version)
        this.$set(this, 'download', responseData.download)
        this.$set(this, 'version', responseData.version)
      }

      const responseListData = await this.axios.get('/hp/apis/sponsors?format=json')
      if (responseListData) {
        this.$set(this, 'list', responseListData.payments)
      }
    },
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'

  .admin__about
    .vditor-reset
      padding-left: 50px

      h2
        font-size: 1.4em

    .about
      &__side
        margin-left: 30px
        width: 162px
        align-self: center

      &__github
        border: 0
        margin-left: 20px

      &__link
        margin: 10px 40px 10px 0

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

      .vditor-reset h2
        font-size: 1.2em
</style>
