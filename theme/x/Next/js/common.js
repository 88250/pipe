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
        document.getElementById(id).innerHTML = count
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

window.increase = Common.increase
window.addLevelToTag = Common.addLevelToTag
Icon()
Common.init()
export default Common