<p align = "center">
<img alt="Pipe" src="https://user-images.githubusercontent.com/873584/52321153-3d6d6900-2a0e-11e9-9ea6-57974a302dbf.png">
<br><br>
小而美的博客平台，为未来而构建
<br><br>
<a title="Build Status" target="_blank" href="https://travis-ci.org/b3log/pipe"><img src="https://img.shields.io/travis/b3log/pipe.svg?style=flat-square"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/b3log/pipe"><img src="https://goreportcard.com/badge/github.com/b3log/pipe"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/repos/github/b3log/pipe/badge.svg?branch=master"><img src="https://img.shields.io/coveralls/github/b3log/pipe.svg?style=flat-square"></a>
<a title="Code Size" target="_blank" href="https://github.com/b3log/pipe"><img src="https://img.shields.io/github/languages/code-size/b3log/pipe.svg?style=flat-square"></a>
<a title="GPLv3" target="_blank" href="https://github.com/b3log/pipe/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-GPLv3-orange.svg?style=flat-square"></a>
<a title="Releases" target="_blank" href="https://github.com/b3log/pipe/releases"><img src="https://img.shields.io/github/release/b3log/pipe.svg?style=flat-square"></a>
<a title="Downloads" target="_blank" href="https://github.com/b3log/pipe/releases"><img src="https://img.shields.io/github/downloads/b3log/pipe/total.svg?style=flat-square"></a>
<a title="Docker Pulls" target="_blank" href="https://hub.docker.com/r/b3log/pipe"><img src="https://img.shields.io/docker/pulls/b3log/pipe.svg?style=flat-square&color=blueviolet"></a>
<a title="Hits" target="_blank" href="https://github.com/b3log/hits"><img src="https://hits.b3log.org/b3log/pipe.svg"></a>
</p>

## 简介

[Pipe](https://github.com/b3log/pipe) 是一款小而美的开源博客平台，专为程序员设计。Pipe 有着非常活跃的[社区](https://hacpai.com)，文章自动推送到社区后可以让很多人看到，产生丰富的交流互动。

## 案例

* [Akkuman 的博客](http://o0o.pub)
* [Vanessa](http://vanessa.b3log.org)
* [图解吧](http://tujie8.net)
* [黑壳博客](http://blog.bhusk.com)
* [zorke 的博客](https://www.zorkelvll.cn)
* [一个码农](http://blog.gitor.org)

## 功能

* 多用户博客平台
* Markdown / Emoji
* 聚合分类 / 标签
* 自定义导航
* 多主题 / 多语言
* Atom / RSS / Sitemap
* 文章搜索
* Hexo/Jekyll 导入 / 导出
* 可配置动静分离
* 支持 SQLite / MySQL

## 界面

### 开始使用

![start](https://user-images.githubusercontent.com/873584/53852374-d044f780-3ffc-11e9-92bb-ea3e6c53ab1c.png)

### 管理后台

![console-index](https://user-images.githubusercontent.com/873584/52255023-a9d36400-294b-11e9-9c9f-a99c208d6b52.png)

### 编辑文章

![article](https://user-images.githubusercontent.com/873584/52255159-4dbd0f80-294c-11e9-980f-bc15cd0dd340.png)

### 主题选择

![theme](https://user-images.githubusercontent.com/873584/52255174-675e5700-294c-11e9-84bf-c04fe39ce5ec.png)

### 主题 Gina

![index](https://user-images.githubusercontent.com/873584/52255024-aa6bfa80-294b-11e9-9cfc-2c6cfc3d687c.png)

## 安装

### 本地试用

* [下载](https://github.com/b3log/pipe/releases)最新的发布包解压，进入解压目录运行 pipe/pipe.exe
* 从源码构建可参考[这里](https://hacpai.com/article/1533965022328)

**请注意**：我们不建议通过发布包或者源码构建部署，因为这样的部署方式在将来有新版本发布时升级会比较麻烦。
这两种方式请仅用于本地试用，线上生产环境建议通过 Docker 部署。

### Docker 部署

获取最新镜像：

```shell
docker pull b3log/pipe
```

* 使用 MySQL

  先手动建库（库名 `pipe`，字符集使用 `utf8mb4`，排序规则 `utf8mb4_general_ci`），然后启动容器：
  
  ```shell
  docker run --detach --name pipe --network=host \
      b3log/pipe --mysql="root:123456@(127.0.0.1:3306)/pipe?charset=utf8mb4&parseTime=True&loc=Local" --runtime_mode=prod --port=5897 --server=http://localhost:5897
  ```
  为了简单，使用了主机网络模式来连接主机上的 MySQL。
  
* 使用 SQLite

  ```shell
  docker run --detach --name pipe --volume ~/pipe.db:/opt/pipe/pipe.db --publish 5897:5897 \
      b3log/pipe --sqlite="/opt/pipe/pipe.db" --runtime_mode=prod --port=5897 --server=http://localhost:5897
  ```
  
启动参数说明：

* `--port`：进程监听端口
* `--server`：访问时的链接

完整启动参数的说明可以使用 `-h` 来查看。

### Docker 升级

1. 拉取最新镜像
2. 重启容器

可参考[这里](https://github.com/b3log/pipe/blob/master/docker-restart.sh)编写一个重启脚本，并通过 crontab 每日凌晨运行来实现自动更新。

## 文档

* [《提问的智慧》精读注解版](https://hacpai.com/article/1536377163156)
* [用户指南](https://hacpai.com/article/1513761942333)
* [开发指南](https://hacpai.com/article/1533965022328)
* [主题开发指南](https://hacpai.com/article/1512550354920)
* [贡献指南](https://github.com/b3log/pipe/blob/master/CONTRIBUTING.md)
* [Postman 测试集](https://www.getpostman.com/collections/900ddef64ad0e60479a6)

## 社区

* [讨论区](https://hacpai.com/tag/pipe)
* [报告问题](https://github.com/b3log/pipe/issues/new/choose)

## 授权

Pipe 使用 [GPLv3](https://www.gnu.org/licenses/gpl-3.0.txt) 开源协议。

## 鸣谢

* [jQuery](https://github.com/jquery/jquery)：JavaScript 工具库，用于主题页面
* [Vue.js](https://github.com/vuejs/vue)：渐进式 JavaScript 框架
* [Nuxt.js](https://github.com/nuxt/nuxt.js)：Vue.js 框架
* [Vuetify](https://github.com/vanessa219/vuetify)：Vue.js 的 Material 组件框架
* [Vditor](https://github.com/b3log/vditor)： 浏览器端的 Markdown 编辑器
* [Gin](https://github.com/gin-gonic/gin)：又快又好用的 golang HTTP web 框架
* [GORM](https://github.com/jinzhu/gorm)：极好的 golang ORM 库
* [Blackfriday](github.com/russross/blackfriday)：golang Markdown 处理器
* [SQLite](https://www.sqlite.org)：使用广泛的嵌入式 SQL 引擎
* [GCache](https://github.com/bluele/gcache)：golang 缓存库
