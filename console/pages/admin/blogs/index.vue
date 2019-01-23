<template>
  <div class="card">
    <blog v-if="showForm" :show.sync="showForm" @addSuccess="addSuccess"></blog>

    <div v-show="!showForm" class="card__body fn__flex">
      <v-text-field
        v-if="list.length > 0"
        @keyup.enter="getList()"
        class="fn__flex-1"
        :label="$t('enterSearch', $store.state.locale)"
        v-model="keyword">
      </v-text-field>
      <v-btn class="btn--success" :class="{'btn--new': list.length > 0}" @click="edit">
        {{ $t('new', $store.state.locale) }}
      </v-btn>
    </div>

    <ul class="list" v-if="list.length > 0">
      <li v-for="item in list" :key="item.id" class="fn__flex">
        <div class="avatar avatar--mid avatar--space"
             :style="`background-image: url(${item.blogLogo})`"></div>
        <div class="fn__flex-1">
          <div class="fn__flex">
            <a class="list__title fn__flex-1" target="_blank" :href="`${item.blogURL}`">
              {{ item.blogTitle }}
            </a>
            <v-btn class="btn--small btn--info" @click="prohibit(item.id, 'unprohibit')" v-if="item.status === 1">
              {{ $t('unProhibit', $store.state.locale) }}
            </v-btn>
            <v-btn class="btn--small btn--danger" @click="prohibit(item.id, 'prohibit')" v-else>
              {{ $t('prohibit', $store.state.locale) }}
            </v-btn>
          </div>
          <div class="list__meta">
            {{ item.blogAdmin }}
            <span v-if="item.status === 1">
               • <span class="ft__danger"> {{ $t('prohibit', $store.state.locale) }} </span>
            </span>
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
</template>

<script>
  import Blog from '~/components/biz/Blog'

  export default {
    components: {
      Blog
    },
    data () {
      return {
        keyword: '',
        showForm: false,
        currentPageNum: 1,
        pageCount: 1,
        windowSize: 1,
        list: []
      }
    },
    head () {
      return {
        title: `${this.$t('blogManage', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      async getList (currentPage = 1) {
        const responseData = await this.axios.get(`/console/blogs?p=${currentPage}&k=${this.keyword}`)
        if (responseData) {
          this.$set(this, 'list', responseData.blogs)
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', document.documentElement.clientWidth < 721 ? 5 : responseData.pagination.windowSize)
        }
      },
      addSuccess () {
        this.getList()
        this.$set(this, 'showForm', false)
      },
      edit () {
        this.$set(this, 'showForm', true)
      },
      async prohibit (id, type) {
        const responseData = await this.axios.put(`/console/blogs/${id}/${type}`)
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.getList()
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      }
    },
    mounted () {
      this.getList()
    }
  }
</script>
