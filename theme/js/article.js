/**
 * @fileoverview article tool for every theme
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.3.1.1, Feb 5, 2019
 */

import $ from 'jquery'
import { LazyLoadCSSImage, LazyLoadImage, ParseMarkdown } from './common'
import config from '../../pipe.json'

/**
 * @description 初始化目录
 * @param {string} tocId 目录 id
 * @param {string} articleId 目录对应正文 id
 */
export const InitToc = (
  tocId, articleId, articleOffset = 0, activeClass = 'toc__item--active') => {
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
        offsetTop: this.offsetTop,
      })
    })

    // 当前目录样式
    let scrollTop = $(window).scrollTop()
    for (let i = 0, iMax = tocPositionArray.length; i < iMax; i++) {
      if (scrollTop < tocPositionArray[i].offsetTop - articleOffset) {
        $toc.find('li').removeClass(activeClass)
        const index = i > 0 ? i - 1 : 0
        $toc.find(`a[href="#${$articleTocs[index].id}"]`).
          parent().
          addClass(activeClass)
        break
      }
    }
    if (scrollTop >= $articleTocs[$articleTocs.length - 1].offsetTop -
      articleOffset) {
      $toc.find('li').removeClass(activeClass)
      $toc.find('li:last').addClass(activeClass)
    }
  })

  $(window).scroll()
}

