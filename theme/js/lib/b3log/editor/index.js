/*
 * Copyright (C) 2012-2017,  b3log.org
 *
 */
/**
 * @fileOverview editor
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 0.3.1.0, Jun 9, 2018
 */
import TurndownService from 'turndown'
import $ from 'jquery'
import hotkeys from '../../hotkeys'
import {ajaxUpload, debounceInput, genUploaded, genUploading, insertTextAtCaret, replaceTextareaValue} from './tool'
import allEmoji from './emoji.json'

hotkeys();
const config = {
  props: {
    staticServePath: {
      type: String,
      required: false
    },
    previewClass: {
      type: String,
      required: false
    },
    uploadURL: {
      type: String,
      required: false
    },
    uploadMax: {
      type: Number,
      required: false
    },
    height: {
      type: Number,
      required: false
    },
    at: {
      type: Function,
      required: false
    },
    change: {
      type: Function,
      required: false
    },
    emoji: {
      type: Function,
      required: false
    },
    ctrlEnter: {
      type: Function,
      required: false
    },
    keyup: {
      type: Function,
      required: false
    },
    fetchUpload: {
      type: Function,
      required: false
    },
    esc: {
      type: Function,
      required: false
    },
    hasView: {
      type: Boolean,
      required: false
    },
    placeholder: {
      type: String,
      required: false
    },
    label: {
      loading: '',
      error: '',
      over: ''
    },
    tipClass: {
      type: String,
      required: false
    },
    emojiURL: {
      type: String,
      required: false
    },
    hintsBottom: {
      type: Number,
      required: false
    },
  }
};

