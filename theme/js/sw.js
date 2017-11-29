/*
 * Symphony - A modern community (forum/SNS/blog) platform written in Java.
 * Copyright (C) 2012-2017,  b3log.org & hacpai.com
 *
 * 本文件属于 Sym 商业版的一部分，请仔细阅读项目根文件夹的 LICENSE 并严格遵守相关约定
 */
/**
 * @fileoverview service work.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.2.1.0, Nov 14, 2017
 * @since 2.2.0
 */

const version = '1511945482232';
const staticServePath = 'https://static.hacpai.com/';
const imgServePath = 'https://img.hacpai.com/';
const servePath = 'https://hacpai.com/';
/**
 * @description add offline cache
 */
self.addEventListener('activate', event => {
  // delete all caches
  event.waitUntil(
    caches.keys().then(function (keyList) {
      return Promise.all(keyList.map(function (key) {
        if (key !== 'hacpai-html' && key !== 'hacpai-image' &&
          key !== 'hacpai-static-' + version) {
          return caches.delete(key);
        }
      }));
    })
  );
});

// 请求截取
self.addEventListener('fetch', event => {
  if (event.request.headers.get('accept').indexOf('text/html') === 0 || (
      event.request.headers.get('accept') === '*/*' &&
      event.request.url.indexOf('/js/') === -1 &&
      event.request.url.indexOf('/notification/unread/count') === -1
    )) {
    // 动态资源
    event.respondWith(
      // 动态资源需要每次进行更新
      fetch(event.request).then(function (response) {
        // 站点以外的需求不缓存
        if (event.request.url.indexOf(servePath) === -1) {
          return response;
        }
        return caches.open('hacpai-html').then(function (cache) {
          // 更新动态资源的缓存
          if (event.request.method !== 'POST' && event.request.method !== 'DELETE' &&
            event.request.method !== 'PUT') {
            // cache is unsupported POST
            cache.put(event.request, response.clone());
          }
          return response;
        });
      }).catch(function () {
        // 动态资源需离线后从缓存中获取
        return caches.match(event.request);
      })
    );
  } else {
    // 静态资源
    event.respondWith(
      caches.match(event.request).then(response => {
        // 指定的静态资源直接从缓存中获取
        return response ||
          // 没有指定的静态资源从服务器拉取
          fetch(event.request).then(function (fetchResponse) {
            if (event.request.url.indexOf(imgServePath) > -1 ||
              event.request.url.indexOf(servePath + 'porter') > -1 ||
              event.request.url.indexOf(staticServePath + 'emoji/') > -1 ||
              event.request.url.indexOf(staticServePath + 'images/emotions/') > -1) {
              // 对用户头像、图片、solo代理图片、emoji、solo emotion 进行缓存
              return caches.open('hacpai-image').then(function (cache) {
                cache.put(event.request, fetchResponse.clone());
                return fetchResponse;
              });
            } else if (event.request.url.indexOf(staticServePath + 'css/') > -1 ||
              event.request.url.indexOf(staticServePath + 'js/') > -1 ||
              event.request.url.indexOf(staticServePath + 'images/') > -1) {
              // 对 css, js, image(不含 emoji) 进行缓存
              return caches.open('hacpai-static-' + version).then(function (cache) {
                cache.put(event.request, fetchResponse.clone());
                return fetchResponse;
              });
            } else {
              return fetchResponse;
            }
          }).catch(function () {
            // 静态资源获取失败
            console.log(`fetch ${event.request.url} error`)
          });
      })
    )
  }
});