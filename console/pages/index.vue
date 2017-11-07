<template>
  <div class="console" id="particles">
    <div class="card" v-if="$store.state.role !== 0 && $store.state.role !== 4">
      <h2 class="card__title">{{ $t('welcome', $store.state.locale) }}</h2>
      <ul class="list">
        <li class="fn-flex" v-for="item in $store.state.blogs">
          <a class="fn-flex-1" :href="item.url">{{ item.title }}</a>
        </li>
      </ul>
    </div>
    <div class="card" v-else>
      <h2 class="card__body">{{ $t('welcome', $store.state.locale) }}</h2>
    </div>
  </div>
</template>

<script>
  import 'particles.js'
  import { initParticlesJS } from '~/plugins/utils'

  export default {
    head () {
      return {
        title: this.$t('welcome', this.$store.state.locale)
      }
    },
    mounted () {
      initParticlesJS('particles')
    }
  }
</script>
