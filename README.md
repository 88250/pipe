# <img src="https://user-images.githubusercontent.com/873584/33324159-c3ea5050-d489-11e7-9f4b-75ee806a7538.png"> [Pipe](https://github.com/b3log/pipe) 

## 简介

小而美的博客平台，通过[黑客派](https://hacpai.com)账号登录即可使用。

### 动机

产品层面：

* 市面上缺乏支持多独立博客的平台级系统
* 实现 [B3log 构思](https://hacpai.com/b3log)

技术层面：

* 在博客系统这个轮子上充分发挥 golang 的优势
* 使用流行的框架和开发方式，比如 Vue.js，前后端分离

作者个人层面：

* 之前的产品 [Solo](https://github.com/b3log/solo) 在一些设计上不够理想
* 希望这是[我](https://github.com/88250)和 [V](https://github.com/vanessa219) 最后一次造博客轮子了

### 案例

TBD

## 特性

- [X] 多用户博客平台
- [X] Markdown / Emoji
- [X] [聚合分类](https://github.com/b3log/solo/issues/12256) / 标签
- [X] 自定义导航
- [X] 多主题 / 多语言
- [X] Atom 订阅
- [X] 可配置动静分离
- [ ] 全文搜索
- [ ] Sitemap
- [ ] Hexo/Jekyll 导入

## 文档

* 用户指南 TBD
* 开发指南 TBD
* 主题开发指南 TBD

## 构建

### 前提

需要预先安装好如下编译环境，请尽量都使用最新版：

1. [Go](https://golang.org)
2. [Node.js](https://nodejs.org)

### 编译后端

```
go build
```

### 编译前端（管理后台）

```
cd console && npm install && npm run build
```

console/config/env.json 中 `clientBaseURL` 为 `/api` 时需启动 `./pipe`，为 `/mock` 时需运行 
```
npm run mock
```

### 前台主题

```
cd theme && npm install && npm install --global gulp && gulp build
```

theme/js 和 theme/scss 下为基础方法和样式，可按需引入使用。主题开发可参照 theme/x/Gina。

## 社区

* 到 Pipe 官方[讨论区](https://hacpai.com/tag/Pipe)发帖（推荐做法）
* 来一发 [issue](https://github.com/b3log/pipe/issues/new)
* 加入 QQ 群 242561391

## 开源协议

Pipe 使用 GPLv3 作为开源授权协议，请尽量遵循，即使是在中国。

## 鸣谢

Pipe 的诞生离不开以下开源项目：

* [Vue.js](https://github.com/vuejs/vue)：渐进式 JavaScript 框架
* [Vuetify](https://github.com/vuetifyjs/vuetify)：Vue.js 的 Material 组件框架
* [jQuery](https://github.com/jquery/jquery)：使用广泛的 JavaScript 工具库
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

