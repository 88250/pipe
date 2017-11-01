/**
 * @fileoverview common tool for every theme
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.2.0.0, Oct 19, 2017
 */

import $ from 'jquery'

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
      it.style.backgroundImage = 'url()'
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
 * @param {string} classes 需要延迟加载的类名
 */
export const LazyLoadCSSImage = (classes) => {
  const loadCSSImage = (it) => {
    const testImage = document.createElement('img')
    testImage.src = it.getAttribute('data-src')
    testImage.addEventListener('load', () => {
      it.style.backgroundImage = 'url(' + testImage.src + ')'
    })
    it.removeAttribute('data-src')
  }

  if (!('IntersectionObserver' in window)) {
    $(classes).each((item) => {
      if (item.getAttribute('data-src')) {
        loadCSSImage(item)
      }
    })
    return
  }

  if (window.CSSImageIntersectionObserver) {
    window.CSSImageIntersectionObserver.disconnect()
    $(classes).each(function () {
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
    $(classes).each(function () {
      window.CSSImageIntersectionObserver.observe(this)
    })
  }
}

/**
 * @description 删除评论
 * @param {string} id 评论 id
 * @param {function} succCB 成功回调
 * @param {function} errorCB 失败回调
 */
export const ReomveComment = (id, succCB, errorCB) => {
  $.ajax({
    url: `${location.origin}/api/console/comments/${id}`,
    type: 'DELETE',
    success: (result) => {
      if (result.code === 0) {
        succCB && succCB()
      } else {
        errorCB && errorCB(result.msg)
      }
    }
  })
}

/**
 * @description 添加评论
 * @param {string} blogURL 博客原始地址
 * @param {string} data 评论内容
 * @param {function} succCB 成功回调
 * @param {function} errorCB 失败回调
 */
export const AddComment = (blogURL, data, succCB, errorCB) => {
  $.ajax({
    url: `${blogURL}/comments`,
    data: JSON.stringify(data),
    type: 'POST',
    success: (result) => {
      if (result.code === 0) {
        localStorage.removeItem('themeCommentContent')
        succCB && succCB(result.data)
      } else {
        errorCB && errorCB(result.msg)
      }
    }
  })
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
 * @description 对输入内容进行本地存储
 * @param {string} id 输入框 id
 */
export const LocalStorageInput = (id) => {
  $(`#${id}`).val(localStorage.getItem(`input${id}`) || '').keyup(function () {
    localStorage.setItem(`input${id}`, $(this).val())
  })
}

/**
 * @description 评论框初始化
 * @param {string} emojiId 表情 id
 * @param {string} commentId 评论框 id
 */
export const InitEditor = (emojiId, commentId) => {
  const _getCursorEndPosition = (textarea) => {
    textarea.focus()
    if (textarea.setSelectionRange) { // W3C
      return textarea.selectionEnd
    } else if (document.selection) { // IE
      let i = 0
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

  $(`#${emojiId}`).find('span').click(function () {
    const comment = document.getElementById(commentId)
    let endPosition = _getCursorEndPosition(comment)
    const key = this.title + ' '
    let textValue = comment.value
    textValue = textValue.substring(0, endPosition) + key + textValue.substring(endPosition, textValue.length)
    comment.value = textValue
    if (document.selection) {
      endPosition -= textValue.split('\n').length - 1
      const oR = comment.createTextRange()
      oR.collapse(true)
      oR.moveStart('character', endPosition + key.length)
      oR.select()
    } else {
      comment.setSelectionRange(endPosition + key.length, endPosition + key.length)
    }
  })
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