export const InitHljs = () => {
  const $codes = $('pre code')
  if ($codes.length === 0) {
    return false
  }

  var parseHljs = function () {
    $codes.each(function (i, block) {
      var $it = $(this)
      if ($it.hasClass('language-flow')) {
        return
      }

      if ($it.parent().find('.pipe-code__copy').length > 0) {
        return
      }

      $it.css({'max-height': $(window).height() - 40}).
        parent().
        prepend(`<textarea>${$it.text()}</textarea><div 
          aria-label="copy" onmouseover="this.setAttribute('aria-label', 'Copy')" 
          class="pipe-code__copy pipe-tooltipped pipe-tooltipped--w" 
          onclick="this.previousElementSibling.select();document.execCommand('copy');this.setAttribute('aria-label','Copied')"><svg><use xlink:href="#copy"></use></svg></div>`)

      if (!$('#pipeLang').data('markedavailable')) {
        hljs.highlightBlock(block)
      }
    })
  }

  if (typeof hljs === 'undefined' && !$('#pipeLang').data('markedavailable')) {
    $.ajax({
      method: 'GET',
      url: (config.StaticServer || config.Server) +
      '/theme/js/lib/highlight.min.js?9.12.0',
      dataType: 'script',
      cache: true,
    }).done(function () {
      parseHljs()
    })
  } else {
    parseHljs()
  }
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
    location.href = `${config.Server}/init`
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
  $('#pipeEditorComment textarea').focus()

  if (typeof Vditor !== 'undefined') {
    return
  }

  $.ajax({
    method: 'GET',
    url: `${config.StaticServer}/dist/vditor/index.min.js`,
    dataType: 'script',
    cache: true,
    success: () => {
      window.vditor = new Vditor('pipeEditorComment', {
        placeholder: $('#pipeEditorComment').data('placeholder'),
        height: 180,
        esc: () => {
          $('#pipeEditorCancel').click()
        },
        ctrlEnter: () => {
          $('#pipeEditorAdd').click()
        },
        preview: {
          delay: 500,
          show: false,
          url: `${config.Server}/api/console/markdown`,
          parse: (element) => {
            if (element.style.display === 'none') {
              return
            }
            LazyLoadImage()
            LazyLoadCSSImage()
            InitHljs()
            ParseMarkdown()
          },
        },
        upload: {
          max: 10 * 1024 * 1024,
          url: `${$('#pipeEditorComment').data('blogurl')}/upload`,
        },
        counter: 2048,
        resize: {
          enable: true,
          position: 'top'
        },
        lang: $('#pipeEditorComment').data('lang'),
        classes: {
          preview: 'pipe-content__reset',
        },
      })
    },
  })

}

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

  // comment null reply
  $('.pipe-comment__null').click(function () {
    const $commentNull = $(this)
    ShowEditor($commentNull.data('title'), $commentNull.data('id'))
  })

  // bottom reply
  $('#pipeCommentsWrap').on('click', '#pipeCommentBottomComment', function () {
    const $bottomComment = $(this)
    ShowEditor($bottomComment.data('text'), $bottomComment.data('id'))
  })

  // comment show reply
  $('#pipeCommentsWrap').on('click', '#pipeComments .fn__pointer', function () {
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
          },
        })
        return
      }
      $.ajax({
        url: `${$('#pipeEditorComment').data('blogurl')}/comments/${$it.data(
          'id')}/replies`,
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
                        <a href="${item.Author.URL}" class="ft__gray">${item.Author.Name}</a>
                        <span class="ft__nowrap ft__12 ft__gray"> • ${item.CreatedAt}</span>
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
              },
            })
            LazyLoadImage()
            LazyLoadCSSImage()
            InitHljs()
            ParseMarkdown()
          } else {
            alert(result.msg)
            $it.removeClass('disabled')
          }
        },
      })
    } else {
      $svg.addClass('pipe-comment__chevron-down')
      $it.next().slideToggle({
        queue: false,
        complete: () => {
          $it.removeClass('disabled')
        },
      })
    }
  })

  // comment remove
  $('#pipeCommentsWrap').
    on('click', '#pipeComments .pipe-comment__btn--danger', function () {
      const $it = $(this)
      if (confirm($it.data('label'))) {
        $.ajax({
          url: `${$('#pipeEditorComment').
            data('blogurl')}/comments/${$it.data('id')}`,
          type: 'DELETE',
          success: (result) => {
            if (result.code === 0) {
              const $commentsCnt = $('#pipeCommentsCnt')
              const $comments = $('#pipeComments')
              const $item = $(`#pipeComment${$it.data('id')}`)

              if ($('#pipeComments > div > section').length === 1) {
                $comments.addClass('pipe-comment__null').html(`<svg><use xlink:href="#icon-reply"></use></svg>
${$it.data('label2')}`).click(function () {
                  ShowEditor($comments.data('title'), $comments.data('id'))
                })
              } else {
                $item.remove()
                $commentsCnt.text(parseInt($commentsCnt.text()) - 1)
              }
            } else {
              alert(result.msg)
            }
          },
        })
      }
    })

  // comment reply
  $('#pipeCommentsWrap').
    on('click', '#pipeComments .pipe-comment__btn--reply', function () {
      const $it = $(this)
      ShowEditor($it.data('title'), $it.data('id'), $it.data('commentid'))
    })

  if ($('#pipeEditorComment').length === 0) {
    return
  }

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
      return
    }

    if ($.trim(vditor.getValue()).length === 0) {
      alert(label2)
      return
    }

    let requestData = {
      'articleID': $editor.data('id'),
      'content': vditor.getValue(),
    }

    if ($editor.data('commentid')) {
      requestData.parentCommentID = $editor.data('commentid')
    }

    $editorAdd.addClass('pipe-btn--disabled')

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
            $comments.removeClass('pipe-comment__null').
              unbind('click').
              html(
                `<div class="pipe-comment__header"><span id="pipeCommentsCnt">1</span>${label}</div><div>${result.data}</div>
<nav class="pipe-comment__pagination fn__clear">
    <span class="fn__right pipe-comment__btn" data-text="${$comments.data(
                  'title')}" data-id="${$editor.data('id')}" id="pipeCommentBottomComment">
         <svg><use xlink:href="#icon-reply"></use></svg> ${label}
    </span>
</nav>`)
          } else {
            $commentsCnt.text(parseInt($commentsCnt.text()) + 1)
            $('#pipeComments > div > section').last().after(result.data)
          }

          if ($(this).data('editable')) {
            $('#pipeComments > div > section').
              last().
              find('.pipe-comment__btn--danger').
              removeClass('pipe-comment__btn--none')
          }

          LazyLoadCSSImage()
          LazyLoadImage()
          ParseMarkdown()
          InitHljs()

          vditor.setValue('')
        } else {
          alert(result.msg)
        }
        $editorAdd.removeClass('pipe-btn--disabled')
      },
    })
  })
}