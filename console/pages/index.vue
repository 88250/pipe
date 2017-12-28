<template>
  <div class="console" id="particles">
    <div class="card" v-if="$store.state.role !== 0">
      <h2 class="card__title">{{ $t('welcome', $store.state.locale) }}</h2>
      <ul class="list">
        <li class="fn-flex" v-for="item in $store.state.blogs">
          <a class="fn-flex-1" :href="item.url">{{ item.title }}</a>
        </li>
      </ul>
    </div>
    <div class="card" v-else>
      <h2 class="card__body">{{ $t('welcome', $store.state.locale) }}</h2>
      <div class="ft-center">
        <div v-html="$t('index2', $store.state.locale)"></div>
        <br>
        <iframe src="https://ghbtns.com/github-btn.html?user=b3log&repo=pipe&type=star&count=true&size=large"
                frameborder="0" scrolling="0" width="160px" height="30px"></iframe><br><br>
      </div>
    </div>
  </div>
</template>

<script>
  import 'particles.js'
  import { initParticlesJS } from '~/plugins/utils'

  export default {
    head () {
      return {
        title: this.$t('welcome', this.$store.state.locale) + ' - Pipe'
      }
    },
    mounted () {
      initParticlesJS('particles')
    }
  }
</script>
