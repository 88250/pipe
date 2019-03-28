const fs = require('fs')
const path = require('path')

const readDirSync = (root) => {
  const pathList = fs.readdirSync(root)
  pathList.forEach(function (file) {
    const info = fs.statSync(root + '/' + file)
    if (info.isDirectory()) {
      readDirSync(root + '/' + file)
    } else {
      if (path.extname(file) === '.js' || path.extname(file) === '.html' || path.extname(file) === '.css') {
        fs.writeFileSync(`${root}/${file}.tpl`,
          fs.readFileSync(`${root}/${file}`))
        console.log(`gen tpl: ${root}/${file}.tpl`)
      }
    }
  })
}
readDirSync('./dist')
