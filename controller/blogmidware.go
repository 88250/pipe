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
	"math"
	"net/http"
	"strconv"
	"time"

	"strings"

	"github.com/b3log/solo.go/i18n"
	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func resolveBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		if "" == username {
			c.AbortWithStatus(http.StatusNotFound)

			return
		}
		blogAdmin := service.User.GetUserByName(username)
		if nil == blogAdmin {
			c.AbortWithStatus(http.StatusNotFound)

			return
		}
		c.Set("blogAdmin", blogAdmin)

		fillCommon(c, &DataModel{})

		c.Next()
	}
}

type DataModel map[string]interface{}

func fillCommon(c *gin.Context, dataModel *DataModel) {
	if "dev" == util.Conf.RuntimeMode {
		i18n.Load()
	}
	localeSetting := service.Setting.GetSetting(model.SettingCategoryI18n, model.SettingNameI18nLocale, 1)
	i18ns := i18n.GetMessages(localeSetting.Value)
	i18nMap := map[string]interface{}{}
	for key, value := range i18ns {
		i18nMap[strings.Title(key)] = value
	}
	(*dataModel)["I18n"] = i18nMap

	settings := service.Setting.GetAllSettings(1)
	settingMap := map[string]string{}
	for _, setting := range settings {
		settingMap[strings.Title(setting.Name)] = setting.Value
	}
	settingMap["SystemPath"] = util.PathBlogs + settingMap["SystemPath"]

	(*dataModel)["Setting"] = settingMap

	statistics := service.Statistic.GetAllStatistics(1)
	statisticMap := map[string]int{}
	for _, statistic := range statistics {
		count, err := strconv.Atoi(statistic.Value)
		if nil != err {
			log.Errorf("statistic [%s] should be an integer, actual is [%v]", statistic.Name, statistic.Value)
		}
		statisticMap[strings.Title(statistic.Name)] = count
	}
	(*dataModel)["Statistic"] = statisticMap
	(*dataModel)["Title"] = settingMap["BasicBlogTitle"]
	(*dataModel)["MetaKeywords"] = settingMap["BasicMetaKeywords"]
	(*dataModel)["FaviconURL"] = settingMap["BasicFaviconURL"]
	(*dataModel)["MetaDescription"] = settingMap["BasicMetaDescription"]
	(*dataModel)["Conf"] = util.Conf
	(*dataModel)["Year"] = time.Now().Year()
	(*dataModel)["BlogURL"] = util.Conf.Server + settingMap["SystemPath"]

	(*dataModel)["Username"] = ""
	session := util.GetSession(c)
	if nil != session {
		(*dataModel)["Username"] = session.UName
	}
	(*dataModel)["UserCount"] = len(service.User.GetBlogUsers(1)) + 2

	navigations := service.Navigation.GetNavigations(1)
	(*dataModel)["Navigations"] = navigations

	tagSize, err := strconv.Atoi(settingMap[model.SettingNamePreferenceMostUseTagListSize])
	if nil != err {
		log.Errorf("setting [%s] shoud be an integer, actual is [%v]", model.SettingNamePreferenceMostUseTagListSize,
			settingMap[model.SettingNamePreferenceMostUseTagListSize])
		tagSize = model.SettingPreferenceMostUseTagListSizeDefault
	}
	tags := service.Tag.GetTags(tagSize)
	themeTags := []*ThemeTag{}
	for _, tag := range tags {
		themeTag := &ThemeTag{
			Title: tag.Title,
			URL:   util.Conf.Server + settingMap["SystemPath"] + "/" + tag.Title,
		}
		themeTags = append(themeTags, themeTag)
	}
	(*dataModel)["MostUseTags"] = themeTags

	categories := service.Category.GetCategories(math.MaxInt8)
	themeCategories := []*ThemeCategory{}
	for _, category := range categories {
		themeCategory := &ThemeCategory{
			Title: category.Title,
			URL:   util.Conf.Server + settingMap["SystemPath"] + "/" + category.Title,
		}
		themeCategories = append(themeCategories, themeCategory)
	}
	(*dataModel)["MostUseCategories"] = themeCategories

	themeArticles := []*ThemeListArticle{}
	for _, tag := range tags {
		author := &ThemeAuthor{
			Name:      "Vanessa",
			URL:       "http://localhost:5879/blogs/solo/vanessa",
			AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
		}
		themeArticle := &ThemeListArticle{
			Title:     tag.Title,
			URL:       util.Conf.Server + settingMap["SystemPath"] + "/" + tag.Title,
			CreatedAt: "1天前",
			Author:    author,
		}
		themeArticles = append(themeArticles, themeArticle)
	}
	(*dataModel)["RecentComments"] = themeArticles
	(*dataModel)["MostViewArticles"] = themeArticles
	(*dataModel)["MostCommentArticles"] = themeArticles
	(*dataModel)["RandomArticles"] = themeArticles

	c.Set("dataModel", dataModel)
}
