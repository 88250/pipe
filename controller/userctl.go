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

	"github.com/gin-gonic/gin"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"strings"
	log "github.com/sirupsen/logrus"
)

func showAuthorsAction(c *gin.Context) {
	dm, _ := c.Get("dataModel")
	dataModel := *(dm.(*DataModel))

	themeAuthorDetail := []*ThemeAuthorDetail{}
	AuthorDetailModels := strings.Split("a, g, c, d", ",")
	for _, authorDetailModel := range AuthorDetailModels {
		authorDetail := &ThemeAuthorDetail{
			Name:      authorDetailModel,
			URL:       "/sss",
			Count:     13,
			AvatarURL: "http://themedesigner.in/demo/admin-press/assets/images/users/2.jpg",
			CreatedAt: "2012-12-12",
		}
		themeAuthorDetail = append(themeAuthorDetail, authorDetail)
	}

	dataModel["Authors"] = themeAuthorDetail
	c.HTML(http.StatusOK, "authors.html", dataModel)
}

func showAuthorArticlesAction(c *gin.Context) {
	dm, _ := c.Get("dataModel")
	dataModel := *(dm.(*DataModel))

	page := c.GetInt("p")
	if 1 > page {
		page = 1
	}
	blogAdmin := getBlogAdmin(c)
	articleModels, pagination := service.Article.GetArticles(page, blogAdmin.BlogID)
	articles := []*ThemeArticle{}
	for _, articleModel := range articleModels {
		themeTags := []*ThemeTag{}
		tagStrs := strings.Split(articleModel.Tags, ",")
		for _, tagStr := range tagStrs {
			themeTag := &ThemeTag{
				Title: tagStr,
				URL:   getSystemPath(c) + util.PathTags + "/" + tagStr,
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
			URL:          getSystemPath(c) + articleModel.Path,
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

	dataModel["AuthorName"] = "Vanessa"
	dataModel["AuthorCount"] = 12

	c.HTML(http.StatusOK, "author-articles.html", dataModel)
}
