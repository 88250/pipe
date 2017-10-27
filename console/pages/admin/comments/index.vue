<template>
  <div class="card">
    <div class="card__body">
      <v-text-field
        v-if="list.length > 0"
        @keyup.enter="getList(1)"
        class="fn-flex-1"
        :label="$t('enterSearch', $store.state.locale)"
        v-model="keyword">
      </v-text-field>
      <div v-else>
        {{ $t('noData', $store.state.locale) }}
      </div>
    </div>
    <ul class="list" v-if="list.length > 0">
      <li v-for="item in list" :key="item.id" class="fn-flex">
        <a class="avatar avatar--mid avatar--space tooltipped tooltipped--s"
           :aria-label="item.author.name"
           :href="item.author.url"
           :style="`background-image: url(${item.author.avatarURL})`"></a>
        <div class="fn-flex-1">
          <div class="fn-flex">
            <div class="fn-flex-1">
              <a :href="item.url" class="list__title">
                {{ item.title }}
              </a>
              <small class="fn-nowrap" v-if="userCount > 1">
                By <a :href="item.articleAuthor.url">{{ item.articleAuthor.name }}</a>
              </small>
            </div>
            <div>
              <v-btn
                v-if="$store.state.name === item.author.name || $store.state.role < 3"
                class="btn btn--danger btn--small"
                @click="remove(item.id)">{{ $t('delete', $store.state.locale) }}
              </v-btn>
            </div>
          </div>
          <div class="content__reset" v-html="item.content"></div>
          <div class="list__meta">
            <time class="fn-nowrap">{{ item.createdAt }}</time>
          </div>
        </div>
      </li>
    </ul>
    <div class="pagination--wrapper fn-clear" v-if="list.length !== 0">
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
</template>

<script>
  export default {
    data () {
      return {
        currentPageNum: 1,
        pageCount: 1,
        windowSize: 1,
        list: [],
        userCount: 1,
        keyword: ''
      }
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t('commentList', this.$store.state.locale)}`
      }
    },
    methods: {
      async getList (currentPage) {
        const responseData = await this.axios.get(`/console/comments?p=${currentPage}&key=${this.keyword}`)
        if (responseData) {
          this.$set(this, 'userCount', responseData.userCount)
          this.$set(this, 'list', responseData.comments)
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', document.documentElement.clientWidth < 721 ? 5 : responseData.pagination.windowSize)
        }
      },
      async remove (id) {
        const responseData = await this.axios.delete(`/console/comments/${id}`)
        if (responseData === null) {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('deleteSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
          this.getList(1)
        }
      }
    },
    mounted () {
      this.getList(1)
    }
  }
</script>
