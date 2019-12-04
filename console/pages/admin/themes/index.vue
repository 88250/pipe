<template>
  <div class="admin__themes">
    <div class="theme__link">
      <a href="https://github.com/88250/pipe/issues/1" target="_blank">新主题推荐</a> •
      <a href="https://hacpai.com/article/1512550354920#toc_h3_20" target="_blank">Pipe 主题开发指南</a>
    </div>
    <div class="fn__clear">
      <div class="card"
           v-for="item in list"
           @click="setup(item.name)"
           :key="item.previewURL"
           :class="{ 'theme--current': item.name === currentName }">
        <div class="theme__name">{{ item.name }}</div>
        <div class="theme__img-wrap">
          <span class="theme__image" :style="`background-image: url('${item.thumbnailURL}')`"/>
          <div class="theme__overlay">
            <v-btn
              v-show="item.name !== currentName"
              class="btn--info">{{ $t('setup', $store.state.locale) }}</v-btn>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        list: [],
        currentName: ''
      }
    },
    head () {
      return {
        title: `${this.$t('themeList', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      async setup (name) {
        if (name === this.currentName) {
          return
        }
        const responseData = await this.axios.put(`/console/themes/${name}`)
        if (responseData.code === 0) {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('setupSuccess', this.$store.state.locale),
            snackModify: 'success'
          })

          this.$set(this, 'currentName', name)
        } else {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: responseData.msg
          })
        }
      }
    },
    async mounted () {
      const responseData = await this.axios.get('/console/themes')
      if (responseData) {
        this.$set(this, 'list', responseData.themes)
        this.$set(this, 'currentName', responseData.currentId)
      }
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'

  .admin__themes
    .card
      margin: 0 30px 62px 0
      float: left
      width: 455px
      cursor: pointer
    .theme
      &__link
        margin-bottom: 52px
        text-align: center
      &__name
        position: absolute
        top: -32px
        text-align: center
        background-color: rgba(255, 255, 255, .7)
        left: 50%
        line-height: 32px
        width: 100px
        margin-left: -50px
      &--current
        background-color: $blue
        cursor: auto
        .theme__overlay
          opacity: 1
        .theme__image
          transform: scale(1.2) translateZ(0)

      &__img-wrap
        overflow: hidden
        position: relative
        &:hover
          .theme__overlay
            opacity: 1
          .theme__image
            transform: scale(1.2) translateZ(0)

      &__overlay
        top: 0
        width: 100%
        height: 100%
        position: absolute
        transition: all .4s ease-in-out
        opacity: 0
        background-color: rgba(0, 0, 0, 0.2)
        text-align: center
        color: #fff
        font-size: 20px
        padding-top: 150px
        box-sizing: border-box

      &__image
        width: 100%
        transition: all .4s linear
        height: 360px
        display: block
        background-repeat: no-repeat
        background-size: cover
        background-position: center center
  @media (max-width: 768px)
    .admin__themes .card
      width: 100%
</style>
