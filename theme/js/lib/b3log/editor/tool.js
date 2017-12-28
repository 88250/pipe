/*
 * Symphony - A modern community (forum/SNS/blog) platform written in Java.
 * Copyright (C) 2012-2017,  b3log.org & hacpai.com
 *
 * 本文件属于 Sym 商业版的一部分，请仔细阅读项目根文件夹的 LICENSE 并严格遵守相关约定
 */
/**
 * @fileOverview editor tool
 *
 * @author <a href="http://vanessa.b3log.org">Liyuan Li</a>
 * @version 1.0.1.0, Dec 21, 2017
 * @since 2.2.0
 */

export const insertTextAtCaret = (textarea, prefix, suffix, replace) => {
  if (typeof textarea.selectionStart === 'number' && typeof textarea.selectionEnd === 'number') {
    const startPos = textarea.selectionStart
    const endPos = textarea.selectionEnd
    const tmpStr = textarea.value
    if (startPos === endPos) {
      // no selection
      textarea.value = tmpStr.substring(0, startPos) + prefix + suffix + tmpStr.substring(endPos, tmpStr.length)
      textarea.selectionStart = startPos + prefix.length
      textarea.selectionEnd = startPos + (prefix.length)
    } else {
      if (replace) {
        textarea.value = tmpStr.substring(0, startPos) + prefix +
          suffix + tmpStr.substring(endPos, tmpStr.length)
        textarea.selectionStart = textarea.selectionEnd = startPos + (endPos - startPos + prefix.length)
      } else {
        if (tmpStr.substring(startPos - prefix.length, startPos) === prefix &&
          tmpStr.substring(endPos, endPos + suffix.length) === suffix) {
          // broke circle, avoid repeat
          textarea.value = tmpStr.substring(0, startPos - prefix.length) +
            tmpStr.substring(startPos, endPos) + tmpStr.substring(endPos + suffix.length, tmpStr.length)
          textarea.selectionStart = startPos - prefix.length
          textarea.selectionEnd = endPos - prefix.length
        } else {
          // insert
          textarea.value = tmpStr.substring(0, startPos) + prefix + tmpStr.substring(startPos, endPos) +
            suffix + tmpStr.substring(endPos, tmpStr.length)
          textarea.selectionStart = startPos + prefix.length
          textarea.selectionEnd = startPos + (endPos - startPos + prefix.length)
        }
      }
    }
  }
  textarea.focus()
}

export const ajaxUpload = (url, files, uploadMax = 5, succCB, errorCB) => {
  const formData = new FormData()
  for (let iMax = files.length, i = 0; i < iMax; i++) {
    if (files[i].size <= 1024 * 1024 * uploadMax) {
      formData.append('file[]', files[i])
    } else if (files.length === 1) {
      errorCB && errorCB()
      return
    }
  }
  const xhr = new XMLHttpRequest()
  xhr.open('POST', url)
  xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE) {
      if (xhr.status === 200) {
        succCB(JSON.parse(xhr.responseText))
      } else {
        errorCB && errorCB(JSON.parse(xhr.responseText))
      }
    }
  }
  xhr.send(formData)
}

export const genUploading = (files, uploadMax = 5, loadingLabel = 'Uploading', overLabel = 'Over') => {
  let uploadingStr = ''
  for (let iMax = files.length, i = 0; i < iMax; i++) {
    const tag = files[i].type.indexOf('image') === -1 ? '' : '!'
    if (files[i].size > 1024 * 1024 * uploadMax) {
      uploadingStr += `\n${tag}[${files[i].name.replace(/\W/g, '')}](${overLabel} ${uploadMax}MB)\n`
    } else {
      uploadingStr += `\n${tag}[${files[i].name.replace(/\W/g, '')}](${loadingLabel})\n`
    }
  }
  return uploadingStr
}

export const genUploaded = (response, text, loadingLabel = 'Uploading', errorLabel = 'Error') => {
  response.errFiles.forEach((data) => {
    text = text.replace(`[${data.replace(/\W/g, '')}](${loadingLabel})\n`,
      `[${data.replace(/\W/g, '')}](${errorLabel})\n`)
  })

  Object.keys(response.succMap).forEach((key) => {
    text = text.replace(`[${key.replace(/\W/g, '')}](${loadingLabel})\n`,
      `[${key.replace(/\W/g, '')}](${response.succMap[key]})\n`)
  })
  return text
}

export const debounceInput = (timerId, configChange, $editor) => {
  const debounce = 500
  if (timerId !== undefined) {
    clearTimeout(timerId)
  }
  timerId = undefined
  timerId = setTimeout(() => {
    configChange && configChange($editor.find('textarea').val(),
      $editor.find('.b3log-editor__icon--current').length === 0 ? undefined : $editor.find('.b3log-editor__markdown'))
  }, debounce)
}