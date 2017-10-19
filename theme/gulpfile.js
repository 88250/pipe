/**
 * @file themes tool.
 *
 * @author <a href='http://vanessa.b3log.org'>Liyuan Li</a>
 * @version 0.3.0.0, Oct 19, 2017
 */

const gulp = require('gulp')
const sass = require('gulp-sass')
const rename = require('gulp-rename')
const composer = require('gulp-uglify/composer')
const uglifyjs = require('uglify-es')
const gulpUtil = require('gulp-util')
const fs = require('fs')
const browserify = require('browserify')
const livereload = require('gulp-livereload')
const source = require('vinyl-source-stream')
const buffer = require('vinyl-buffer')
const sourcemaps = require('gulp-sourcemaps')

gulp.task('sass', function () {
  return gulp.src('./x/*/css/*.scss')
    .pipe(sass({outputStyle: 'compressed'}).on('error', sass.logError))
    .pipe(gulp.dest('./x'))
})

gulp.task('dev', function (cb) {
  const file = gulpUtil.env.theme || 'gina'
  return browserify({entries: [`./x/${file}/js/common.js`, `./x/${file}/js/article.js`], debug: true})
    .transform('babelify', {presets: ['es2015']})
    .bundle()
    .pipe(source())
    .pipe(buffer())
    .pipe(sourcemaps.init())
    .pipe(sourcemaps.write('.'))
    .pipe(gulp.dest(`./x/${file}/js/`))
    .pipe(livereload())
})

gulp.task('watch', function () {
  livereload.listen()
  gulp.watch(['./x/*/js/article.js', './x/*/js/common.js', './x/*/js/symbol.js'], ['dev'])
  gulp.watch('./x/*/css/*.scss', ['sass'])
})

gulp.task('build', function () {
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

  const minify = composer(uglifyjs)
  // min sw.js
  gulp.src('./sw.js')
    .pipe(minify().on('error', gulpUtil.log))
    .pipe(rename({suffix: '.min'}))
    .pipe(gulp.dest('.'))

  // theme js
  fs.readdirSync('./x').forEach(function (file) {
    const jsPath = `./x/${file}/js/`
    browserify({entries: [`./x/${file}/js/common.js`, `./x/${file}/js/article.js`]})
      .transform('babelify', {presets: ['es2015']})
      .bundle()
      .pipe(source('common.min.js'))
      .pipe(buffer())
      .pipe(minify().on('error', gulpUtil.log))
      .pipe(gulp.dest(jsPath))
      .pipe(livereload())
  })
})

gulp.task('default', ['sass', 'build'])