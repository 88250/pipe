// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2018, b3log.org
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
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	model "github.com/b3log/pipe/model"
)

func showIndexAction(c *gin.Context) {
	t, err := template.ParseFiles(filepath.ToSlash(filepath.Join(model.Conf.StaticRoot, "console/dist/index.html")))
	if nil != err {
		logger.Errorf("load index page failed: " + err.Error())
		c.String(http.StatusNotFound, "load index page failed")

		return
	}

	t.Execute(c.Writer, nil)
}

func showPlatInfo(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	data := map[string]interface{}{}
	data["version"] = model.Version
	data["database"] = service.Database()
	data["mode"] = model.Conf.RuntimeMode
	data["server"] = model.Conf.Server
	data["staticServer"] = model.Conf.StaticServer
	data["staticResourceVer"] = model.Conf.StaticResourceVersion

	result.Data = data
}

func showTopBlogs(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	blogs := service.User.GetTopBlogs(10)
	for _, blog := range blogs {
		blog.ID = 0
		blog.UserID = 0
		blog.UserRole = 0
	}

	result.Data = blogs
}
