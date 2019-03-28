const fs = require('fs')
const path = require(path)
fs.readdirSync('./dist').forEach(function (file) {
  if (path.extname(file) === 'js') {
    fs.writeFileSync(`${jsPath}/common.min.js.tpl`,
      fs.readFileSync(`${jsPath}/common.min.js`))
  }

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
