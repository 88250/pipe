<template>
  <div class="card__body fn__clear">
    <v-form ref="form">
      <v-text-field
        :label="$t('blogAdmin', $store.state.locale)"
        v-model="name"
        :counter="32"
        :rules="requiredRules"
      ></v-text-field>

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
  import {required, maxSize} from '~/plugins/validate'

  export default {
    data () {
      return {
        errorMsg: '',
        error: false,
        name: '',
        requiredRules: [
          (v) => required.call(this, v),
          (v) => maxSize.call(this, v, 32)
        ]
      }
    },
    methods: {
      async created () {
        if (!this.$refs.form.validate()) {
          return
        }
        const responseData = await this.axios.post('/console/blogs', {
          name: this.name
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
