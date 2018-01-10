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
    $('.sidebar').click(function () {
      if ($(this).hasClass('sidebar--active')) {
        $(this).removeClass('sidebar--active')
        $('.main').css('margin-right', '0')
        $('.side').css('right', '-320px')
      } else {
        $(this).addClass('sidebar--active')
        $('.main').css('margin-right', '320px')
        $('.side').css('right', '0')
      }
    })

    $('.top__btn').click(function () {
      $("html, body").animate({
        scrollTop: 0
      }, 800)
    })
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