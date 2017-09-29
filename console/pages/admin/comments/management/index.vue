<template>
  <div class="card fn-clear">
    <ul class="list">
      <li v-for="item in list" :key="item.id" class="fn-flex">
        <div class="avatar avatar--mid avatar--space"
             :style="`background-image: url(${item.author.avatarURL})`"></div>
        <div class="fn-flex-1">
          <div class="list__title">
            <nuxt-link :to="item.permalink">{{ item.title }}</nuxt-link>
          </div>
          <div class="content-reset" v-html="item.content"></div>
          <div class="list__meta">
            <time>{{ item.createdAt }}</time>
            • {{ item.author.name }}
          </div>
        </div>
        <div>
          <button
            v-if="$store.state.name === item.author.name || $store.state.role < 2"
            class="btn btn--danger"
            @click="remove(item.id)">{{ $t('delete', $store.state.locale) }}
          </button>
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
</template>

<script>
  export default {
    data () {
      return {
        currentPageNum: 1,
        pageCount: 1,
        windowSize: 1,
        list: [],
        userCount: 1
      }
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t('commentList', this.$store.state.locale)}`
      }
    },
    methods: {
      async getList (currentPage) {
        const responseData = await this.axios.get(`/console/comments?p=${currentPage}`)
        if (responseData) {
          this.$set(this, 'userCount', responseData.userCount)
          this.$set(this, 'list', responseData.comments)
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', responseData.pagination.windowSize)
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
