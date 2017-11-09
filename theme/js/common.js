/**
 * @fileoverview common tool for every theme
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.2.0.0, Oct 19, 2017
 */

import $ from 'jquery'
import hljs from 'highlight.js'

/**
 * @description 图片预览
 */
export const PreviewImg = () => {
  const _previewImg = (it) => {
    const $it = $(it);
    var top = it.offsetTop,
      left = it.offsetLeft;
    if ($it.closest('.comments').length === 1) {
      top = top + $it.closest('li')[0].offsetTop;
      left = left + $('.comments')[0].offsetLeft + 15;
    }

    $('body').append('<div class="pipe-preview__img" onclick="this.remove()"><img style="transform: translate3d(' +
      Math.max(0, left) + 'px, ' + Math.max(0, (top - $(window).scrollTop())) + 'px, 0)" src="' +
      ($it.attr('src').split('?imageView2')[0]) + '"></div>');

    $('.pipe-preview__img').css({
      'background-color': '#fff',
      'position': 'fixed'
    });

    $('.pipe-preview__img img')[0].onload = function () {
      const $previewImage = $('.pipe-preview__img');
      $previewImage.find('img').css('transform', 'translate3d(' +
        (Math.max(0, $(window).width() - $previewImage.find('img').width()) / 2) + 'px, ' +
        (Math.max(0, $(window).height() - $previewImage.find('img').height()) / 2) + 'px, 0)');

      // fixed chrome render transform bug
      setTimeout(function () {
        $previewImage.width($(window).width());
      }, 300);
    }
  }
  // init
  $('body').on('click', '.pipe-content__reset img', function () {
    _previewImg(this)
  });
}

/**
 * @description 图片延迟加载
 * @returns {boolean}
 */
export const LazyLoadImage = () => {
  const loadImg = (it) => {
    const testImage = document.createElement('img')
    testImage.src = it.getAttribute('data-src')
    testImage.addEventListener('load', () => {
      it.src = testImage.src
      it.style.backgroundImage = 'none'
      it.style.backgroundColor = 'transparent'
    })
    it.removeAttribute('data-src')
  }

  if (!('IntersectionObserver' in window)) {
    $('img').each(() => {
      if (this.getAttribute('data-src')) {
        loadImg(this)
      }
    })
    return false
  }

  if (window.imageIntersectionObserver) {
    window.imageIntersectionObserver.disconnect()
    $('img').each(function () {
      window.imageIntersectionObserver.observe(this)
    })
  } else {
    window.imageIntersectionObserver = new IntersectionObserver((entries) => {
      entries.forEach((entrie) => {
        if ((typeof entrie.isIntersecting === 'undefined' ? entrie.intersectionRatio !== 0 : entrie.isIntersecting)
          && entrie.target.getAttribute('data-src')) {
          loadImg(entrie.target)
        }
      })
    })
    $('img').each(function () {
      window.imageIntersectionObserver.observe(this)
    })
  }
}

/**
 * @description CSS 背景图延迟加载
 */
export const LazyLoadCSSImage = () => {
  const loadCSSImage = (it) => {
    const testImage = document.createElement('img')
    testImage.src = it.getAttribute('data-src')
    testImage.addEventListener('load', () => {
      it.style.backgroundImage = 'url(' + testImage.src + ')'
    })
    it.removeAttribute('data-src')
  }

  const $cssImage = $('*[data-src]')
  if (!('IntersectionObserver' in window)) {
    $cssImage.each((item) => {
      if (item.tagName.toLowerCase() === 'img') {
        return
      }
      if (item.getAttribute('data-src')) {
        loadCSSImage(item)
      }
    })
    return
  }

  if (window.CSSImageIntersectionObserver) {
    window.CSSImageIntersectionObserver.disconnect()
    $cssImage.each(function () {
      if (this.tagName.toLowerCase() === 'img') {
        return
      }
      window.CSSImageIntersectionObserver.observe(this)
    })
  } else {
    window.CSSImageIntersectionObserver = new IntersectionObserver((entries) => {
      entries.forEach((entrie) => {
        if ((typeof entrie.isIntersecting === 'undefined' ? entrie.intersectionRatio !== 0 : entrie.isIntersecting)
          && entrie.target.getAttribute('data-src') && entrie.target.tagName.toLocaleLowerCase() !== 'img') {
          loadCSSImage(entrie.target)
        }
      })
    })
    $cssImage.each(function () {
      if (this.tagName.toLowerCase() === 'img') {
        return
      }
      window.CSSImageIntersectionObserver.observe(this)
    })
  }
}

/**
 * @description 不兼容 IE 9 以下
 */
