<template>
  <div>
    <div class="card fn-clear card__body card--space" v-if="$store.state.role <= 1">
      <h1>{{ $t('import', $store.state.locale) }}</h1>
      <button class="btn btn--info">JSON</button>
      <button class="btn btn--info btn--space">{{ $t('staticBlog', $store.state.locale) }}</button>
    </div>

    <div class="card fn-clear card__body card--space">
      <h1>{{ $t('export', $store.state.locale) }}</h1>
      <button class="btn btn--info">JSON</button>
      <button class="btn btn--info btn--space">HTML</button>
    </div>

    <div class="card fn-clear card__body card--space">
      <h1>{{ $t('tags', $store.state.locale) }}</h1>
      <button class="btn btn--danger">{{ $t('removeUnusedTags', $store.state.locale) }}</button>
    </div>

    <div class="card fn-clear card__body">
      <h1>B3log key</h1>
      {{ key }}
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        key: ''
      }
    },
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t('others', this.$store.state.locale)}`
      }
    },
    async mounted () {
      const responseData = await this.axios.get('/console/b3log-key')
      if (responseData) {
        this.$set(this, 'key', responseData)
      }
    }
  }
</script>
