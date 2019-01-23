<template>
  <div class="card">
    <category v-if="showForm" :show.sync="showForm" @addSuccess="addSuccess" :id="editId"></category>
    <div v-show="!showForm" class="card__body fn__clear">
      <v-btn class="btn--success" :class="{'fn__right': list.length > 0}" @click="edit(0)">{{ $t('new', $store.state.locale) }}</v-btn>
    </div>
    <ul class="list" v-if="list.length > 0">
      <li v-for="item in list" :key="item.id" class="fn__flex">
        <div class="fn__flex-1">
          <div class="fn__flex">
            <a class="list__title fn__flex-1"
               @click.stop="openURL(item.url)"
               href="javascript:void(0)">
              {{ item.title }}
            </a>
            <v-menu
              v-if="$store.state.role < 3"
              :nudge-bottom="28"
              :nudge-width="60"
              :nudge-left="60"
              :open-on-hover="true">
              <v-toolbar-title slot="activator">
                <v-btn class="btn--info btn--small" @click="edit(item.id)">
                  {{ $t('edit', $store.state.locale) }}
                  <v-icon>arrow_drop_down</v-icon>
                </v-btn>
              </v-toolbar-title>
              <v-list>
                <v-list-tile @click="edit(item.id)" class="list__tile--link">
                  {{ $t('edit', $store.state.locale) }}
                </v-list-tile>
                <v-list-tile @click="remove(item.id)" class="list__tile--link">
                  {{ $t('delete', $store.state.locale) }}
                </v-list-tile>
              </v-list>
            </v-menu>
          </div>
          <div class="list__meta">
            {{ item.description }}
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
  import Category from '~/components/biz/Category'

  export default {
    components: {
      Category
    },
    data () {
      return {
        editId: '',
        showForm: false,
        currentPageNum: 1,
        pageCount: 1,
        windowSize: 1,
        list: [],
      }
    },
    head () {
      return {
        title: `${this.$t('categoryList', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      openURL (url) {
        window.location.href = url
      },
      async getList (currentPage = 1) {
        const responseData = await this.axios.get(`/console/categories?p=${currentPage}`)
        if (responseData) {
          this.$set(this, 'list', responseData.categories || [])
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', document.documentElement.clientWidth < 721 ? 5 : responseData.pagination.windowSize)
        }
      },
      async remove (id) {
        const responseData = await this.axios.delete(`/console/categories/${id}`)
        if (responseData === null) {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: this.$t('deleteSuccess', this.$store.state.locale),
            snackModify: 'success'
          })
          this.getList()
          this.$set(this, 'showForm', false)
        }
      },
      addSuccess () {
        this.getList()
        this.$set(this, 'showForm', false)
      },
      edit (id) {
        this.$set(this, 'showForm', true)
        this.$set(this, 'editId', id)
      }
    },
    mounted () {
      this.getList()
    }
  }
</script>
