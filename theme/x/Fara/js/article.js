/**
 * @fileoverview article.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.1.0.0, Sep 10, 2018
 */

import $ from 'jquery'
import {InitComment, InitToc, ShowEditor, InitHljs} from '../../../js/article'
import './common'

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
    InitHljs()
  },
}

Article.init()