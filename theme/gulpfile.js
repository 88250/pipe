/**
 * @file themes tool.
 *
 * @author <a href='http://vanessa.b3log.org'>Liyuan Li</a>
 * @version 0.1.0.0, Oct 11, 2017
 */

const gulp = require('gulp')
const concat = require('gulp-concat')
const cleanCSS = require('gulp-clean-css')
const uglify = require('gulp-uglify')
const sass = require('gulp-sass')
const clean = require('gulp-clean')
const rename = require('gulp-rename')
const composer = require('gulp-uglify/composer')
const uglifyjs = require('uglify-es')
const gulpUtil = require('gulp-util')
const fs = require('fs')

gulp.task('sass', function () {
  return gulp.src('./*/css/*.scss')
    .pipe(sass({ outputStyle: 'compressed' }).on('error', sass.logError))
    .pipe(gulp.dest('.'))
})

gulp.task('sass:watch', function () {
  gulp.watch('./*/css/*.scss', ['sass'])
})

gulp.task('clean', ['sass'], function () {
  // set static version
  const swFile = fs.readFileSync('./src/main/webapp/sw.js', 'UTF-8')
  const oldVersion = /const version = '(\d{13})'/.exec(swFile)[1]
  const newVersion = (new Date()).getTime()
  // set sw.js
  fs.writeFileSync('./src/main/webapp/sw.js',
    swFile.replace(`const version = ${oldVersion}, const version = ${newVersion}`), 'UTF-8')
  // set latke.properties
  fs.writeFileSync('./src/main/resources/latke.properties',
    fs.readFileSync('./src/main/resources/latke.properties', 'UTF-8')
      .replace('staticResourceVersion=' + oldVersion, 'staticResourceVersion=' + newVersion), 'UTF-8')

  // remove min js
  return gulp.src(['./src/main/webapp/js/*.min.js', './src/main/webapp/sw.min.js'], { read: false })
    .pipe(clean())
})

gulp.task('build', ['sass', 'clean'], function (cb) {
  const minify = composer(uglifyjs)

  // min css
  gulp.src('./src/main/webapp/js/lib/editor/codemirror.css')
    .pipe(cleanCSS())
    .pipe(concat('codemirror.min.css'))
    .pipe(gulp.dest('./src/main/webapp/js/lib/editor/'))

  // min js
  gulp.src('./src/main/webapp/js/*.js')
    .pipe(uglify().on('error', gulpUtil.log))
    .pipe(rename({ suffix: '.min' }))
    .pipe(gulp.dest('./src/main/webapp/js/'))

  // min sw.js
  gulp.src('./src/main/webapp/sw.js')
    .pipe(minify().on('error', gulpUtil.log))
    .pipe(rename({ suffix: '.min' }))
    .pipe(gulp.dest('./src/main/webapp/'))

  // concat js
  const jsJqueryUpload = ['./src/main/webapp/js/lib/jquery/file-upload-9.10.1/vendor/jquery.ui.widget.js',
    './src/main/webapp/js/lib/jquery/file-upload-9.10.1/jquery.iframe-transport.js',
    './src/main/webapp/js/lib/jquery/file-upload-9.10.1/jquery.fileupload.js',
    './src/main/webapp/js/lib/jquery/file-upload-9.10.1/jquery.fileupload-process.js',
    './src/main/webapp/js/lib/jquery/file-upload-9.10.1/jquery.fileupload-validate.js']
  gulp.src(jsJqueryUpload)
    .pipe(uglify())
    .pipe(concat('jquery.fileupload.min.js'))
    .pipe(gulp.dest('./src/main/webapp/js/lib/jquery/file-upload-9.10.1/'))

  const jsCodemirror = ['./src/main/webapp/js/lib/editor/diff_match_patch.js',
    './src/main/webapp/js/lib/editor/codemirror.js',
    './src/main/webapp/js/lib/editor/placeholder.js',
    './src/main/webapp/js/lib/editor/merge.js',
    './src/main/webapp/js/overwrite/codemirror/addon/hint/show-hint.js',
    './src/main/webapp/js/lib/editor/editor.js',
    './src/main/webapp/js/lib/to-markdown.js']
  gulp.src(jsCodemirror)
    .pipe(uglify())
    .pipe(concat('codemirror.min.js'))
    .pipe(gulp.dest('./src/main/webapp/js/lib/editor/'))

  const jsCommonLib = ['./src/main/webapp/js/lib/jquery/jquery-3.1.0.min.js',
    './src/main/webapp/js/lib/md5.js',
    './src/main/webapp/js/lib/reconnecting-websocket.min.js',
    './src/main/webapp/js/lib/jquery/jquery.bowknot.min.js',
    './src/main/webapp/js/lib/ua-parser.min.js',
    './src/main/webapp/js/lib/jquery/jquery.hotkeys.js',
    './src/main/webapp/js/lib/jquery/jquery.pjax.js',
    './src/main/webapp/js/lib/nprogress/nprogress.js']
  gulp.src(jsCommonLib)
    .pipe(uglify())
    .pipe(concat('libs.min.js'))
    .pipe(gulp.dest('./src/main/webapp/js/lib/compress/'))

  const jsArticleLib = [
    // start codemirror.min.js
    './src/main/webapp/js/lib/editor/diff_match_patch.js',
    './src/main/webapp/js/lib/editor/codemirror.js',
    './src/main/webapp/js/lib/editor/placeholder.js',
    './src/main/webapp/js/lib/editor/merge.js',
    './src/main/webapp/js/overwrite/codemirror/addon/hint/show-hint.js',
    './src/main/webapp/js/lib/editor/editor.js',
    './src/main/webapp/js/lib/to-markdown.js',
    // end codemirror.min.js
    './src/main/webapp/js/lib/highlight.js-9.6.0/highlight.pack.js',
    // start jquery.fileupload.min.js
    './src/main/webapp/js/lib/jquery/file-upload-9.10.1/vendor/jquery.ui.widget.js',
    './src/main/webapp/js/lib/jquery/file-upload-9.10.1/jquery.iframe-transport.js',
    './src/main/webapp/js/lib/jquery/file-upload-9.10.1/jquery.fileupload.js',
    './src/main/webapp/js/lib/jquery/file-upload-9.10.1/jquery.fileupload-process.js',
    './src/main/webapp/js/lib/jquery/file-upload-9.10.1/jquery.fileupload-validate.js',
    // end jquery.fileupload.min.js
    './src/main/webapp/js/lib/sound-recorder/SoundRecorder.js',
    './src/main/webapp/js/lib/jquery/jquery.qrcode.min.js',
    './src/main/webapp/js/lib/aplayer/APlayer.min.js']
  gulp.src(jsArticleLib)
    .pipe(uglify())
    .pipe(concat('article-libs.min.js'))
    .pipe(gulp.dest('./src/main/webapp/js/lib/compress/'))
})
gulp.task('default', ['sass', 'clean', 'build'])
