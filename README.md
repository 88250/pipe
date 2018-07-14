# <img src="https://user-images.githubusercontent.com/873584/33324159-c3ea5050-d489-11e7-9f4b-75ee806a7538.png"> [Pipe](https://github.com/b3log/pipe) [![Build Status](https://img.shields.io/travis/b3log/pipe.svg?style=flat)](https://travis-ci.org/b3log/pipe) [![Go Report Card](https://goreportcard.com/badge/github.com/b3log/pipe)](https://goreportcard.com/report/github.com/b3log/pipe) [![Coverage Status](https://coveralls.io/repos/github/b3log/pipe/badge.svg?branch=master)](https://coveralls.io/github/b3log/pipe?branch=master) [![GitHub license](https://img.shields.io/github/license/b3log/pipe.svg)](https://github.com/b3log/pipe/blob/master/LICENSE)

<p align="center">
<a href="https://github.com/b3log/pipe/blob/master/README.md"><strong>English</strong></a> | <a href="https://github.com/b3log/pipe/blob/master/README_zh_CN.md"><strong>中文</strong></a>
</p>

* [Introduction](#introduction)
* [Features](#features)
* [Installation](#installation)
* [Documents](#documents)
* [Screenshots](#screenshots)
* [Build](#build)
* [Contributions](#contributions)
* [License](#license)
* [Credits](#credits)

## Introduction

[Pipe](https://github.com/b3log/pipe) is a small and beautiful blogging platform, login via [HacPai](https://hacpai.com) to use.

### Motivation

Product level:

* There is a lack of platform-level systems on the market that support multiple independent blogs
* Implements [B3log Idea](https://hacpai.com/b3log)

Technical level:

* Take full advantage of golang on the wheel of blogging system
* Blog console using Vue.js for frontend-backend separation

### Use cases

* [B3log Pipe](http://pipe.b3log.org)
* [Akkuman 的博客](http://o0o.pub)
* [Vanessa](http://vanessa.b3log.org)
* [Domolo Reader](http://www.domolo.com)
* [图解吧](http://tujie8.net)

You are also welcome to update this list through PR.

## Features

* Multi-user blog platform
* Markdown / Emoji
* Aggregate classification / tag
* Custom navigation
* Themes / I18n
* Atom subscription
* Search
* Hexo/Jekyll Import/Export
* Configurable static separation
* Supports SQLite / MySQL

## Installation

[Download](https://pan.baidu.com/s/1jHPtHLO) the latest release package, enter the decompressed directory and run the pipe/pipe.exe executable file.

**For more details, refer to the [Pipe User Guide](https://hacpai.com/article/1513761942333). In addition, if you do not want to maintain the server yourself, you can directly use the [Pipe Service](http://pipe.b3log.org) of our operation and maintenance (for domain name binding, please contact QQ845765).**

## Documents

* [Pipe User Guide](https://hacpai.com/article/1513761942333)
* [Pipe Theme Dev Guide](https://hacpai.com/article/1512550354920)
* [Postman Collection](https://www.getpostman.com/collections/c466e81beb7acd5685ec)

## Screenshots

### Init

![Init](https://user-images.githubusercontent.com/873584/34195698-e860c0c4-e599-11e7-9d4f-32307712324d.jpg)

### Console

![Console](https://user-images.githubusercontent.com/873584/34195907-b390adf4-e59a-11e7-8ef7-97f8393c770d.jpg)

### Edit

![Edit](https://user-images.githubusercontent.com/873584/34195873-975c07dc-e59a-11e7-83ca-c07272c5933c.jpg)

### Theme

![Theme](https://user-images.githubusercontent.com/873584/34195948-d2b0106c-e59a-11e7-922d-b85e7a172eef.jpg)

## Build

Need to pre-install the following compilation environment, please use the latest version as far as possible:

1. [Go](https://golang.org)
2. [Node.js](https://nodejs.org)

Build.sh can be automatically built on Linux, and other platforms can be manually built using the following steps.

### Compile the server

```
go build -i -v
```

### Compile console frond-end

Enter console directory then execute:

```
npm install && npm run build
```

* The development environment executes `npm run dev` and the access port is :3000
* In pipe.json `./pipe` needs to be started when `AxiosBaseURL` is `/api`, `npm run mock` is required when `/mock` is specified

### Pack the theme

Go to the theme directory and then:

```
npm install && npm install --global gulp && gulp
```

* Development environment implementation `gulp watch --theme=ThemeName`
* Basic methods and styles under theme/js and theme/scss can be introduced as needed
* Theme development please refer to theme/x/Gina

## Contributions

### Authors

The main authors of Pipe are [Daniel](https://github.com/88250) and [Vanessa](https://github.com/Vanessa219), and all contributors can be found [here](https://github.com/b3log/pipe/graphs/contributors).

We are very much looking forward to you joining this project. Whether it is using feedback or code patches, it is a full of love for Pipe :heart:

### Discussion

* [Official forum](https://hacpai.com/tag/Pipe)
* [Issue](https://github.com/b3log/pipe/issues/new)

## License

Pipe uses GPLv3 as an open source license, please follow it as much as possible, even in China.

## Credits

The birth of Pipe is inseparable from the following items:

* [jQuery](https://github.com/jquery/jquery): JavaScript library for themes
* [Vue.js](https://github.com/vuejs/vue): Progressive JavaScript framework
* [Nuxt.js](https://github.com/nuxt/nuxt.js): Vue.js framewok
* [Vuetify](https://github.com/vanessa219/vuetify): Vue.js Material component
* [Gin](https://github.com/gin-gonic/gin): golang HTTP web framework
* [GORM](https://github.com/jinzhu/gorm): fantastic golang ORM library
* [Blackfriday](github.com/russross/blackfriday): golang Markdown processor
* [SQLite](https://www.sqlite.org): the most used database enging in the world
* [GCache](https://github.com/bluele/gcache): golang cache
* [GoLand](https://www.jetbrains.com/go): a wonderful IDE

----

<p align = "center">
<strong>A small and beautiful blogging platform, build for the future</strong>
<br><br>
<img src="https://user-images.githubusercontent.com/873584/33324033-441773da-d489-11e7-8d39-78abbeb563f0.png">
</p>
