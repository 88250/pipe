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
	"math"
	"net/http"
	"strings"

	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

func showCategoriesAction(c *gin.Context) {
	dataModel := getDataModel(c)
	blogID := getBlogID(c)
	categoryModels := service.Category.GetCategories(math.MaxInt8, blogID)
	themeCategories := []*ThemeCategory{}
	for _, categoryModel := range categoryModels {
		themeTags := []*ThemeTag{}
		tagStrs := strings.Split(categoryModel.Tags, ",")
		for _, tagTitle := range tagStrs {
			themeTag := &ThemeTag{
				Title: tagTitle,
				URL:   getBlogURL(c) + util.PathTags + "/" + tagTitle,
			}
			themeTags = append(themeTags, themeTag)
		}

		themeCategory := &ThemeCategory{
			Title:        categoryModel.Title,
			URL:          getBlogURL(c) + categoryModel.Path,
			Description:  categoryModel.Description,
			Tags:         themeTags,
			ArticleCount: 8,
		}
		themeCategories = append(themeCategories, themeCategory)
	}

	dataModel["Categories"] = themeCategories
	c.HTML(http.StatusOK, getTheme(c)+"/categories.html", dataModel)
}

func showCategoryArticlesArticlesAction(c *gin.Context) {
	dm, _ := c.Get("dataModel")
	dataModel := *(dm.(*DataModel))

	page := util.GetPage(c)
	blogID := getBlogID(c)
	articleModels, pagination := service.Article.GetArticles(page, blogID)
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
			logger.Errorf("not found author of article [id=%d, authorID=%d]", articleModel.ID, articleModel.AuthorID)

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

	dataModel["CategoryName"] = "JavaScript"
	dataModel["CategoryCount"] = 12

	c.HTML(http.StatusOK, getTheme(c)+"/category-articles.html", dataModel)
}
