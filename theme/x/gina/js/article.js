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
  init: function () {
    LocalStorageComment()
  }, removeComment (id, label, label2) {
    if (confirm(label)) {
      ReomveComment(id, function () {
        const $commentsCnt = $('#commentsCnt')
        const $comments = $('#comments')
        $('#comment' + id).remove()
        $commentsCnt.text(parseInt($commentsCnt.text()) - 1)
        if ($comments.find('section').length === 0) {
          $comments.addClass('ft-center comment__null fn-bottom').attr('onclick', 'Common.showComment()')
            .html(`${label2} <svg><use xlink:href="#comment"></use></svg>`)
        }
      }, function (msg) {
        alert(msg)
      })
    }
  },
  showComment: function (reply, id, commentId) {
    const $editor = $('#editor')
    $editor.css('bottom', '0').data('id', id).data('commentid', commentId)
    $('body').css('padding-bottom', $editor.outerHeight() + 'px')
    $('#replyObject').text(reply)
  },
  hideComment: function () {
    $('#editor').css('bottom', '-400px')
    $('body').css('padding-bottom', 0)
  },
  addComment: function (label) {
    const $editor = $('#editor')
    AddComment($editor.data('commentid') || $editor.data('id'), function (data) {
      Article.hideComment()
      const $commentsCnt = $('#commentsCnt')
      $commentsCnt.text(parseInt($commentsCnt.text()) + 1)

      // TODO contentHTML
      if (parseInt($commentsCnt.text()) === 1) {
        $('#comments').removeClass('ft-center comment__null').removeAttr('onclick')
          .html(`<div class="module__header"><span id="commentsCnt">1</span>${label}</div>${data}`)
      } else {
        $('#comments').find('section').last().after(data)
      }
    }, function (msg) {
      alert(msg)
    })
  }
}

Article.init()