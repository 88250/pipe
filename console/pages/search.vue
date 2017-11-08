<template>
  <div class="content">
    <div class="card card__body fn-clear">
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
          <a class="search__title" :href="item.url"><b>{{ item.title }}</b></a>
          <div>
            {{item.content}}
          </div>
        </li>
      </ul>
      <v-pagination
        v-if="currentPageNum > 1"
        :length="pageCount"
        v-model="currentPageNum"
        :total-visible="windowSize"
        class="fn-right"
        circle
        next-icon="angle-right"
        prev-icon="angle-left"
        @input="getList"
      ></v-pagination>
    </div>
  </div>
</template>

<script>
  import {required} from '~/plugins/validate'

  export default {
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
      async getList (currentPage = 1) {
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
      this.getList()
      setTimeout(() => {
        this.$set(this, 'keyword', this.$route.query.k)
      })
    }
  }
</script>