export const Editor = (config) => {
  if (/Mac/.test(navigator.platform)) {
    Object.keys(config.label).forEach((key) => {
      if (config.label[key]) {
        config.label[key] = config.label[key].replace('ctrl', '⌘').replace('shift', '⇧');
      }
    })
  }
  const html = `<div
    class="b3log-editor"
    style="height: ${config.height || 200}px">
    <div class="b3log-editor__toolbar">
      <span data-type="emoji" aria-label="${config.label.emoji || 'emoji'}" class="${config.tipClass || 'tooltipped tooltipped-e'}"><svg><use xlink:href="#emoji"></use></svg></span>
      <span aria-label="${config.label.bold || 'bold'}" class="${config.tipClass || 'tooltipped tooltipped-e'}" data-prefix="**" data-suffix="**"><svg><use xlink:href="#bold"></use></svg></span>
      <span aria-label="${config.label.italic || 'italic'}" class="${config.tipClass || 'tooltipped tooltipped-e'}" data-prefix="*" data-suffix="*"><svg><use xlink:href="#italic"></use></svg></span>
      <span aria-label="${config.label.quote || 'quote'}" class="${config.tipClass || 'tooltipped tooltipped-e'}" data-prefix="> " data-suffix=""><svg><use xlink:href="#quote"></use></svg></span>
      <span aria-label="${config.label.link || 'link'}" class="${config.tipClass || 'tooltipped tooltipped-e'}" data-prefix="[" data-suffix="](http://)"><svg><use xlink:href="#link"></use></svg></span>
      <span aria-label="${config.label.upload || 'upload'}" class="${config.tipClass || 'tooltipped tooltipped-e'}">
        <label>
          <svg><use xlink:href="#upload"></use></svg>
          <input multiple="multiple" type="file"/>
        </label>
      </span>
      <span aria-label="${config.label.unorderedList || 'unorderedList'}" class="${config.tipClass || 'tooltipped tooltipped-e'}" data-prefix="* " data-suffix=""><svg><use xlink:href="#unordered-list"></use></svg></span>
      <span aria-label="${config.label.orderedList || 'orderedList'}" class="${config.tipClass || 'tooltipped tooltipped-e'}" data-prefix="1. " data-suffix=""><svg><use xlink:href="#ordered-list"></use></svg></span>
      <span aria-label="${config.label.view || 'view'}" class="${config.tipClass || 'tooltipped tooltipped-e'}" data-type="view"><svg><use xlink:href="#view"></use></svg></span>
      <span aria-label="${config.label.fullscreen || 'fullscreen'}" class="${config.tipClass || 'tooltipped tooltipped-e'}" data-type="fullscreen"><svg><use xlink:href="#fullscreen"></use></svg></span>
      <a aria-label="${config.label.question || 'question'}" class="${config.tipClass || 'tooltipped tooltipped-e'}" target="_blank" href="https://hacpai.com/guide/markdown">
        <svg><use xlink:href="#question"></use></svg>
      </a> 
      <div class="b3log-editor__emoji">
        <div class="fn-clear">
        </div>
        <div class="b3log-editor__emoji-tip">
          <a href="${config.emojiURL || 'https://hacpai.com/settings/function'}" target="_blank">${config.label.emojiTip || 'Tip'}</a>
        </div>
      </div>
    </div>
    <div class="b3log-editor__content">
      <div class="b3log-editor__textarea">
        <textarea
          placeholder="${config.placeholder || ''}"></textarea>
      </div>
      <div class="b3log-editor__markdown ${config.previewClass} fn-none" ref="pipeView"></div>
    </div>
  </div>`;

  let timerId = undefined;
  const $editor = $(`#${config.id}`);
  $editor.html(html);

  const $textarea = $editor.find('textarea');
  const textarea = $editor.find('textarea')[0];

  // hasView
  if (config.hasView) {
    $editor.find('.b3log-editor__toolbar > span[data-type="view"]').addClass('b3log-editor__icon--current');
    $editor.find('.b3log-editor__markdown').show();
    timerId = debounceInput(timerId, config.change, $editor)
  }

  // emoji
  const $emoji = $editor.find('.b3log-editor__emoji')
  let emojiHTML = '';

  if (config.emoji) {
    const genEmoji = (data) => {
      Object.keys(data).forEach((key) => {
        emojiHTML += `<span data-value="${data[key]} "><img src="${config.staticServePath || ''}/emoji/graphics/${key}.png"></span>`
      });
      $emoji.find('.fn-clear').html(emojiHTML);
      config._emojiData = data;
    }
    config.emoji(genEmoji)
  } else {
    config._emojiData = {
      "smile": "😄",
      "joy": "😂",
      "+1": "👍",
      "scream": "😱",
      "smiling_imp": "😈",
      "sob": "😭",
      "yum": "😋",
      "tada": "🎉",
      "ok_hand": "👌",
      "pray": "🙏",
      "punch": "👊",
      "heart_eyes": "😍"
    }
    Object.keys(config._emojiData).forEach((key) => {
      emojiHTML += `<span data-value="${config._emojiData[key]} ">${config._emojiData[key]}</span>`
    });
  }
  $emoji.find('.fn-clear').html(emojiHTML);
  $emoji.on('click', 'span', function () {
    insertTextAtCaret(textarea, $(this).data('value'), '', true)
    if (/firefox/i.test(navigator.userAgent)) {
      timerId = debounceInput(timerId, config.change, $editor)
    }
  })

  $editor.find('.b3log-editor__toolbar > span').click(function () {
    const $it = $(this);
    if ($it.data('prefix')) {
      insertTextAtCaret(textarea, $it.data('prefix'), $it.data('suffix'));
      if (/firefox/i.test(navigator.userAgent)) {
        timerId = debounceInput(timerId, config.change, $editor)
      }
    } else if ($it.data('type') === 'view') {
      if ($it.hasClass('b3log-editor__icon--current')) {
        $it.removeClass('b3log-editor__icon--current');
        $editor.find('.b3log-editor__markdown').hide();
      } else {
        $it.addClass('b3log-editor__icon--current');
        $editor.find('.b3log-editor__markdown').show();
        timerId = debounceInput(timerId, config.change, $editor)
      }
    } else if ($it.data('type') === 'fullscreen') {
      if ($it.find('use').attr('xlink:href') === '#fullscreen') {
        $it.find('use').attr('xlink:href', '#contract')
        $editor.find('.b3log-editor').addClass('b3log-editor--fullscreen')
      } else {
        $it.find('use').attr('xlink:href', '#fullscreen')
        $editor.find('.b3log-editor').removeClass('b3log-editor--fullscreen')
      }
    } else if ($it.data('type') === 'emoji') {
      $emoji.toggle()
    }
  })

  $editor.find('.b3log-editor__toolbar > span input').change(function (event) {
    insertTextAtCaret(textarea,
      genUploading(event.target.files, config.uploadMax, config.label.loading, config.label.over), '')
    if (/firefox/i.test(navigator.userAgent)) {
      timerId = debounceInput(timerId, config.change, $editor)
    }
    ajaxUpload(config.uploadURL, event.target.files, config.uploadMax, (response) => {
      genUploaded(response.data, textarea, config.label.loading, config.label.error)
      if (/firefox/i.test(navigator.userAgent)) {
        timerId = debounceInput(timerId, config.change, $editor)
      }
      event.target.value = ''
    }, (response) => {
      event.target.value = ''
      response && alert(response.msg)
    });
  })

  $textarea.focus(() => {
    $emoji.hide()
  }).scroll(function (event) {
    if ($editor.find('.b3log-editor__icon--current').length === 0) {
      return
    }
    const textScrollTop = event.target.scrollTop
    const textHeight = event.target.clientHeight
    const textScrollHeight = event.target.scrollHeight
    const view = $editor.find('.b3log-editor__markdown')[0]
    if ((textScrollTop / textHeight > 0.5)) {
      view.scrollTop = (textScrollTop + textHeight) *
        view.scrollHeight / textScrollHeight - textHeight
    } else {
      view.scrollTop = textScrollTop *
        view.scrollHeight / textScrollHeight
    }
  }).bind('paste', function (event) {
    if (event.originalEvent.clipboardData.getData('text/html').replace(/(^\s*)|(\s*)$/g, '') !== '') {
      let hasCode = false
      let turndownService = new TurndownService({
        converters: [{
          filter: ['pre', 'code'],
          replacement: function (content) {
            if (content.split('\n').length > 1) {
              hasCode = true
            }
            return '`' + content + '`'
          }
        }, {
          filter: ['img'],
          replacement: function (content, target) {
            if (1 === target.attributes.length) {
              return '';
            }

            config.fetchUpload && config.fetchUpload(target.src, (originalURL, url) => {
              replaceTextareaValue(textarea, originalURL, url)
              if (/firefox/i.test(navigator.userAgent)) {
                timerId = debounceInput(timerId, config.change, $editor)
              }
            });

            return `![${target.alt}](${target.src})`;
          }
        }],
        gfm: true
      })
      let markdownStr = turndownService.turndown(event.originalEvent.clipboardData.getData('text/html'))

      if (hasCode) {
        insertTextAtCaret(event.target, event.originalEvent.clipboardData.getData('text/plain'), '', true)
        if (/firefox/i.test(navigator.userAgent)) {
          timerId = debounceInput(timerId, config.change, $editor)
        }
      } else {
        const div = document.createElement('div')
        div.innerHTML = markdownStr
        markdownStr = div.innerText.replace(/\n{2,}/g, '\n\n').replace(/(^\s*)|(\s*)$/g, '')
        insertTextAtCaret(event.target, markdownStr, '', true)
        if (/firefox/i.test(navigator.userAgent)) {
          timerId = debounceInput(timerId, config.change, $editor)
        }
        div.remove()
      }
    } else if (event.originalEvent.clipboardData.getData('text/plain').replace(/(^\s*)|(\s*)$/g, '') !== '' &&
      event.originalEvent.clipboardData.files.length === 0) {
      insertTextAtCaret(event.target, event.originalEvent.clipboardData.getData('text/plain'), '', true)
      if (/firefox/i.test(navigator.userAgent)) {
        timerId = debounceInput(timerId, config.change, $editor)
      }
    } else if (event.originalEvent.clipboardData.files.length > 0) {
      // upload file
      if (config.uploadURL) {
        insertTextAtCaret(textarea,
          genUploading(event.originalEvent.clipboardData.files, config.uploadMax, config.label.loading, config.label.over),
          '', true)
        if (/firefox/i.test(navigator.userAgent)) {
          timerId = debounceInput(timerId, config.change, $editor)
        }
        ajaxUpload(config.uploadURL, event.originalEvent.clipboardData.files, config.uploadMax, (response) => {
          genUploaded(response.data, event.target, config.label.loading, config.label.error)
          if (/firefox/i.test(navigator.userAgent)) {
            timerId = debounceInput(timerId, config.change, $editor)
          }
        }, (response) => {
          response && alert(response.msg)
        })
      }
    }
    event.preventDefault();
  }).bind('drop', function (event) {
    event.stopPropagation();
    event.preventDefault();

    const files = event.originalEvent.dataTransfer.files
    if (files.length === 0) {
      return
    }
    insertTextAtCaret(textarea,
      genUploading(files, config.uploadMax, config.label.loading, config.label.over), '')
    if (/firefox/i.test(navigator.userAgent)) {
      timerId = debounceInput(timerId, config.change, $editor)
    }
    ajaxUpload(config.uploadURL, files, config.uploadMax, (response) => {
      genUploaded(response.data, textarea, config.label.loading, config.label.error)
      if (/firefox/i.test(navigator.userAgent)) {
        timerId = debounceInput(timerId, config.change, $editor)
      }
    }, (response) => {
      response && alert(response.msg)
    })
  }).bind('input', function (event) {
    // at and emoji hints
    const valueArray = this.value.substr(0, this.selectionStart).split('\n')
    const xValue = valueArray.slice(-1).pop()
    let $hints = $editor.find('.b3log-editor__hints')

    const genHintsHTML = (data) => {
      if (data.length === 0) {
        $hints.hide()
        return
      }
      let y = valueArray.length * 18 + 47 - $textarea.scrollTop()
      const zhReg = xValue.match(/[\u4E00-\u9FA5\uF900-\uFA2D\uFF00-\uFFEF]/g)
      const zhLength = zhReg === null ? 0 : zhReg.length
      const x = zhLength * 15 + (xValue.length - zhLength) * 9 + 10 + $textarea.scrollLeft()
      let hintsHTML = ''
      data.forEach((hintData, i) => {
        hintsHTML += `<li data-value="${hintData.value} " class="${i || 'b3log-editor__hints--current'}">
${hintData.imageURL ? '<img src="' + hintData.imageURL + '"/>' : hintData.value} ${hintData.label}</li>`
      })

      if ($hints.length === 0) {
        $textarea.after(`<ul class="b3log-editor__hints" style="left: ${x}px; top: ${y}px">${hintsHTML}</ul>`)
        $hints = $editor.find('.b3log-editor__hints')
      } else {
        $hints.html(hintsHTML).css({
          top: `${y}px`,
          left: `${x}px`,
          display: 'block'
        })
      }

      $hints.find('li').click(function () {
        $hints.hide()
        $textarea.focus()

        const $it = $(this)
        let splitChar = ':'
        if ($it.data('value').indexOf('@') === 0) {
          splitChar = '@'
        }
        if (/firefox/i.test(navigator.userAgent)) {
          const valueArray = textarea.value.substr(0, textarea.selectionStart).split(splitChar)
          valueArray.pop()
          textarea.value = valueArray.join(splitChar) + $it.data('value') + textarea.value.substr(textarea.selectionStart)
          textarea.selectionEnd = textarea.selectionStart = (valueArray.join(splitChar) + $it.data('value')).length
          timerId = debounceInput(timerId, config.change, $editor)
          return;
        }
        while (!textarea.value.substr(0, textarea.selectionEnd).endsWith(splitChar) &&
        textarea.value.substr(0, textarea.selectionEnd) !== '') {
          document.execCommand('delete', false)
        }
        document.execCommand('delete', false)
        document.execCommand('insertText', false, $it.data('value'))
      });

      if (y + $hints.outerHeight() > $editor.outerHeight() + (config.hintsBottom || 68)) {
        // hint 展现在上部
        $hints.css('top', `${y - $hints.outerHeight() - 18}px`)
      }
    }

    const getSearchKey = (splitChar) => {
      const xAtArray = xValue.split(splitChar)
      let searchKey = undefined
      if (xAtArray.length > 1) {
        if (xAtArray.length === 2 && xAtArray[0] === '') {
          if ((xAtArray[1] === '' || $.trim(xAtArray[1]) !== '') && xAtArray[1].indexOf(' ') === -1 &&
            xAtArray[1].length < 33) {
            searchKey = xAtArray[1]
          }
        } else {
          const prefAt = xAtArray[xAtArray.length - 2]
          const currentAt = xAtArray.slice(-1).pop()
          if (prefAt.slice(-1) === ' ' && currentAt.indexOf(' ') === -1 &&
            ((currentAt === '' || $.trim(currentAt) !== '') && currentAt.length < 33)) {
            searchKey = currentAt
          }
        }
      }
      return searchKey
    }

    const searchAt = getSearchKey('@')
    const searchEmoji = getSearchKey(':')
    if (searchAt === undefined && searchEmoji === undefined) {
      $hints.hide()
    } else {
      if (searchAt !== undefined) {
        config.at && config.at(searchAt, genHintsHTML)
      }
      if (searchEmoji !== undefined) {
        const matchEmoji = []
        let emojies = allEmoji
        if (searchEmoji === '') {
          emojies = config._emojiData
        }
        Object.keys(emojies).forEach((key) => {
          if (matchEmoji.length > 4) {
            return
          }
          if (key.indexOf(searchEmoji.toLowerCase()) > -1) {
            const emojiItem = {
              value: emojies[key],
              label: key
            }
            if (config.emoji) {
              emojiItem.imageURL = `${config.staticServePath || ''}/emoji/graphics/${key}.png`
            }
            matchEmoji.push(emojiItem)
          }
        })
        genHintsHTML(matchEmoji)
      }
    }
    timerId = debounceInput(timerId, config.change, $editor)
  }).bind('keyup', function (event) {
    config.keyup && config.keyup(event)
  }).bind('keydown', function (event) {
    // at hints action
    const $hints = $editor.find('.b3log-editor__hints')
    if ($hints.length === 0 || $hints.find('li').length === 0 || $hints[0].style.display === 'none') {
      return
    }
    const $currentHint = $hints.find('.b3log-editor__hints--current')
    if (event.keyCode === 40) {
      // down
      event.preventDefault();
      if ($currentHint.next().length === 0) {
        $hints.find('li:eq(0)').addClass('b3log-editor__hints--current')
      } else {
        $currentHint.next().addClass('b3log-editor__hints--current')
      }
      $currentHint.removeClass('b3log-editor__hints--current')
    } else if (event.keyCode === 38) {
      // up
      event.preventDefault();
      if ($currentHint.prev().length === 0) {
        $hints.find('li').last().addClass('b3log-editor__hints--current')
      } else {
        $currentHint.prev().addClass('b3log-editor__hints--current')
      }
      $currentHint.removeClass('b3log-editor__hints--current')
    } else if (event.keyCode === 13) {
      // enter
      event.preventDefault()
      $hints.hide()
      $currentHint.removeClass('b3log-editor__hints--current')

      let splitChar = ':'
      if ($currentHint.data('value').indexOf('@') === 0) {
        splitChar = '@'
      }

      if (/firefox/i.test(navigator.userAgent)) {
        const valueArray = this.value.substr(0, this.selectionStart).split(splitChar)
        valueArray.pop()
        this.value = valueArray.join(splitChar) + $currentHint.data('value') + this.value.substr(this.selectionStart)
        this.selectionEnd = this.selectionStart = (valueArray.join(splitChar) + $currentHint.data('value')).length
        timerId = debounceInput(timerId, config.change, $editor)
        return;
      }
      while (!this.value.substr(0, this.selectionEnd).endsWith(splitChar) &&
      this.value.substr(0, this.selectionEnd) !== '') {
        document.execCommand('delete', false)
      }
      document.execCommand('delete', false)
      document.execCommand('insertText', false, $currentHint.data('value'))
    }
  }).bind('keydown', 'esc', function (event) {
    config.esc && config.esc(event);
    return false;
  }).bind('keydown', 'Ctrl+/', function (event) {
    event.preventDefault();
    event.stopPropagation();
    $emoji.toggle()
    return false;
  }).bind('keydown', 'Ctrl+return', function (event) {
    event.preventDefault();
    event.stopPropagation();
    config.ctrlEnter();
    return false;
  }).bind('keydown', 'Ctrl+b', function (event) {
    event.preventDefault();
    event.stopPropagation();
    $editor.find('.b3log-editor__toolbar > span:eq(1)').click()
    return false;
  }).bind('keydown', 'Ctrl+i', function (event) {
    event.preventDefault();
    event.stopPropagation();
    $editor.find('.b3log-editor__toolbar > span:eq(2)').click()
    return false;
  }).bind('keydown', 'Ctrl+e', function (event) {
    event.preventDefault();
    event.stopPropagation();
    $editor.find('.b3log-editor__toolbar > span:eq(3)').click()
    return false;
  }).bind('keydown', 'Ctrl+k', function (event) {
    event.preventDefault();
    event.stopPropagation();
    $editor.find('.b3log-editor__toolbar > span:eq(4)').click()
    return false;
  }).bind('keydown', 'Ctrl+l', function (event) {
    event.preventDefault();
    event.stopPropagation();
    $editor.find('.b3log-editor__toolbar > span:eq(6)').click()
    return false;
  }).bind('keydown', 'Ctrl+d', function (event) {
    event.preventDefault();
    event.stopPropagation();
    $editor.find('.b3log-editor__toolbar > span:eq(8)').click()
    return false;
  }).bind('keydown', 'Ctrl+Shift+l', function (event) {
    event.preventDefault();
    event.stopPropagation();
    $editor.find('.b3log-editor__toolbar > span:eq(7)').click()
    return false;
  }).bind('keydown', 'Ctrl+Shift+a', function (event) {
    event.preventDefault();
    event.stopPropagation();
    $editor.find('.b3log-editor__toolbar > span:eq(9)').click()
    return false;
  })

  if (typeof Audio.init !== 'undefined') {
    $textarea.bind('keydown', 'Alt+s', function (event) {
      event.preventDefault();
      event.stopPropagation();
      if (!Audio.availabel) {
        Audio.init(function () {
          Audio.handleStartRecording();
          insertTextAtCaret(textarea, '\n[Start Recording]\n', '');
          if (/firefox/i.test(navigator.userAgent)) {
            timerId = debounceInput(timerId, config.change, $editor)
          }
        })
      } else {
        Audio.handleStartRecording();
      }
      return false;
    }).bind('keydown', 'Alt+r', function (event) {
      event.preventDefault();
      event.stopPropagation();

      if (!Audio.availabel) {
        return;
      }

      Audio.handleStopRecording();

      replaceTextareaValue(textarea, '\n[Start Recording]\n', '\n[End Recording, Start Uploading]\n');
      if (/firefox/i.test(navigator.userAgent)) {
        timerId = debounceInput(timerId, config.change, $editor)
      }

      ajaxUpload(config.uploadURL, [Audio.wavFileBlob.getDataBlob()], config.uploadMax, (response) => {
        if (response.data.errFiles.length > 0) {
          replaceTextareaValue(textarea, '\n[End Recording, Start Uploading]\n', `\n[Record Upload Error]\n`)
          if (/firefox/i.test(navigator.userAgent)) {
            timerId = debounceInput(timerId, config.change, $editor)
          }
        } else if (response.data.succMap) {
          replaceTextareaValue(textarea, '\n[End Recording, Start Uploading]\n',
            `\n<audio controls="controls" src="${response.data.succMap.blob}">\n`)
          if (/firefox/i.test(navigator.userAgent)) {
            timerId = debounceInput(timerId, config.change, $editor)
          }
        }
      }, (response) => {
        response && alert(response.msg)
      })
      return false;
    })
  }

  return $textarea
}