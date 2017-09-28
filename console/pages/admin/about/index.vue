<template>
  <div class="card">
    <div class="card__body">
      <div class="fn-flex admin__about">
        <div class="about__side">
          <img src="~/static/images/logo.jpg"/> <br>
          <nuxt-link class="btn btn--info btn--margin-t30 btn--block" to="http://b3log.org/donate.html">{{ $t('becomeSponsor', $store.state.locale) }}</nuxt-link>
        </div>
        <div class="fn-flex-1 content-reset">
          <h2 v-if="isLatest">
            {{ $t('about1', $store.state.locale) }}
            <a :href="download" target="_blank">{{ version }}</a>
          </h2>
          <h2 v-else>
            {{ $t('about2', $store.state.locale) }}
            <a :href="download" target="_blank">{{ version }}</a>
          </h2>
          <div v-html="$t('about3', $store.state.locale)"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        isLatest: true,
        download: '',
        version: ''
      }
    },
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('about', this.$store.state.locale)}`
      }
    },
    async mounted () {
      const responseData = await this.axios.get(`/hp/apis/check-version`)
      if (responseData) {
        this.$set(this, 'isLatest', this.$store.state.version === responseData.version)
        this.$set(this, 'download', responseData.download)
        this.$set(this, 'version', responseData.version)
      }
    }
  }
</script>

<style lang="sass">
  .admin__about .about__side
    margin: 40px 50px 0 30px
</style>
