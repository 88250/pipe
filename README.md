<p align = "center">
<img alt="Pipe" src="https://b3log.org/images/brand/pipe-128.png">
<br><br>
小而美的博客平台，为未来而构建
<br><br>
<a title="Build Status" target="_blank" href="https://github.com/88250/pipe/actions/workflows/gotest.yml"><img src="https://img.shields.io/github/actions/workflow/status/88250/pipe/gotest.yml?style=flat-square"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/88250/pipe"><img src="https://goreportcard.com/badge/github.com/88250/pipe?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/github/88250/pipe"><img src="https://img.shields.io/coveralls/github/88250/pipe.svg?style=flat-square&color=CC9933"></a>
<a title="Code Size" target="_blank" href="https://github.com/88250/pipe"><img src="https://img.shields.io/github/languages/code-size/88250/pipe.svg?style=flat-square"></a>
<a title="GPLv3" target="_blank" href="https://github.com/88250/pipe/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-GPLv3-orange.svg?style=flat-square"></a>
<br>
<a title="Releases" target="_blank" href="https://github.com/88250/pipe/releases"><img src="https://img.shields.io/github/release/88250/pipe.svg?style=flat-square"></a>
<a title="Release Date" target="_blank" href="https://github.com/88250/pipe/releases"><img src="https://img.shields.io/github/release-date/88250/pipe.svg?style=flat-square&color=99CCFF"></a>
<a title="Docker Image CI" target="_blank" href="https://github.com/88250/pipe/actions"><img src="https://img.shields.io/github/actions/workflow/status/88250/pipe/dockerimage.yml?label=Actions&logo=github&style=flat-square"></a>
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
<a href="https://github.com/88250/pipe/blob/master/README_en_US.md">English</a>
</p>

## 💡 简介

