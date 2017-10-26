<template>
  <div class="card">
    <user v-if="showForm" :show.sync="showForm" @addSuccess="addSuccess" :id="editId"></user>

    <div v-show="!showForm" class="card__body" >
      <v-text-field
        v-if="list.length > 0"
        @keyup.enter="getList(1)"
        class="fn-flex-1"
        :label="$t('enterSearch', $store.state.locale)"
        v-model="keyword">
      </v-text-field>
      <template v-else>
        {{ $t('noData', $store.state.locale) }}
      </template>
    </div>

    <ul class="list" v-if="list.length > 0">
      <li v-for="item in list" :key="item.id" class="fn-flex"
          v-if="($store.state.role === 2 && item.name === $store.state.name) || $store.state.role < 2">
        <a :href="item.url"
           :aria-label="item.name"
           class="avatar avatar--mid avatar--space tooltipped tooltipped--s"
           :style="`background-image: url(${item.avatarURL})`"></a>
        <div class="fn-flex-1">
          <div class="fn-flex">
            <a class="list__title fn-flex-1" :href="item.url">
              {{ item.nickname || item.name }}
            </a>
            <v-menu
              :nudge-bottom="28"
              :nudge-width="60"
              :nudge-left="60"
              :open-on-hover="true">
              <v-toolbar-title slot="activator">
                <v-btn class="btn--small btn--info" @click="edit(item.id)">
                  {{ $t('edit', $store.state.locale) }}
                  <v-icon>arrow_drop_down</v-icon>
                </v-btn>
              </v-toolbar-title>
              <v-list>
                <v-list-tile @click="edit(item.id)">
                  {{ $t('edit', $store.state.locale) }}
                </v-list-tile>
                <v-list-tile @click="remove(item.id)">
                  {{ $t('delete', $store.state.locale) }}
                </v-list-tile>
              </v-list>
            </v-menu>
          </div>
          <div class="list__meta">
            <span class="fn-nowrap">{{ item.email }}</span> •
            <span class="fn-nowrap">{{ item.PublishedArticleCount }} {{ $t('article', $store.state.locale) }}</span> •
            <span class="fn-nowrap">{{ getRoleName(item.role) }}</span>
          </div>
        </div>
      </li>
    </ul>
    <div class="pagination--wrapper fn-clear" v-if="list.length > 0">
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
  import User from '~/components/biz/User'

  export default {
    components: {
      User
    },
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
        title: `${this.$store.state.blogTitle} - ${this.$t('userList', this.$store.state.locale)}`
      }
    },
    methods: {
      getRoleName (role) {
        let roleName = this.$t('commonUser', this.$store.state.locale)
        switch (role) {
          case 0:
            roleName = this.$t('superAdmin', this.$store.state.locale)
            break
          case 1:
            roleName = this.$t('blogAdmin', this.$store.state.locale)
            break
          default:
            break
        }
        return roleName
      },
      async getList (currentPage) {
        const responseData = await this.axios.get(`/console/users?p=${currentPage}&key=${this.keyword}`)
        if (responseData) {
          this.$set(this, 'list', responseData.navigation)
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', document.documentElement.clientWidth < 721 ? 5 : responseData.pagination.windowSize)
        }
      },
      async remove (id) {
        const responseData = await this.axios.delete(`/console/users/${id}`)
        if (responseData === null) {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('deleteSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
          this.getList(1)
          this.$set(this, 'showForm', false)
        }
      },
      addSuccess () {
        this.getList(1)
        this.$set(this, 'showForm', false)
      },
      edit (id) {
        this.$set(this, 'showForm', true)
        this.$set(this, 'editId', id)
      }
    },
    mounted () {
      this.getList(1)
    }
  }
</script>
