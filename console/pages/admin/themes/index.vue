<template>
  <div class="admin__themes fn-clear">
    <div class="card"
         v-for="item in list"
         :key="item.previewURL"
         :class="{ 'theme--current': item.id === currentId }">
      <div class="theme__img-wrap">
        <img :src="item.previewURL"/>
        <div class="theme__overlay">
          <div>
            <v-btn
              v-show="item.id !== currentId"
              class="btn--info"
              @click="setup(item.id)">{{ $t('setup', $store.state.locale) }}</v-btn>
            <a
              class="btn btn--danger btn--space"
              target="_blank"
              :href="item.thumbnailURL">{{ $t('preview', $store.state.locale) }}</a>
          </div>
        </div>
      </div>
      <h3>{{ item.title }}</h3>
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        list: [],
        currentId: ''
      }
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t('themeList', this.$store.state.locale)}`
      }
    },
    methods: {
      async setup (id) {
        const responseData = await this.axios.put(`/console/themes/${id}`)
        if (responseData.code === 0) {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('setupSuccess', this.$store.state.locale),
            snackModify: 'success'
          })

          this.$set(this, 'currentId', id)
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
        this.$set(this, 'currentId', responseData.currentId)
      }
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'

  .admin__themes
    .card
      margin: 0 30px 30px 0
      float: left
      width: 360px

      &.theme--current
        background-color: $blue

        h3
          color: #fff

      .theme__img-wrap
        overflow: hidden
        margin-bottom: 15px
        position: relative
        min-height: 250px
        &:hover
          .theme__overlay
            opacity: 1
          img
            transform: scale(1.2) translateZ(0)

      .theme__overlay
        top: 0
        width: 100%
        height: 100%
        position: absolute
        transition: all .4s ease-in-out
        opacity: 0
        background-color: rgba(0, 0, 0, 0.7)

        & > div
          top: 50%
          left: 0
          right: 0
          transform: translateY(-50%) translateZ(0)
          position: absolute
          text-align: center
      img
        width: 100%
        transition: all .4s linear
      h3
        margin-bottom: 25px
        text-align: center
</style>
