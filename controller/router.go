// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
// Copyright (C) 2017, b3log.org
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

// Package controller includes controllers.
package controller

import (
	"github.com/b3log/solo.go/controller/console"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
)

// MapRoutes returns a gin engine and binds controllers with request URLs.
func MapRoutes() *gin.Engine {
	ret := gin.New()
	//ret.Use(favicon.New("./favicon.ico"))
	ret.Use(gin.Recovery())

	ret.Any("/hp/*apis", util.HacPaiAPI())

	ret.POST("/init", initCtl)

	status := ret.Group("/status")
	status.GET("", statusCtl)
	status.GET("/ping", pingCtl)

	adminConsole := ret.Group("/console")
	adminConsole.GET("/articles", console.GetArticlesCtl)
	adminConsole.GET("/articles/:id", console.GetArticleCtl)
	adminConsole.DELETE("/articles/:id", console.RemoveArticleCtl)
	adminConsole.PUT("/articles/:id", console.UpdateArticleCtl)

	return ret
}
