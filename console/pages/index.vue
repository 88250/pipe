<template>
  <div class="console" :style="flex" id="particles">
    <div class="card ft__center" ref="content">
      <h2 class="card__body"></h2>
      <div v-if="$store.state.role === 0" v-html="$t('index2', $store.state.locale)"></div>
      <div v-html="$t('index3', $store.state.locale)"></div>
      <br>
      <iframe src="https://ghbtns.com/github-btn.html?user=88250&repo=pipe&type=star&count=true&size=large"
              frameborder="0" scrolling="0" width="160px" height="30px"></iframe>
      <br><br>
      <h2 class="card__title">{{ $t('popularBlog', $store.state.locale) }}</h2>
      <ul class="list ">
        <li class="fn__flex" v-for="item in list">
          <a class="fn__flex-1" :href="item.url" target="_blank">{{ item.title }}</a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  import 'particles.js'
  import Vue from 'vue'
  import { initParticlesJS } from '~/plugins/utils'

  export default {
    head () {
      return {
        title: this.$t('welcome', this.$store.state.locale) + ' - Pipe'
      }
    },
    data () {
      return {
        flex: '',
        list: []
      }
    },
    async mounted () {
      const responseTopData = await this.axios.get('/blogs/top')
      if (responseTopData) {
        this.$set(this, 'list', responseTopData)
        Vue.nextTick(() => {
          if (this.$refs.content.scrollHeight > this.$refs.content.clientHeight) {
            this.$set(this, 'flex', 'flex:none')
          }
        })
      }
      initParticlesJS('particles')
    }
  }
</script>

<style lang="sass">
  .card__title
    display: block
</style>
