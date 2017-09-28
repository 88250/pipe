<template>
  <div class="admin__skins">
    <div class="card" v-for="item in list" :key="item.previewURL">
      <div class="skin__img-wrap">
        <img :src="item.previewURL"/>
        <div class="skin__overlay">
          <a href="javascript:void(0)"><icon icon="setting"/></a>
          <nuxt-link :to="item.previewURL"><icon icon="info"/></nuxt-link>
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

        a
          border: 1px solid #fff
          color: #fff
          padding: 12px 15px 10px

        .icon
          color: #fff
          cursor: pointer
      img
        height: 250px
        width: 360px
        transition: all .4s linear
      h3
        margin-bottom: 25px
        text-align: center
</style>
