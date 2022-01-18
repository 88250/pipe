<p align = "center">
<img alt="Pipe" src="https://b3log.org/images/brand/pipe-128.png">
<br><br>
Small and beautiful blogging platform, built for the future
<br><br>
<a title="Build Status" target="_blank" href="https://github.com/88250/pipe/actions/workflows/gotest.yml"><img src="https://img.shields.io/github/workflow/status/88250/pipe/Go%20Test?style=flat-square"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/88250/pipe"><img src="https://goreportcard.com/badge/github.com/88250/pipe?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/github/88250/pipe"><img src="https://img.shields.io/coveralls/github/88250/pipe.svg?style=flat-square&color=CC9933"></a>
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

<p align="center">
<a href="https://github.com/88250/pipe/blob/master/README.md">中文</a>
</p>

## 💡 Introduction

[Pipe](https://github.com/88250/pipe) is a small and beautiful open source blog platform designed for programmers. Pipe has a very active [community](https://ld246.com), which can push articles as posts to the community, and replies from the community will be linked as blog comments (for details, please visit [B3log Ideas - Distributed Community Network](https://ld246.com/article/1546941897596)).

> This is a brand new online community experience, so that you who love recording and sharing no longer feel lonely!

Welcome to [Pipe Official Discussion Forum](https://ld246.com/tag/pipe) to learn more.

## 🗃 Showcases

* [Vanessa](https://vanessa.b3log.org)

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
* Support SQLite / MySQL / PostgreSQL

## 🎨 Screenshots

### Start

![start.png](https://b3logfile.com/file/2020/04/start-7fb7b415.png)

### Console

![console.png](https://b3logfile.com/file/2020/04/console-047922de.png)

### Post

![post.png](https://b3logfile.com/file/2020/04/post-f52cbd5c.png)

### Theme

![theme.png](https://b3logfile.com/file/2020/04/theme-d2799005.png)

### Theme Gina

![gina.png](https://b3logfile.com/file/2020/04/gina-d7fe2313.png)

## 🛠️ Setup

Pipe only supports deployment via Docker. If you need to build from source, please refer to [here](https://ld246.com/article/1533965022328).

### Docker deploy

Get the latest image: 

```shell
docker pull b3log/pipe
```

* Use MySQL
  First create database schema manually (schema name `pipe`, character set use` utf8mb4`, sorting rule `utf8mb4_general_ci`), and then start the container:

  ```shell
  docker run --detach --name pipe --network=host \
      b3log/pipe --mysql="root:123456@(127.0.0.1:3306)/pipe?charset=utf8mb4&parseTime=True&loc=Local&timeout=1s" --runtime_mode=prod --port=5897 --server=http://localhost:5897
  ```

  For simplicity, the host network mode is used to connect to MySQL on the host.
* Use SQLite

  ```shell
  docker run --detach --name pipe --volume ~/pipe.db:/opt/pipe/pipe.db --publish 5897:5897 \
      b3log/pipe --sqlite="/opt/pipe/pipe.db" --runtime_mode=prod --port=5897 --server=http://localhost:5897
  ```

Start command line arguments description:

* `--port`: process listen port
* `--server`: the URL for the final visiting

The description of the complete startup arguments can be viewed using `-h`.

### Docker upgrade

1. Pull the latest image
2. Restart the container

You can refer to [here](https://github.com/88250/pipe/blob/master/docker-restart.sh) to write a restart script and run it through crontab every morning to achieve automatic update.

### NGINX reverse proxy

```nginx
upstream pipe {
    server localhost:5897;
}

server {
    listen 80;
    server_name pipe.b3log.org; # blog domain

    location / {
        proxy_pass http://pipe$request_uri;
        proxy_set_header  Host $host:$server_port;
        proxy_set_header  X-Real-IP  $remote_addr;
        client_max_body_size  10m;
    }
}
```

In addition, you can refer to https://ld246.com/article/1517474627971 for configuration.

## 📜 Documentation

* [Pipe User Guide](https://ld246.com/article/1513761942333)
* [Pipe Developer Guide](https://ld246.com/article/1533965022328)
* [Pipe Theme Development Guide](https://ld246.com/article/1512550354920)
* [Pipe Postman Test Collection](https://www.getpostman.com/collections/900ddef64ad0e60479a6)

## 🏘️ Community

* [Forum](https://ld246.com/tag/pipe)
* [Issues](https://github.com/88250/pipe/issues/new/choose)

## 📄 License

Pipe uses the [Mulan Permissive Software License, Version 2](http://license.coscl.org.cn/MulanPSL2) open source license.

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
* [Gulu](https://github.com/88250/gulu): Go commons utilities
* [Lute](https://github.com/88250/lute): A structured Markdown engine that supports Go and JavaScript
