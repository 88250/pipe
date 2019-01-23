<template>
  <div>
    <div class="card card--space">
      <ul class="list">
        <li class="fn__flex" v-if="$store.state.role <= 3">
          <div class="fn__flex-1">
            {{ $t('import', $store.state.locale) }}
          </div>
          <a href="https://hacpai.com/article/1498490209748" target="_blank"><v-icon>question</v-icon></a>
          <label class="btn btn--small btn--info other__upload">
            {{ $t('staticBlog', $store.state.locale) }}
            <input @change="importMD" type="file"/>
          </label>
        </li>
        <li class="fn__flex">
          <div class="fn__flex-1">
            {{ $t('export', $store.state.locale) }}
          </div>
          <a class="btn--small btn--info btn" href="/api/console/export/md">
            {{ $t('staticBlog', $store.state.locale) }}
          </a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    head () {
      return {
        title: `${this.$t('others', this.$store.state.locale)} - ${this.$store.state.blogTitle}`
      }
    },
    methods: {
      async importMD (event) {
        const formData = new FormData()
        formData.append('file', event.target.files[0])
        const responseData = await this.axios.post('/console/import/md', formData)
        if (responseData.code === 0) {
          event.target.value = ''
        } else {
          this.$store.commit('setSnackBar', {
            snackBar: true,
            snackMsg: responseData.msg
          })
        }
      }
    }
  }
</script>
<style lang="sass">
  .other__upload
    cursor: pointer
    padding: 0 10px
    margin-left: 20px
    input
      position: absolute
      width: 1px
      opacity: .001
      height: 16px
      overflow: hidden
</style>
