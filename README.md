<p align = "center">
<img alt="Pipe" src="https://static.b3log.org/images/brand/pipe-128.png">
<br><br>
小而美的博客平台，为未来而构建
<br><br>
<a title="Build Status" target="_blank" href="https://travis-ci.org/88250/pipe"><img src="https://img.shields.io/travis/88250/pipe.svg?style=flat-square"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/88250/pipe"><img src="https://goreportcard.com/badge/github.com/88250/pipe?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/repos/github/88250/pipe/badge.svg?branch=master"><img src="https://img.shields.io/coveralls/github/88250/pipe.svg?style=flat-square&color=CC9933"></a>
<a title="Code Size" target="_blank" href="https://github.com/88250/pipe"><img src="https://img.shields.io/github/languages/code-size/88250/pipe.svg?style=flat-square"></a>
<a title="GPLv3" target="_blank" href="https://github.com/88250/pipe/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-GPLv3-orange.svg?style=flat-square"></a>
<br>
<a title="Releases" target="_blank" href="https://github.com/88250/pipe/releases"><img src="https://img.shields.io/github/release/88250/pipe.svg?style=flat-square"></a>
<a title="Release Date" target="_blank" href="https://github.com/88250/pipe/releases"><img src="https://img.shields.io/github/release-date/88250/pipe.svg?style=flat-square&color=99CCFF"></a>
<a title="Docker Image CI" target="_blank" href="https://github.com/88250/pipe/actions"><img src="https://img.shields.io/github/workflow/status/88250/pipe/Docker%20Image%20CI?label=Actions&logo=github&style=flat-square"></a>
<a title="Docker Pulls" target="_blank" href="https://hub.docker.com/r/b3log/pipe"><img src="https://img.shields.io/docker/pulls/b3log/pipe.svg?style=flat-square&color=blueviolet"></a>
<br>
<a title="GitHub Commits" target="_blank" href="https://github.com/88250/pipe/commits/master"><img src="https://img.shields.io/github/commit-activity/m/88250/pipe.svg?style=flat-square"></a>
<a title="Last Commit" target="_blank" href="https://github.com/88250/pipe/commits/master"><img src="https://img.shields.io/github/last-commit/88250/pipe.svg?style=flat-square&color=FF9900"></a>
<a title="GitHub Pull Requests" target="_blank" href="https://github.com/88250/pipe/pulls"><img src="https://img.shields.io/github/issues-pr-closed/88250/pipe.svg?style=flat-square&color=FF9966"></a>
<a title="Hits" target="_blank" href="https://github.com/88250/hits"><img src="https://hits.b3log.org/88250/pipe.svg"></a>
<br><br>
<a title="GitHub Watchers" target="_blank" href="https://github.com/88250/pipe/watchers"><img src="https://img.shields.io/github/watchers/88250/pipe.svg?label=Watchers&style=social"></a>  
<a title="GitHub Stars" target="_blank" href="https://github.com/88250/pipe/stargazers"><img src="https://img.shields.io/github/stars/88250/pipe.svg?label=Stars&style=social"></a>  
<a title="GitHub Forks" target="_blank" href="https://github.com/88250/pipe/network/members"><img src="https://img.shields.io/github/forks/88250/pipe.svg?label=Forks&style=social"></a>  
<a title="Author GitHub Followers" target="_blank" href="https://github.com/88250"><img src="https://img.shields.io/github/followers/88250.svg?label=Followers&style=social"></a>
</p>

## 💡 简介

