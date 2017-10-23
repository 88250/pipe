<template>
  <div class="search">
    <div class="card fn-clear">
      <v-form ref="form" class="search__form" @submit.prevent="goSearch">
        <v-text-field
          :label="$t('enterSearch', $store.state.locale)"
          v-model="keyword"
          :rules="keywordRules"
          required
          @keyup.enter.prevent="goSearch"
        ></v-text-field>
      </v-form>
      <ul class="list">
        <li v-for="item in list" :key="item.id">
          <nuxt-link class="search__title" :to="item.url"><b>{{ item.title }}</b></nuxt-link>
          <div>
            {{item.content}}
          </div>
        </li>
      </ul>
      <v-pagination
        :length="pageCount"
        v-model="currentPageNum"
        :total-visible="windowSize"
        class="fn-right"
        circle
        next-icon=">"
        prev-icon="<"
        @input="getList"
      ></v-pagination>
    </div>
  </div>
</template>

<script>
  import { required } from '~/plugins/validate'

  export default {
    layout: 'console',
    head () {
      return {
        title: this.$t('search', this.$store.state.locale)
      }
    },
    data () {
      return {
        keyword: '',
        keywordRules: [
          (v) => required.call(this, v)
        ],
        currentPageNum: 1,
        pageCount: 1,
        windowSize: 1,
        list: []
      }
    },
    methods: {
      async getList (currentPage) {
        const responseData = await this.axios.get(`/search?key=${this.$route.query.k}&p=${currentPage}`)
        if (responseData) {
          this.$set(this, 'list', responseData.articles)
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', document.documentElement.clientWidth < 721 ? 5 : responseData.pagination.windowSize)
        }
      },
      async goSearch () {
        if (!this.$refs.form.validate()) {
          return
        }
        this.$router.push(`${location.pathname}?k=${this.keyword}`)
      }
    },
    mounted () {
      this.getList(1)
      this.$set(this, 'keyword', this.$route.query.k)
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'
  .search
    padding: 50px
    background-color: $blue-lighter

    &__form
      margin: 30px 100px 0

</style>
