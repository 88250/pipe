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
import { InitComment, InitToc, ShowEditor } from '../../../js/article'
import QRious from 'qrious'

const Common = {
  /**
   * @description 页面初始化
   */
  init: () => {
    PreviewImg()
    KillBrowser()
    initPjax(() => {
      if ($('#pipeComments').length === 1) {
        Article.init()
      }

      if ($('#pipeComments').length !== 1 || $('#toc').length !== 1) {
        $('#sideBar').removeClass('sidebar--active')
        $('.main').css('margin-right', '0')
        $('.side').removeClass('side--active')
      }
    })
    $('#sideBar').click(function () {
      if ($(this).hasClass('sidebar--active')) {
        $(this).removeClass('sidebar--active')
        $('.main').css('margin-right', '0')
        $('.side').removeClass('side--active')
      } else {
        $(this).addClass('sidebar--active')
        $('.main').css('margin-right', '320px')
        $('.side').addClass('side--active')
      }
    })

    $('.sidebar--header').click(function () {
      if ($(this).hasClass('sidebar--active')) {
        $(this).removeClass('sidebar--active')
        $('.header nav').hide()

      } else {
        $(this).addClass('sidebar--active')
        $('.header nav').show()
      }
    })

    $('.top__btn').click(function () {
      $("html, body").animate({
        scrollTop: 0
      }, 800)
    })

    $(window).scroll(function () {
      if ($(window).scrollTop() > $(window).height()) {
        $('.top__btn').css({
          'bottom': '34px',
          'opacity': '1'
        });
      } else {
        $('.top__btn').css({
          'bottom': '-25px',
          'opacity': '0'
        });
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

    InitComment()

    Article._share();

    if ($('#sideToc').length === 1) {
      if ($(window).width() > 768) {
        $('#sideBar').click();
      }
      InitToc('toc', 'articleContent')

      $('.side__tab').click(function () {
        $('.side__tab').removeClass('side__tab--current')
        $(this).addClass('side__tab--current')

        if ($(this).data('type') === 'toc') {
          $('#sideGist').removeClass('current');
          setTimeout(function () {
            $('#sideGist').hide();
            $('#sideToc').show(function () {
              $(this).addClass('current');
            });
          }, 0);
        } else {
          $('#sideToc').removeClass('current');
          setTimeout(function () {
            $('#sideToc').hide();
            $('#sideGist').show(function () {
              $(this).addClass('current');
            });
          }, 0)
        }
      });
    }
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

    $this.find('span').click(function () {
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
  window.addLevelToTag = Common.addLevelToTag
  Icon()
  Common.init()
  if ($('#pipeComments').length === 1) {
    Article.init()
  }
}
