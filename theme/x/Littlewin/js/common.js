/**
 * @fileoverview util and every page should be used.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 1.0.0.0, May 28, 2019
 */

import $ from 'jquery'
import APlayer from 'aplayer'
import Icon from './symbol'
import {
  initPjax,
  KillBrowser,
  PreviewImg,
} from '../../../js/common'
import { InitComment, ShowEditor } from '../../../js/article'
import QRious from 'qrious'

const Common = {
  /**
   * @description 页面初始化
   */
  init: () => {
    $('.header__menu').click(function () {
      $('.header__links').slideToggle()
    })

    new APlayer({
      container: document.getElementById('player'),
      autoplay: true,
      theme: '#ccc',
      preload: 'auto',
      lrcType: 3,
      listFolded: true,
      audio: [
        {
          name: 'Kiss The Rain',
          artist: '이루마',
          url: `${$('#pipeLang').data('staticserver')}/theme/x/Littlewin/images/kisstherain.mp3`,
          cover: `${$('#pipeLang').data('staticserver')}/theme/x/Littlewin/images/kisstherain.jpeg`,
          lrc: `${$('#pipeLang').data('staticserver')}/theme/x/Littlewin/images/kisstherain.lrc`,
          theme: '#60b044',
        },
        {
          name: 'L\'ESPOIR',
          artist: 'Richard',
          url: `${$('#pipeLang').data('staticserver')}/theme/x/Littlewin/images/LESPOIR.mp3`,
          cover: `${$('#pipeLang').data('staticserver')}/theme/x/Littlewin/images/LESPOIR.jpeg`,
          lrc: `${$('#pipeLang').data('staticserver')}/theme/x/Littlewin/images/LESPOIR.lrc`,
          theme: '#60b044',
        },
        {
          name: 'ACOMMAMOUR',
          artist: 'Richard',
          url: `${$('#pipeLang').data('staticserver')}/theme/x/Littlewin/images/ACOMMAMOUR.mp3`,
          cover: `${$('#pipeLang').data('staticserver')}/theme/x/Littlewin/images/ACOMMAMOUR.jpeg`,
          lrc: `${$('#pipeLang').data('staticserver')}/theme/x/Littlewin/images/ACOMMAMOUR.lrc`,
          theme: '#60b044',
        },
      ],
    })

    initPjax(() => {
      if ($('#pipeComments').length === 1) {
        Article.init()
      }
    })

    PreviewImg()
    KillBrowser()
  },
  increase (max, time, id, count) {
    if (count < max) {
      setTimeout(() => {
        increase(max, time, id, ++count)
        if (document.getElementById(id)) {
          document.getElementById(id).innerHTML = count
        }
      }, time / max)
    }
  },
  addLevelToTag () {
    const $tags = $('#tags')
    const tagsArray = $tags.find('a')
    // 根据引用次数添加样式，产生云效果
    const max = parseInt(tagsArray.first().data('count'))
    const distance = Math.ceil(max / 5)
    for (let i = 0; i < tagsArray.length; i++) {
      const count = parseInt($(tagsArray[i]).data('count'))
      // 算出当前 tag 数目所在的区间，加上 class
      for (let j = 0; j < 5; j++) {
        if (count > j * distance && count <= (j + 1) * distance) {
          tagsArray[i].className = `tag tag--level${j}`
          break
        }
      }
    }

    // 按字母或者中文拼音进行排序
    $tags.html(tagsArray.get().sort(function (a, b) {
      var valA = $(a).text().toLowerCase()
      var valB = $(b).text().toLowerCase()
      // 对中英文排序的处理
      return valA.localeCompare(valB)
    }))
  },
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
  },
  _share: () => {
    const $this = $('.post__share')
    const $qrCode = $this.find('.post__code')
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
    urls.qqz =`https://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url=${url}&sharesource=qzone&title=${title}&pics=${avatarURL}`,

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
