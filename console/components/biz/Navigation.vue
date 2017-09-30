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
      <v-text-field
        :label="$t('openMethod', $store.state.locale)"
        v-model="openMethod"
        :rules="titleRules"
        :counter="32"
        required
        @keyup.enter="created"
      ></v-text-field>
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
        titleRules: [
          (v) => !!v || this.$t('required', this.$store.state.locale),
          (v) => v.length <= 32 || this.$t('validateRule', this.$store.state.locale)
        ],
        iconURLRules: [
          (v) => v.length <= 32 || this.$t('validateRule', this.$store.state.locale)
        ]
      }
    },
    methods: {
      async created () {
        if (!this.$refs.form.validate()) {
          return
        }
        let responseData = {}
        if (this.id === '') {
          responseData = await this.axios.post('/console/navigation', {
            title: this.title,
            permalink: this.permalink,
            iconURL: this.iconURL,
            openMethod: this.openMethod
          })
        } else {
          responseData = await this.axios.put(`/console/navigation/${this.id}`, {
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
      }
    },
    async mounted () {
      if (this.id === '') {
        return
      }
      const responseData = await this.axios.get(`/console/navigation/${this.id}`)
      if (responseData) {
        this.$set(this, 'title', responseData.title)
        this.$set(this, 'permalink', responseData.permalink)
        this.$set(this, 'iconURL', responseData.iconURL)
        this.$set(this, 'openMethod', responseData.openMethod)
      }
    }
  }
</script>
