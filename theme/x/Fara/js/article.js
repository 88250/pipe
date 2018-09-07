/**
 * @fileoverview article.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.1.0.0, Oct 19, 2017
 */

import $ from 'jquery'
import QRious from 'qrious'
import {InitComment, InitToc, ShowEditor, InitHljs} from '../../../js/article'
import './common'

const Article = {
  /**
   * @description 页面初始化
   */
  init: () => {
    $('#articleCommentBtn').click(function () {
      const $this = $(this)
      ShowEditor($this.data('title'), $this.data('id'))
    })

    if ($('#toc').length === 1) {
      InitToc('toc', 'articleContent')

      // https://github.com/b3log/pipe/issues/114
      $('#toc a').each(function () {
        $(this).data('href', $(this).attr('href')).attr('href', 'javascript:void(0)')
      }).click(function () {
        const hash = $(this).data('href')
        location.hash = hash
        $(window).scrollTop($(hash)[0].offsetTop - 50)
      })

      $('.side__toc').width($('.side').width()).css('max-height', ($(window).height() - 170) + 'px')

      $(window).scroll(function () {
        if ($(window).scrollTop() > 75) {
          $('.side__toc').addClass('side__toc--fix')
        } else {
          $('.side__toc').removeClass('side__toc--fix')
        }
      })

      $('#hideToc').click(function () {
        $(this).hide()
        $('#showToc').show()
        $('.side__toc').animate({
          'margin-top': '300px'
        }, 300, function () {
          $(this).hide()
          $('.side > .fn-none').show()
        })
      })

      $('#showToc').click(function () {
        $(this).hide()
        $('#hideToc').show()
        $('.side__toc').show().animate({
          'margin-top': '0'
        }, 100)
        $('.side > .fn-none').hide()
      })

      if ($(window).width() < 768) {
        $('#hideToc').click()
      }
    }

    InitComment()
    InitHljs()

    Article._share()
  },
  _share: () => {
    const $this = $('.share__btns')
    const $qrCode = $this.find('.share__code')
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

    $this.find('.share__btn').click(function () {
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