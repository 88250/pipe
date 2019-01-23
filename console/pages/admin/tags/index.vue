<template>
  <div>
    <div class="card fn__clear">
      <div class="card__body">
        <v-text-field
          @keyup.enter="getList()"
          class="fn__flex-1"
          :label="$t('enterSearch', $store.state.locale)"
          v-model="keyword">
        </v-text-field>
      </div>
      <ul class="list" v-if="list.length > 0">
        <li v-for="item in list" :key="item.id" class="fn__flex">
          <div class="fn__flex-1">
            <div class="fn__flex">
              <a target="_blank" class="list__title fn__flex-1"
                 @click.stop="openURL(item.url)"
                 href="javascript:void(0)">
                {{ item.title }}
              </a>
              <v-btn class="btn--danger btn--small" @click="remove(item.id)">
                {{ $t('delete', $store.state.locale) }}
              </v-btn>
            </div>
          </div>
        </li>
      </ul>
      <div class="pagination--wrapper fn__clear" v-if="pageCount > 1">
        <v-pagination
          :length="pageCount"
          v-model="currentPageNum"
          :total-visible="windowSize"
          class="fn__right"
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
        list: [],
        keyword: ''
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
        const responseData = await this.axios.get(`/console/taglist?p=${currentPage}&key=${this.keyword}`)
        if (responseData) {
          this.$set(this, 'list', responseData.tags || [])
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', document.documentElement.clientWidth < 721 ? 5 : responseData.pagination.windowSize)
        }
      },
      async remove (id) {
        const responseData = await this.axios.delete(`/console/tags/${id}`)
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
