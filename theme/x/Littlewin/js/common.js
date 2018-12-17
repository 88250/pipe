/**
 * @fileoverview util and every page should be used.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.3.0.0, Oct 31, 2018
 */

import $ from 'jquery'
import Icon from './symbol'
import {
  initPjax,
  KillBrowser,
  PreviewImg,
} from '../../../js/common'
import config from '../../../../pipe.json'

const Common = {
  /**
   * @description 页面初始化
   */
  init: () => {
    initPjax(() => {
      if ($('#pipeComments').length === 1) {
        $.ajax({
          method: 'GET',
          url: `${config.StaticServer}/theme/x/Gina/js/article.min.js`,
          dataType: 'script',
          cache: true,
        })
      }
      if ($('#pipeComments').length === 1 && $('#toc').length === 1) {
        $('body').addClass('body--side')
      } else {
        $('body').removeClass('body--side')
      }
      setTimeout(() => {
        $('.header__logo').width($('.header a').get(1).offsetLeft - 30)
      }, 301)
    })

    $('.header__logo').width($('.header a').get(1).offsetLeft - 30)

    PreviewImg()
    KillBrowser()
    $('#sidebarIcon').click(() => {
      Common.toggleSide()
    })
  },
  toggleSide: () => {
    const $body = $('body')
    $body.toggleClass('body--side')
    setTimeout(() => {
      $('.header__logo').width($('.header a').get(1).offsetLeft - 30)
    }, 301)
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

if (!window.increase) {
  window.increase = Common.increase
  window.addLevelToTag = Common.addLevelToTag
  Icon()
  Common.init()
}
export default Common