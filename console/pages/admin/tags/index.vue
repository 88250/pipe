<template>
  <div>
    <div class="card fn-clear">
      <ul class="list" v-if="list.length > 0">
        <li v-for="item in list" :key="item.id" class="fn-flex">
          <div class="fn-flex-1">
            <div class="fn-flex">
              <a target="_blank" class="list__title fn-flex-1"
                 @click.stop="openURL(item.url)"
                 href="javascript:void(0)">
                {{ item.title }}
              </a>
              <v-btn class="btn--danger btn--small" @click="remove(item.title)">
                {{ $t('delete', $store.state.locale) }}
              </v-btn>
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
  export default {
    data () {
      return {
        editId: '',
        showForm: false,
        currentPageNum: 1,
        pageCount: 1,
        windowSize: 1,
        list: []
      }
    },
    head () {
      return {
        title: `${this.$t('tagList', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      openURL (url) {
        window.location.href = url
      },
      async getList (currentPage = 1) {
        const responseData = await this.axios.get(`/console/taglist?p=${currentPage}`)
        if (responseData) {
          this.$set(this, 'list', responseData.tags || [])
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', document.documentElement.clientWidth < 721 ? 5 : responseData.pagination.windowSize)
        }
      },
      async remove (title) {
        const responseData = await this.axios.delete(`/console/tags/${title}`)
        if (responseData === null) {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('deleteSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
          this.getList()
          this.$set(this, 'showForm', false)
        }
      }
    },
    mounted () {
      this.getList()
    }
  }
</script>
