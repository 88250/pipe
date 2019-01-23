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
        label="URI"
        v-model="url"
        :rules="URIRules"
        :counter="255"
        required
      ></v-text-field>
      <v-text-field
        :label="$t('description', $store.state.locale)"
        v-model="description"
        :rules="descriptionRules"
        :counter="255"
      ></v-text-field>
      <v-select
        v-model="tags"
        :label="$t('tags', $store.state.locale)"
        chips
        tags
        :items="$store.state.tagsItems"
        required
        :rules="tagsRules"
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
        description: '',
        tags: '',
        titleRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 128)
        ],
        descriptionRules: [
          (v) => maxSize.call(this, v, 255)
        ],
        URIRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 255)
        ],
        tagsRules: [
          (v) => this.tags.length > 0 || this.$t('required', this.$store.state.locale)
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
          path: this.url,
          description: this.description,
          tags: this.tags.join(',')
        }
        if (this.id === 0) {
          responseData = await this.axios.post('/console/categories', requestData)
        } else {
          responseData = await this.axios.put(`/console/categories/${this.id}`, requestData)
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
        const responseData = await this.axios.get(`/console/categories/${this.id}`)
        if (responseData) {
          this.$set(this, 'title', responseData.title)
          this.$set(this, 'url', responseData.path)
          this.$set(this, 'description', responseData.description)
          this.$set(this, 'tags', responseData.tags.split(','))
        }
      }
    },
    mounted () {
      this.init()
      // get tags
      this.$store.dispatch('getTags')
    }
  }
</script>
