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
	"strconv"
	"strings"
	"time"

	"github.com/b3log/solo.go/controller/console"
	"github.com/b3log/solo.go/i18n"
	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/theme"
	"github.com/b3log/solo.go/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type DataModel map[string]interface{}

// MapRoutes returns a gin engine and binds controllers with request URLs.
func MapRoutes() *gin.Engine {
	ret := gin.New()

	ret.Use(gin.Recovery())

	store := sessions.NewCookieStore([]byte(util.Conf.SessionSecret))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   util.Conf.SessionMaxAge,
		Secure:   strings.HasPrefix(util.Conf.Server, "https"),
		HttpOnly: true,
	})
	ret.Use(sessions.Sessions("solo.go", store))

	api := ret.Group(util.PathAPI)
	api.POST("/init", initAction)
	api.POST("/login", loginAction)
	api.POST("/logout", logoutAction)
	api.Any("/hp/*apis", util.HacPaiAPI())
	api.GET("/status", getStatusAction)

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

	consoleGroup.POST("/blogs/switch/:id", console.BlogSwitchAction)

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

	ret.StaticFile(util.PathFavicon, "console/static/favicon.ico")

	themePath := "theme/x/" + theme.DefaultTheme
	ret.Static("/"+themePath+"/css", themePath+"/css")
	ret.Static("/"+themePath+"/js", themePath+"/js")
	ret.Static("/"+themePath+"/images", themePath+"/images")
	ret.Static(util.PathTheme+"/css", "theme/css")
	ret.Static(util.PathTheme+"/js", "theme/js")
	ret.LoadHTMLGlob(themePath + "/*.html")
	themeGroup := ret.Group("/blogs/:username")
	themeGroup.Use(resolveBlog())
	themeGroup.GET("/", showArticlesAction)
	themeGroup.GET(util.PathActivities, showActivitiesAction)
	themeGroup.GET(util.PathArchives+"/:archive", showArchiveArticlesAction)
	themeGroup.GET(util.PathArchives, showArchivesAction)
	themeGroup.GET(util.PathArticles+"/:id", showArticleAction)
	themeGroup.GET(util.PathAuthors+"/:name", showAuthorArticlesAction)
	themeGroup.GET(util.PathAuthors, showAuthorsAction)
	themeGroup.GET(util.PathCategories, showCategoriesAction)
	themeGroup.GET(util.PathCategories+"/:category", showCategoriesAction)
	themeGroup.GET(util.PathTags, showTagsAction)
	themeGroup.GET(util.PathTags+"/:tag", showTagArticlesAction)

	ret.GET(util.PathAdmin+"/*path", console.ShowPageAction)
	ret.GET("/login", console.ShowLoginAction)
	ret.Static(util.PathAssets, "./console/dist")

	return ret
}

func fillCommon(c *gin.Context, dataModel *DataModel) {
	if "dev" == util.Conf.RuntimeMode {
		i18n.Load()
	}
	localeSetting := service.Setting.GetSetting(model.SettingCategoryI18n, model.SettingNameI18nLocale, 1)
	(*dataModel)["i18n"] = i18n.GetMessages(localeSetting.Value)

	settings := service.Setting.GetAllSettings(1)
	settingMap := map[string]string{}
	for _, setting := range settings {
		settingMap[setting.Name] = setting.Value
	}
	settingMap[model.SettingNameSystemPath] = "/blog" + settingMap[model.SettingNameSystemPath]
	(*dataModel)["setting"] = settingMap
	statistics := service.Statistic.GetAllStatistics(1)
	statisticMap := map[string]int{}
	for _, statistic := range statistics {
		count, err := strconv.Atoi(statistic.Value)
		if nil != err {
			log.Errorf("statistic [%s] should be an integer, actual is [%v]", statistic.Name, statistic.Value)
		}
		statisticMap[statistic.Name] = count
	}
	(*dataModel)["statistic"] = statisticMap
	(*dataModel)["title"] = settingMap["basicBlogTitle"]
	(*dataModel)["metaKeywords"] = settingMap["basicMetaKeywords"]
	(*dataModel)["metaDescription"] = settingMap["basicMetaDescription"]
	(*dataModel)["conf"] = util.Conf
	(*dataModel)["year"] = time.Now().Year()

	(*dataModel)["username"] = ""
	session := util.GetSession(c)
	if nil != session {
		(*dataModel)["username"] = session.UName
	}
	(*dataModel)["userCount"] = len(service.User.GetBlogUsers(1))

	navigations := service.Navigation.GetNavigations(1)
	(*dataModel)["navigations"] = navigations
}
