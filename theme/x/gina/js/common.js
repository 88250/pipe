/**
 * @fileoverview util and every page should be used.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.2.0.0, Oct 19, 2017
 */

import $ from 'jquery'
import hljs from 'highlight.js'
import QRious from 'qrious'
import Icon from './symbol'
import { KillBrowser, LazyLoadCSSImage, LazyLoadImage } from '../../../js/common'

const Common = {
  /**
   * @description 页面初始化
   */
  init: () => {
    KillBrowser()
    LazyLoadCSSImage('.avatar, .article__thumb')
    LazyLoadImage()
    Common._header()
    Common._share()
    $('#sidebarIcon').click(() => {
      Common.toggleSide()
    })

    $('pre > code').each(function(i, block) {
      hljs.highlightBlock(block);
    });
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

      $this.find('.action__btns span').click(function () {
        const key = $(this).data('type')

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
    if ($body.hasClass('body--side')) {
      $('#editor').width($(window).width() - 340)
    } else {
      $('#editor').width('100%')
    }
  }
}

Icon()
Common.init()

export default Common