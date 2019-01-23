<template>
  <div class="card">
    <user v-if="showForm" :show.sync="showForm" @addSuccess="addSuccess"></user>

    <div v-show="!showForm" class="card__body fn__flex">
      <v-text-field
        v-if="list.length > 0"
        @keyup.enter="getList()"
        class="fn__flex-1"
        :label="$t('enterSearch', $store.state.locale)"
        v-model="keyword">
      </v-text-field>
      <v-btn class="btn--success" :class="{'btn--new': list.length > 0}" @click="edit">{{ $t('new', $store.state.locale) }}</v-btn>
    </div>

    <ul class="list" v-if="list.length > 0">
      <li v-for="item in list" :key="item.id" class="fn__flex"
          v-if="($store.state.role === 3 && item.name === $store.state.name) || $store.state.role < 3">
        <a :href="item.url"
           :aria-label="item.name"
           class="avatar avatar--mid avatar--space pipe-tooltipped pipe-tooltipped--n"
           :style="`background-image: url(${item.avatarURL})`"></a>
        <div class="fn__flex-1">
          <div class="fn__flex">
            <a class="list__title fn__flex-1" :href="item.url">
              {{ item.nickname || item.name }}
            </a>
            <v-btn class="btn--small btn--info" @click="prohibit(item.id, 'unprohibit')" v-if="item.role === 4">
              {{ $t('unProhibit', $store.state.locale) }}
            </v-btn>
            <v-btn class="btn--small btn--danger" @click="prohibit(item.id, 'prohibit')" v-else>
              {{ $t('prohibit', $store.state.locale) }}
            </v-btn>
          </div>
          <div class="list__meta">
            <span class="fn-nowrap">{{ item.articleCount }} {{ $t('article', $store.state.locale) }}</span> •
            <span class="fn-nowrap" :class="{'ft__danger': item.role === 4}">{{ getRoleName(item.role) }}</span>
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
  import User from '~/components/biz/User'

  export default {
    components: {
      User
    },
    data () {
      return {
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
        title: `${this.$t('userList', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      getRoleName (role) {
        let roleName = this.$t('blogUser', this.$store.state.locale)
        switch (role) {
          case 1:
            roleName = this.$t('superAdmin', this.$store.state.locale)
            break
          case 2:
            roleName = this.$t('blogAdmin', this.$store.state.locale)
            break
          case 3:
            roleName = this.$t('blogUser', this.$store.state.locale)
            break
          case 4:
            roleName = this.$t('prohibitUser', this.$store.state.locale)
            break
          default:
            break
        }
        return roleName
      },
      async getList (currentPage = 1) {
        const responseData = await this.axios.get(`/console/users?p=${currentPage}&key=${this.keyword}`)
        if (responseData) {
          this.$set(this, 'list', responseData.users)
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', document.documentElement.clientWidth < 721 ? 5 : responseData.pagination.windowSize)
        }
      },
      async prohibit (id, type) {
        const responseData = await this.axios.put(`/console/users/${id}/${type}`)
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.getList()
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      },
      addSuccess () {
        this.getList()
        this.$set(this, 'showForm', false)
      },
      edit () {
        this.$set(this, 'showForm', true)
      }
    },
    mounted () {
      this.getList()
    }
  }
</script>
