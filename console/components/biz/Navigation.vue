<template>
  <div class="card__body fn-clear">
    <v-form ref="form">
      <v-text-field
        :label="$t('title', $store.state.locale)"
        v-model="title"
        :counter="32"
        :rules="titleRules"
        required
      ></v-text-field>
      <v-text-field
        :label="$t('links', $store.state.locale)"
        v-model="permalink"
        :rules="titleRules"
        :counter="32"
        required
      ></v-text-field>
      <v-text-field
        :label="$t('iconPath', $store.state.locale)"
        v-model="iconURL"
        :rules="iconURLRules"
        :counter="32"
      ></v-text-field>
      <v-select
        :label="$t('openMethod', $store.state.locale)"
        :items="openMethods"
        v-model="openMethod"
        append-icon=""
      ></v-select>
      <div class="alert alert--danger" v-show="error">
        <icon icon="danger"/>
        <span>{{ errorMsg }}</span>
      </div>
    </v-form>
    <button class="fn-right btn btn--margin-t30 btn--info btn--space" @click="created">
      {{ $t('confirm', $store.state.locale) }}
    </button>
    <button class="fn-right btn btn--margin-t30 btn--danger btn--space" @click="$emit('update:show', false)">
      {{ $t('cancel', $store.state.locale) }}
    </button>
  </div>
</template>

<script>
  export default {
    props: ['id'],
    data () {
      return {
        errorMsg: '',
        error: false,
        title: '',
        permalink: '',
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
          (v) => !!v || this.$t('required', this.$store.state.locale),
          (v) => v.length <= 32 || this.$t('validateRule2', this.$store.state.locale)
        ],
        iconURLRules: [
          (v) => v.length <= 32 || this.$t('validateRule2', this.$store.state.locale)
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
        if (this.id === '') {
          responseData = await this.axios.post('/console/navigations', {
            title: this.title,
            permalink: this.permalink,
            iconURL: this.iconURL,
            openMethod: this.openMethod
          })
        } else {
          responseData = await this.axios.put(`/console/navigations/${this.id}`, {
            title: this.title,
            permalink: this.permalink,
            iconURL: this.iconURL,
            openMethod: this.openMethod
          })
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
        if (this.id === '') {
          return
        }
        const responseData = await this.axios.get(`/console/navigations/${this.id}`)
        if (responseData) {
          this.$set(this, 'title', responseData.title)
          this.$set(this, 'permalink', responseData.permalink)
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
