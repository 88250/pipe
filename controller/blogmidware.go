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
	"html/template"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/88250/pipe/i18n"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
)

// DataModel represents data model.
type DataModel map[string]interface{}

func fillUser(c *gin.Context) {
	inited := service.Init.Inited()
	if !inited && util.PathInit != c.Request.URL.Path {
		c.Redirect(http.StatusSeeOther, model.Conf.Server+util.PathInit)
		c.Abort()

		return
	}

	dataModel := &DataModel{}
	c.Set("dataModel", dataModel)
	session := util.GetSession(c)
	(*dataModel)["User"] = session

	c.Next()
}

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
	go service.Statistic.IncViewCount(userBlog.ID)

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
	if "dev" == model.Conf.RuntimeMode {
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
		v := setting.Value
		if model.SettingNameBasicHeader == setting.Name || model.SettingNameBasicFooter == setting.Name || model.SettingNameBasicNoticeBoard == setting.Name {
			mdResult := util.Markdown(v)
			v = mdResult.ContentHTML
			v = strings.TrimSpace(v)
			v = strings.TrimPrefix(v, "<p>")
			v = strings.TrimSuffix(v, "</p>")
			v = strings.TrimSpace(v)
		}

		settingMap[strings.Title(setting.Name)] = v
		settingMap[setting.Name] = v
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
	(*dataModel)["Conf"] = model.Conf
	(*dataModel)["Year"] = time.Now().Year()
	users, _ := service.User.GetBlogUsers(1, blogID)
	(*dataModel)["UserCount"] = len(users)
	(*dataModel)["BlogAdmin"] = service.User.GetBlogAdmin(blogID)
	(*dataModel)["Navigations"] = service.Navigation.GetNavigations(blogID)

	fillMostUseCategories(&settingMap, dataModel, blogID)
	fillMostUseTags(&settingMap, dataModel, blogID)
	fillMostViewArticles(c, &settingMap, dataModel, blogID)
	fillRecentComments(c, &settingMap, dataModel, blogID)
	fillMostCommentArticles(c, &settingMap, dataModel, blogID)

	c.Set("dataModel", dataModel)
}

func fillMostUseCategories(settingMap *map[string]interface{}, dataModel *DataModel, blogID uint64) {
	categories := service.Category.GetCategories(math.MaxInt8, blogID)
	var themeCategories []*model.ThemeCategory
	for _, category := range categories {
		themeCategory := &model.ThemeCategory{
			Title: category.Title,
			URL:   (*settingMap)[model.SettingNameBasicBlogURL].(string) + util.PathCategories + category.Path,
		}
		themeCategories = append(themeCategories, themeCategory)
	}
	(*dataModel)["MostUseCategories"] = themeCategories
}

func fillMostUseTags(settingMap *map[string]interface{}, dataModel *DataModel, blogID uint64) {
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

func fillMostViewArticles(c *gin.Context, settingMap *map[string]interface{}, dataModel *DataModel, blogID uint64) {
	mostViewArticleSize, err := strconv.Atoi((*settingMap)[model.SettingNamePreferenceMostViewArticleListSize].(string))
	if nil != err {
		logger.Errorf("setting [%s] should be an integer, actual is [%v]", model.SettingNamePreferenceMostViewArticleListSize,
			(*settingMap)[model.SettingNamePreferenceMostViewArticleListSize])
		mostViewArticleSize = model.SettingPreferenceMostViewArticleListSizeDefault
	}
	mostViewArticles := service.Article.GetMostViewArticles(mostViewArticleSize, blogID)
	var themeMostViewArticles []*model.ThemeArticle
	for _, article := range mostViewArticles {
		authorModel := service.User.GetUser(article.AuthorID)
		if nil == authorModel {
			logger.Errorf("not found author of article [id=%d, authorID=%d]", article.ID, article.AuthorID)

			continue
		}
		author := &model.ThemeAuthor{
			Name:      authorModel.Name,
			URL:       getBlogURL(c) + util.PathAuthors + "/" + authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
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

func fillRecentComments(c *gin.Context, settingMap *map[string]interface{}, dataModel *DataModel, blogID uint64) {
	recentCommentSize, err := strconv.Atoi((*settingMap)[model.SettingNamePreferenceRecentCommentListSize].(string))
	if nil != err {
		logger.Errorf("setting [%s] should be an integer, actual is [%v]", model.SettingNamePreferenceRecentCommentListSize,
			(*settingMap)[model.SettingNamePreferenceRecentCommentListSize])
		recentCommentSize = model.SettingPreferenceRecentCommentListSizeDefault
	}
	recentComments := service.Comment.GetRecentComments(recentCommentSize, blogID)
	var themeRecentComments []*model.ThemeComment
	for _, comment := range recentComments {
		author := &model.ThemeAuthor{}
		if model.SyncCommentAuthorID == comment.AuthorID {
			author.URL = comment.AuthorURL
			author.Name = comment.AuthorName
			author.AvatarURL = comment.AuthorAvatarURL
		} else {
			commentAuthor := service.User.GetUser(comment.AuthorID)
			commentAuthorBlog := service.User.GetOwnBlog(comment.AuthorID)
			author.URL = service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, commentAuthorBlog.ID).Value + util.PathAuthors + "/" + commentAuthor.Name
			author.Name = commentAuthor.Name
			author.AvatarURL = commentAuthor.AvatarURL
		}

		page := service.Comment.GetCommentPage(comment.ArticleID, comment.ID, blogID)
		article := service.Article.ConsoleGetArticle(comment.ArticleID)

		title := util.Markdown(comment.Content).AbstractText
		if "" == title {
			continue
		}
		themeComment := &model.ThemeComment{
			Title:     title,
			URL:       getBlogURL(c) + article.Path + "?p=" + strconv.Itoa(page) + "#pipeComment" + strconv.Itoa(int(comment.ID)),
			CreatedAt: humanize.Time(comment.CreatedAt),
			Author:    author,
		}
		themeRecentComments = append(themeRecentComments, themeComment)
	}

	(*dataModel)["RecentComments"] = themeRecentComments
}

func fillMostCommentArticles(c *gin.Context, settingMap *map[string]interface{}, dataModel *DataModel, blogID uint64) {
	mostCommentArticleSize, err := strconv.Atoi((*settingMap)[model.SettingNamePreferenceMostCommentArticleListSize].(string))
	if nil != err {
		logger.Errorf("setting [%s] should be an integer, actual is [%v]", model.SettingNamePreferenceMostCommentArticleListSize,
			(*settingMap)[model.SettingNamePreferenceMostCommentArticleListSize])
		mostCommentArticleSize = model.SettingPreferenceMostCommentArticleListSizeDefault
	}
	mostCommentArticles := service.Article.GetMostCommentArticles(mostCommentArticleSize, blogID)
	var themeMostCommentArticles []*model.ThemeArticle
	for _, article := range mostCommentArticles {
		authorModel := service.User.GetUser(article.AuthorID)
		if nil == authorModel {
			logger.Errorf("not found author of article [id=%d, authorID=%d]", article.ID, article.AuthorID)

			continue
		}
		author := &model.ThemeAuthor{
			Name:      authorModel.Name,
			URL:       getBlogURL(c) + util.PathAuthors + "/" + authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
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

func getBlogID(c *gin.Context) uint64 {
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
	t, err := template.ParseFiles("console/dist/start/index.html")
	if nil != err {
		logger.Errorf("load 404 page failed: " + err.Error())
		c.String(http.StatusNotFound, "load 404 page failed")

		return
	}

	c.Status(http.StatusNotFound)
	t.Execute(c.Writer, nil)
}
