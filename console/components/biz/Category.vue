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
        v-model="permalink"
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
    <button class="fn-right btn btn--margin-t30 btn--info btn--space" @click="created">
      {{ $t('confirm', $store.state.locale) }}
    </button>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        errorMsg: '',
        error: false,
        title: '',
        permalink: '',
        description: '',
        tags: '',
        titleRules: [
          (v) => !!v || this.$t('required', this.$store.state.locale),
          (v) => v.length <= 32 || this.$t('validateRule', this.$store.state.locale)
        ],
        descriptionRules: [
          (v) => v.length <= 32 || this.$t('validateRule', this.$store.state.locale)
        ]
      }
    },
    methods: {
      async created () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.post('/console/categories', {
          title: this.title,
          permalink: this.permalink,
          description: this.description,
          tags: this.tags
        })
        if (responseData.code === 0) {
          this.$set(this, 'error', false)
          this.$set(this, 'errorMsg', '')
          this.$emit('addSuccess')
        } else {
          this.$set(this, 'error', true)
          this.$set(this, 'errorMsg', responseData.msg)
        }
      }
    }
  }
</script>
