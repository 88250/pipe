/**
 * @fileoverview article.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.1.0.0, Oct 19, 2017
 */

import $ from 'jquery'
import { AddComment, LocalStorageInput, ReomveComment } from '../../../js/common'
import './common'

const Article = {
  /**
   * @description 页面初始化
   */
  init: () => {
    LocalStorageInput('commentContent')
    $('#articleCommentBtn').click(function () {
      const $this = $(this)
      Article.showComment($this.data('title'), $this.data('id'))
    })

    $('.comment__item').each(function () {
      const $this = $(this)
      $this.find('.commentDelete').click(function () {
        // remove comment
        const $firstBtn = $(this)
        Article.removeComment($firstBtn.data('id'), $firstBtn.data('label'), $firstBtn.data('label2'))
      })
      $this.find('.commentAdd').click(function () {
        // add reply comment
        const $lastBtn = $(this)
        Article.showComment($lastBtn.data('title'), $lastBtn.data('id'), $lastBtn.data('commentid'))
      })
    })

    $('#comments.comment__null').click(function () {
      const $commentNull = $(this)
      Article.showComment($commentNull.data('title'), $commentNull.data('id'))
    })

    $('#editorAdd').click(function () {
      Article.addComment($(this).data('label'))
    })

    $('#editorCancel').click(function () {
      Article.hideComment()
    })
  },
  removeComment: (id, label, label2) => {
    if (confirm(label)) {
      ReomveComment(id, () => {
        const $commentsCnt = $('#commentsCnt')
        const $comments = $('#comments')
        const $item = $('#comment' + id)

        if ($comments.find('section').length === 1) {
          $comments.addClass('ft-center comment__null fn-bottom')
            .html(`${label2} <svg><use xlink:href="#comment"></use></svg>`).click(function () {
              const $itemReplyBtn = $item.find('.comment__btn:last')
            Article.showComment($itemReplyBtn.data('title'), $itemReplyBtn.data('id'))
          })
        } else {
          $item.remove()
          $commentsCnt.text(parseInt($commentsCnt.text()) - 1)
        }
      }, (msg) => {
        alert(msg)
      })
    }
  },
  showComment: (reply, id, commentId) => {
    const $editor = $('#editor')
    if ($editor.length === 0) {
      location.href = `${location.origin}/login`
      return false
    }
    $editor.css('bottom', '0').data('id', id).data('commentid', commentId)
    $('body').css('padding-bottom', $editor.outerHeight() + 'px')
    $('#replyObject').text(reply)
  },
  hideComment: () => {
    const $editor = $('#editor')
    $editor.css('bottom', `-${$editor.outerHeight()}px`)
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

      if ($commentsCnt.length === 0) {
        $('#comments').removeClass('ft-center comment__null').removeAttr('onclick')
          .html(`<div class="module__header"><span id="commentsCnt">1</span>${label}</div>${data}`)
      } else {
        $commentsCnt.text(parseInt($commentsCnt.text()) + 1)
        $('#comments').find('section').last().after(data)
      }
    }, (msg) => {
      alert(msg)
    })
  }
}

Article.init()