// Pipe - A small and beautiful blogging platform written in golang.
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
	"html/template"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/b3log/pipe/i18n"
	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
)

func resolveBlog(c *gin.Context) {
	username := c.Param("username")
	if "" == username {
		notFound(c)
		c.Abort()

		return
	}
	user := service.User.GetUserByName(username)
	if nil == user {
		notFound(c)
		c.Abort()

		return
	}
	userBlog := service.User.GetOwnBlog(user.ID)
	if nil == userBlog {
		notFound(c)
		c.Abort()

		return
	}
	c.Set("userBlog", userBlog)

	fillCommon(c)
	service.Statistic.IncViewCount(userBlog.ID)

	path := strings.Split(c.Request.RequestURI, username)[1]
	path = strings.TrimSpace(path)
	if end := strings.Index(path, "?"); 0 < end {
		path = path[:end]
	}
	article := service.Article.GetArticleByPath(path, userBlog.ID)
	if nil == article {
		c.Next()

		return
	}

	c.Set("article", article)
	showArticleAction(c)
	c.Abort()
}

func fillCommon(c *gin.Context) {
	if "dev" == util.Conf.RuntimeMode {
		i18n.Load()
	}

	userBlogVal, _ := c.Get("userBlog")
	userBlog := userBlogVal.(*service.UserBlog)
	blogID := userBlog.ID

	dataModelVal, _ := c.Get("dataModel")
	dataModel := dataModelVal.(*DataModel)

	localeSetting := service.Setting.GetSetting(model.SettingCategoryI18n, model.SettingNameI18nLocale, blogID)
	i18ns := i18n.GetMessages(localeSetting.Value)
	i18nMap := map[string]interface{}{}
	for key, value := range i18ns {
		i18nMap[strings.Title(key)] = value
		i18nMap[key] = value
	}
	(*dataModel)["I18n"] = i18nMap

	settings := service.Setting.GetAllSettings(blogID)
	settingMap := map[string]interface{}{}
	for _, setting := range settings {
		settingMap[strings.Title(setting.Name)] = setting.Value
		settingMap[setting.Name] = setting.Value
	}
	settingMap[strings.Title(model.SettingNameBasicHeader)] = template.HTML(settingMap[model.SettingNameBasicHeader].(string))
	settingMap[strings.Title(model.SettingNameBasicFooter)] = template.HTML(settingMap[model.SettingNameBasicFooter].(string))
	settingMap[strings.Title(model.SettingNameBasicNoticeBoard)] = template.HTML(settingMap[model.SettingNameBasicNoticeBoard].(string))
	settingMap[strings.Title(model.SettingNameArticleSign)] = template.HTML(settingMap[model.SettingNameArticleSign].(string))
	(*dataModel)["Setting"] = settingMap

	statistics := service.Statistic.GetAllStatistics(blogID)
	statisticMap := map[string]int{}
	for _, statistic := range statistics {
		count, err := strconv.Atoi(statistic.Value)
		if nil != err {
			logger.Errorf("statistic [%s] should be an integer, actual is [%v]", statistic.Name, statistic.Value)
		}
		statisticMap[strings.Title(statistic.Name)] = count
		statisticMap[statistic.Name] = count
	}
	(*dataModel)["Statistic"] = statisticMap
	(*dataModel)["FaviconURL"] = settingMap[model.SettingNameBasicFaviconURL]
	(*dataModel)["LogoURL"] = settingMap[model.SettingNameBasicLogoURL]
	(*dataModel)["BlogURL"] = settingMap[model.SettingNameBasicBlogURL]
	(*dataModel)["Title"] = settingMap[model.SettingNameBasicBlogTitle]
	(*dataModel)["MetaKeywords"] = settingMap[model.SettingNameBasicMetaKeywords]
	(*dataModel)["MetaDescription"] = settingMap[model.SettingNameBasicMetaDescription]
	(*dataModel)["Conf"] = util.Conf
	(*dataModel)["Year"] = time.Now().Year()
	users, _ := service.User.GetBlogUsers(1, blogID)
	(*dataModel)["UserCount"] = len(users)

	(*dataModel)["Navigations"] = service.Navigation.GetNavigations(blogID)

	fillMostUseCategories(&settingMap, dataModel, blogID)
	fillMostUseTags(&settingMap, dataModel, blogID)
	fillMostViewArticles(&settingMap, dataModel, blogID)
	fillRecentComments(&settingMap, dataModel, blogID)
	fillMostCommentArticles(&settingMap, dataModel, blogID)

	c.Set("dataModel", dataModel)
}

func fillMostUseCategories(settingMap *map[string]interface{}, dataModel *DataModel, blogID uint) {
	categories := service.Category.GetCategories(math.MaxInt8, blogID)
	var themeCategories []*model.ThemeCategory
	for _, category := range categories {
		themeCategory := &model.ThemeCategory{
			Title: category.Title,
			URL:   (*settingMap)[model.SettingNameBasicBlogURL].(string) + "/" + category.Title,
		}
		themeCategories = append(themeCategories, themeCategory)
	}
	(*dataModel)["MostUseCategories"] = themeCategories
}

func fillMostUseTags(settingMap *map[string]interface{}, dataModel *DataModel, blogID uint) {
	tagSize, err := strconv.Atoi((*settingMap)[model.SettingNamePreferenceMostUseTagListSize].(string))
	if nil != err {
		logger.Errorf("setting [%s] should be an integer, actual is [%v]", model.SettingNamePreferenceMostUseTagListSize,
			(*settingMap)[model.SettingNamePreferenceMostUseTagListSize])
		tagSize = model.SettingPreferenceMostUseTagListSizeDefault
	}
	tags := service.Tag.GetTags(tagSize, blogID)
	var themeTags []*model.ThemeTag
	for _, tag := range tags {
		themeTag := &model.ThemeTag{
			Title: tag.Title,
			URL:   (*settingMap)[model.SettingNameBasicBlogURL].(string) + "/tags/" + tag.Title,
		}
		themeTags = append(themeTags, themeTag)
	}
	(*dataModel)["MostUseTags"] = themeTags
}

