// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package controller

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/88250/pipe/util"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/gin-gonic/gin"
)

func showIndexAction(c *gin.Context) {
	t, err := template.ParseFiles("console/dist/index.html")
	if nil != err {
		logger.Errorf("load index page failed: " + err.Error())
		c.String(http.StatusNotFound, "load index page failed")

		return
	}

	t.Execute(c.Writer, nil)
}

func showChangelogsAction(c *gin.Context) {
	data, err := ioutil.ReadFile("changelogs.md")
	if nil != err {
		logger.Errorf("load changelogs.md failed: " + err.Error())
		c.String(http.StatusNotFound, "load changelogs failed")

		return
	}

	result := util.Markdown(string(data))
	c.Data(http.StatusOK, "text/html; charset=utf-8", gulu.Str.ToBytes(result.ContentHTML))
}

func showStartPageAction(c *gin.Context) {
	t, err := template.ParseFiles("console/dist/start/index.html")
	if nil != err {
		logger.Errorf("load start page failed: " + err.Error())
		c.String(http.StatusNotFound, "load start page failed")

		return
	}

	t.Execute(c.Writer, nil)
}

func showPlatInfoAction(c *gin.Context) {
	result := blogInfo(c)
	c.JSON(http.StatusOK, result)
}

func showBlogInfoAction(c *gin.Context) {
	result := blogInfo(c)
	blogID := getBlogID(c)
	blogAdmin := service.User.GetBlogAdmin(blogID)
	result.Data.(map[string]interface{})["userName"] = blogAdmin.Name
	c.JSON(http.StatusOK, result)
}

func blogInfo(c *gin.Context) (ret *gulu.Result) {
	ret = gulu.Ret.NewResult()
	platformAdmin := service.User.GetPlatformAdmin()

	ret.Data = map[string]interface{}{
		"version":         util.Version,
		"servePath":       model.Conf.Server,
		"staticServePath": model.Conf.StaticServer,
		"runtimeMode":     model.Conf.RuntimeMode,
		"runtimeDatabase": service.Database(),
		"platformAdmin":   platformAdmin.Name,
	}

	return
}

func showTopBlogsAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	blogs := service.User.GetTopBlogs(10)
	for _, blog := range blogs {
		blog.ID = 0
		blog.UserID = 0
		blog.UserRole = 0
	}

	result.Data = blogs
}

func showManifestAction(c *gin.Context) {
	data, err := ioutil.ReadFile(filepath.FromSlash("theme/js/manifest.json"))
	if nil != err {
		notFound(c)

		return
	}

	manifest := string(data)
	manifest = strings.ReplaceAll(manifest, "{server}", model.Conf.Server)

	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.Write([]byte(manifest))
}
