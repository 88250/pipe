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
    window.Util.killBrowser()
    window.Util.lazyLoadCSSImage()
    this._header()
  },
  _header: function () {
    var $input = $('#headerSearch input')
    $('#headerSearch').click(function () {
      $input.width(95).focus()
    })
    $input.blur(function () {
      $(this).width(0)
    })
  }
}

Common.init()