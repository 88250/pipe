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

// Package controller is the "controller" layer.
package controller

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func showArticlesAction(c *gin.Context) {
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
			URL:       "http://localhost:5879/blogs/solo/vanessa",
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

	dataModel := getDataModel(c)

	dataModel["Articles"] = articles
	dataModel["Pagination"] = pagination
	c.HTML(http.StatusOK, "index.html", dataModel)
}

func showArticleAction(c *gin.Context) {
	dataModel := getDataModel(c)
	blogAdmin := getBlogAdmin(c)

	a, _ := c.Get("article")
	article := a.(*model.Article)

	themeTags := []*ThemeTag{}
	tagStrs := strings.Split(article.Tags, ",")
	for _, tagStr := range tagStrs {
		themeTag := &ThemeTag{
			Title: tagStr,
			URL:   getSystemPath(c) + util.PathTags + "/" + tagStr,
		}
		themeTags = append(themeTags, themeTag)
	}

	dataModel["Article"] = &ThemeArticle{
		Author: &ThemeAuthor{
			Name:      "Vanessa",
			URL:       "http://localhost:5879/blogs/solo/vanessa",
			AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
		},
		CreatedAt:    article.CreatedAt.Format("2006-01-02"),
		Title:        article.Title,
		Tags:         themeTags,
		URL:          getSystemPath(c) + article.Path,
		Topped:       article.Topped,
		ViewCount:    article.ViewCount,
		CommentCount: article.CommentCount,
		Content:      template.HTML(util.Markdown(article.Content)),
		Editable:     true,
	}

	articleModels, pagination := service.Article.GetArticles(1, blogAdmin.BlogID)
	articles := []*ThemeArticle{}
	for _, articleModel := range articleModels {
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
			URL:       "http://localhost:5879/blogs/solo/vanessa",
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
			Content:      "sfasdfsf",
			Editable:     false,
		}

		articles = append(articles, article)
	}

	page := c.GetInt("p")
	if 1 > page {
		page = 1
	}
	commentModels, pagination := service.Comment.GetArticleComments(article.ID, page, blogAdmin.BlogID)
	comments := []*ThemeComment{}
	for _, commentModel := range commentModels {
		author := &ThemeAuthor{
			Name:      "test name",
			URL:       "http://localhost:5879/blogs/solo/vanessa",
			AvatarURL: "https://img.hacpai.com/20170818zhixiaoyun.jpeg",
		}

		comment := &ThemeComment{
			ID:        commentModel.ID,
			Content:   template.HTML(util.Markdown(commentModel.Content)),
			Author:    author,
			Removable: false,
		}

		comments = append(comments, comment)
	}

	dataModel["Comments"] = comments
	dataModel["Pagination"] = pagination

	dataModel["RelevantArticles"] = articles
	dataModel["ExternalRelevantArticles"] = articles
	fillPreviousArticle(c, article, &dataModel)
	fillNextArticle(c, article, &dataModel)
	dataModel["ToC"] = template.HTML("<ul id='toc' class='toc'><li class='toc__1'><a data-id='toc_1_0'>ToC</a></li><li class='toc__2'><a data-id='toc_2_0'>ToC</a></li><li class='toc__3'><a data-id='toc_3_0'>ToC</a></li><li class='toc__4'><a data-id='toc_4_0'>ToC</a></li><li class='toc__5'><a data-id='toc_5_0'>ToC</a></li><li class='toc__6'><a data-id='toc_6_0'>ToC</a></li><li class='toc__6'><a data-id='toc_6_1'>ToC</a></li></ul>")
	c.HTML(http.StatusOK, "article.html", dataModel)
}

func fillPreviousArticle(c *gin.Context, article *model.Article, dataModel *DataModel) {
	previous := service.Article.GetPreviousArticle(article.ID, article.BlogID)
	if nil == previous {
		return
	}

	previousArticle := &ThemeArticle{
		Title: previous.Title,
		URL:   getSystemPath(c) + previous.Path,
	}
	(*dataModel)["PreviousArticle"] = previousArticle
}

func fillNextArticle(c *gin.Context, article *model.Article, dataModel *DataModel) {
	next := service.Article.GetNextArticle(article.ID, article.BlogID)
	if nil == next {
		return
	}

	nextArticle := &ThemeArticle{
		Title: next.Title,
		URL:   getSystemPath(c) + next.Path,
	}
	(*dataModel)["NextArticle"] = nextArticle
}
