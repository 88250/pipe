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

package console

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AddArticleCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	article := &model.Article{}
	if err := c.BindJSON(article); nil != err {
		result.Code = -1
		result.Msg = "parses add article request failed"

		return
	}

	if err := service.Article.AddArticle(article); nil != err {
		log.Error("add article failed: " + err.Error())
		result.Code = -1
	}
}

func GetArticleCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.Atoi(idArg)
	if nil != err {
		result.Code = -1

		return
	}

	data := service.Article.ConsoleGetArticle(uint(id))
	if nil == data {
		result.Code = -1

		return
	}

	result.Data = data
}

type ConsoleArticle struct {
	ID           uint            `json:"id"`
	Author       Author          `json:"author"`
	CreatedAt    string          `json:"createdAt"`
	Title        string          `gorm:"size:128" json:"title"`
	Tags         []*TagPermalink `json:"tags"`
	Permalink    string          `json:"permalink"`
	Topped       bool            `json:"topped"`
	ViewCount    int             `json:"viewCount"`
	CommentCount int             `json:"commentCount"`
}

type TagPermalink struct {
	Title     string `json:"title"`
	Permalink string `json:"permalink"`
}

type Author struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatarURL"`
}

func GetArticlesCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	articleModels, pagination := service.Article.ConsoleGetArticles(c.GetInt("p"))

	articles := []ConsoleArticle{}
	for _, articleModel := range articleModels {
		tagPermalinks := []*TagPermalink{}
		tagStrs := strings.Split(articleModel.Tags, ",")
		for _, tagStr := range tagStrs {
			tagPermalink := &TagPermalink{
				Title:     tagStr,
				Permalink: "context/" + tagStr,
			}
			tagPermalinks = append(tagPermalinks, tagPermalink)
		}

		article := ConsoleArticle{
			ID:           articleModel.ID,
			CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
			Title:        articleModel.Title,
			Tags:         tagPermalinks,
			Permalink:    articleModel.Permalink,
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

func RemoveArticleCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.Atoi(idArg)
	if nil != err {
		result.Code = -1

		return
	}

	if err := service.Article.RemoveArticle(uint(id)); nil != err {
		log.Error("remove article failed: " + err.Error())
		result.Code = -1
	}
}

func UpdateArticleCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.Atoi(idArg)
	if nil != err {
		result.Code = -1

		return
	}

	article := &model.Article{Model: model.Model{ID: uint(id)}}
	if err := c.BindJSON(article); nil != err {
		result.Code = -1
		result.Msg = "parses update article request failed"

		return
	}

	if err := service.Article.UpdateArticle(article); nil != err {
		log.Error("update article failed: " + err.Error())
		result.Code = -1
	}
}
