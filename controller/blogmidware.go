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
	"strings"
	"time"

	"github.com/b3log/solo.go/i18n"
	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/dustin/go-humanize"
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
	blogAdminVal, _ := c.Get("blogAdmin")
	blogAdmin := blogAdminVal.(*model.User)
	blogID := blogAdmin.BlogID

	if "dev" == util.Conf.RuntimeMode {
		i18n.Load()
	}
	localeSetting := service.Setting.GetSetting(model.SettingCategoryI18n, model.SettingNameI18nLocale, blogID)
	i18ns := i18n.GetMessages(localeSetting.Value)
	i18nMap := map[string]interface{}{}
	for key, value := range i18ns {
		i18nMap[strings.Title(key)] = value
		i18nMap[key] = value
	}
	(*dataModel)["I18n"] = i18nMap

	settings := service.Setting.GetAllSettings(blogID)
	settingMap := map[string]string{}
	for _, setting := range settings {
		settingMap[strings.Title(setting.Name)] = setting.Value
		settingMap[setting.Name] = setting.Value
	}
	settingMap[strings.Title(model.SettingNameSystemPath)] = util.PathBlogs + settingMap[strings.Title(model.SettingNameSystemPath)]
	settingMap[model.SettingNameSystemPath] = util.PathBlogs + settingMap[model.SettingNameSystemPath]

	(*dataModel)["Setting"] = settingMap

	statistics := service.Statistic.GetAllStatistics(blogID)
	statisticMap := map[string]int{}
	for _, statistic := range statistics {
		count, err := strconv.Atoi(statistic.Value)
		if nil != err {
			log.Errorf("statistic [%s] should be an integer, actual is [%v]", statistic.Name, statistic.Value)
		}
		statisticMap[strings.Title(statistic.Name)] = count
		statisticMap[statistic.Name] = count
	}
	(*dataModel)["Statistic"] = statisticMap
	(*dataModel)["Title"] = settingMap["BasicBlogTitle"]
	(*dataModel)["MetaKeywords"] = settingMap["BasicMetaKeywords"]
	(*dataModel)["FaviconURL"] = settingMap["BasicFaviconURL"]
	(*dataModel)["MetaDescription"] = settingMap["BasicMetaDescription"]
	(*dataModel)["Conf"] = util.Conf
	(*dataModel)["Year"] = time.Now().Year()
	(*dataModel)["BlogURL"] = util.Conf.Server + settingMap[model.SettingNameSystemPath]

	(*dataModel)["Username"] = ""
	session := util.GetSession(c)
	if nil != session {
		(*dataModel)["Username"] = session.UName
	}
	(*dataModel)["UserCount"] = len(service.User.GetBlogUsers(blogID)) + 2

	(*dataModel)["Navigations"] = service.Navigation.GetNavigations(blogID)

	fillMostUseCategories(&settingMap, dataModel, blogID)
	fillMostUseTags(&settingMap, dataModel, blogID)
	fillMostViewArticles(&settingMap, dataModel, blogID)
	fillRecentComments(&settingMap, dataModel, blogID)
	fillMostCommentArticles(&settingMap, dataModel, blogID)
	fillRandomArticles(&settingMap, dataModel, blogID)

	c.Set("dataModel", dataModel)
}

func fillMostUseCategories(settingMap *map[string]string, dataModel *DataModel, blogID uint) {
	categories := service.Category.GetCategories(math.MaxInt8, blogID)
	themeCategories := []*ThemeCategory{}
	for _, category := range categories {
		themeCategory := &ThemeCategory{
			Title: category.Title,
			URL:   util.Conf.Server + (*settingMap)[model.SettingNameSystemPath] + "/" + category.Title,
		}
		themeCategories = append(themeCategories, themeCategory)
	}
	(*dataModel)["MostUseCategories"] = themeCategories
}

func fillMostUseTags(settingMap *map[string]string, dataModel *DataModel, blogID uint) {
	tagSize, err := strconv.Atoi((*settingMap)[model.SettingNamePreferenceMostUseTagListSize])
	if nil != err {
		log.Errorf("setting [%s] should be an integer, actual is [%v]", model.SettingNamePreferenceMostUseTagListSize,
			(*settingMap)[model.SettingNamePreferenceMostUseTagListSize])
		tagSize = model.SettingPreferenceMostUseTagListSizeDefault
	}
	tags := service.Tag.GetTags(tagSize, blogID)
	themeTags := []*ThemeTag{}
	for _, tag := range tags {
		themeTag := &ThemeTag{
			Title: tag.Title,
			URL:   util.Conf.Server + (*settingMap)[model.SettingNameSystemPath] + "/" + tag.Title,
		}
		themeTags = append(themeTags, themeTag)
	}
	(*dataModel)["MostUseTags"] = themeTags
}

