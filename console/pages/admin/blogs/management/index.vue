<template>
  <div>
    <div class="card fn-clear">
      <blog v-if="showForm" :show.sync="showForm" @addSuccess="addSuccess" :id="editId"></blog>

      <div v-show="!showForm" class="card__body fn-clear">
        <v-btn class="btn btn--success fn-right" @click="edit('')">{{ $t('new', $store.state.locale) }}</v-btn>
      </div>
      <ul class="list">
        <li v-for="item in list" :key="item.id" class="fn-flex">
          <div class="avatar avatar--mid avatar--space" :style="`background-image: url(${item.blogAdmin.avatarURL})`"></div>
          <div class="fn-flex-1">
            <div class="list__title">
              <nuxt-link target="_blank" :to="`/${item.blogPath}`">
                {{ item.blogTitle }}
              </nuxt-link>
            </div>
            <div class="list__meta">
              {{ item.blogSubtitle }}
            </div>
          </div>
          <v-menu
            v-if="$store.state.role < 2"
            :nudge-bottom="38"
            :nudge-right="24"
            :nudge-width="100"
            :open-on-hover="true">
            <v-toolbar-title slot="activator">
              <v-btn class="btn btn--info" @click="edit(item.id)">
                {{ $t('edit', $store.state.locale) }}
                <icon icon="chevron-down"/>
              </v-btn>
            </v-toolbar-title>
            <v-list>
              <v-list-tile>
                <v-list-tile-title>
                  <div @click="edit(item.id)">{{ $t('edit', $store.state.locale) }}</div>
                </v-list-tile-title>
                <v-list-tile-title>
                  <div @click="remove(item.id)">{{ $t('delete', $store.state.locale) }}</div>
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
        title: `${this.$store.state.blogTitle} - ${this.$t('blogManage', this.$store.state.locale)}`
      }
    },
    methods: {
      async getList (currentPage) {
        const responseData = await this.axios.get(`/console/blogs?p=${currentPage}`)
        if (responseData) {
          this.$set(this, 'list', responseData.navigation)
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', responseData.pagination.windowSize)
        }
      },
      async remove (id) {
        const responseData = await this.axios.delete(`/console/blogs/${id}`)
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
