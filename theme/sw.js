/**
 * @fileoverview service work.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.1.0.0, Oct 11, 2017
 * @since 2.2.0
 */

const version = '1509461684551'
const staticServePath = 'https://static.hacpai.com'

/**
 * @description add offline cache
 */
self.addEventListener('install', e => {
  e.waitUntil(
    caches.open('hacpai-static-' + version).then(cache => {
      return cache.addAll([
        '.',
        `${staticServePath}/js/symbol-defs.min.js?${version}`,
        `${staticServePath}/js/lib/compress/libs.min.js?${version}`,
        `${staticServePath}/js/common.min.js?${version}`,
        `${staticServePath}/js/lib/compress/article-libs.min.js?${version}`,
        `${staticServePath}/js/m-article.min.js?${version}`,
        `${staticServePath}/js/settings.min.js?${version}`,
        `${staticServePath}/js/article.min.js?${version}`,
        `${staticServePath}/js/channel.min.js?${version}`,
        `${staticServePath}/css/mobile-base.css?${version}`,
        `${staticServePath}/css/responsive.css?${version}`,
        `${staticServePath}/css/base.css?${version}`,
        `${staticServePath}/css/index.css?${version}`,
        `${staticServePath}/js/lib/highlight.js-9.6.0/styles/github.css`,
        `${staticServePath}/js/lib/editor/codemirror.min.css`
      ])
        .then(() => self.skipWaiting())
    })
  )
})

self.addEventListener('activate', event => {
  // delete all caches
  event.waitUntil(
    caches.keys().then(function (keyList) {
      return Promise.all(keyList.map(function (key) {
        if (key !== 'hacpai-html' && key !== 'hacpai-avatar-emoji' &&
          key !== 'hacpai-static-' + version) {
          return caches.delete(key)
        }
      }))
    })
  )
})

// 请求截取
self.addEventListener('fetch', event => {
  if (event.request.headers.get('accept').indexOf('text/html') === 0) {
    // 动态资源
    event.respondWith(
      // 动态资源需要每次进行更新
      fetch(event.request).then(function (response) {
        return caches.open('hacpai-html').then(function (cache) {
          // 更新动态资源的缓存
          cache.put(event.request, response.clone())
          return response
        })
      }).catch(function () {
        // 动态资源需离线后从缓存中获取
        return caches.match(event.request)
      })
    )
  } else {
    // 静态资源
    event.respondWith(
      caches.match(event.request).then(response => {
        // 指定的静态资源直接从缓存中获取
        return response ||
          // 没有指定的静态资源从服务器拉取
          fetch(event.request).then(function (fetchResponse) {
            if (event.request.url.indexOf('//img.hacpai.com/avatar/') > -1 ||
              event.request.url.indexOf('//static.hacpai.com/emoji/graphics/') > -1 ||
              event.request.url.indexOf('//static.hacpai.com/images/emotions/') > -1) {
              // 对用户头像、emoji、emotion 进行缓存
              return caches.open('hacpai-avatar-emoji').then(function (cache) {
                // 缓存没有指定的静态资源
                cache.put(event.request, fetchResponse.clone())
                return fetchResponse
              })
            } else {
              return fetchResponse
            }
          }).catch(function () {
            // 静态资源获取失败
            console.log(`fetch ${event.request.url} error`)
          })
      })
    )
  }
})
