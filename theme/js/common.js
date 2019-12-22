/**
 * @fileoverview common tool for every theme
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.8.0.0, Nov 19, 2019
 */

import $ from 'jquery'
import NProgress from 'nprogress'
import pjax from './lib/pjax'

export const ParseHljs = () => {
  Vditor.highlightRender({
    style: 'github',
    enable: false,
  }, document)
}
/**
 * @description 初始化 markdown 解析
 */
export const ParseMarkdown = () => {
  Vditor.codeRender(document.body, $('#pipeLang').data('lang'))
  Vditor.mathRender(document.body)
  Vditor.abcRender()
  Vditor.chartRender()
  Vditor.mediaRender(document)
  Vditor.mermaidRender(document.body)
}

/**
 * @description 图片预览
 */
export const PreviewImg = () => {
  const _previewImg = (it) => {
    const $it = $(it)
    var top = it.offsetTop,
      left = it.offsetLeft

    $('body').
      append('<div class="pipe-preview__img" onclick="this.remove()"><img style="transform: translate3d(' +
        Math.max(0, left) + 'px, ' +
        Math.max(0, (top - $(window).scrollTop())) + 'px, 0)" src="' +
        ($it.attr('src').split('?imageView2')[0]) + '"></div>')

    $('.pipe-preview__img').css({
      'background-color': '#fff',
      'position': 'fixed',
    })

    $('.pipe-preview__img img')[0].onload = function () {
      const $previewImage = $('.pipe-preview__img')
      $previewImage.find('img').css('transform', 'translate3d(' +
        (Math.max(0, $(window).width() - $previewImage.find('img').width()) /
          2) + 'px, ' +
        (Math.max(0, $(window).height() - $previewImage.find('img').height()) /
          2) + 'px, 0)')

      // fixed chrome render transform bug
      setTimeout(function () {
        $previewImage.width($(window).width())
      }, 300)
    }
  }
  // init
  $('body').on('click', '.vditor-reset img', function () {
    _previewImg(this)
  })
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
    $('img').each(function () {
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
        if ((typeof entrie.isIntersecting === 'undefined'
          ? entrie.intersectionRatio !== 0
          : entrie.isIntersecting)
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
    $cssImage.each(function () {
      if (this.tagName.toLowerCase() === 'img') {
        return
      }
      if (this.getAttribute('data-src')) {
        loadCSSImage(this)
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
    window.CSSImageIntersectionObserver = new IntersectionObserver(
      (entries) => {
        entries.forEach((entrie) => {
          if ((typeof entrie.isIntersecting === 'undefined'
            ? entrie.intersectionRatio !== 0
            : entrie.isIntersecting)
            && entrie.target.getAttribute('data-src') &&
            entrie.target.tagName.toLocaleLowerCase() !== 'img') {
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
 * @description 登出
 */
export const Logout = () => {
  $.ajax({
    url: `${$('#pipeLang').data('server')}/api/logout`,
    type: 'POST',
  })
}

/**
 * @description 去除查询字符串中的 'b3id=xxx' 参数
 */
const TrimB3Id = () => {
  const search = location.search
  if (search.indexOf('b3id') === -1) {
    return
  }
  history.replaceState('', '',
    window.location.href.replace(/(&b3id=\w{8})|(b3id=\w{8}&)|(\?b3id=\w{8})/,
      ''))
}

/**
 * @description 添加版权
 */
const addCopyright = () => {
  const genCopy = (author, link) => {
    return [
      '',
      '',
      `作者：${author}`,
      `链接：${link}`,
      '来源：Pipe',
      '著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。',
    ]
  }

  $('body').on('copy', '.vditor-reset', function (event) {
    if (!window.getSelection) {
      return
    }

    let selectionObj = window.getSelection()
    const author = $(this).data('author') || 'Pipe'
    let link = $(this).data('link') || location.href

    if (selectionObj.toString().length < 128) {
      return
    }

    if (selectionObj.rangeCount) {
      var container = document.createElement('div')
      for (var i = 0, len = selectionObj.rangeCount; i < len; ++i) {
        container.appendChild(selectionObj.getRangeAt(i).cloneContents())
      }
    }

    if ('object' === typeof event.originalEvent.clipboardData) {
      event.originalEvent.clipboardData.setData(
        'text/html', container.innerHTML + genCopy(author, link).join('<br>'))
      event.originalEvent.clipboardData.setData(
        'text/plain', selectionObj.toString() +
        genCopy(author, link).join('\n'))
      container.remove()
      event.preventDefault()
      return
    }

    $('body').append(`<div id="pipeFixCopy" style="position: fixed; left: -9999px;">
${selectionObj.toString()}${genCopy(author, link).join('<br>')}</div>`)
    window.getSelection().selectAllChildren($('#pipeFixCopy')[0])
    setTimeout(function () {
      $('#pipeFixCopy').remove()
    }, 200)
  })
}

/**
 * @description 初始化 pjax
 * @param {function} cb pjax 成功加载后的回调函数
 */
export const initPjax = (cb) => {
  if ($('#pjax').length === 1) {
    pjax({
      selector: 'a',
      container: '#pjax',
      show: '',
      cache: false,
      storage: true,
      titleSuffix: '',
      filter: function (href) {
        if (href.indexOf('/atom') > -1 ||
          href.indexOf($('#pipeLang').data('server') + '/admin') > -1) {
          return true
        } else if (href.indexOf($('#script').data('blogurl')) > -1) {
          return false
        }
        return true
      },
      callback: function () {
        LazyLoadCSSImage()
        LazyLoadImage()
        ParseMarkdown()
        ParseHljs()
        cb && cb()
      },
    })
    NProgress.configure({showSpinner: false})
    $('#pjax').bind('pjax.start', function () {
      NProgress.start()
    })
    $('#pjax').bind('pjax.end', function () {
      window.scroll(window.scrollX, 0)
      NProgress.done()
    })
  }
}

(() => {
  $.ajax({
    method: 'GET',
    url: 'https://cdn.jsdelivr.net/npm/vditor@1.10.11/dist/index.min.js',
    dataType: 'script',
    cache: true,
    success: () => {
      ParseMarkdown()
      ParseHljs()
    },
  })
  TrimB3Id()
  LazyLoadCSSImage()
  LazyLoadImage()
  addCopyright()

  // if ('serviceWorker' in navigator && 'caches' in window && 'fetch' in window && config.RuntimeMode === 'prod') {
  // navigator.serviceWorker.register(`${config.Server}/sw.min.js?${config.StaticResourceVersion}`, {scope: '/'})
  // }
})()
