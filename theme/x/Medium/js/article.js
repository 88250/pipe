/**
 * @fileoverview article.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.1.1.0, Jan 30, 2018
 */

import $ from 'jquery'
import QRious from 'qrious'
import { InitComment, InitToc, ShowEditor, InitHljs } from '../../../js/article'
import './common'

const Article = {
  /**
   * @description 页面初始化
   */
  init: () => {
    InitComment()
    InitHljs()

    Article._share('#articleShare')
    Article._share('#articleSideShare')
    Article._share('#articleBottomShare')

    $('#articleCommentBtn, #articleSideCommentBtn, #articleCommentBottomBtn').click(function () {
      const $this = $(this)
      ShowEditor($this.data('title'), $this.data('id'))
    })

    const $postSide = $('.post__side')
    if ($(window).height() >= $('.post').height()) {
      $postSide.css('opacity', 1)
    }

    $postSide.css('left', (($('.post').offset().left - 20) / 2 - 27) + 'px')

    const sideAbsoluteTop = ($(window).height() - 207) / 2 + 105
    let beforScrollTop = $(window).scrollTop()
    $(window).scroll(function () {
      const scrollTop = $(window).scrollTop()
      const bottomTop = $('.article__bottom').offset().top
      if (scrollTop > 65) {
        $postSide.css('opacity', 1)

        if (beforScrollTop - scrollTop > 0) {
          // up
          $('.header').addClass('header--fixed').css({'top': '0'})
          $('.main').css('padding-top', '64px')
          if ($(window).height() <= $('.post').height() && scrollTop < bottomTop - $(window).height()) {
            $('.article__toolbar').css({
              'bottom': 0,
              'opacity': 1
            })
          }
        } else if (beforScrollTop - scrollTop < 0) {
          // down
          $('.header').css({'top': '-64px'}).removeClass('header--fixed')
          $('.main').css('padding-top', '0')
          $('.article__toolbar').css({
            'bottom': '-44px',
            'opacity': 0
          })
        }

      } else {
        if ($(window).height() <= $('.post').height()) {
          $postSide.css('opacity', 0)
        }

        $('.header').removeClass('header--fixed').css('top', '-64px')
        $('.main').css('padding-top', '0')
      }

      if (scrollTop > bottomTop - $(window).height()) {
        if (bottomTop < $(window).height()) {
          $postSide.css({
            'position': 'absolute',
            'top': (bottomTop - 105) + 'px'
          })
        } else {
          $postSide.css({
            'position': 'absolute',
            'top': (bottomTop - sideAbsoluteTop) + 'px'
          })
        }
      } else {
        $postSide.css({
          'position': 'fixed',
          'top': '50%'
        })
      }

      beforScrollTop = scrollTop
    })

    $(window).scroll()
  },
  _share: (id) => {
    const $this = $(id)
    const $qrCode = $this.find('.article__code')
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

    $this.find('span').click(function () {
      const key = $(this).data('type')

      if (!key) {
        return
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
  }
}

Article.init()