[Pipe](https://github.com/88250/pipe) 是一款小而美的开源博客平台，专为程序员设计。Pipe 有着非常活跃的[社区](https://ld246.com)，可将文章作为帖子推送到社区，来自社区的回帖将作为博客评论进行联动（具体细节请浏览 [B3log 构思 - 分布式社区网络](https://ld246.com/article/1546941897596)）。

> 这是一种全新的网络社区体验，让热爱记录和分享的你不再感到孤单！

欢迎到 [Pipe 官方讨论区](https://ld246.com/tag/pipe)了解更多。同时也欢迎关注 B3log 开源社区微信公众号 `B3log开源`：

![b3logos.jpg](https://b3logfile.com/file/2020/08/b3logos-032af045.jpg)

## 🗃 案例

* [Vanessa](https://vanessa.b3log.org)

## ✨ 功能

* 多用户博客平台
* [Markdown 编辑器](https://github.com/Vanessa219/vditor)支持三种编辑模式：所见即所得 / 即时渲染 / 分屏预览
* 聚合分类 / 标签
* 自定义导航
* 多主题 / 多语言
* Atom / RSS / Sitemap
* 文章搜索
* Hexo/Jekyll 导入 / 导出
* 可配置动静分离
* 支持 SQLite / MySQL / PostgreSQL

## 🎨 界面

### 开始使用

![start.png](https://b3logfile.com/file/2020/04/start-7fb7b415.png)

### 管理后台

![console.png](https://b3logfile.com/file/2020/04/console-047922de.png)

### 编辑文章

![post.png](https://b3logfile.com/file/2020/04/post-f52cbd5c.png)

### 主题选择

![theme.png](https://b3logfile.com/file/2020/04/theme-d2799005.png)

### 主题 Gina

![gina.png](https://b3logfile.com/file/2020/04/gina-d7fe2313.png)

## 🛠️ 安装

Pipe 仅支持通过 Docker 进行部署，如果你需要从源码构建可参考[这里](https://ld246.com/article/1533965022328)。

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
  
  注意：需先确保 SQLite 数据库文件已存在。如果数据库文件不存在时，docker run --volume 参数默认将宿主路径识别为目录，并自动创建这个目录，这可能导致 pipe 创建 sqlite 数据库文件失败。新建数据库文件可以简单用 touch 命令，如：
  
* ```shell
  $ touch ~/pipe.db
  ```

启动参数说明：

* `--port`：进程监听端口
* `--server`：访问时的链接

完整启动参数的说明可以使用 `-h` 来查看。

### Docker 升级

1. 拉取最新镜像
2. 重启容器

可参考[这里](https://github.com/88250/pipe/blob/master/docker-restart.sh)编写一个重启脚本，并通过 crontab 每日凌晨运行来实现自动更新。

### NGINX 反代

```nginx
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

另外，可以参考 https://ld246.com/article/1517474627971 进行配置。

## 📜 文档

* [《提问的智慧》精读注解版](https://ld246.com/article/1536377163156)
* [Pipe 用户指南](https://ld246.com/article/1513761942333)
* [Pipe 开发指南](https://ld246.com/article/1533965022328)
* [Pipe 主题开发指南](https://ld246.com/article/1512550354920)
* [Pipe Postman 测试集](https://www.getpostman.com/collections/900ddef64ad0e60479a6)

## 🏘️ 社区

* [讨论区](https://ld246.com/tag/pipe)
* [报告问题](https://github.com/88250/pipe/issues/new/choose)

## 📄 授权

Pipe 使用 [木兰宽松许可证, 第2版](http://license.coscl.org.cn/MulanPSL2) 开源协议。

## 🙏 鸣谢

* [jQuery](https://github.com/jquery/jquery)：JavaScript 工具库，用于主题页面
* [Vue.js](https://github.com/vuejs/vue)：渐进式 JavaScript 框架
* [Nuxt.js](https://github.com/nuxt/nuxt.js)：Vue.js 框架
* [Vuetify](https://github.com/vanessa219/vuetify)：Vue.js 的 Material 组件框架
* [Vditor](https://github.com/Vanessa219/vditor)： 浏览器端的 Markdown 编辑器
* [Gin](https://github.com/gin-gonic/gin)：又快又好用的 golang HTTP Web 框架
* [GORM](https://github.com/jinzhu/gorm)：极好的 golang ORM 库
* [SQLite](https://www.sqlite.org)：使用广泛的嵌入式 SQL 引擎
* [GCache](https://github.com/bluele/gcache)：golang 缓存库
* [Gulu](https://github.com/88250/gulu)：Go 语言常用工具库，这个轱辘还算圆
* [Lute](https://github.com/88250/lute)：一款结构化的 Markdown 引擎，支持 Go 和 JavaScript

---

## 特性说明

### 发布文章

Pipe 的文章编辑器支持 Markdown，并支持复制/粘贴图片、粘贴 HTML 自动转换 Markdown、流程图、数学公式等。

另外，可以为文章启用自动配图，会自动在文章最前面插入所选择的配图。

### 聚合分类

Pipe 使用“自底向上”的分类方式：

1. 定义分类，并配置该分类包含的标签
2. 查询某个分类文章列表时通过分类-> 标签集-> 标签关联的文章进行聚合

也就是说一篇文章在编辑时只需要打标签，访问分类时会根据该分类包含的标签将文章关联出来。这是一个自底向上的信息架构，在使用时更灵活一些，可以随时调整分类而不必重新更新文章。

### 域名绑定

在 Pipe 平台上的每个博客都可以配置域名，需要博主和服务器运维者分别操作：

1. 博主在设置 -> 基础配置 -> 博客地址一栏填写域名
2. 运维者通过配置 NGINX 实现域名到 /blogs/{username} 的反向代理

```
server {
    listen 80;
    server_name vanessa.b3log.org;

    location / {
        proxy_pass http://pipe/blogs/Vanessa/;
    }
}
```

### 导入 / 导出

Pipe 支持导入 Hexo/Jekyll 的 Markdown 文件，将需要导入的 Markdown 文件使用 zip 压缩上传即可。导入时会按标题去重，并自动按原文章的创建时间生成存档。

同样地，Pipe 也支持 Markdown 导出，格式为 Hexo。

### 链滴

在 Pipe 上发布文章时可选择是否自动推送到链滴上，这样能让更多人看到你创作的内容，更容易引起大家的关注和互动。

## 运维

### 数据库

Pipe 使用 SQLite3 数据库引擎，数据文件默认情况下存放在 ~/pipe.db，可以通过修改 pipe.json 的 `DataFilePath` 指定新的存放路径。

建议定期备份数据文件，避免意外情况导致数据丢失。

### 版本升级

在管理后台的关于中可以检查版本更新，如果提示有更新请尽快升级，一般来说升级只需要下载新的发布包然后部署重启，实际升级方式以每次版本发布公告为准。

## FAQ

### 如何做友链页面？

Pipe 没有单独的友链管理功能。可以通过发一篇文章，然后在导航管理中新建一个友链导航跳转过去。

## 结语

* 如果你在使用 Pipe 的过程中碰到问题或者有需求要提，欢迎跟帖，我们会在第一时间回复 😄
* 如果你想自己开发 Pipe，请参考 https://ld246.com/article/1533965022328
* 如果你想自己开发 Pipe 主题，请参考 https://ld246.com/article/1512550354920
