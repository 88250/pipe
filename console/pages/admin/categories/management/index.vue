<template>
  <div>
    <div class="card fn-clear">
      <category v-if="showForm" :show.sync="showForm" @addSuccess="addSuccess" :id="editId"></category>

      <div v-show="!showForm" class="card__body fn-clear">
        <button class="btn btn--success fn-right" @click="netCategory">{{ $t('new', $store.state.locale) }}</button>
      </div>
      <ul class="list">
        <li v-for="item in list" :key="item.id" class="fn-flex">
          <div class="fn-flex-1">
            <div class="list__title">
              <nuxt-link :to="`${$store.state.blogPath}/${item.permalink}`">
                {{ item.title }}
              </nuxt-link>
            </div>
            <div class="list__meta">
              {{ item.description }}
            </div>
          </div>
          <v-menu
            v-if="$store.state.role < 2"
            :nudge-bottom="38"
            :nudge-right="24"
            :nudge-width="100"
            :open-on-hover="true">
            <v-toolbar-title slot="activator">
              <button class="btn btn--info" @click="edit(item.id)">
                {{ $t('edit', $store.state.locale) }}
                <icon icon="chevron-down"/>
              </button>
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
        userCount: 1
      }
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t('categoryList', this.$store.state.locale)}`
      }
    },
    methods: {
      async getList (currentPage) {
        const responseData = await this.axios.get(`/console/categories?p=${currentPage}`)
        if (responseData) {
          this.$set(this, 'list', responseData.categories)
          this.$set(this, 'currentPageNum', responseData.pagination.currentPageNum)
          this.$set(this, 'pageCount', responseData.pagination.pageCount)
          this.$set(this, 'windowSize', responseData.pagination.windowSize)
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
          this.getList(1)
        }
      },
      addSuccess () {
        this.getList(1)
        this.$set(this, 'showForm', false)
      },
      edit (id) {
        this.$set(this, 'showForm', true)
        this.$set(this, 'editId', id)
      },
      netCategory () {
        this.$set(this, 'showForm', true)
        this.$set(this, 'editId', '')
      }
    },
    mounted () {
      this.getList(1)
    }
  }
</script>
