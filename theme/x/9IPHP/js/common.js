/**
 * @fileoverview util and every page should be used.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 1.0.0.0, May 28, 2019
 */

import $ from 'jquery'
import Icon from './symbol'
import {
  KillBrowser,
  PreviewImg,
  initPjax
} from '../../../js/common'
import { InitComment, InitToc, ShowEditor } from '../../../js/article'
import QRious from 'qrious'

const Common = {
  /**
   * @description 页面初始化
   */
  init: () => {
    initPjax(() => {
      if ($('#pipeComments').length === 1) {
        Article.init()
      }
      $('.nav a, .mobile__nav a').removeClass('nav--current')
      $('.nav a, .mobile__nav a').each(function (i) {
        const $it = $(this)
        if (i === 0 || i === $('.mobile__nav a').length) {
          if (location.origin + location.pathname === $it.attr('href')) {
            $it.addClass('nav--current')
          }
        } else if (location.href.indexOf($it.attr('href')) > -1) {
          $it.addClass('nav--current')
        }
      })
    })
    PreviewImg()
    KillBrowser()

    $('.nav a, .mobile__nav a').each(function (i) {
      const $it = $(this)
      if (i === 0 || i === $('.mobile__nav a').length) {
        if (location.origin + location.pathname === $it.attr('href')) {
          $it.addClass('nav--current')
        }
      } else if (location.href.indexOf($it.attr('href')) > -1) {
        $it.addClass('nav--current')
      }
    })

    $(window).scroll(function () {
      if ($(window).scrollTop() > $(window).height()) {
        $('#goTop').show()
      } else {
        $('#goTop').hide()
      }

      if ($(window).width() < 768) {
        return
      }
      if ($(window).scrollTop() > 75) {
        $('.nav').addClass('nav--fix').next().css('margin-top', '75px')
      } else {
        $('.nav').removeClass('nav--fix').next().css('margin-top', 0)
      }
    })

    $('#goTop').click(function () {
      $("html, body").animate({
        scrollTop: 0
      }, 800)
    })

    $('.header .fn__none').click(function () {
      $('.mobile__nav').slideToggle()
    })

    $(window).scroll()

    // https://github.com/b3log/pipe/issues/115
    const ua = window.navigator.userAgent;
    if (/MicroMessenger/i.test(ua)) {
      $('body').css('display', 'block')
    }
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
  },
  addLevelToTag() {
    const $tags = $('#tags');
    const tagsArray = $tags.find('.tag')
    // 根据引用次数添加样式，产生云效果
    const max = parseInt(tagsArray.first().data('count'));
    const distance = Math.ceil(max / 5);
    for (let i = 0; i < tagsArray.length; i++) {
      const count = parseInt($(tagsArray[i]).data('count'));
      // 算出当前 tag 数目所在的区间，加上 class
      for (let j = 0; j < 5; j++) {
        if (count > j * distance && count <= (j + 1) * distance) {
          tagsArray[i].className = `tag tag__level${j}`;
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
          $('.side > .fn__none').show()
        })
      })

      $('#showToc').click(function () {
        $(this).hide()
        $('#hideToc').show()
        $('.side__toc').show().animate({
          'margin-top': '0'
        }, 100)
        $('.side > .fn__none').hide()
      })

      if ($(window).width() < 768) {
        $('#hideToc').click()
      }
    }

    InitComment()

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

if (!window.increase) {
  window.increase = Common.increase
  window.addLevelToTag = Common.addLevelToTag
  Icon()
  Common.init()
  if ($('#pipeComments').length === 1) {
    Article.init()
  }
}
