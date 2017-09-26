# Solo.go

## 简介

小而美的 golang 博客平台，[Solo](https://github.com/b3log/solo) 的 golang 版。

你以前没有用过 Solo 也没关系，只需要相信这是史上最好的博客平台！

### 动机

* 很受欢迎的 [Solo](https://github.com/b3log/solo) 已经证明了 Java 实现博客系统的可能性，但无法解决由 Java 带来的问题
* 市面上缺乏支持多独立博客的平台级系统
* 在博客这个轮子上充分发挥 golang 的优势，易部署、性能好、跨平台、省资源
* 争取击败所有已知的博客系统，包括各种静态博客系统

### 案例

TBD

## 特性

* Markdown / Emoji
* [聚合分类](https://github.com/b3log/solo/issues/12256) / 标签
* 自定义导航（页面、链接）
* 多皮肤 / 多语言
* 上传七牛云
* 多博客、团队博客
* SQL、JSON、MD 格式导出
* [Hexo/Jekyll 导入](https://hacpai.com/article/1498490209748)
* CDN 静态资源分离

总得来说就是去除了 Solo 中部分鸡肋的功能，并对一些特性进行了改进。另外最重要的一点是 Solo.go 是一个支持多博客的平台，超级管理员可以创建多个博客，每个博客都是独立的团队博客。

## 文档

数据浏览可以使用 SQLite 的图形化工具 [SQLiteBrowser](http://sqlitebrowser.org)。Solo.go 有内存缓存，**切勿直接修改数据库**。

## 构建

因为目录名是 solo.go ，所以在不带 `-o` 指定输出文件名 build 时在非 Windows 上会有点小尴尬，这样会生成二进制 solo.go，而下次编译的时候就会被编译器当作源码，进而出现如下类似的报错：

```
can't load package: package github.com/b3log/solo.go: read d:\gogogo\src\github.com\b3log\solo.go\solo.go: unexpected NUL in input
```

而 Windows 平台上不带 `-o` 编译后会生成 solo.go.exe，没毛病。总之，强烈建议在非 Windows 平台上构建时使用如下命令：

```
go build -o solo
```

## 技术栈

### 皮肤

jQuery

TBD

### 管理后台

Vue.js

TBD

### 后端框架

* Web 层使用 [Gin](https://github.com/gin-gonic/gin) 框架
* 持久层使用 [GORM](https://github.com/jinzhu/gorm) ORM 库

## 作者

Solo.go 的主要作者是 [Daniel](https://github.com/88250) 与 [Vanessa](https://github.com/Vanessa219)，所有贡献者可以在[这里](https://github.com/b3log/solo.go/graphs/contributors)看到。

我们非常期待你加入到这个项目中，无论是使用反馈还是代码补丁，都是对 Solo.go 一份满满的爱 :heart:

## 社区

* 到 Solo.go 官方[讨论区](https://hacpai.com/tag/Solo.go)发帖（推荐做法）
* 来一发 [issue](https://github.com/b3log/solo.go/issues/new)
* 加入 Solo.go 社区 Q 群 242561391

## 开源协议

Solo.go 使用 GPLv3 作为开源授权协议，请尽量遵循，即使是在中国。

## 鸣谢

Solo.go 的诞生离不开以下开源项目：

* [jQuery](https://github.com/jquery/jquery)：使用最广泛的 JavaScript 工具库
* [Gin](https://github.com/gin-gonic/gin)：又快又好用的 golang HTTP web 框架
* [GORM](https://github.com/jinzhu/gorm)：梦幻般的 golang ORM 库
* [IntelliJ IDEA](https://www.jetbrains.com/idea)：全宇宙暂时排名第二的 IDE