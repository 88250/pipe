# <img src="https://user-images.githubusercontent.com/873584/33324159-c3ea5050-d489-11e7-9f4b-75ee806a7538.png"> [Pipe](https://github.com/b3log/pipe) [![Build Status](https://img.shields.io/travis/b3log/pipe.svg?style=flat)](https://travis-ci.org/b3log/pipe) [![Go Report Card](https://goreportcard.com/badge/github.com/b3log/pipe)](https://goreportcard.com/report/github.com/b3log/pipe) [![Coverage Status](https://coveralls.io/repos/github/b3log/pipe/badge.svg?branch=master)](https://coveralls.io/github/b3log/pipe?branch=master) [![GitHub license](https://img.shields.io/github/license/b3log/pipe.svg)](https://github.com/b3log/pipe/blob/master/LICENSE)

<p align="center">
<a href="https://github.com/b3log/pipe/blob/master/README.md"><strong>English</strong></a> | <a href="https://github.com/b3log/pipe/blob/master/README_zh_CN.md"><strong>中文</strong></a>
</p>

* [简介](#简介)
* [特性](#特性)
* [安装](#安装)
* [文档](#文档)
* [界面](#界面)
* [构建](#构建)
* [贡献](#贡献)
* [开源协议](#开源协议)
* [鸣谢](#鸣谢)

## 简介

[Pipe](https://github.com/b3log/pipe) 是一款小而美的开源博客平台，通过[黑客派](https://hacpai.com)账号登录即可使用。

### 动机

产品层面：

* 市面上缺乏支持多独立博客的平台级系统
* 实现 [B3log 构思](https://hacpai.com/b3log)

技术层面：

* 在博客系统这个轮子上充分发挥 golang 的优势
* 博客管理后台界面使用 Vue.js 进行前后端分离

### 案例

* [B3log Pipe](http://pipe.b3log.org)
* [Akkuman 的博客](http://o0o.pub)
* [Vanessa](http://vanessa.b3log.org)
* [Domolo Reader](http://www.domolo.com)
* [图解吧](http://tujie8.net)

你也在使用的话欢迎通过 PR 更新该列表。

## 特性

* 多用户博客平台
* Markdown / Emoji
* 聚合分类 / 标签
* 自定义导航
* 多主题 / 多语言
* Atom 订阅
* 搜索
* Hexo/Jekyll 导入 / 导出
* 可配置动静分离
* 支持 SQLite / MySQL

## 安装

[下载](https://pan.baidu.com/s/1jHPtHLO)最新的发布包解压，进入解压目录运行 pipe/pipe.exe 可执行文件即可。

**更多细节请参考 [Pipe 用户指南](https://hacpai.com/article/1513761942333)。另外，如果你不想自己维护服务器，可以直接使用我们运维的 [Pipe 服务](http://pipe.b3log.org)（域名绑定请联系 QQ845765）。**

## 文档

* [用户指南](https://hacpai.com/article/1513761942333)
* [主题开发指南](https://hacpai.com/article/1512550354920)
* [Postman 测试集](https://www.getpostman.com/collections/c466e81beb7acd5685ec)

## 界面

### 初始化

![初始化](https://user-images.githubusercontent.com/873584/34195698-e860c0c4-e599-11e7-9d4f-32307712324d.jpg)

### 管理后台

![管理后台](https://user-images.githubusercontent.com/873584/34195907-b390adf4-e59a-11e7-8ef7-97f8393c770d.jpg)

### 编辑文章

![编辑文章](https://user-images.githubusercontent.com/873584/34195873-975c07dc-e59a-11e7-83ca-c07272c5933c.jpg)

### 默认主题

![默认主题](https://user-images.githubusercontent.com/873584/34195948-d2b0106c-e59a-11e7-922d-b85e7a172eef.jpg)

## 构建

需要预先安装好如下编译环境，请尽量都使用最新版：

1. [Go](https://golang.org)
2. [Node.js](https://nodejs.org)

在 Linux 上可以执行 build.sh 进行自动构建，其他平台可按照下面步骤进行手动构建。

### 编译后端

```
go build -i -v
```

### 编译管理后台前端

进入 console 目录，然后：

```
npm install && npm run build
```

* 开发环境执行 `npm run dev`，访问端口为 :3000
* pipe.json 中 `AxiosBaseURL` 为 `/api` 时需启动 `./pipe`，为 `/mock` 时需运行 `npm run mock`

### 打包前台主题

进入 theme 目录，然后：

```
npm install && npm install --global gulp && gulp
```

* 开发环境执行 `gulp watch --theme=ThemeName`
* theme/js 和 theme/scss 下为基础方法和样式，可按需引入使用
* 主题开发请参考 theme/x/Gina

## 贡献

### 作者

Pipe 的主要作者是 [Daniel](https://github.com/88250) 与 [Vanessa](https://github.com/Vanessa219)，所有贡献者可以在[这里](https://github.com/b3log/pipe/graphs/contributors)看到。

我们非常期待你加入到这个项目中，无论是使用反馈还是代码补丁，都是对 Pipe 一份满满的爱 :heart:

### 讨论区

* 到 Pipe 官方[论坛](https://hacpai.com/tag/Pipe)发帖（推荐做法）
* 来一发 [issue](https://github.com/b3log/pipe/issues/new)

## 开源协议

Pipe 使用 GPLv3 作为开源授权协议，请尽量遵循，即使是在中国。

## 鸣谢

Pipe 的诞生离不开以下项目：

* [jQuery](https://github.com/jquery/jquery)：JavaScript 工具库，用于主题页面
* [Vue.js](https://github.com/vuejs/vue)：渐进式 JavaScript 框架
* [Nuxt.js](https://github.com/nuxt/nuxt.js)：Vue.js 框架
* [Vuetify](https://github.com/vanessa219/vuetify)：Vue.js 的 Material 组件框架
* [Gin](https://github.com/gin-gonic/gin)：又快又好用的 golang HTTP web 框架
* [GORM](https://github.com/jinzhu/gorm)：极好的 golang ORM 库
* [Blackfriday](github.com/russross/blackfriday)：golang Markdown 处理器
* [SQLite](https://www.sqlite.org)：使用广泛的嵌入式 SQL 引擎
* [GCache](https://github.com/bluele/gcache)：golang 缓存库
* [GoLand](https://www.jetbrains.com/go)：全宇宙暂时排名第一的 golang IDE

----

<p align = "center">
<strong>小而美的博客平台，为未来而构建</strong>
<br><br>
<img src="https://user-images.githubusercontent.com/873584/33324033-441773da-d489-11e7-8d39-78abbeb563f0.png">
</p>
