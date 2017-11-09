/**
 * @fileoverview article.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.1.0.0, Oct 19, 2017
 */

import $ from 'jquery'
import hljs from 'highlight.js'
import {
  InitToc,
  ShowEditor,
  InitComment,
  ParseMarkdown
} from '../../../js/article'
import './common'

const Article = {
  /**
   * @description 页面初始化
   */
  init: () => {
    $('#articleCommentBtn').click(function () {
      const $this = $(this)
      ShowEditor($this.data('title'), $this.data('id'))
    })

    if ($('#toc').length === 1) {
      InitToc('toc', 'articleContent')
      if ($('body').width() > 768) {
        $('body').addClass('body--side')
      }
    }

    InitComment()

    $('pre > code').each(function (i, block) {
      hljs.highlightBlock(block);
    });

    ParseMarkdown()
  }
}

Article.init()