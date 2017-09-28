<template>
  <div class="admin__skins">
    <div class="card" v-for="item in list" :key="item.previewURL">
      <div class="skin__img-wrap">
        <img :src="item.previewURL"/>
        <div class="skin__overlay">
          <div>
            <a class="btn btn--info" href="javascript:void(0)"><icon icon="setting"/></a>
            <nuxt-link class="btn btn--info btn--space" :to="item.previewURL"><icon icon="info"/></nuxt-link>
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
        list: []
      }
    },
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('about', this.$store.state.locale)}`
      }
    },
    async mounted () {
      const responseData = await this.axios.get('/console/skins')
      if (responseData) {
        this.$set(this, 'list', responseData.skins)
      }
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'

  .admin__skins
    display: flex
    .card
      margin: 0 30px 30px 0
      align-content: flex-start

      .skin__img-wrap
        overflow: hidden
        height: 250px
        width: 360px
        margin-bottom: 15px
        position: relative

        &:hover
          .skin__overlay
            opacity: 1
          img
            transform: scale(1.2) translateZ(0)

      .skin__overlay
        top: 0
        width: 100%
        height: 100%
        position: absolute
        transition: all .4s ease-in-out
        opacity: 0
        background-color: rgba(0, 0, 0, 0.7)

        div
          top: 50%
          left: 0
          right: 0
          transform: translateY(-50%) translateZ(0)
          position: absolute
          text-align: center

        a
          border: 1px solid #fff
          background-color: transparent

        .icon
          cursor: pointer
          height: 20px
          width: 20px
      img
        height: 250px
        width: 360px
        transition: all .4s linear
      h3
        margin-bottom: 25px
        text-align: center
</style>
