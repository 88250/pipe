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
	"net/http"

	"strings"

	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func showTagsAction(c *gin.Context) {
	dataModel := getDataModel(c)
	blogAdmin := getBlogAdmin(c)
	tags := service.Tag.ConsoleGetTags(blogAdmin.BlogID)
	themeTags := []*ThemeTag{}
	for _, tag := range tags {
		themeTag := &ThemeTag{
			Title: tag.Title,
			URL:   getBlogURL(c) + util.PathTags + "/" + tag.Title,
			Count: tag.ArticleCount,
		}
		themeTags = append(themeTags, themeTag)
	}
	dataModel["Tags"] = themeTags

	c.HTML(http.StatusOK, getTheme(c)+"/tags.html", dataModel)
}

func showTagArticlesAction(c *gin.Context) {
	page := c.GetInt("p")
	if 1 > page {
		page = 1
	}
	dataModel := getDataModel(c)
	tagTitle := strings.SplitAfter(c.Request.URL.Path, util.PathTags+"/")[1]
	blogAdmin := getBlogAdmin(c)
	articleModels, pagination := service.Article.GetTagArticles(tagTitle, page, blogAdmin.BlogID)
	articles := []*ThemeArticle{}
	for _, articleModel := range articleModels {
		themeTags := []*ThemeTag{}
		tagStrs := strings.Split(articleModel.Tags, ",")
		for _, tagStr := range tagStrs {
			themeTag := &ThemeTag{
				Title: tagStr,
				URL:   getBlogURL(c) + util.PathTags + "/" + tagStr,
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
			URL:       "http://localhost:5879/blogs/pipe/vanessa",
			AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
		}

		article := &ThemeArticle{
			ID:           articleModel.ID,
			Author:       author,
			CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
			Title:        articleModel.Title,
			Tags:         themeTags,
			URL:          getBlogURL(c) + articleModel.Path,
			Topped:       articleModel.Topped,
			ViewCount:    articleModel.ViewCount,
			CommentCount: articleModel.CommentCount,
			ThumbnailURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
			Editable:     false,
		}

		articles = append(articles, article)
	}
	dataModel["Articles"] = articles
	dataModel["Pagination"] = pagination

	dataModel["TagName"] = "Vanessa"
	dataModel["TagArticlesCount"] = 12

	c.HTML(http.StatusOK, getTheme(c)+"/tag-articles.html", dataModel)
}
