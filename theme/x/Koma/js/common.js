/**
 * @fileoverview util and every page should be used.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.1.0.0, Dec 11, 2017
 */

import $ from 'jquery'
import Icon from './symbol'
import {
  KillBrowser,
  PreviewImg,
} from '../../../js/common'

const Common = {
  /**
   * @description 页面初始化
   */
  init: () => {
    PreviewImg()
    KillBrowser()

    $("a.tag").each(function (i) {
      $(this).addClass("tag--color" + Math.ceil(Math.random() * 4));
    });

    $('#goTop').click(function () {
      $("html, body").animate({
        scrollTop: 0
      }, 800)
    });

    $(window).scroll(function () {
      if ($('.article__item').length > 0) {
        $('.article__item:not(.article__item--show)').each(function () {
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