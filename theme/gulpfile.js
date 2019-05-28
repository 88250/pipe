/**
 * @file themes tool.
 *
 * @author <a href='http://vanessa.b3log.org'>Liyuan Li</a>
 * @version 1.0.0.1, May 28, 2019
 */

const gulp = require('gulp')
const sass = require('gulp-sass')
const rename = require('gulp-rename')
var uglify = require('gulp-uglify')
const fs = require('fs')
const browserify = require('browserify')
const source = require('vinyl-source-stream')
const buffer = require('vinyl-buffer')

function themeSassProcess () {
  return gulp.src('./x/*/css/*.scss').
    pipe(sass({outputStyle: 'compressed'}).on('error', sass.logError)).
    pipe(gulp.dest('./x'))
}

function themeSassProcessWatch () {
  gulp.watch('./x/*/css/*.scss', themeSassProcess)
}

function baseSassProcess () {
  return gulp.src('./scss/*.scss').
    pipe(sass({outputStyle: 'compressed'}).on('error', sass.logError)).
    pipe(gulp.dest('./scss'))
}

function baseSassProcessWatch () {
  gulp.watch('./scss/*.scss', baseSassProcess)
}

function devJSProcessWatch () {
  const watcher = gulp.watch([
    './x/*/js/common.js'])

  watcher.on('change', function (entry) {
    browserify({entries: [entry]}).
      transform('babelify', {presets: ['@babel/preset-env']}).
      bundle().
      on('error', function (err) { console.error(err) }).
      pipe(source(entry)).
      pipe(rename({suffix: '.min'})).
      pipe(buffer()).
      pipe(gulp.dest('.'))
    console.log('[Packed] ' + entry)
  })
}

gulp.task('watch',
  gulp.parallel(themeSassProcessWatch, baseSassProcessWatch, devJSProcessWatch))

function updateVersion () {
  // set static version
  const newVersion = (new Date()).getTime()

  // set pipe.json
  fs.writeFileSync('../pipe.json',
    fs.readFileSync('../pipe.json', 'UTF-8').
      replace(/"StaticResourceVersion": "\d{13}"/,
        `"StaticResourceVersion": "${newVersion}"`), 'UTF-8')

  // min sw.min.js.tpl
  browserify({entries: `./sw.js`}).
    transform('babelify', {presets: ['@babel/preset-env']}).
    bundle().
    on('error', function (err) { console.error(err) }).
    pipe(source('sw.min.js.tpl')).
    pipe(buffer()).
    pipe(uglify()).
    pipe(gulp.dest('.'))

  // min sw.min.js
  return browserify({entries: `./sw.js`}).
    transform('babelify', {presets: ['@babel/preset-env']}).
    bundle().
    on('error', function (err) { console.error(err) }).
    pipe(source('sw.min.js')).
    pipe(buffer()).
    pipe(uglify()).
    pipe(gulp.dest('.'))
}

function minThemeJS () {
  fs.readdirSync('./x').forEach(function (file) {
    const jsPath = `./x/${file}/js/`
    try {
      fs.statSync(`${jsPath}/common.js`)
      browserify({entries: [`${jsPath}/common.js`]}).
        transform('babelify', {presets: ['@babel/preset-env']}).
        bundle().
        on('error', function (err) { console.error(err) }).
        pipe(source(`${jsPath}/common.js`)).
        pipe(rename({suffix: '.min'})).
        pipe(buffer()).
        pipe(uglify()).
        pipe(gulp.dest('.'))
    } catch (e) {
    }
  })
  // fixed gulp task return error
  return gulp.src('./sw.js');
}

gulp.task('default',
  gulp.series(updateVersion,
    gulp.parallel(themeSassProcess, baseSassProcess, minThemeJS)))