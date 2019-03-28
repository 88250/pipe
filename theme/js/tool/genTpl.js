const fs = require('fs')
fs.readdirSync('./x').forEach(function (file) {
  const jsPath = `./x/${file}/js/`
  if (file === '.DS_Store') {
    return
  }
  try {
    fs.writeFileSync(`${jsPath}/common.min.js.tpl`,
      fs.readFileSync(`${jsPath}/common.min.js`))
    fs.writeFileSync(`${jsPath}/article.min.js.tpl`,
      fs.readFileSync(`${jsPath}/article.min.js`))
  } catch (e) {
    console.log(e)
  }
})