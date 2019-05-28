/**
 * @fileoverview util and every page should be used.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 1.0.0.0, May 28, 2019
 */

import $ from 'jquery'
import Icon from './symbol'
import {
  initPjax,
  KillBrowser,
  PreviewImg,
} from '../../../js/common'
import { InitComment } from '../../../js/article'
import QRious from 'qrious'

const Common = {
  initEvent: () => {
    $("a.tag").each(function () {
      $(this).addClass("tag--color" + Math.ceil(Math.random() * 4));
    });

    $('#goTop').click(function () {
      $("html, body").animate({
        scrollTop: 0
      }, 800)
    });

    $('#hideToc').click(function () {
      if ($('.side > .fn__none').css('display') !== 'none') {
        return
      }
      $('.side > .fn__none').show()
      $('.toc__panel').addClass('toc__panel--hide')
      setTimeout(function () {
        $('.toc__panel').hide();
      }, 300)
    });

    $('#showToc').click(function () {
      if ($('.toc__panel').css('display') !== 'none') {
        return
      }
      $('.toc__panel').show().removeClass('toc__panel--hide')
      setTimeout(function () {
        $('.side > .fn__none').hide()
      }, 300)
    });
  },
  /**
   * @description 页面初始化
   */
  init: () => {
    PreviewImg()
    KillBrowser()
    initPjax(() => {
      $(window).scroll();
      Common.initEvent()
    })

    Common.initEvent()

    $(window).scroll(function () {
      if ($('.article__item').length > 0) {
        $('.article__item:not(.article__item--show)').each(function (index) {
          if (index === 0) {
            $(this).addClass('article__item--show');
            return
          }
          if ($(this).offset().top <= $(window).scrollTop() + $(window).height() - $(this).height() / 7) {
            $(this).addClass('article__item--show');
          }
        });
      }

      if ($(window).scrollTop() > $(window).height()) {
        $('#goTop').show();
      } else {
        $('#goTop').hide();
      }
    });

    $(window).scroll();
  },
  increase(max, time, id, count) {
    if (count < max) {
      setTimeout(() => {
        increase(max, time, id, ++count)
        if (document.getElementById(id)) {
          document.getElementById(id).innerHTML = count
        }
      }, time / max)
    }
  }
}


const Article = {
  /**
   * @description 页面初始化
   */
  init: () => {
    InitComment()

    Article._share();
  },
  _share: () => {
    const $this = $('.article__share')
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

    $this.find('.article__share-btn').click(function () {
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
  }
}

if (!window.increase) {
  window.increase = Common.increase
  Icon()
  Common.init()
  if ($('#pipeComments').length === 1) {
    Article.init()
  }
}