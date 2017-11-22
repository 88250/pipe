<template>
  <div>
    <div class="card card--space">
      <ul class="list">
        <li class="fn-flex" v-if="$store.state.role <= 3">
          <div class="fn-flex-1">
            {{ $t('import', $store.state.locale) }}
          </div>
          <v-menu
            v-if="$store.state.role < 2"
            :nudge-bottom="28"
            :nudge-width="60"
            :nudge-left="60"
            :open-on-hover="true">
            <v-toolbar-title slot="activator">
              <v-btn class="btn--small btn--info" @click="">
                JSON
                <v-icon>arrow_drop_down</v-icon>
              </v-btn>
            </v-toolbar-title>
            <v-list>
              <v-list-tile @click="">
                JSON
              </v-list-tile>
              <v-list-tile @click="">
                {{ $t('staticBlog', $store.state.locale) }}
              </v-list-tile>
            </v-list>
          </v-menu>
        </li>
        <li class="fn-flex">
          <div class="fn-flex-1">
            {{ $t('export', $store.state.locale) }}
          </div>
          <v-menu
            v-if="$store.state.role < 3"
            :nudge-bottom="28"
            :nudge-width="60"
            :nudge-left="60"
            :open-on-hover="true">
            <v-toolbar-title slot="activator">
              <v-btn class="btn--small btn--info" @click="">
                JSON
                <v-icon>arrow_drop_down</v-icon>
              </v-btn>
            </v-toolbar-title>
            <v-list>
              <v-list-tile @click="">
                JSON
              </v-list-tile>
              <v-list-tile @click="">
                {{ $t('staticBlog', $store.state.locale) }}
              </v-list-tile>
              <v-list-tile @click="">
                SQL
              </v-list-tile>
            </v-list>
          </v-menu>
        </li>
        <!--
        <li class="fn-flex" v-if="$store.state.role <= 3">
          <div class="fn-flex-1">
            {{ $t('clearData', $store.state.locale) }}
          </div>
          <v-menu
            v-if="$store.state.role < 3"
            :nudge-bottom="28"
            :nudge-width="60"
            :nudge-left="60"
            :open-on-hover="true">
            <v-toolbar-title slot="activator">
              <v-btn class="btn--small btn--info" @click="">
                {{ $t('removeUnusedTags', $store.state.locale) }}
                <v-icon>arrow_drop_down</v-icon>
              </v-btn>
            </v-toolbar-title>
            <v-list>
              <v-list-tile @click="">
                {{ $t('removeUnusedTags', $store.state.locale) }}
              </v-list-tile>
              <v-list-tile @click="">
                {{ $t('removeData', $store.state.locale) }}
              </v-list-tile>
            </v-list>
          </v-menu>
        </li>
        -->
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    head () {
      return {
        title: `${this.$store.state.blogTitle} - ${this.$t('others', this.$store.state.locale)}`
      }
    }
  }
</script>
