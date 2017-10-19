/**
 * @fileoverview article.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.1.0.0, Oct 19, 2017
 */

import { AddComment, LocalStorageComment, ReomveComment } from '../../../js/common'
import Common from './common'

const Article = {
  /**
   * @description 页面初始化
   */
  init: () => {
    LocalStorageComment()
  },
  removeComment: (id, label, label2) => {
    if (confirm(label)) {
      ReomveComment(id, () => {
        const $commentsCnt = $('#commentsCnt')
        const $comments = $('#comments')
        $('#comment' + id).remove()
        $commentsCnt.text(parseInt($commentsCnt.text()) - 1)
        if ($comments.find('section').length === 0) {
          $comments.addClass('ft-center comment__null fn-bottom').attr('onclick', 'Common.showComment()')
            .html(`${label2} <svg><use xlink:href="#comment"></use></svg>`)
        }
      }, (msg) => {
        alert(msg)
      })
    }
  },
  showComment: (reply, id, commentId) => {
    const $editor = $('#editor')
    $editor.css('bottom', '0').data('id', id).data('commentid', commentId)
    $('body').css('padding-bottom', $editor.outerHeight() + 'px')
    $('#replyObject').text(reply)
  },
  hideComment: () => {
    $('#editor').css('bottom', '-400px')
    $('body').css('padding-bottom', 0)
  },
  addComment: (label) => {
    const $editor = $('#editor')
    AddComment({
      'articleID': $editor.data('id'),
      'content': $('#commentContent').val(),
      'parentCommentID': $editor.data('commentid')
    }, (data) => {
      Article.hideComment()
      const $commentsCnt = $('#commentsCnt')
      $commentsCnt.text(parseInt($commentsCnt.text()) + 1)

      if (parseInt($commentsCnt.text()) === 1) {
        $('#comments').removeClass('ft-center comment__null').removeAttr('onclick')
          .html(`<div class="module__header"><span id="commentsCnt">1</span>${label}</div>${data}`)
      } else {
        $('#comments').find('section').last().after(data)
      }
    }, (msg) => {
      alert(msg)
    })
  }
}

Article.init()