func fillMostViewArticles(settingMap *map[string]string, dataModel *DataModel, blogID uint) {
	mostViewArticleSize, err := strconv.Atoi((*settingMap)[model.SettingNamePreferenceMostViewArticleListSize])
	if nil != err {
		log.Errorf("setting [%s] should be an integer, actual is [%v]", model.SettingNamePreferenceMostViewArticleListSize,
			(*settingMap)[model.SettingNamePreferenceMostViewArticleListSize])
		mostViewArticleSize = model.SettingPreferenceMostViewArticleListSizeDefault
	}
	mostViewArticles := service.Article.GetMostViewArticles(mostViewArticleSize, blogID)
	themeMostViewArticles := []*ThemeArticle{}
	for _, article := range mostViewArticles {
		author := &ThemeAuthor{
			Name:      "Vanessa",
			URL:       "http://localhost:5879/blogs/solo/vanessa",
			AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
		}
		themeArticle := &ThemeArticle{
			Title:     article.Title,
			URL:       util.Conf.Server + (*settingMap)[model.SettingNameSystemPath] + article.Path,
			CreatedAt: humanize.Time(article.CreatedAt),
			Author:    author,
		}
		themeMostViewArticles = append(themeMostViewArticles, themeArticle)
	}

	(*dataModel)["MostViewArticles"] = themeMostViewArticles
}

func fillRecentComments(settingMap *map[string]string, dataModel *DataModel, blogID uint) {
	recentCommentSize, err := strconv.Atoi((*settingMap)[model.SettingNamePreferenceRecentCommentListSize])
	if nil != err {
		log.Errorf("setting [%s] should be an integer, actual is [%v]", model.SettingNamePreferenceRecentCommentListSize,
			(*settingMap)[model.SettingNamePreferenceRecentCommentListSize])
		recentCommentSize = model.SettingPreferenceRecentCommentListSizeDefault
	}
	recentComments := service.Comment.GetRecentComments(recentCommentSize, blogID)
	themeRecentComments := []*ThemeListComment{}
	for _, comment := range recentComments {
		themeComment := &ThemeListComment{
			Title:     comment.Content,
			Content:   comment.Content,
			URL:       "todo",
			CreatedAt: humanize.Time(comment.CreatedAt),
			Author: &ThemeAuthor{
				Name:      "Vanessa",
				URL:       "http://localhost:5879/blogs/solo/vanessa",
				AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
			},
		}
		themeRecentComments = append(themeRecentComments, themeComment)
	}

	(*dataModel)["RecentComments"] = themeRecentComments
}

func fillMostCommentArticles(settingMap *map[string]string, dataModel *DataModel, blogID uint) {
	mostCommentArticleSize, err := strconv.Atoi((*settingMap)[model.SettingNamePreferenceMostCommentArticleListSize])
	if nil != err {
		log.Errorf("setting [%s] should be an integer, actual is [%v]", model.SettingNamePreferenceMostCommentArticleListSize,
			(*settingMap)[model.SettingNamePreferenceMostCommentArticleListSize])
		mostCommentArticleSize = model.SettingPreferenceMostCommentArticleListSizeDefault
	}
	mostCommentArticles := service.Article.GetMostCommentArticles(mostCommentArticleSize, blogID)
	themeMostCommentArticles := []*ThemeArticle{}
	for _, article := range mostCommentArticles {
		author := &ThemeAuthor{
			Name:      "Vanessa",
			URL:       "http://localhost:5879/blogs/solo/vanessa",
			AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
		}
		themeArticle := &ThemeArticle{
			Title:     article.Title,
			URL:       util.Conf.Server + (*settingMap)[model.SettingNameSystemPath] + article.Path,
			CreatedAt: humanize.Time(article.CreatedAt),
			Author:    author,
		}
		themeMostCommentArticles = append(themeMostCommentArticles, themeArticle)
	}

	(*dataModel)["MostCommentArticles"] = themeMostCommentArticles
}

func fillRandomArticles(settingMap *map[string]string, dataModel *DataModel, blogID uint) {
	randomArticleSize, err := strconv.Atoi((*settingMap)[model.SettingNamePreferenceRandomArticleListSize])
	if nil != err {
		log.Errorf("setting [%s] should be an integer, actual is [%v]", model.SettingNamePreferenceRandomArticleListSize,
			(*settingMap)[model.SettingNamePreferenceRandomArticleListSize])
		randomArticleSize = model.SettingPreferenceRandomArticleListSizeDefault
	}
	randomArticles := service.Article.GetRandomArticles(randomArticleSize, blogID)
	themeRandomArticles := []*ThemeArticle{}
	for _, article := range randomArticles {
		author := &ThemeAuthor{
			Name:      "Vanessa",
			URL:       "http://localhost:5879/blogs/solo/vanessa",
			AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
		}
		themeArticle := &ThemeArticle{
			Title:     article.Title,
			URL:       util.Conf.Server + (*settingMap)[model.SettingNameSystemPath] + article.Path,
			CreatedAt: humanize.Time(article.CreatedAt),
			Author:    author,
		}
		themeRandomArticles = append(themeRandomArticles, themeArticle)
	}

	(*dataModel)["RandomArticles"] = themeRandomArticles
}
