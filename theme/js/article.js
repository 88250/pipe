/**
 * @fileoverview common tool for every theme
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.2.0.0, Oct 19, 2017
 */

import $ from 'jquery'
import hljs from 'highlight.js'
import {Editor} from "./lib/b3log/editor/index";
import {LazyLoadCSSImage, LazyLoadImage, ParseMarkdown} from './common'
import config from '../../pipe.json'

/**
 * @description 初始化目录
 * @param {string} tocId 目录 id
 * @param {string} articleId 目录对应正文 id
 */
export const InitToc = (tocId, articleId, articleOffset = 0, activeClass = 'toc__item--active') => {
  const $toc = $(`#${tocId}`)
  const $articleTocs = $(`#${articleId} [id^=toc]`)
  let tocPositionArray = []

  // 目录点击
  $toc.find('li').click(function () {
    const $it = $(this)
    setTimeout(function () {
      $toc.find('li').removeClass(activeClass)
      $it.addClass(activeClass)
    }, 50)
  })

  $(window).scroll(function (event) {
    // 界面各种图片加载会导致帖子目录定位
    tocPositionArray = []
    $articleTocs.each(function () {
      tocPositionArray.push({
        id: this.id,
        offsetTop: this.offsetTop
      })
    })

    // 当前目录样式
    let scrollTop = $(window).scrollTop()
    for (let i = 0, iMax = tocPositionArray.length; i < iMax; i++) {
      if (scrollTop < tocPositionArray[i].offsetTop - articleOffset) {
        $toc.find('li').removeClass(activeClass)
        const index = i > 0 ? i - 1 : 0
        $toc.find(`a[href="#${$articleTocs[index].id}"]`).parent().addClass(activeClass)
        break
      }
    }
    if (scrollTop >= $articleTocs[$articleTocs.length - 1].offsetTop - articleOffset) {
      $toc.find('li').removeClass(activeClass)
      $toc.find('li:last').addClass(activeClass)
    }
  })

  $(window).scroll()
}

export const InitHljs = () => {
  $('pre > code').each(function (i, block) {
    hljs.highlightBlock(block);
  });
}

/**
 * @description 展开编辑器
 * @param {String} reply 回复对象的名称
 * @param {String} id 文章 id
 * @param {String} commentId 回复 id
 * @private
 */
export const ShowEditor = (reply, id, commentId) => {
  const $editor = $('#pipeEditor')
  if ($editor.length === 0) {
    location.href = 'https://hacpai.com/login'
    return
  }
  if (commentId) {
    $editor.data('commentid', commentId)
  } else {
    $editor.removeData('commentid')
  }

  $editor.css({'bottom': '0', 'opacity': 1}).data('id', id)
  $('body').css('padding-bottom', $editor.outerHeight() + 'px')
  $('#pipeEditorReplyTarget').text(reply)
  $('#pipeEditorComment').focus()
};

