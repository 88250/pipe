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

// Package console defines console controllers.
package console

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/b3log/pipe/log"
	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

// Logger
var logger = log.NewLogger(os.Stdout)

func MarkdownAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "parses markdown request failed"

		return
	}

	mdText := arg["mdText"].(string)
	mdResult := util.Markdown(mdText)
	data := map[string]interface{}{}
	data["html"] = mdResult.ContentHTML
	result.Data = data
}

func AddArticleAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)

	article := &model.Article{}
	if err := c.BindJSON(article); nil != err {
		result.Code = -1
		result.Msg = "parses add article request failed"

		return
	}

	article.IP = util.GetRemoteAddr(c)
	article.BlogID = session.BID
	article.AuthorID = session.UID

	if err := service.Article.AddArticle(article); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

func GetArticleAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Code = -1

		return
	}

	data := service.Article.ConsoleGetArticle(uint64(id))
	if nil == data {
		result.Code = -1

		return
	}

	result.Data = data
}

func GetArticlesAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	articleModels, pagination := service.Article.ConsoleGetArticles(c.Query("key"), util.GetPage(c), session.BID)
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, session.BID)

	var articles []*ConsoleArticle
	for _, articleModel := range articleModels {
		var consoleTags []*ConsoleTag
		tagStrs := strings.Split(articleModel.Tags, ",")
		for _, tagStr := range tagStrs {
			consoleTag := &ConsoleTag{
				Title: tagStr,
				URL:   blogURLSetting.Value + util.PathTags + "/" + tagStr,
			}
			consoleTags = append(consoleTags, consoleTag)
		}

		authorModel := service.User.GetUser(articleModel.AuthorID)
		author := &ConsoleAuthor{
			Name:      authorModel.Name,
			URL:       blogURLSetting.Value + util.PathAuthors + "/" + authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
		}

		article := &ConsoleArticle{
			ID:           articleModel.ID,
			Author:       author,
			CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
			Title:        articleModel.Title,
			Tags:         consoleTags,
			URL:          blogURLSetting.Value + articleModel.Path,
			Topped:       articleModel.Topped,
			ViewCount:    articleModel.ViewCount,
			CommentCount: articleModel.CommentCount,
		}

		articles = append(articles, article)
	}

	data := map[string]interface{}{}
	data["articles"] = articles
	data["pagination"] = pagination
	result.Data = data
}

func RemoveArticleAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}

	session := util.GetSession(c)
	blogID := session.BID

	if err := service.Article.RemoveArticle(id, blogID); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

func RemoveArticlesAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "parses batch remove articles request failed"

		return
	}

	session := util.GetSession(c)
	blogID := session.BID

	ids := arg["ids"].([]interface{})
	for _, id := range ids {
		if err := service.Article.RemoveArticle(uint64(id.(float64)), blogID); nil != err {
			logger.Errorf("remove article failed: " + err.Error())
		}
	}
}

func UpdateArticleAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}

	article := &model.Article{Model: model.Model{ID: id}}
	if err := c.BindJSON(article); nil != err {
		result.Code = -1
		result.Msg = "parses update article request failed"

		return
	}

	article.IP = util.GetRemoteAddr(c)
	session := util.GetSession(c)
	article.BlogID = session.BID

	if err := service.Article.UpdateArticle(article); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

func GetArticleThumbsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	n, _ := strconv.Atoi(c.Query("n"))
	urls := util.RandImages(n)

	// original: 1920*1080

	w, _ := strconv.Atoi(c.Query("w"))
	if w < 1 {
		w = 960
	}
	h, _ := strconv.Atoi(c.Query("h"))
	if h < 1 {
		h = 520
	}

	var styledURLs []string
	for _, url := range urls {
		styledURLs = append(styledURLs, url+"?imageView2/1/w/"+strconv.Itoa(w)+
			"/h/"+strconv.Itoa(h)+"/interlace/1/q/100")
	}

	result.Data = styledURLs
}
