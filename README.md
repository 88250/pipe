# Pipe

## 简介

小而美的 golang 博客平台。

### 动机

* 在博客这个轮子上充分发挥 golang 的优势，易部署、性能好、跨平台、省资源
* [Solo](https://github.com/b3log/solo) 重制版
* 市面上缺乏支持多独立博客的平台级系统
* 实现 [B3log 构思](https://hacpai.com/b3log)

### 案例

TBD

## 特性

* Markdown / Emoji
* [聚合分类](https://github.com/b3log/solo/issues/12256) / 标签
* 自定义导航（页面、链接）
* 多主题 / 多语言
* 内置云存储
* 多博客、团队博客
* SQL、JSON、MD 格式导出
* [Hexo/Jekyll 导入](https://hacpai.com/article/1498490209748)
* CDN 静态资源分离

总得来说就是去除了 [Solo](https://github.com/b3log/solo) 中部分鸡肋的功能，并对一些特性进行了改进。另外最重要的一点是 Pipe 是一个支持多博客的平台，超级管理员可以创建多个博客，每个博客都是独立的团队博客。

## 文档

* 用户指南 数据浏览可以使用 SQLite 的图形化工具 [SQLiteBrowser](http://sqlitebrowser.org)。
* 开发指南
* 主题开发指南

## 构建

```
go build
```

## 技术栈

### 管理后台

Vue.js, Nuxt, Vuetify

安装依赖
```
cd console && npm install
```
console/config/env.json 中 `clientBaseURL` 为 `/api` 时需启动 `./pipe`，为 `/mock` 时需运行 
```
npm run mock
```
开发
```
npm run dev
```

打包
```
npm run build
```

### 服务器端

* Web 层使用 [Gin](https://github.com/gin-gonic/gin) 框架
* 持久层使用 [GORM](https://github.com/jinzhu/gorm) ORM 库

### 主题

jQuery

TBD

## 作者

Pipe 的主要作者是 [Daniel](https://github.com/88250) 与 [Vanessa](https://github.com/Vanessa219)，所有贡献者可以在[这里](https://github.com/b3log/pipe/graphs/contributors)看到。

我们非常期待你加入到这个项目中，无论是使用反馈还是代码补丁，都是对 Pipe 一份满满的爱 :heart:

## 社区

* 到 Pipe 官方[讨论区](https://hacpai.com/tag/Pipe)发帖（推荐做法）
* 来一发 [issue](https://github.com/b3log/pipe/issues/new)
* 加入 Pipe 社区 Q 群 242561391

## 开源协议

Pipe 使用 GPLv3 作为开源授权协议，请尽量遵循，即使是在中国。

## 鸣谢

Pipe 的诞生离不开以下开源项目：

* [Vue.js](https://github.com/vuejs/vue)：渐进式 JavaScript 框架
* [jQuery](https://github.com/jquery/jquery)：使用广泛的 JavaScript 工具库
* [Gin](https://github.com/gin-gonic/gin)：又快又好用的 golang HTTP web 框架
* [GORM](https://github.com/jinzhu/gorm)：梦幻般的 golang ORM 库
* [Logrus](https://github.com/sirupsen/logrus)：结构化的 golang 日志组件
* [Blackfriday](github.com/russross/blackfriday)：golang Markdown 处理器
* [SQLite](https://www.sqlite.org)：使用广泛的嵌入式 SQL 引擎
* [GCache](github.com/bluele/gcache)：golang 缓存库
* [IntelliJ IDEA](https://www.jetbrains.com/idea)：全宇宙暂时排名第二的 IDE