func fillMostViewArticles(settingMap *map[string]interface{}, dataModel *DataModel, blogID uint) {
	mostViewArticleSize, err := strconv.Atoi((*settingMap)[model.SettingNamePreferenceMostViewArticleListSize].(string))
	if nil != err {
		logger.Errorf("setting [%s] should be an integer, actual is [%v]", model.SettingNamePreferenceMostViewArticleListSize,
			(*settingMap)[model.SettingNamePreferenceMostViewArticleListSize])
		mostViewArticleSize = model.SettingPreferenceMostViewArticleListSizeDefault
	}
	mostViewArticles := service.Article.GetMostViewArticles(mostViewArticleSize, blogID)
	var themeMostViewArticles []*model.ThemeArticle
	for _, article := range mostViewArticles {
		author := &model.ThemeAuthor{
			Name:      "Vanessa",
			URL:       "http://localhost:5879/blogs/pipe/vanessa",
			AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
		}
		themeArticle := &model.ThemeArticle{
			Title:     article.Title,
			URL:       (*settingMap)[model.SettingNameBasicBlogURL].(string) + article.Path,
			CreatedAt: humanize.Time(article.CreatedAt),
			Author:    author,
		}
		themeMostViewArticles = append(themeMostViewArticles, themeArticle)
	}

	(*dataModel)["MostViewArticles"] = themeMostViewArticles
}

func fillRecentComments(settingMap *map[string]interface{}, dataModel *DataModel, blogID uint) {
	recentCommentSize, err := strconv.Atoi((*settingMap)[model.SettingNamePreferenceRecentCommentListSize].(string))
	if nil != err {
		logger.Errorf("setting [%s] should be an integer, actual is [%v]", model.SettingNamePreferenceRecentCommentListSize,
			(*settingMap)[model.SettingNamePreferenceRecentCommentListSize])
		recentCommentSize = model.SettingPreferenceRecentCommentListSizeDefault
	}
	recentComments := service.Comment.GetRecentComments(recentCommentSize, blogID)
	var themeRecentComments []*model.ThemeComment
	for _, comment := range recentComments {
		themeComment := &model.ThemeComment{
			Title:     util.Markdown(comment.Content).AbstractText,
			Content:   "",
			URL:       "todo",
			CreatedAt: humanize.Time(comment.CreatedAt),
			Author: &model.ThemeAuthor{
				Name:      "Vanessa",
				URL:       "http://localhost:5879/blogs/pipe/vanessa",
				AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
			},
		}
		themeRecentComments = append(themeRecentComments, themeComment)
	}

	(*dataModel)["RecentComments"] = themeRecentComments
}

func fillMostCommentArticles(settingMap *map[string]interface{}, dataModel *DataModel, blogID uint) {
	mostCommentArticleSize, err := strconv.Atoi((*settingMap)[model.SettingNamePreferenceMostCommentArticleListSize].(string))
	if nil != err {
		logger.Errorf("setting [%s] should be an integer, actual is [%v]", model.SettingNamePreferenceMostCommentArticleListSize,
			(*settingMap)[model.SettingNamePreferenceMostCommentArticleListSize])
		mostCommentArticleSize = model.SettingPreferenceMostCommentArticleListSizeDefault
	}
	mostCommentArticles := service.Article.GetMostCommentArticles(mostCommentArticleSize, blogID)
	var themeMostCommentArticles []*model.ThemeArticle
	for _, article := range mostCommentArticles {
		author := &model.ThemeAuthor{
			Name:      "Vanessa",
			URL:       "http://localhost:5879/blogs/pipe/vanessa",
			AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
		}
		themeArticle := &model.ThemeArticle{
			Title:     article.Title,
			URL:       (*settingMap)[model.SettingNameBasicBlogURL].(string) + article.Path,
			CreatedAt: humanize.Time(article.CreatedAt),
			Author:    author,
		}
		themeMostCommentArticles = append(themeMostCommentArticles, themeArticle)
	}

	(*dataModel)["MostCommentArticles"] = themeMostCommentArticles
}

func getBlogURL(c *gin.Context) string {
	dataModel := getDataModel(c)

	return dataModel["Setting"].(map[string]interface{})[model.SettingNameBasicBlogURL].(string)
}

func getBlogID(c *gin.Context) uint {
	userBlogVal, _ := c.Get("userBlog")

	return userBlogVal.(*service.UserBlog).ID
}

func getTheme(c *gin.Context) string {
	dataModel := getDataModel(c)

	return dataModel["Setting"].(map[string]interface{})[model.SettingNameThemeName].(string)
}

func getDataModel(c *gin.Context) DataModel {
	dataModelVal, _ := c.Get("dataModel")

	return *(dataModelVal.(*DataModel))
}

func getLocale(c *gin.Context) string {
	dataModel := getDataModel(c)

	return dataModel["Setting"].(map[string]interface{})[model.SettingNameI18nLocale].(string)
}

func notFound(c *gin.Context) {
	t, err := template.ParseFiles("console/dist/init/index.html")
	if nil != err {
		logger.Errorf("load 404 page failed: " + err.Error())
		c.String(http.StatusNotFound, "load 404 page failed")

		return
	}

	c.Status(http.StatusNotFound)
	t.Execute(c.Writer, nil)
}