export const KillBrowser = () => {
  const index = navigator.userAgent.indexOf('MSIE ')
  if (index > -1) {
    if (parseInt(navigator.userAgent.substr(index + 5).split(';')[0]) < 9) {
      document.body.innerHTML = `<div>为了让浏览器能更好的发展，您能拥有更好的体验，让我们放弃使用那些过时的、不安全的浏览器吧！</div><br>
        <ul>
          <li><a href="http://www.google.com/chrome" target="_blank" rel="noopener">谷歌浏览器</a></li>
          <li><a href="http://www.mozilla.com/" target="_blank" rel="noopener">火狐</a></li>
          <li>
            <a href="http://se.360.cn/" target="_blank" rel="noopener">360</a>或者
            <a href="https://www.baidu.com/s?wd=%E6%B5%8F%E8%A7%88%E5%99%A8" target="_blank" rel="noopener">其它浏览器</a>
          </li>
        </ul>`
    }
  }
}

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

/**
 * @description 登出
 */
export const Logout = () => {
  $.ajax({
    url: `${location.origin}/api/logout`,
    type: 'POST',
    success: (result) => {
      window.location.href = 'https://hacpai.com/logout'
    }
  })
}

/**
 * @description 去除查询字符串中的 'b3id=xxx' 参数
 */
export const TrimB3Id = () => {
  const search = location.search
  if (search.indexOf('b3id') === -1) {
    return
  }
  history.replaceState('', '', window.location.href.replace(/(&b3id=\w{8})|(b3id=\w{8}&)|(\?b3id=\w{8}$)/, ''))
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
   * @description 获取 textarea 光标位置
   * @param {Bom} textarea textarea 对象
   * @private
   */
  const _getCursorEndPosition = (textarea) => {
    textarea.focus()
    if (textarea.setSelectionRange) { // W3C
      return textarea.selectionEnd
    } else if (document.selection) { // IE
      let i
      const oS = document.selection.createRange()
      const oR = document.body.createTextRange()
      oR.moveToElementText(textarea)
      oS.getBookmark()
      for (i = 0; oR.compareEndPoints('StartToStart', oS) < 0 && oS.moveStart("character", -1) !== 0; i++) {
        if (textarea.value.charAt(i) === '\n') {
          i++
        }
      }
      return i
    }
  }

  /**
   * @description 隐藏编辑器
   * @private
   */
  const _hideEditor = () => {
    const $editor = $('#pipeEditor')
    $editor.css({'bottom': `-${$editor.outerHeight()}px`, 'opacity': 0})
    $('body').css('padding-bottom', 0)
  }

  // comment local storage
  $('#pipeEditorComment').val(localStorage.getItem('pipeEditorComment') || '').keyup(function () {
    localStorage.setItem('pipeEditorComment', $(this).val())
  })

  // comment null reply
  $('.pipe-comment__null').click(function () {
    const $commentNull = $(this)
    ShowEditor($commentNull.data('title'), $commentNull.data('id'))
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
      if  ($it.next().find('.pipe-comment__item').length > 0) {
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

  // comment remove
  $('body').on('click', '#pipeComments .pipe-comment__btn--reply', function () {
    const $it = $(this)
    ShowEditor($it.data('title'), $it.data('id'), $it.data('commentid'))
  })

  // editor emoji
  $('.pipe-editor__emotion').find('span').click(function () {
    const comment = document.getElementById('pipeEditorComment')
    const endPosition = _getCursorEndPosition(comment)
    const key = this.title + ' '
    let textValue = comment.value
    textValue = textValue.substring(0, endPosition) + key + textValue.substring(endPosition, textValue.length)
    comment.value = textValue
  })

  // editor hot key
  $('#pipeEditorComment').keydown(function (event) {
    if (event.metaKey && 13 === event.keyCode) {
      $('#pipeEditorAdd').click()
    }

    if (27 === event.keyCode) {
      _hideEditor()
    }
  }).keypress(function (event) {
    if (event.ctrlKey && 10 === event.charCode) {
      $('#pipeEditorAdd').click()
    }
  });

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

    if ($.trim($commentContent.val()).length === 0) {
      alert(label2)
      return;
    }

    let requestData = {
      'articleID': $editor.data('id'),
      'content': $commentContent.val()
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

          $comments.find('pre > code').each(function (i, block) {
            hljs.highlightBlock(block)
          })
          LazyLoadCSSImage()
          LazyLoadImage()
          $commentContent.val('')
          localStorage.removeItem('pipeEditorComment')
        } else {
          alert(result.msg)
        }
        $editorAdd.removeClass('disabled')
      }
    })
  })
}