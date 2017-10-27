<template>
  <div class="card">
    <div class="card__body fn-flex">
      <v-text-field v-if="list.length > 0"
        @keyup.enter="getList(1)"
        class="fn-flex-1"
        :label="$t('enterSearch', $store.state.locale)"
        v-model="keyword">
      </v-text-field>
      <nuxt-link to="/admin/articles/post" class="btn btn--success" :class="{'btn--new': list.length > 0}">{{ $t('new', $store.state.locale)
        }}
      </nuxt-link>
    </div>
    <ul class="list" v-if="list.length > 0">
      <li v-for="item in list" :key="item.id" class="fn-flex">
        <a class="avatar avatar--mid avatar--space tooltipped tooltipped--s"
           v-if="userCount > 1"
           :aria-label="item.author.name"
           :href="item.author.url"
           :style="`background-image: url(${item.author.avatarURL})`"></a>
        <div class="fn-flex-1">
          <div class="fn-flex">
            <a class="fn-flex-1 list__title" :href="item.url">{{ item.title }}</a>
            <v-menu
              v-if="$store.state.name === item.author.name || $store.state.role < 3"
              :nudge-bottom="28"
              :nudge-width="60"
              :nudge-left="60"
              :open-on-hover="true">
              <v-toolbar-title slot="activator">
                <nuxt-link class="btn btn--small btn--info" :to="`/admin/articles/post?id=${item.id}`">
                  {{ $t('edit', $store.state.locale) }}
                  <v-icon>arrow_drop_down</v-icon>
                </nuxt-link>
              </v-toolbar-title>
              <v-list>
                <v-list-tile class="list__tile--link" @click="goEdit(item.id)">
                  {{ $t('edit', $store.state.locale) }}
                </v-list-tile>
                <v-list-tile class="list__tile--link" @click="remove(item.id)">
                  {{ $t('delete', $store.state.locale) }}
                </v-list-tile>
                <v-list-tile class="list__tile--link" @click="top(item.id)">
                  {{ $t('top', $store.state.locale) }}
                </v-list-tile>
              </v-list>
            </v-menu>
          </div>
          <div class="list__meta">
            <span class="tags">
              <a class="tag" :key="tag.title" v-for="tag in item.tags" :href="tag.url">{{ tag.title
                }}</a>
            </span>
            <span class="fn-nowrap">{{ item.commentCount }} {{ $t('comment', $store.state.locale) }}</span> •
            <span class="fn-nowrap">{{ item.viewCount }} {{ $t('view', $store.state.locale) }}</span> •
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
        title: `${this.$store.state.blogTitle} - ${this.$t('articleList', this.$store.state.locale)}`
      }
    },
    methods: {
      async getList (currentPage) {
        const responseData = await this.axios.get(`/console/articles?p=${currentPage}&key=${this.keyword}`)
        if (responseData) {
          this.$set(this, 'userCount', responseData.userCount)
          this.$set(this, 'list', responseData.articles)
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', document.documentElement.clientWidth < 721 ? 5 : responseData.pagination.windowSize)
        }
      },
      async remove (id) {
        const responseData = await this.axios.delete(`/console/articles/${id}`)
        if (responseData === null) {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('deleteSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
          this.getList(1)
        }
      },
      async top (id) {
        const responseData = await this.axios.put(`/console/articles/${id}`)
        if (responseData.code === 0) {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('topSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
          this.getList(1)
        } else {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: responseData.msg
          })
        }
      },
      goEdit (id) {
        this.$router.push(`/admin/articles/post?id=${id}`)
      }
    },
    mounted () {
      this.getList(1)
    }
  }
</script>
