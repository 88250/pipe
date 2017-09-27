<template>
  <div class="card">
    <div class="card__body">
      <v-form>
        <v-text-field
          :label="$t('title', $store.state.locale)"
          v-model="title"
          :rules="titleRules"
          :counter="10"
          required
        ></v-text-field>

        <mavon-editor v-model="content"/>

        <v-text-field
          :label="$t('tags', $store.state.locale)"
          v-model="title"
          :rules="titleRules"
          :counter="10"
          required
        ></v-text-field>

        <mavon-editor v-model="abstract"/>

        <v-text-field
          :label="$t('links', $store.state.locale)"
          v-model="links"
        ></v-text-field>

        <v-text-field
          type="password"
          :label="$t('visitPassword', $store.state.locale)"
          v-model="password"
        ></v-text-field>

        <label class="checkbox">
          <input type="checkbox" disabled checked/><span class="checkbox__icon"></span>
          {{ $t('allowComment', $store.state.locale) }}
        </label><br>

        <label class="checkbox">
          <input type="checkbox"/><span class="checkbox__icon"></span>
          {{ $t('useSign', $store.state.locale) }}
        </label>
      </v-form>
    </div>
  </div>
</template>

<script>
  import 'mavon-editor/dist/css/index.css'

  export default {
    data () {
      return {
        content: '',
        abstract: '',
        title: '',
        titleRules: [
          (v) => !!v || this.$t('titleIsRequired', this.$store.state.locale),
          (v) => v.length <= 10 || this.$t('titleRule', this.$store.state.locale)
        ],
        links: '',
        password: '',
        isUseSign: false,
        isComment: false
      }
    },
    head () {
      return {
        title: `${this.$store.state.userName} - ${this.$t('postArticle', this.$store.state.locale)}`
      }
    },
    mounted () {
    }
  }
</script>

<style>
</style>
