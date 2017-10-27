<template>
  <div>
    <div class="card fn-clear">
      <blog v-if="showForm" :show.sync="showForm" @addSuccess="addSuccess" :id="editId"></blog>

      <div v-show="!showForm" class="card__body fn-clear">
        <v-btn class="btn--success fn-right" @click="edit(0)">{{ $t('new', $store.state.locale) }}</v-btn>
      </div>
      <ul class="list">
        <li v-for="item in list" :key="item.id" class="fn-flex">
          <div class="avatar avatar--mid avatar--space" :style="`background-image: url(${item.blogAdmin.avatarURL})`"></div>
          <div class="fn-flex-1">
            <div class="fn-flex">
              <nuxt-link class="list__title fn-flex-1" target="_blank" :to="`${item.blogURL}`">
                {{ item.blogTitle }}
              </nuxt-link>
              <v-menu
                v-if="$store.state.role < 3"
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
            </div>
            <div class="list__meta">
              {{ item.blogSubtitle }}
            </div>
          </div>
        </li>
      </ul>
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
          this.$set(this, 'windowSize', document.documentElement.clientWidth < 721 ? 5 : responseData.pagination.windowSize)
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
