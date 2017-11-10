/**
 * @fileoverview util and every page should be used.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.2.0.0, Oct 19, 2017
 */

import $ from 'jquery'
import QRious from 'qrious'
import Icon from './symbol'
import {
  KillBrowser,
  LazyLoadCSSImage,
  LazyLoadImage,
  Logout,
  TrimB3Id,
  PreviewImg,
  ParseMarkdown
} from '../../../js/common'

const Common = {
  /**
   * @description 页面初始化
   */
  init: () => {
    TrimB3Id()
    KillBrowser()
    LazyLoadCSSImage()
    LazyLoadImage()
    PreviewImg()
    ParseMarkdown()
    Common._header()
    Common._share()
    $('#sidebarIcon').click(() => {
      Common.toggleSide()
    })

    $('#logout').click(function () {
      Logout()
    })
  },
  _header: () => {
    const $headerSearch = $('#headerSearch')
    const $input = $headerSearch.find('input')
    $headerSearch.click(() => {
      $input.width(95).focus()
    })
    $input.blur(function () {
      $(this).width(0)
    })
  },
  _share: () => {
    $('.action__share').each(function () {
      const $this = $(this)
      const $qrCode = $this.find('.action__code')
      const shareURL = $qrCode.data('url')
      const avatarURL = $qrCode.data('avatar')
      const title = encodeURIComponent($qrCode.data('title') + ' - ' + $qrCode.data('blogtitle')),
        url = encodeURIComponent(shareURL)

      const urls = {}
      urls.tencent = 'http://share.v.t.qq.com/index.php?c=share&a=index&title=' + title +
        '&url=' + url + '&pic=' + avatarURL
      urls.weibo = 'http://v.t.sina.com.cn/share/share.php?title=' +
        title + '&url=' + url + '&pic=' + avatarURL
      urls.google = 'https://plus.google.com/share?url=' + url
      urls.twitter = 'https://twitter.com/intent/tweet?status=' + title + ' ' + url

      $this.find('.action__btn').click(function () {
        const key = $(this).data('type')

        if (!key) {
          return;
        }

        if (key === 'wechat') {
          if ($qrCode.css('background-image') === 'none') {
            const qr = new QRious({
              element: $qrCode[0],
              value: shareURL,
              size: 128
            })
            $qrCode.css('background-image', `url(${qr.toDataURL('image/jpeg')})`).hide()
          }
          $qrCode.slideToggle()
          return false
        }

        window.open(urls[key], '_blank', 'top=100,left=200,width=648,height=618')
      })
    })
  },
  toggleSide: () => {
    const $body = $('body')
    $body.toggleClass('body--side')
  },
  increase(max, time, id, count) {
    if (count < max) {
      setTimeout(() => {
        increase(max, time, id, ++count)
        document.getElementById(id).innerHTML = count
      }, time / max)
    }
  },
  addLevelToTag() {
    const $tags = $('#tags');
    const tagsArray = $tags.find('.tag')
    // 根据引用次数添加样式，产生云效果
    const max = parseInt(tagsArray.last().data('count'));
    const distance = Math.ceil(max / 5);
    for (let i = 0; i < tagsArray.length; i++) {
      const count = parseInt(tagsArray.data('count'));
      // 算出当前 tag 数目所在的区间，加上 class
      for (let j = 0; j < 5; j++) {
        if (count > j * distance && count <= (j + 1) * distance) {
          tagsArray[i].className = `tag tags__level${j}`;
          break;
        }
      }
    }

    // 按字母或者中文拼音进行排序
    $tags.html(tagsArray.get().sort(function (a, b) {
      var valA = $(a).text().toLowerCase();
      var valB = $(b).text().toLowerCase();
      // 对中英文排序的处理
      return valA.localeCompare(valB);
    }));
  }
}

window.increase = Common.increase
window.addLevelToTag = Common.addLevelToTag
Icon()
Common.init()
export default Common