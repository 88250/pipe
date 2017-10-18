/**
 * @fileoverview util and every page should be used.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.1.0.0, Oct 12, 2017
 */
var Common = {
  /**
   * @description 页面初始化
   */
  init: function () {
    Util.killBrowser()
    Util.lazyLoadCSSImage('.avatar, .article__thumb')
    Util.lazyLoadImage()
    this._header()
    this._share()
  },
  _header: function () {
    var $input = $('#headerSearch input')
    $('#headerSearch').click(function () {
      $input.width(95).focus()
    })
    $input.blur(function () {
      $(this).width(0)
    })
  },
  _share: function () {
    var shareURL = $('#qrCode').data('url')
    var avatarURL = $('#qrCode').data('avatar')

    new QRCode(document.getElementById('qrCode'), {
      text: shareURL,
      width: 128,
      height: 128
    })
    $('#qrCode').hide()

    var title = encodeURIComponent($('#qrCode').data('title') + ' - ' + $('#qrCode').data('blogtitle')),
      url = encodeURIComponent(shareURL)

    var urls = {}
    urls.tencent = 'http://share.v.t.qq.com/index.php?c=share&a=index&title=' + title +
      '&url=' + url + '&pic=' + avatarURL
    urls.weibo = 'http://v.t.sina.com.cn/share/share.php?title=' +
      title + '&url=' + url + '&pic=' + avatarURL
    urls.google = 'https://plus.google.com/share?url=' + url
    urls.twitter = 'https://twitter.com/intent/tweet?status=' + title + ' ' + url

    $('#share > span').click(function () {
      var key = $(this).data('type')
      if (key === 'wechat') {
        $('#qrCode').slideToggle()
        return false
      }
      window.open(urls[key], '_blank', 'top=100,left=200,width=648,height=618')
    })
  },
  toggleSide: function () {
    $('body').toggleClass('body--side')
    if ($('body').hasClass('body--side')) {
      $('#editor').width(940)
    } else {
      $('#editor').width('100%')
    }
  },
  removeComment(id, label, label2) {
    if (confirm(label)) {
      Util.reomveComment(id, function () {
        $('#comment' + id).remove()
        $('#commentsCnt').text(parseInt($('#commentsCnt').text()) - 1)
        if ($('#comments section').length === 0) {
          $('#comments').addClass('ft-center comment__null fn-bottom').attr('onclick', 'Common.showComment()')
            .html(label2 + ' <svg><use xlink:href="#comment"></use></svg>')
        }
      }, function (msg) {
        alert(msg)
      })
    }
  },
  showComment: function (reply, id, commentId) {
    $('#editor').css('bottom', '0').data('id', id).data('commentid', commentId)
    $('body').css('padding-bottom', $('#editor').outerHeight() + 'px')
    $('#replyObject').text(reply)
  },
  hideComment: function () {
    $('#editor').css('bottom', '-400px')
    $('body').css('padding-bottom', 0)
  },
  addComment: function (label) {
    Util.addComment($('#editor').data('commentid') || $('#editor').data('id'), function (data) {
      Common.hideComment();
      $('#commentsCnt').text(parseInt($('#commentsCnt').text()) + 1);

      var contentHTML = '<section></section>';
      // TODO contentHTML
      if (parseInt($('#commentsCnt').text()) === 1) {
        $('#comments').removeClass('ft-center comment__null').removeAttr('onclick')
          .html('<div class="module__header"><span id="commentsCnt">1</span> ' +
            label + '</div>' + contentHTML);
      } else {
        $('#comments section:last').after(contentHTML);
      }
    }, function (msg) {
      alert(msg)
    })
  }
}

Common.init()