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
  }
}

window.increase = Common.increase
Icon()
Common.init()
export default Common