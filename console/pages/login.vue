<template>
  <div class="login" id="login">
    <div class="card">
      <h3 class="card__title">{{ $t('login', $store.state.locale) }}</h3>
      <div class="card__body fn-clear">
        <v-form class="init__center" ref="form">
          <v-text-field
            :label="$t('accountOrEmail', $store.state.locale)"
            v-model="accountOrEmail"
            :counter="16"
            :rules="userNameRules"
            required
          ></v-text-field>
          <v-text-field
            :label="$t('password', $store.state.locale)"
            v-model="password"
            :rules="userNameRules"
            :counter="16"
            required
            type="password"
            @keyup.enter="login"
          ></v-text-field>
          <div class="alert alert--danger" v-show="error">
            <v-icon>danger</v-icon>
            <span>{{ errorMsg }}</span>
          </div>
        </v-form>
        <v-btn class="fn-right btn btn--margin-t30 btn--info" @click="login">
          {{ $t('confirm', $store.state.locale) }}
        </v-btn>
      </div>
    </div>
  </div>
</template>

<script>
  import md5 from 'blueimp-md5'
  import 'particles.js'
  import { required, maxSize } from '~/plugins/validate'

  export default {
    layout: 'console',
    head () {
      return {
        title: this.$t('login', this.$store.state.locale)
      }
    },
    data () {
      return {
        accountOrEmail: '',
        password: '',
        userNameRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 16)
        ],
        error: false,
        errorMsg: ''
      }
    },
    methods: {
      async login () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.post('/login', {
          nameOrEmail: this.accountOrEmail,
          passwordHashed: md5(this.password)
        })
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$store.commit('setUserInfo', responseData.data)
          this.$router.push(this.$route.query.goto || '/admin')
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      }
    },
    mounted () {
      window.particlesJS('login', {
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
      }, function () {
        console.log('callback - particles.js config loaded')
      })
    }
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'
  .login
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
    .login
      padding: 0 15px
</style>
