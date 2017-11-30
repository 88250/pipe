<template>
  <div id="particles" class="console">
    <div class="card">
      <v-form ref="form" @submit.prevent="goSearch" class="card__body">
        <v-text-field
          :label="$t('enterSearch', $store.state.locale)"
          v-model="keyword"
          @keyup.enter.prevent="goSearch"
        ></v-text-field>
      </v-form>
      <ul class="list">
        <li v-for="item in list" :key="item.id" class="fn-flex">
          <div class="fn-flex-1">
            <a :href="item.url" class="list__title">{{ item.title }}</a>
            <div class="tags">
              <a v-for="tag in item.tags" :href="tag.url" class="tag">{{tag.title}}</a>
            </div>
            <div>
              {{item.abstract}}
            </div>
          </div>
        </li>
      </ul>
      <div class="pagination--wrapper fn-clear" v-if="pageCount > 1">
        <v-pagination
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
  </div>
</template>

<script>
  import 'particles.js'
  import { initParticlesJS } from '~/plugins/utils'

  export default {
    head () {
      return {
        title: this.$t('search', this.$store.state.locale) + ' - ' + this.$store.state.blogTitle
      }
    },
    data () {
      return {
        keyword: '',
        currentPageNum: 1,
        pageCount: 1,
        windowSize: 1,
        list: []
      }
    },
    methods: {
      async getList (currentPage = 1) {
        const responseData = await this.axios.post(`/search`, {
          key: this.$route.query.key,
          p: currentPage
        })
        if (responseData.code === 0) {
          this.$set(this, 'list', responseData.data.articles)
          this.$set(this, 'currentPageNum', responseData.data.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.data.pagination.pageCount)
          this.$set(this, 'windowSize', 5)
        }
      },
      async goSearch () {
        if (!this.$refs.form.validate()) {
          return
        }
        this.$router.push(`${location.pathname}?key=${this.keyword}`)
      }
    },
    mounted () {
      this.getList()
      initParticlesJS('particles')
      setTimeout(() => {
        this.$set(this, 'keyword', this.$route.query.key)
      }, 0)
    }
  }
</script>
