<template>
  <div class="card__body fn__clear">
    <v-form ref="form">
      <v-text-field
        :label="$t('title', $store.state.locale)"
        v-model="title"
        :counter="128"
        :rules="titleRules"
        required
      ></v-text-field>
      <v-text-field
        :label="$t('links', $store.state.locale)"
        v-model="url"
        :rules="linkRules"
        :counter="255"
        required
      ></v-text-field>
      <v-text-field
        :label="$t('iconPath', $store.state.locale)"
        v-model="iconURL"
        :rules="iconURLRules"
        :counter="255"
      ></v-text-field>
      <v-select
        :label="$t('openMethod', $store.state.locale)"
        :items="openMethods"
        v-model="openMethod"
        append-icon=""
      ></v-select>
      <div class="alert alert--danger" v-show="error">
        <v-icon>danger</v-icon>
        <span>{{ errorMsg }}</span>
      </div>
    </v-form>
    <v-btn class="fn__right btn--margin-t30 btn--info btn--space" @click="created">
      {{ $t('confirm', $store.state.locale) }}
    </v-btn>
    <v-btn class="fn__right btn--margin-t30 btn--danger btn--space" @click="$emit('update:show', false)">
      {{ $t('cancel', $store.state.locale) }}
    </v-btn>
  </div>
</template>

<script>
  import { required, maxSize } from '~/plugins/validate'

  export default {
    props: {
      id: {
        type: Number,
        required: true
      }
    },
    data () {
      return {
        errorMsg: '',
        error: false,
        title: '',
        url: '',
        iconURL: '',
        openMethod: '',
        openMethods: [
          {
            'text': this.$t('openMethod1', this.$store.state.locale),
            'value': ''
          },
          {
            'text': this.$t('openMethod2', this.$store.state.locale),
            'value': '_blank'
          },
          {
            'text': this.$t('openMethod3', this.$store.state.locale),
            'value': '_parent'
          },
          {
            'text': this.$t('openMethod4', this.$store.state.locale),
            'value': '_top'
          }
        ],
        titleRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 128)
        ],
        linkRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 255)
        ],
        iconURLRules: [
          (v) => maxSize.call(this, v, 255)
        ]
      }
    },
    watch: {
      id: function () {
        this.init()
      }
    },
    methods: {
      async created () {
        if (!this.$refs.form.validate()) {
          return
        }
        let responseData = {}
        const requestData = {
          title: this.title,
          url: this.url,
          iconURL: this.iconURL,
          openMethod: this.openMethod
        }
        if (this.id === 0) {
          responseData = await this.axios.post('/console/navigations', requestData)
        } else {
          responseData = await this.axios.put(`/console/navigations/${this.id}`, requestData)
        }

        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$emit('addSuccess')
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      },
      async init () {
        if (this.id === 0) {
          return
        }
        const responseData = await this.axios.get(`/console/navigations/${this.id}`)
        if (responseData) {
          this.$set(this, 'title', responseData.title)
          this.$set(this, 'url', responseData.url)
          this.$set(this, 'iconURL', responseData.iconURL)
          this.$set(this, 'openMethod', responseData.openMethod)
        }
      }
    },
    mounted () {
      this.init()
    }
  }
</script>
