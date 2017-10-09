// Solo.go - A small and beautiful blogging platform written in golang.
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

package controller

import (
	"net/http"
	"strings"

	"github.com/b3log/solo.go/controller/console"
	"github.com/b3log/solo.go/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// MapRoutes returns a gin engine and binds controllers with request URLs.
func MapRoutes() *gin.Engine {
	ret := gin.New()
	// TODO: D, ret.Use(favicon.New("./favicon.ico"))
	ret.Use(gin.Recovery())

	store := sessions.NewCookieStore([]byte(util.Conf.SessionSecret))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   util.Conf.SessionMaxAge,
		Secure:   strings.HasPrefix(util.Conf.Server, "https"),
		HttpOnly: true,
	})
	ret.Use(sessions.Sessions("sologo", store))

	api := ret.Group("/api")
	api.POST("/init", initCtl)
	api.POST("/login", loginCtl)
	api.POST("/logout", logoutCtl)
	api.Any("/hp/*apis", util.HacPaiAPI())
	statusGroup := api.Group("/status")
	statusGroup.GET("", GetPlatformStatusCtl)
	statusGroup.GET("/ping", pingCtl)
	consoleGroup := api.Group("/console")
	consoleGroup.Use(console.LoginCheck())
	consoleGroup.POST("/articles", console.AddArticleCtl)
	consoleGroup.GET("/articles", console.GetArticlesCtl)
	consoleGroup.GET("/articles/:id", console.GetArticleCtl)
	consoleGroup.DELETE("/articles/:id", console.RemoveArticleCtl)
	consoleGroup.PUT("/articles/:id", console.UpdateArticleCtl)
	consoleGroup.GET("/comments", console.GetCommentsCtl)
	consoleGroup.DELETE("/comments/:id", console.RemoveCommentCtl)
	consoleGroup.GET("/navigations", console.GetNavigationsCtl)
	consoleGroup.GET("/navigation/:id", console.GetNavigationCtl)
	consoleGroup.GET("/tags", console.GetTagsCtl)
	consoleSettingsGroup := consoleGroup.Group("/settings")
	consoleSettingsGroup.GET("/basic", console.GetBasicSettingsCtl)
	consoleGroup.POST("/blog/switch/:id", console.BlogSwitchCtl)
	themeGroup := ret.Group("")
	themeGroup.GET("/", indexCtl)

	ret.LoadHTMLFiles("console/dist/admin/index.html")
	ret.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	ret.Static("/assets", "./console/dist")

	return ret
}
