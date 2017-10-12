/**
 * @file themes tool.
 *
 * @author <a href='http://vanessa.b3log.org'>Liyuan Li</a>
 * @version 0.2.0.0, Oct 12, 2017
 */

const gulp = require('gulp')
const concat = require('gulp-concat')
const sass = require('gulp-sass')
const clean = require('gulp-clean')
const rename = require('gulp-rename')
const composer = require('gulp-uglify/composer')
const uglifyjs = require('uglify-es')
const gulpUtil = require('gulp-util')
const fs = require('fs')

gulp.task('sass', function () {
  return gulp.src('./x/*/css/*.scss')
    .pipe(sass({ outputStyle: 'compressed' }).on('error', sass.logError))
    .pipe(gulp.dest('.'))
})

gulp.task('sass:watch', function () {
  gulp.watch('./x/*/css/*.scss', ['sass'])
})

gulp.task('clean', ['sass'], function () {
  // set static version
  const newVersion = (new Date()).getTime()
  // set sw.js
  fs.writeFileSync('./sw.js',
    fs.readFileSync('./sw.js', 'UTF-8')
      .replace(/const version = '\d{13}'/, `const version = '${newVersion}'`), 'UTF-8')
  // set solo.json
  fs.writeFileSync('../solo.json',
    fs.readFileSync('../solo.json', 'UTF-8')
      .replace(/"StaticResourceVersion": "\d{13}"/, `"StaticResourceVersion": "${newVersion}"`), 'UTF-8')
  // remove min js
  return gulp.src(['./x/*/js/*.min.js', './sw.min.js'], { read: false })
    .pipe(clean())
})

gulp.task('build', ['sass', 'clean'], function (cb) {
  const minify = composer(uglifyjs)
  // min sw.js
  gulp.src('./sw.js')
    .pipe(minify().on('error', gulpUtil.log))
    .pipe(rename({ suffix: '.min' }))
    .pipe(gulp.dest('.'))
  // concat js
  const eachThemeJS = ['./theme/js/jquery.min.js',
    './theme/js/common.js']

  gulp.src(eachThemeJS)
    .pipe(minify().on('error', gulpUtil.log))
    .pipe(concat('jquery.fileupload.min.js'))
    .pipe(gulp.dest('./src/main/webapp/'))
  // editor js
})
gulp.task('default', ['sass', 'clean', 'build'])
