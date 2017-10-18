var Util = {
  imageIntersectionObserver: undefined,
  CSSImageIntersectionObserver: undefined,
  /**
   * @description 不兼容 IE 9 以下
   */
  killBrowser: function () {
    var index = navigator.userAgent.indexOf('MSIE ')
    if (index > -1) {
      if (parseInt(navigator.userAgent.substr(index + 5).split(';')[0]) < 9) {
        document.body.innerHTML = '<div>为了让浏览器能更好的发展，您能拥有更好的体验，让我们放弃使用那些过时的、不安全的浏览器吧！</div><br>' +
          '<ul>' +
          '<li><a href="http://www.google.com/chrome" target="_blank" rel="noopener">谷歌浏览器</a></li>' +
          '<li><a href="http://www.mozilla.com/" target="_blank" rel="noopener">火狐</a></li>' +
          '<li><a href="http://se.360.cn/" target="_blank" rel="noopener">360</a>或者<a href="https://www.baidu.com/s?wd=%E6%B5%8F%E8%A7%88%E5%99%A8" target="_blank" rel="noopener">其它浏览器</a></li>' +
          '</ul>'
      }
    }
  },
  /**
   * @description 图片延迟加载
   * @returns {boolean}
   */
  lazyLoadImage: function () {
    var loadImg = function (it) {
      var testImage = document.createElement('img');
      testImage.src = it.getAttribute('data-src');
      testImage.addEventListener('load', function () {
        it.src = testImage.src;
        it.style.backgroundImage = 'url()';
        it.style.backgroundColor = 'transparent';
      });
      it.removeAttribute('data-src');
    };

    if (!('IntersectionObserver' in window)) {
      $('img').each(function () {
        if (this.getAttribute('data-src')) {
          loadImg(this);
        }
      });
      return false;
    }

    if (Util.imageIntersectionObserver) {
      Util.imageIntersectionObserver.disconnect();
      $('img').each(function () {
        Util.imageIntersectionObserver.observe(this);
      });
    } else {
      Util.imageIntersectionObserver = new IntersectionObserver(function (entries) {
        entries.forEach(function (entrie) {
          if ((typeof entrie.isIntersecting === 'undefined' ? entrie.intersectionRatio !== 0 : entrie.isIntersecting)
            && entrie.target.getAttribute('data-src')) {
            loadImg(entrie.target);
          }
        });
      });
      $('img').each(function () {
        Util.imageIntersectionObserver.observe(this);
      });
    }
  },
  /**
   * @description CSS 背景图延迟加载
   * @param classes{string} 需要延迟加载的类名
   * @returns {boolean}
   */
  lazyLoadCSSImage: function (classes) {
    var loadCSSImage = function (it) {
      var testImage = document.createElement('img');
      testImage.src = it.getAttribute('data-src');
      testImage.addEventListener('load', function () {
        it.style.backgroundImage = 'url(' + testImage.src + ')';
      });
      it.removeAttribute('data-src');
    };

    if (!('IntersectionObserver' in window)) {
      $(classes).each(function () {
        if (this.getAttribute('data-src')) {
          loadCSSImage(this);
        }
      });
      return false;
    }

    if (Util.CSSImageIntersectionObserver) {
      Util.CSSImageIntersectionObserver.disconnect();
      $(classes).each(function () {
        Util.CSSImageIntersectionObserver.observe(this);
      });
    } else {
      Util.CSSImageIntersectionObserver = new IntersectionObserver(function (entries) {
        entries.forEach(function (entrie) {
          if ((typeof entrie.isIntersecting === 'undefined' ? entrie.intersectionRatio !== 0 : entrie.isIntersecting)
            && entrie.target.getAttribute('data-src') && entrie.target.tagName.toLocaleLowerCase() !== 'img') {
            loadCSSImage(entrie.target);
          }
        });
      });
      $(classes).each(function () {
        Util.CSSImageIntersectionObserver.observe(this);
      });
    }
  },
  reomveComment: function (id, succCB, errorCB) {
    $.ajax({
      url: conf.server + '/api/console/comments/' + id,
      type: 'DELETE',
      success: function (result) {
        if (result.code === 0) {
          succCB && succCB();
        } else {
          errorCB && errorCB(result.msg);
        }
      }
    });
  },
  addComment: function (id, succCB, errorCB) {
    $.ajax({
      url: conf.server + '/api/console/comments/' + id,
      type: 'POST',
      success: function (result) {
        if (result.code === 0) {
          localStorage.removeItem('themeCommentName');
          localStorage.removeItem('themeCommentEmail');
          localStorage.removeItem('themeCommentURL');
          localStorage.removeItem('themeCommentContent');
          succCB && succCB(result.data);
        } else {
          errorCB && errorCB(result.msg);
        }
      }
    });
  },
  localStorageComment: function () {
    $('#commentName').val(localStorage.getItem('themeCommentName') || '').keyup(function () {
      localStorage.setItem('themeCommentName', $(this).val())
    });
    $('#commentEmail').val(localStorage.getItem('themeCommentEmail') || '').keyup(function () {
      localStorage.setItem('themeCommentEmail', $(this).val())
    });
    $('#commentURL').val(localStorage.getItem('themeCommentURL') || '').keyup(function () {
      localStorage.setItem('themeCommentURL', $(this).val())
    });
    $('#commentContent').val(localStorage.getItem('themeCommentContent') || '').keyup(function () {
      localStorage.setItem('themeCommentContent', $(this).val())
    });
  }
};