export const InitComment = () => {
  /**
   * @description 隐藏编辑器
   * @private
   */
  const _hideEditor = () => {
    const $editor = $('#pipeEditor')
    $editor.css({'bottom': `-${$editor.outerHeight()}px`, 'opacity': 0})
    $('body').css('padding-bottom', 0)
  }

  // init Editor
  const label = $('#pipeEditorComment').data('label').split(',')
  const $b3logEditor = Editor({
    id: 'pipeEditorComment',
    placeholder: $('#pipeEditorComment').data('placeholder'),
    label: {
      emoji: label[0]+ ' <ctrl+&frasl;>',
      bold: label[1]+ ' <ctrl+b>',
      italic: label[2]+ ' <ctrl+i>',
      quote: label[3]+ ' <ctrl+e>',
      link: label[4]+ ' <ctrl+k>',
      upload: label[5],
      unorderedList: label[6]+ ' <ctrl+l>',
      orderedList: label[7]+ ' <ctrl+shift+l>',
      view: label[8]+ ' <ctrl+d>',
      question: label[9],
      fullscreen: label[10]+ ' <ctrl+shift+a>',
      emojiTip: 'EMOJI CHEAT SHEET'
    },
    height: 200,
    keyup: (event) => {
      localStorage.setItem('pipeEditorComment', event.target.value)
    },
    esc: _hideEditor,
    ctrlEnter: () => {
      $('#pipeEditorAdd').click();
    },
    hasView: false,
    uploadURL: `${$('#pipeEditorComment').data('blogurl')}/upload`,
    fetchUpload: (url, succCB) => {
      $.ajax({
        url: `${$('#pipeEditorComment').data('blogurl')}/fetch-upload`,
        type: "POST",
        data: JSON.stringify({
          url
        }),
        success: function (result) {
          succCB(result.data.originalURL, result.data.url)
        }
      });
    },
    previewClass: 'pipe-content__reset',
    staticServePath: config.StaticServer || config.Server,
    change: (value, $preview) => {
      if ($.trim(value) === '' || !$preview) {
        return;
      }
      $.ajax({
        url: `${config.Server}/api/console/markdown`,
        type: "POST",
        data: JSON.stringify({
          mdText: value
        }),
        success: function (result) {
          $preview.html(result.data.html);
          LazyLoadImage()
          LazyLoadCSSImage()
          InitHljs()
          ParseMarkdown()
        }
      });
    }
  })

  $b3logEditor.val(localStorage.getItem('pipeEditorComment'))

  // comment null reply
  $('.pipe-comment__null').click(function () {
    const $commentNull = $(this)
    ShowEditor($commentNull.data('title'), $commentNull.data('id'))
  })

  // bottom reply
  $('#pipeCommentBottomComment').click(function () {
    const $bottomComment = $(this)
    ShowEditor($bottomComment.text(), $bottomComment.data('id'))
  })

  // comment show reply
  $('body').on('click', '#pipeComments .fn-pointer', function () {
    const $it = $(this)
    if ($it.hasClass('disabled')) {
      return
    }
    $it.addClass('disabled')
    const $svg = $it.find('svg')
    if ($svg.hasClass('pipe-comment__chevron-down')) {
      $svg.removeClass('pipe-comment__chevron-down')
      if ($it.next().find('.pipe-comment__item').length > 0) {
        $it.next().slideToggle({
          queue: false,
          complete: () => {
            $it.removeClass('disabled')
          }
        })
        return
      }
      $.ajax({
        url: `${$('#pipeEditorComment').data('blogurl')}/comments/${$it.data('id')}/replies`,
        type: 'GET',
        success: (result) => {
          if (result.code === 0) {
            let commentHTML = ''
            result.data.forEach((item) => {
              commentHTML += `<section class="pipe-comment__item">
                    <a rel="nofollow"
                       class="pipe-comment__avatar"
                       style="background-image: url(${item.Author.AvatarURL})"
                       href="${item.Author.URL}">
                    </a>
                    <div class="pipe-comment__body">
                        <a href="${item.Author.URL}" class="ft-gray">${item.Author.Name}</a>
                        <span class="ft-nowrap ft-12 ft-gray"> • ${item.CreatedAt}</span>
                        <div class="pipe-content__reset">
                             ${item.Content}
                        </div>
                    </div>
                </section>`
            })
            $it.next().html(commentHTML).slideToggle({
              queue: false,
              complete: () => {
                $it.removeClass('disabled')
              }
            })
            LazyLoadImage()
            LazyLoadCSSImage()
            InitHljs()
            ParseMarkdown()
          } else {
            alert(result.msg)
            $it.removeClass('disabled')
          }
        }
      })
    } else {
      $svg.addClass('pipe-comment__chevron-down')
      $it.next().slideToggle({
        queue: false,
        complete: () => {
          $it.removeClass('disabled')
        }
      })
    }
  });

  // comment remove
  $('body').on('click', '#pipeComments .pipe-comment__btn--danger', function () {
    const $it = $(this)
    if (confirm($it.data('label'))) {
      $.ajax({
        url: `/api/console/comments/${$it.data('id')}`,
        type: 'DELETE',
        success: (result) => {
          if (result.code === 0) {
            const $commentsCnt = $('#pipeCommentsCnt')
            const $comments = $('#pipeComments')
            const $item = $(`#pipeComment${$it.data('id')}`)

            if ($('#pipeComments > div > section').length === 1) {
              $comments.addClass('pipe-comment__null')
                .html(`${$it.data('label2')} <svg><use xlink:href="#comment"></use></svg>`).click(function () {
                ShowEditor($comments.data('title'), $comments.data('id'))
              })
            } else {
              $item.remove()
              $commentsCnt.text(parseInt($commentsCnt.text()) - 1)
            }
          } else {
            alert(result.msg)
          }
        }
      })
    }
  })

  // comment reply
  $('body').on('click', '#pipeComments .pipe-comment__btn--reply', function () {
    const $it = $(this)
    ShowEditor($it.data('title'), $it.data('id'), $it.data('commentid'))
  })

  // editor cancel
  $('#pipeEditorCancel').click(function () {
    _hideEditor()
  })

  // editor add
  $('#pipeEditorAdd').click(function () {
    const label = $(this).data('label')
    const label2 = $(this).data('label2')
    const $editor = $('#pipeEditor')
    const $editorAdd = $(this)
    const $commentContent = $('#pipeEditorComment')

    if ($editorAdd.hasClass('disabled')) {
      return;
    }

    if ($.trim($b3logEditor.val()).length === 0) {
      alert(label2)
      return;
    }

    let requestData = {
      'articleID': $editor.data('id'),
      'content': $b3logEditor.val()
    }

    if ($editor.data('commentid')) {
      requestData.parentCommentID = $editor.data('commentid')
    }

    $editorAdd.addClass('disabled')

    $.ajax({
      url: `${$commentContent.data('blogurl')}/comments`,
      data: JSON.stringify(requestData),
      type: 'POST',
      success: (result) => {
        if (result.code === 0) {
          _hideEditor()
          const $commentsCnt = $('#pipeCommentsCnt')
          const $comments = $('#pipeComments')

          if ($commentsCnt.length === 0) {
            $comments.removeClass('pipe-comment__null').unbind('click')
              .html(`<div class="pipe-comment__header"><span id="pipeCommentsCnt">1</span>${label}</div><div>${result.data}</div>`)
          } else {
            $commentsCnt.text(parseInt($commentsCnt.text()) + 1)
            $('#pipeComments > div > section').last().after(result.data)
          }

          LazyLoadCSSImage()
          LazyLoadImage()
          ParseMarkdown()
          InitHljs()

          $b3logEditor.val('')
          localStorage.removeItem('pipeEditorComment')
        } else {
          alert(result.msg)
        }
        $editorAdd.removeClass('disabled')
      }
    })
  })
}