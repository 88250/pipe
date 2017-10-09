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
        label="URI"
        v-model="url"
        :rules="titleRules"
        :counter="32"
        required
      ></v-text-field>
      <v-text-field
        :label="$t('description', $store.state.locale)"
        v-model="description"
        :rules="descriptionRules"
        :counter="32"
      ></v-text-field>
      <v-text-field
        :label="$t('tags', $store.state.locale)"
        v-model="tags"
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
    <v-btn class="fn-right btn btn--margin-t30 btn--info btn--space" @click="created">
      {{ $t('confirm', $store.state.locale) }}
    </v-btn>
    <v-btn class="fn-right btn btn--margin-t30 btn--danger btn--space" @click="$emit('update:show', false)">
      {{ $t('cancel', $store.state.locale) }}
    </v-btn>
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
        url: '',
        description: '',
        tags: '',
        titleRules: [
          (v) => !!v || this.$t('required', this.$store.state.locale),
          (v) => v.length <= 32 || this.$t('validateRule2', this.$store.state.locale)
        ],
        descriptionRules: [
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
          responseData = await this.axios.post('/console/categories', {
            title: this.title,
            url: this.url,
            description: this.description,
            tags: this.tags
          })
        } else {
          responseData = await this.axios.put(`/console/categories/${this.id}`, {
            title: this.title,
            url: this.url,
            description: this.description,
            tags: this.tags
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
        const responseData = await this.axios.get(`/console/categories/${this.id}`)
        if (responseData) {
          this.$set(this, 'title', responseData.title)
          this.$set(this, 'url', responseData.url)
          this.$set(this, 'description', responseData.description)
          this.$set(this, 'tags', responseData.tags)
        }
      }
    },
    mounted () {
      this.init()
    }
  }
</script>
