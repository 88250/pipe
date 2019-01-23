<template>
  <div class="card">
    <div class="card__title">
      <h3>{{ $t('communityRecommend', $store.state.locale)}}</h3>
    </div>
    <ul class="list admin__index">
      <li v-for="item in list" :key="item.articleCreateTime" class="fn__flex">
        <a class="fn__flex-1" :href="item.articlePermalink" target="_blank">{{ item.articleTitle }}</a>
        <div>{{ item.articleAuthorName }}</div>
      </li>
    </ul>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        list: [],
        tags: ''
      }
    },
    head () {
      return {
        title: `${this.$t('home', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    async mounted () {
      const responseTagsData = await this.axios.get('/console/tags')
      if (responseTagsData) {
        let tags = ''
        responseTagsData.map((v, i) => {
          if (i < 11) {
            tags += `${v.title},`
          }
        })
        this.$set(this, 'tags', tags.substr(0, tags.length - 1))
      }
      const responseData = await this.axios.get(`/hp/apis/articles?tags=${this.tags}&format=json`)
      if (responseData) {
        this.$set(this, 'list', responseData)
      }
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'
  .admin__index
    li
      line-height: 20px
      color: $text-color
      a
        display: block
        padding-right: 15px
        word-break: break-all
        &:hover
          text-decoration: none
</style>
