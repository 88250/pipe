/**
 * @fileoverview util and every page should be used.
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.3.0.0, Oct 31, 2018
 */

import $ from 'jquery'
import APlayer from 'aplayer'
import Icon from './symbol'
import {
  initPjax,
  KillBrowser,
  PreviewImg,
} from '../../../js/common'
import config from '../../../../pipe.json'

const Common = {
  /**
   * @description 页面初始化
   */
  init: () => {
    $('.header__menu').click(function () {
      $('.header__links').slideToggle()
    })

    new APlayer({
      container: document.getElementById('player'),
      autoplay: true,
      theme: '#ccc',
      preload: 'auto',
      lrcType: 3,
      listFolded: true,
      audio: [
        {
          name: 'Kiss The Rain',
          artist: '이루마',
          url: `${config.StaticServer}/theme/x/Littlewin/images/kisstherain.mp3`,
          cover: `${config.StaticServer}/theme/x/Littlewin/images/kisstherain.jpeg`,
          lrc: `${config.StaticServer}/theme/x/Littlewin/images/kisstherain.lrc`,
          theme: '#60b044',
        },
      ],
    })

    initPjax(() => {
      if ($('#pipeComments').length === 1) {
        $.ajax({
          method: 'GET',
          url: `${config.StaticServer}/theme/x/Littlewin/js/article.min.js`,
          dataType: 'script',
          cache: true,
        })
      }
    })

    PreviewImg()
    KillBrowser()
  },
  increase (max, time, id, count) {
    if (count < max) {
      setTimeout(() => {
        increase(max, time, id, ++count)
        if (document.getElementById(id)) {
          document.getElementById(id).innerHTML = count
        }
      }, time / max)
    }
  },
  addLevelToTag () {
    const $tags = $('#tags')
    const tagsArray = $tags.find('a')
    // 根据引用次数添加样式，产生云效果
    const max = parseInt(tagsArray.first().data('count'))
    const distance = Math.ceil(max / 5)
    for (let i = 0; i < tagsArray.length; i++) {
      const count = parseInt($(tagsArray[i]).data('count'))
      // 算出当前 tag 数目所在的区间，加上 class
      for (let j = 0; j < 5; j++) {
        if (count > j * distance && count <= (j + 1) * distance) {
          tagsArray[i].className = `tag tag--level${j}`
          break
        }
      }
    }

    // 按字母或者中文拼音进行排序
    $tags.html(tagsArray.get().sort(function (a, b) {
      var valA = $(a).text().toLowerCase()
      var valB = $(b).text().toLowerCase()
      // 对中英文排序的处理
      return valA.localeCompare(valB)
    }))
  },
}

if (!window.increase) {
  window.increase = Common.increase
  window.addLevelToTag = Common.addLevelToTag
  Icon()
  Common.init()
}
export default Common