[Pipe](https://github.com/88250/pipe) is a small and beautiful open source blog platform designed for programmers. Pipe has a very active [community](https://hacpai.com), which can push articles as posts to the community, and replies from the community will be linked as blog comments (for details, please visit [B3log Ideas - Distributed Community Network](https://hacpai.com/article/1546941897596)).

> This is a brand new online community experience, so that you who love recording and sharing no longer feel lonely!

Welcome to [Pipe Official Discussion Forum](https://hacpai.com/tag/pipe) to learn more.

## 🗃 Showcases

* http://vanessa.b3log.org
* http://blog.bhusk.com
* https://www.zorkelvll.cn
* http://o0o.pub
* http://blog.gitor.org

## ✨ Features

* Multi-user blog platform
* [Markdown editor](https://github.com/Vanessa219/vditor) supports three editing modes: WYSIWYG/Instant Rendering/Split View
* Tag aggregation classification
* Custom navigation links
* Multiple themes / multiple languages
* Atom / RSS / Sitemap
* Article search
* Hexo/Jekyll import / export
* CDN static resource separation
* Support SQLite / MySQL

## 🎨 Screenshots

### Start

![start.png](https://img.hacpai.com/file/2020/04/start-3064240e.png)

### Console

![console.png](https://img.hacpai.com/file/2020/04/console-047922de.png)

### Post

![post.png](https://img.hacpai.com/file/2020/04/post-f52cbd5c.png)

### Theme

![theme.png](https://img.hacpai.com/file/2020/04/theme-d2799005.png)

### Theme Gina

![gina.png](https://img.hacpai.com/file/2020/04/gina-d7fe2313.png)

## 🛠️ Setup

### 本地试用

* [下载](https://github.com/88250/pipe/releases)最新的发布包解压，进入解压目录运行 pipe/pipe.exe
* 从源码构建可参考[这里](https://hacpai.com/article/1533965022328)

**请注意**：我们不建议通过发布包或者源码构建部署，因为这样的部署方式在将来有新版本发布时升级会比较麻烦。
这两种方式请仅用于本地试用，线上生产环境建议通过 Docker 部署。

### Docker 部署

获取最新镜像：

```shell
docker pull b3log/pipe
```

* 使用 MySQL
  先手动建库（库名 `pipe` ，字符集使用 `utf8mb4` ，排序规则 `utf8mb4_general_ci` ），然后启动容器：

  ```shell
  docker run --detach --name pipe --network=host \
      b3log/pipe --mysql="root:123456@(127.0.0.1:3306)/pipe?charset=utf8mb4&parseTime=True&loc=Local&timeout=1s" --runtime_mode=prod --port=5897 --server=http://localhost:5897
  ```


  为了简单，使用了主机网络模式来连接主机上的 MySQL。
* 使用 SQLite

  ```shell
  docker run --detach --name pipe --volume ~/pipe.db:/opt/pipe/pipe.db --publish 5897:5897 \
      b3log/pipe --sqlite="/opt/pipe/pipe.db" --runtime_mode=prod --port=5897 --server=http://localhost:5897
  ```

启动参数说明：

* `--port` ：进程监听端口
* `--server` ：访问时的链接

完整启动参数的说明可以使用 `-h` 来查看。

### Docker 升级

1. 拉取最新镜像
2. 重启容器

可参考[这里](https://github.com/88250/pipe/blob/master/docker-restart.sh)编写一个重启脚本，并通过 crontab 每日凌晨运行来实现自动更新。

### NGINX 反代

```
upstream pipe {
    server localhost:5897;
}

server {
    listen 80;
    server_name pipe.b3log.org; # 配置为你自己的域名

    location / {
        proxy_pass http://pipe$request_uri;
        proxy_set_header  Host $host:$server_port;
        proxy_set_header  X-Real-IP  $remote_addr;
        client_max_body_size  10m;
    }
}
```

另外，可以参考 https://hacpai.com/article/1517474627971 进行配置。

## 📜 文档

* [《提问的智慧》精读注解版](https://hacpai.com/article/1536377163156)
* [用户指南](https://hacpai.com/article/1513761942333)
* [开发指南](https://hacpai.com/article/1533965022328)
* [主题开发指南](https://hacpai.com/article/1512550354920)
* [贡献指南](https://github.com/88250/pipe/blob/master/CONTRIBUTING.md)
* [Postman 测试集](https://www.getpostman.com/collections/900ddef64ad0e60479a6)

## 🏘️ 社区

* [讨论区](https://hacpai.com/tag/pipe)
* [报告问题](https://github.com/88250/pipe/issues/new/choose)

## 📄 License

Pipe uses the [Mulan Permissive Software License，Version 2](http://license.coscl.org.cn/MulanPSL2) open source license.

## 🙏 Acknowledgement

* [jQuery](https://github.com/jquery/jquery): A JavaScript tool library for theme pages
* [Vue.js](https://github.com/vuejs/vue): A progressive, incrementally-adoptable JavaScript framework
* [Nuxt.js](https://github.com/nuxt/nuxt.js): The Vue.js Framework
* [Vuetify](https://github.com/vanessa219/vuetify): Material Component Framework for Vue
* [Vditor](https://github.com/Vanessa219/vditor): An In-browser Markdown editor
* [Gin](https://github.com/gin-gonic/gin): A HTTP web framework written in Go
* [GORM](https://github.com/jinzhu/gorm): The fantastic ORM library for Golang
* [SQLite](https://www.sqlite.org): The most used database engine in the world
* [GCache](https://github.com/bluele/gcache): Cache library for golang
* [Gulu](https://github.com/88250/gulu)：Go commons utilities
* [Lute](https://github.com/88250/lute): A structured Markdown engine that supports Go and JavaScript
