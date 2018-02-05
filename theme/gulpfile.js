/**
 * @file themes tool.
 *
 * @author <a href='http://vanessa.b3log.org'>Liyuan Li</a>
 * @version 0.1.0.1, Dec 27, 2017
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
const es = require('event-stream')

gulp.task('sass', function () {
  gulp.src('./scss/*.scss')
    .pipe(sass({outputStyle: 'compressed'}).on('error', sass.logError))
    .pipe(gulp.dest('./scss'))

  return gulp.src('./x/*/css/*.scss')
    .pipe(sass({outputStyle: 'compressed'}).on('error', sass.logError))
    .pipe(gulp.dest('./x'))
})

gulp.task('dev', function () {
  const theme = gulpUtil.env.theme || '9IPHP'
  const files = [`./x/${theme}/js/common.js`, `./x/${theme}/js/article.js`]

  const tasks = files.map(function (entry) {
    return browserify({entries: [entry]})
      .transform('babelify', {presets: ['es2015']})
      .bundle()
      .pipe(source(entry))
      .pipe(rename({suffix: '.min'}))
      .pipe(buffer())
      .pipe(sourcemaps.init())
      .pipe(sourcemaps.write('.'))
      .pipe(gulp.dest('.'))
      .pipe(livereload())
  })
  return es.merge.apply(null, tasks)
})

gulp.task('watch', function () {
  livereload.listen()
  gulp.watch(['./js/common.js', './js/article.js', './x/*/js/article.js', './x/*/js/common.js', './x/*/js/symbol.js'], ['dev'])
  gulp.watch(['./x/*/css/*.scss', './scss/*.scss'], ['sass'])
})

gulp.task('build', function () {
  // set static version
  const newVersion = (new Date()).getTime()

  // set pipe.json
  fs.writeFileSync('../pipe.json',
    fs.readFileSync('../pipe.json', 'UTF-8')
      .replace(/"StaticResourceVersion": "\d{13}"/, `"StaticResourceVersion": "${newVersion}"`), 'UTF-8')

  const minify = composer(uglifyjs)
  // min sw.js
  browserify({entries: `./sw.js`})
    .transform('babelify', {presets: ['es2015']})
    .bundle()
    .pipe(source('sw.min.js'))
    .pipe(buffer())
    .pipe(minify().on('error', gulpUtil.log))
    .pipe(gulp.dest('.'))

  // theme js
  fs.readdirSync('./x').forEach(function (file) {
    const jsPath = `./x/${file}/js/`
    try {
      fs.statSync(`${jsPath}/common.js`)
      browserify({entries: `${jsPath}/common.js`})
        .transform('babelify', {presets: ['es2015']})
        .bundle()
        .pipe(source('common.min.js'))
        .pipe(buffer())
        .pipe(minify().on('error', gulpUtil.log))
        .pipe(gulp.dest(jsPath))
      browserify({entries: `${jsPath}/article.js`})
        .transform('babelify', {presets: ['es2015']})
        .bundle()
        .pipe(source('article.min.js'))
        .pipe(buffer())
        .pipe(minify().on('error', gulpUtil.log))
        .pipe(gulp.dest(jsPath))
    } catch (e) {
    }
  })
})

gulp.task('default', ['sass', 'build'])