<template>
  <div class="card">
    <div class="card__body fn__flex" v-show="!isBatch">
      <v-text-field v-if="list.length > 0 || isSearch"
        @keyup.enter="getList();isSearch = true"
        class="fn__flex-1"
        :label="$t('enterSearch', $store.state.locale)"
        v-model="keyword">
      </v-text-field>
      <nuxt-link to="/admin/articles/post" class="btn btn--success" :class="{'btn--new': list.length > 0 || isSearch}">{{ $t('new', $store.state.locale)
        }}
      </nuxt-link>
    </div>
    <div class="card__batch-action fn__flex" v-show="isBatch">
      <label class="checkbox fn__flex-1">
        <input
          type="checkbox"
          :checked="isSelectAll"
          @click="selectAll"/>
        <span class="checkbox__icon"></span>
        {{ $t('hasChoose', $store.state.locale) }}
        {{selectedIds.length}}
        {{ $t('cntContent', $store.state.locale) }}
      </label>
      <div>
        <v-btn
          :class="{'btn--disabled': selectedIds.length === 0}"
          class="btn--danger"
          @click="batchAction">
          {{ $t('delete', $store.state.locale) }}
        </v-btn>
        <v-btn class="btn--success btn--space" @click="isBatch = false; selectedIds = []">
          {{ $t('cancel', $store.state.locale) }}
        </v-btn>
      </div>
    </div>
    <ul class="list" v-if="list.length > 0">
      <li
        v-for="item in list"
        :key="item.id"
        class="fn__flex"
        :class="{'selected': isSelected(item.id)}"
        @click="setSelectedId(item.id)">
        <a class="avatar avatar--mid avatar--space pipe-tooltipped pipe-tooltipped--s"
           v-if="userCount > 1"
           :aria-label="item.author.name"
           @click.stop="openURL(item.author.url)"
           href="javascript:void(0)"
           :style="`background-image: url(${item.author.avatarURL})`"></a>
        <div class="fn__flex-1">
          <div class="fn__flex">
            <span class="fn__flex-1">
              <a class="list__title" @click.stop="openURL(item.url)" href="javascript:void(0)">{{ item.title }}</a>
            </span>
            <v-menu
              v-show="!isBatch"
              v-if="$store.state.name === item.author.name || $store.state.role < 3"
              :nudge-bottom="28"
              :nudge-width="60"
              :nudge-left="60"
              :open-on-hover="true">
              <v-toolbar-title slot="activator">
                <v-btn class="btn--small btn--info" @click.stop="goEdit(item.id)">
                  {{ $t('edit', $store.state.locale) }}
                  <v-icon>arrow_drop_down</v-icon>
                </v-btn>
              </v-toolbar-title>
              <v-list>
                <v-list-tile class="list__tile--link" @click.stop="goEdit(item.id)">
                  {{ $t('edit', $store.state.locale) }}
                </v-list-tile>
                <v-list-tile class="list__tile--link" @click.stop="syncToCommunity(item.id)">
                  {{ $t('syncToCommunity', $store.state.locale) }}
                </v-list-tile>
                <v-list-tile class="list__tile--link" @click.stop="remove(item.id)">
                  {{ $t('delete', $store.state.locale) }}
                </v-list-tile>
              </v-list>
            </v-menu>
          </div>
          <div class="list__meta">
            <span class="tags">
              <a class="tag" :key="tag.title" v-for="tag in item.tags"
                 @click.stop="openURL(tag.url)"
                 href="javascript:void(0)">{{ tag.title
                }}</a>
            </span>
            <span class="fn-nowrap">{{ item.commentCount }} {{ $t('comment', $store.state.locale) }}</span> •
            <span class="fn-nowrap">{{ item.viewCount }} {{ $t('view', $store.state.locale) }}</span> •
            <time class="fn-nowrap">{{ item.createdAt }}</time>
          </div>
        </div>
      </li>
    </ul>
    <ul class="list" v-else-if="isSearch && !isBatch">
      <li class="fn__flex fn__pointer" @click="isSearch = false;keyword = '';getList();">
        {{ $t('searchNull', $store.state.locale) }} &nbsp;
        <span class="ft__danger">{{keyword}}</span>
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
  export default {
    data () {
      return {
        isSearch: false,
        isSelectAll: false,
        isBatch: false,
        selectedIds: [],
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
        title: `${this.$t('articleList', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      openURL (url) {
        window.location.href = url
      },
      selectAll () {
        this.$set(this, 'isSelectAll', !this.isSelectAll)
        if (!this.isSelectAll) {
          this.$set(this, 'selectedIds', [])
          return
        }

        const selectedIds = []
        this.list.forEach((data) => {
          selectedIds.push(data.id)
        })
        this.$set(this, 'selectedIds', selectedIds)
      },
      isSelected (id) {
        let isSelected = false
        this.selectedIds.forEach((data) => {
          if (data === id) {
            isSelected = true
          }
        })
        return isSelected
      },
      async batchAction () {
        if (this.selectedIds.length === 0) {
          return
        }
        if (!confirm(this.$t('confirmDelete', this.$store.state.locale))) {
          return
        }
        const responseData = await this.axios.post('/console/articles/batch-delete', {
          ids: this.selectedIds
        })
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$set(this, 'isSelectAll', false)
          this.$set(this, 'selectedIds', [])
          this.getList()
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      },
      setSelectedId (id) {
        let isSelected = false
        this.selectedIds.forEach((data) => {
          if (data === id) {
            isSelected = true
          }
        })

        if (isSelected) {
          this.$set(this, 'selectedIds', this.selectedIds.filter((data) => id !== data))
          if (this.selectedIds.length < 1) {
            this.$set(this, 'isBatch', false)
          }
        } else {
          this.$set(this, 'isBatch', true)
          this.selectedIds.push(id)
        }

        if (this.selectedIds.length === this.list.length) {
          this.$set(this, 'isSelectAll', true)
        } else {
          this.$set(this, 'isSelectAll', false)
        }
      },
      async getList (currentPage = 1) {
        const responseData = await this.axios.get(`/console/articles?p=${currentPage}&key=${this.keyword}`)
        if (responseData) {
          this.$set(this, 'userCount', responseData.userCount)
          this.$set(this, 'list', responseData.articles || [])
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', document.documentElement.clientWidth < 721 ? 5 : responseData.pagination.windowSize)
        }
      },
      async remove (id) {
        if (!confirm(this.$t('confirmDelete', this.$store.state.locale))) {
          return
        }
        const responseData = await this.axios.delete(`/console/articles/${id}`)
        if (responseData === null) {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('deleteSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
          this.getList()
        }
      },
      async syncToCommunity (id) {
        const responseData = await this.axios.get(`/console/articles/${id}/push`)
        if (responseData === null) {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('syncSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
        }
      },
      goEdit (id) {
        this.$router.push(`/admin/articles/post?id=${id}`)
      }
    },
    mounted () {
      this.getList()
    }
  }
</script>
