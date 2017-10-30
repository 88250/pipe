<template>
  <div class="welcome" id="index">
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

  export default {
    layout: 'console',
    head () {
      return {
        title: this.$t('welcome', this.$store.state.locale)
      }
    },
    mounted () {
      window.particlesJS('index', {
        'particles': {
          'number': {
            'value': 6,
            'density': {
              'enable': true,
              'value_area': 200
            }
          },
          'color': {
            'value': '#bbb'
          },
          'opacity': {
            'value': 0.5,
            'anim': {
              'speed': 1,
              'opacity_min': 0.1
            }
          },
          'size': {
            'value': 10,
            'random': true,
            'anim': {
              'enable': false,
              'speed': 80,
              'size_min': 0.1,
              'sync': false
            }
          },
          'line_linked': {
            'enable': true,
            'distance': 300,
            'color': '#bbb',
            'opacity': 0.4,
            'width': 1
          },
          'move': {
            'enable': true,
            'speed': 2,
            'direction': 'none',
            'straight': false,
            'out_mode': 'out',
            'bounce': false,
            'attract': {
              'enable': false,
              'rotateX': 300,
              'rotateY': 600
            }
          }
        },
        'interactivity': {
          'detect_on': 'canvas',
          'events': {
            'onclick': {
              'enable': false
            },
            'resize': true
          },
          'modes': {
            'grab': {
              'distance': 400,
              'line_linked': {
                'opacity': 0.7
              }
            },
            'bubble': {
              'distance': 800,
              'size': 80,
              'duration': 2,
              'opacity': 0.8,
              'speed': 3
            },
            'repulse': {
              'distance': 400,
              'duration': 0.4
            },
            'push': {
              'particles_nb': 4
            },
            'remove': {
              'particles_nb': 2
            }
          }
        },
        'retina_detect': true
      })
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'
  .welcome
    background-color: $blue-lighter
    position: relative
    flex: 1
    display: flex
    align-items: center
    overflow: hidden
    .particles-js-canvas-el
      position: absolute
      top: 0
    .card
      width: 650px
      margin: 0 auto
      position: relative
      z-index: 1

  @media (max-width: 768px)
    .welcome
      padding: 0 15px
</style>
