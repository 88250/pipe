# <img src="https://user-images.githubusercontent.com/873584/33324159-c3ea5050-d489-11e7-9f4b-75ee806a7538.png"> [Pipe](https://github.com/b3log/pipe) 

* [简介](#简介)
* [特性](#特性)
* [文档](#文档)
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
* 后台使用 Vue.js 进行前后端分离

### 案例

* http://pipe.b3log.org

你也在使用的话欢迎通过 PR 更新该列表。

## 特性

* 多用户博客平台
* Markdown / Emoji
* 聚合分类 / 标签
* 自定义导航
* 多主题 / 多语言
* Atom 订阅
* 可配置动静分离
* 搜索
* Hexo/Jekyll 导入 / 导出

## 文档

* 用户指南 TBD
* 开发指南 TBD
* [主题开发指南](https://hacpai.com/article/1512550354920)

## 构建

### 前提

需要预先安装好如下编译环境，请尽量都使用最新版：

1. [Go](https://golang.org)
2. [Node.js](https://nodejs.org)

### 编译后端

```
go build -i -v
```

### 编译管理后台前端

进入 console 目录，然后：

```
npm install && npm run build
```

* 开发环境执行 `npm run dev`
* pipe.json 中 `AxiosBaseURL` 为 `/api` 时需启动 `./pipe`，为 `/mock` 时需运行 `npm run mock`

### 打包前台主题

进入 theme 目录，然后：

```
npm install && npm install --global gulp && gulp
```

* 开发环境执行 `gulp watch`
* theme/js 和 theme/scss 下为基础方法和样式，可按需引入使用
* 主题开发请参考 theme/x/Gina

## 贡献

### 作者

Pipe 的主要作者是 [Daniel](https://github.com/88250) 与 [Vanessa](https://github.com/Vanessa219)，所有贡献者可以在[这里](https://github.com/b3log/pipe/graphs/contributors)看到。

我们非常期待你加入到这个项目中，无论是使用反馈还是代码补丁，都是对 Pipe 一份满满的爱 :heart:

### 讨论区

* 到 Pipe 官方[讨论区](https://hacpai.com/tag/Pipe)发帖（推荐做法）
* 来一发 [issue](https://github.com/b3log/pipe/issues/new)
* 加入 QQ 群 242561391

### 算力

Pipe 默认会通过浏览者的浏览器进行挖矿（只会使用空闲的 CPU 资源，并且占用很低），收益将用于维持项目运维。原理请参考[使用访问者浏览器挖矿的方法](https://hacpai.com/article/1512269880744)。

如果你不方便帮助我们，可以将 common.js、utils.js 中的 `miner` 相关代码注释掉。我们恳请你尽量保留，谢谢。

## 开源协议

Pipe 使用 GPLv3 作为开源授权协议，请尽量遵循，即使是在中国。

## 鸣谢

Pipe 的诞生离不开以下开源项目：

* [jQuery](https://github.com/jquery/jquery)：JavaScript 工具库，用于主题页面
* [Vue.js](https://github.com/vuejs/vue)：渐进式 JavaScript 框架
* [Nuxt.js](https://github.com/nuxt/nuxt.js)：Vue.js 框架
* [Vuetify](https://github.com/vanessa219/vuetify)：Vue.js 的 Material 组件框架
* [Gin](https://github.com/gin-gonic/gin)：又快又好用的 golang HTTP web 框架
* [GORM](https://github.com/jinzhu/gorm)：极好的 golang ORM 库
* [Blackfriday](github.com/russross/blackfriday)：golang Markdown 处理器
* [SQLite](https://www.sqlite.org)：使用广泛的嵌入式 SQL 引擎
* [GCache](github.com/bluele/gcache)：golang 缓存库

----

<p align = "center">
<strong>小而美的博客平台，为未来而构建</strong>
<br><br>
<img src="https://user-images.githubusercontent.com/873584/33324033-441773da-d489-11e7-8d39-78abbeb563f0.png">
</p>

