<template>
  <aside class="side">
    <nav>
      <ul>
        <li>
          <nuxt-link to="/"><icon icon="home"></icon> {{ $t('home') }}</nuxt-link>
        </li>
        <li>
          <nuxt-link to="/articles/post"><icon icon="add"></icon> {{ $t('postArticle') }}</nuxt-link>
        </li>
        <li>
          <nuxt-link to="/skins"><icon icon="skin"></icon> {{ $t('skinList') }}</nuxt-link>
        </li>
        <li>
          <icon icon="manage"></icon> {{ $t('manage') }}
        </li>
        <li>
          <ul>
            <li>
              <nuxt-link to="/articles/management">{{ $t('articleList') }}</nuxt-link>
            </li>
            <li>
              <nuxt-link to="/comments/management">{{ $t('commentList') }}</nuxt-link>
            </li>
            <li>
              <nuxt-link to="/categories/management">{{ $t('categoryList') }}</nuxt-link>
            </li>
            <li>
              <nuxt-link to="/navigation/management">{{ $t('navigationList') }}</nuxt-link>
            </li>
            <li>
              <nuxt-link to="/links/management">{{ $t('linkList') }}</nuxt-link>
            </li>
            <li>
              <nuxt-link to="/users/management">{{ $t('userList') }}</nuxt-link>
            </li>
            <li>
              <nuxt-link to="/plugins/management">{{ $t('pluginsList') }}</nuxt-link>
            </li>
            <li>
              <nuxt-link to="/others/management">{{ $t('others') }}</nuxt-link>
            </li>
          </ul>
        </li>
        <li><icon icon="setting"></icon> {{ $t('setting') }}</li>
        <li>
          <ul>
            <li>
              <nuxt-link to="/configurations/setting">{{ $t('configuration') }}</nuxt-link>
            </li>
            <li>
              <nuxt-link to="/signs/setting">{{ $t('signs') }}</nuxt-link>
            </li>
            <li>
              <nuxt-link to="/parameters/setting">{{ $t('parameters') }}</nuxt-link>
            </li>
            <li>
              <nuxt-link to="/qiniu/setting">{{ $t('qiniu') }}</nuxt-link>
            </li>
            <li>
              <nuxt-link to="/b3log/setting">B3log</nuxt-link>
            </li>
          </ul>
        </li>
        <li>
          <nuxt-link to="/about"><icon icon="info"></icon> {{ $t('about') }}</nuxt-link>
        </li>
      </ul>
    </nav>
  </aside>
</template>

<script>
</script>

<style lang="sass">
  @import "~assets/scss/_variables.scss"

  .side
    width: 240px
    background-color: $white
    position: fixed
    height: 100%
    top: 60px
    box-shadow: 1px 0px 20px rgba(0, 0, 0, 0.08)
    overflow: auto

  .side::-webkit-scrollbar
    display: none
</style>
