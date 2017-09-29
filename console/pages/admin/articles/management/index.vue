<template>
  <div class="card fn-clear">
    <ul class="list">
      <li v-for="item in list" :key="item.id" class="fn-flex">
        <div v-if="userCount > 1">
          <div class="avatar avatar--mid avatar--space"
               :style="`background-image: url(${item.author.avatarURL})`"></div>
        </div>
        <div class="fn-flex-1">
          <div class="list__title">
            <nuxt-link :to="item.permalink">{{ item.title }}</nuxt-link>
          </div>
          <div class="list__meta">
            <nuxt-link :key="tag.title" v-for="tag in item.tags" :to="tag.permalink">{{ tag.title }}</nuxt-link>&nbsp;
            {{ item.commentCount }} {{ $t('comment', $store.state.locale) }} •
            {{ item.viewCount }} {{ $t('view', $store.state.locale) }} •
            <time>{{ item.createdAt }}</time>
            <span v-if="userCount > 1">
               • {{ item.author.name }}
            </span>
          </div>
        </div>
        <v-menu
          v-if="$store.state.name === item.author.name || $store.state.role < 2"
          :nudge-bottom="38"
          :nudge-right="24"
          :nudge-width="100"
          :open-on-hover="true">
          <v-toolbar-title slot="activator">
            <nuxt-link class="btn btn--info" :to="`/admin/articles/post?id=${item.id}`">
              {{ $t('edit', $store.state.locale) }}
              <icon icon="chevron-down"/>
            </nuxt-link>
          </v-toolbar-title>
          <v-list>
            <v-list-tile>
              <v-list-tile-title>
                <nuxt-link :to="`/admin/articles/post?id=${item.id}`">{{ $t('edit', $store.state.locale) }}</nuxt-link>
              </v-list-tile-title>
              <v-list-tile-title>
                <div @click="remove(item.id)">{{ $t('delete', $store.state.locale) }}</div>
              </v-list-tile-title>
              <v-list-tile-title>
                <div @click="top(item.id)">{{ $t('top', $store.state.locale) }}</div>
              </v-list-tile-title>
            </v-list-tile>
          </v-list>
        </v-menu>
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
        title: `${this.$store.state.blogTitle} - ${this.$t('articleList', this.$store.state.locale)}`
      }
    },
    methods: {
      async getList (currentPage) {
        const responseData = await this.axios.get(`/console/articles?p=${currentPage}`)
        if (responseData) {
          this.$set(this, 'userCount', responseData.userCount)
          this.$set(this, 'list', responseData.articles)
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', responseData.pagination.windowSize)
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
      }
    },
    mounted () {
      this.getList(1)
    }
  }
</script>
