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
	"strconv"
	"strings"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/vinta/pangu"
)

func showArticlesAction(c *gin.Context) {
	page := util.GetPage(c)
	dataModel := getDataModel(c)
	blogAdmin := getBlogAdmin(c)
	session := util.GetSession(c)
	articleModels, pagination := service.Article.GetArticles(page, blogAdmin.BlogID)
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
			URL:       getBlogURL(c) + util.PathAuthors + "/" + authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
		}

		abstract, thumb := util.MarkdownAbstract(articleModel.Content)
		article := &ThemeArticle{
			ID:           articleModel.ID,
			Abstract:     abstract,
			Author:       author,
			CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
			Title:        pangu.SpacingText(articleModel.Title),
			Tags:         themeTags,
			URL:          getBlogURL(c) + articleModel.Path,
			Topped:       articleModel.Topped,
			ViewCount:    articleModel.ViewCount,
			CommentCount: articleModel.CommentCount,
			ThumbnailURL: thumb,
			Editable:     session.UID == authorModel.ID,
		}

		articles = append(articles, article)
	}

	dataModel["Articles"] = articles
	dataModel["Pagination"] = pagination
	c.HTML(http.StatusOK, getTheme(c)+"/index.html", dataModel)
}

func showArticleAction(c *gin.Context) {
	dataModel := getDataModel(c)
	blogAdmin := getBlogAdmin(c)
	session := util.GetSession(c)

	a, _ := c.Get("article")
	article := a.(*model.Article)

	themeTags := []*ThemeTag{}
	tagStrs := strings.Split(article.Tags, ",")
	for _, tagStr := range tagStrs {
		themeTag := &ThemeTag{
			Title: tagStr,
			URL:   getBlogURL(c) + util.PathTags + "/" + tagStr,
		}
		themeTags = append(themeTags, themeTag)
	}

	authorModel := service.User.GetUser(article.AuthorID)
	dataModel["Article"] = &ThemeArticle{
		Author: &ThemeAuthor{
			Name:      authorModel.Name,
			URL:       getBlogURL(c) + util.PathAuthors + "/" + authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
		},
		ID:           article.ID,
		CreatedAt:    article.CreatedAt.Format("2006-01-02"),
		Title:        pangu.SpacingText(article.Title),
		Tags:         themeTags,
		URL:          getBlogURL(c) + article.Path,
		Topped:       article.Topped,
		ViewCount:    article.ViewCount,
		CommentCount: article.CommentCount,
		Content:      template.HTML(util.Markdown(article.Content)),
		Editable:     session.UID == authorModel.ID,
	}

	page := util.GetPage(c)
	commentModels, pagination := service.Comment.GetArticleComments(article.ID, page, blogAdmin.BlogID)
	comments := []*ThemeComment{}
	for _, commentModel := range commentModels {
		commentAuthor := service.User.GetUser(commentModel.AuthorID)
		if nil == commentAuthor {
			logger.Errorf("not found comment author [userID=%d]", commentModel.AuthorID)

			continue
		}
		commentAuthorURL := util.HacPaiURL + "/member/" + commentAuthor.Name
		blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, commentAuthor.BlogID)
		if nil != blogURLSetting {
			commentAuthorURL = blogURLSetting.Value + util.PathAuthors + "/" + commentAuthor.Name
		}

		author := &ThemeAuthor{
			Name:      commentAuthor.Name,
			URL:       commentAuthorURL,
			AvatarURL: commentAuthor.AvatarURLWithSize(64),
		}

		comment := &ThemeComment{
			ID:         commentModel.ID,
			Content:    template.HTML(util.Markdown(commentModel.Content)),
			Author:     author,
			CreatedAt:  commentModel.CreatedAt.Format("2006-01-02"),
			Removable:  session.UID == authorModel.ID,
			ReplyCount: service.Comment.GetRepliesCount(commentModel.ID, commentModel.BlogID),
		}
		if 0 != commentModel.ParentCommentID {
			parentCommentModel := service.Comment.GetComment(commentModel.ParentCommentID)
			parentCommentAuthorModel := service.User.GetUser(parentCommentModel.AuthorID)
			page := service.Comment.GetCommentPage(commentModel.ArticleID, commentModel.ID, commentModel.BlogID)

			parentComment := &ThemeComment{
				ID: parentCommentModel.ID,
				Author: &ThemeAuthor{
					Name: parentCommentAuthorModel.Name,
				},
				URL: getBlogURL(c) + article.Path + "?p=" + strconv.Itoa(page) + "#pipeComment" + strconv.Itoa(int(parentCommentModel.ID)),
			}
			comment.Parent = parentComment
		}

		comments = append(comments, comment)
	}

	dataModel["Comments"] = comments
	dataModel["Pagination"] = pagination

	articleModels, pagination := service.Article.GetArticles(1, blogAdmin.BlogID)
	articles := []*ThemeArticle{}
	for _, articleModel := range articleModels {
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
			Content:      "sfasdfsf",
			Editable:     false,
		}

		articles = append(articles, article)
	}
	dataModel["RelevantArticles"] = articles
	fillPreviousArticle(c, article, &dataModel)
	fillNextArticle(c, article, &dataModel)
	dataModel["ToC"] = "todo"
	c.HTML(http.StatusOK, getTheme(c)+"/article.html", dataModel)

	service.Article.IncArticleViewCount(article)
}

func fillPreviousArticle(c *gin.Context, article *model.Article, dataModel *DataModel) {
	previous := service.Article.GetPreviousArticle(article.ID, article.BlogID)
	if nil == previous {
		return
	}

	previousArticle := &ThemeArticle{
		Title: previous.Title,
		URL:   getBlogURL(c) + previous.Path,
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
		URL:   getBlogURL(c) + next.Path,
	}
	(*dataModel)["NextArticle"] = nextArticle
}
