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
  initPjax,
} from '../../../js/common'
import { InitComment, InitToc } from '../../../js/article'
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
      if ($('#pipeComments').length === 1 && $('#toc').length === 1) {
        $('.body__content').addClass('body__content--toc')
      } else {
        $('.body__content').removeClass('body__content--toc')
      }
    })

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

    $('#goTop').click(function () {
      $('html, body').animate({
        scrollTop: 0,
      }, 800)
    })
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
    const tagsArray = $tags.find('.tag')
    // 根据引用次数添加样式，产生云效果
    const max = parseInt(tagsArray.first().data('count'))
    const distance = Math.ceil(max / 5)
    for (let i = 0; i < tagsArray.length; i++) {
      const count = parseInt($(tagsArray[i]).data('count'))
      // 算出当前 tag 数目所在的区间，加上 class
      for (let j = 0; j < 5; j++) {
        if (count > j * distance && count <= (j + 1) * distance) {
          tagsArray[i].className = `tag tag__level${j}`
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
    if ($('#toc').length === 1) {
      InitToc('toc', 'articleContent')

      $('#toc a').each(function () {
        $(this).data('href', $(this).attr('href')).attr('href', 'javascript:void(0)')
      }).click(function () {
        const hash = $(this).data('href')
        location.hash = hash
        $(window).scrollTop($(hash)[0].offsetTop)
      })
    }

    InitComment()
  },
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