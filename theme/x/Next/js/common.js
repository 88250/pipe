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

    $('#hideToc').click(function () {
      if ($('.side > .fn-none').css('display') !== 'none') {
        return
      }
      $('.side > .fn-none').show()
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
        $('.side > .fn-none').hide()
      }, 300)
    });

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
        document.getElementById(id).innerHTML = count
      }, time / max)
    }
  }
}

window.increase = Common.increase
Icon()
Common.init()
export default Common