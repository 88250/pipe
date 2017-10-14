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

// Package controller is the "controller" layer.
package controller

import (
	"net/http"
	"strings"

	"github.com/b3log/solo.go/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ThemeListArticle struct {
	ID           uint
	Abstract     string
	Author       *ThemeAuthor
	CreatedAt    string
	Title        string
	Tags         []*ThemeTag
	URL          string
	Topped       bool
	ViewCount    int
	CommentCount int
}

type ThemeTag struct {
	Title string
	URL   string
}

type ThemeAuthor struct {
	Name      string
	AvatarURL string
}

func showArticlesAction(c *gin.Context) {
	dataModel := DataModel{}
	fillCommon(c, &dataModel)

	page := c.GetInt("p")
	if 1 > page {
		page = 1
	}

	articleModels, pagination := service.Article.GetArticles(page, 1)
	articles := []*ThemeListArticle{}
	for _, articleModel := range articleModels {
		themeTags := []*ThemeTag{}
		tagStrs := strings.Split(articleModel.Tags, ",")
		for _, tagStr := range tagStrs {
			themeTag := &ThemeTag{
				Title: tagStr,
				URL:   "/todotagpath",
			}
			themeTags = append(themeTags, themeTag)
		}

		authorModel := service.User.GetUser(articleModel.AuthorID)
		if nil == authorModel {
			log.Errorf("not found author of article [id=%d, authorID=%d]", articleModel.ID, articleModel.AuthorID)

			continue
		}

		author := &ThemeAuthor{
			Name:      authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
		}

		article := &ThemeListArticle{
			ID:           articleModel.ID,
			Abstract:     articleModel.Abstract,
			Author:       author,
			CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
			Title:        articleModel.Title,
			Tags:         themeTags,
			URL:          "/todoarticlepath",
			Topped:       articleModel.Topped,
			ViewCount:    articleModel.ViewCount,
			CommentCount: articleModel.CommentCount,
		}

		articles = append(articles, article)
	}

	dataModel["articles"] = articles
	dataModel["pagination"] = pagination
	c.HTML(http.StatusOK, "index.html", dataModel)
}

func showArticleAction(c *gin.Context) {
	dataModel := DataModel{}
	fillCommon(c, &dataModel)
	c.HTML(http.StatusOK, "article.html", dataModel)
}
