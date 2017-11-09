/**
 * @fileoverview common tool for every theme
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.2.0.0, Oct 19, 2017
 */

import $ from 'jquery'

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