// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package controller

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"

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
		"version":         model.Version,
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
