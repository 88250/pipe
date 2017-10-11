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
	api.POST("/init", initAction)
	api.POST("/login", loginAction)
	api.POST("/logout", logoutAction)
	api.Any("/hp/*apis", util.HacPaiAPI())

	statusGroup := api.Group("/status")
	statusGroup.GET("", GetPlatformStatusAction)
	statusGroup.GET("/ping", pingAction)

	consoleGroup := api.Group("/console")
	consoleGroup.Use(console.LoginCheck())

	consoleGroup.POST("/articles", console.AddArticleAction)
	consoleGroup.GET("/articles", console.GetArticlesAction)
	consoleGroup.GET("/articles/:id", console.GetArticleAction)
	consoleGroup.DELETE("/articles/:id", console.RemoveArticleAction)
	consoleGroup.PUT("/articles/:id", console.UpdateArticleAction)

	consoleGroup.GET("/comments", console.GetCommentsAction)
	consoleGroup.DELETE("/comments/:id", console.RemoveCommentAction)

	consoleGroup.GET("/navigations", console.GetNavigationsAction)
	consoleGroup.GET("/navigations/:id", console.GetNavigationAction)
	consoleGroup.PUT("/navigations/:id", console.UpdateNavigationAction)
	consoleGroup.POST("/navigations", console.AddNavigationAction)
	consoleGroup.DELETE("/navigations/:id", console.RemoveNavigationAction)

	consoleGroup.GET("/categories", console.GetCategoriesAction)
	consoleGroup.POST("/categories", console.AddCategoryAction)

	consoleGroup.GET("/tags", console.GetTagsAction)

	consoleGroup.POST("/blog/switch/:id", console.BlogSwitchAction)

	consoleSettingsGroup := consoleGroup.Group("/settings")
	consoleSettingsGroup.GET("/basic", console.GetBasicSettingsAction)
	consoleSettingsGroup.PUT("/basic", console.UpdateBasicSettingsAction)
	consoleSettingsGroup.GET("/preference", console.GetPreferenceSettingsAction)
	consoleSettingsGroup.PUT("/preference", console.UpdatePreferenceSettingsAction)
	consoleSettingsGroup.GET("/sign", console.GetSignSettingsAction)
	consoleSettingsGroup.PUT("/sign", console.UpdateSignSettingsAction)
	consoleSettingsGroup.GET("/i18n", console.GetI18nSettingsAction)
	consoleSettingsGroup.PUT("/i18n", console.UpdateI18nSettingsAction)
	consoleSettingsGroup.GET("/feed", console.GetFeedSettingsAction)
	consoleSettingsGroup.PUT("/feed", console.UpdateFeedSettingsAction)

	themeGroup := ret.Group("")
	themeGroup.GET("/", indexAction)

	ret.LoadHTMLFiles("console/dist/admin/index.html")
	ret.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	ret.Static("/assets", "./console/dist")

	return ret
}
