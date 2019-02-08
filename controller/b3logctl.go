// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2019, b3log.org & hacpai.com
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
	"fmt"
	"net/http"
	"strconv"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

// addSymCommentAction adds a comment come from Sym. Sees https://hacpai.com/article/1457158841475 for more details.
func addSymCommentAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "parses add comment request failed"

		return
	}

	client := arg["client"].(map[string]interface{})
	b3Key := client["userB3Key"].(string)
	articleAuthorName := client["userName"].(string)
	articleAuthor := service.User.GetUserByName(articleAuthorName)
	if articleAuthor.B3Key != b3Key {
		result.Code = -1
		result.Msg = "wrong B3 Key"

		return
	}

	requestCmt := arg["comment"].(map[string]interface{})
	articleId, err := strconv.ParseUint(requestCmt["articleId"].(string), 10, 64)
	if nil != err {
		result.Code = -1
		result.Msg = "parses add comment request failed"

		return
	}

	blogID := getBlogID(c)
	comment := &model.Comment{
		BlogID:          blogID,
		ArticleID:       articleId,
		AuthorID:        model.SyncCommentAuthorID,
		Content:         requestCmt["content"].(string),
		IP:              requestCmt["ip"].(string),
		UserAgent:       requestCmt["ua"].(string),
		AuthorName:      requestCmt["authorName"].(string),
		AuthorURL:       requestCmt["authorURL"].(string),
		AuthorAvatarURL: requestCmt["authorAvatarURL"].(string),
	}

	if err := service.Comment.AddComment(comment); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}

	if err := service.Comment.UpdatePushedAt(comment); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

// addSymArticleAction adds an article come from Sym. Sees https://hacpai.com/article/1457158841475 for more details.
func addSymArticleAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "parses add article request failed"

		return
	}

	requestArticle := arg["article"].(map[string]interface{})
	articleId, err := strconv.ParseUint(requestArticle["oId"].(string), 10, 64)
	if nil != err {
		result.Code = -1
		result.Msg = "parses add article request failed"

		return
	}

	blogID := getBlogID(c)
	b3Key := requestArticle["userB3Key"].(string)
	blogAdmin := service.User.GetBlogAdmin(blogID)
	if b3Key != blogAdmin.B3Key {
		result.Code = -1
		result.Msg = "B3 key not match, ignored add article"

		return
	}

	article := &model.Article{
		BlogID:   blogID,
		AuthorID: blogAdmin.ID,
		Title:    requestArticle["articleTitle"].(string),
		Tags:     requestArticle["articleTags"].(string),
		Content:  requestArticle["articleContent"].(string),
		Commentable: true,
	}
	article.ID = articleId

	if err := service.Article.AddArticle(article); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}

	if err := service.Article.UpdatePushedAt(article); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

// updateSymArticleAction updates an article come from Sym. Sees https://hacpai.com/article/1457158841475 for more details.
func updateSymArticleAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "parses update article request failed"

		return
	}

	requestArticle := arg["article"].(map[string]interface{})
	articleId, err := strconv.ParseUint(requestArticle["oId"].(string), 10, 64)
	if nil != err {
		result.Code = -1
		result.Msg = "parses update article request failed"

		return
	}

	blogID := getBlogID(c)
	b3Key := requestArticle["userB3Key"].(string)
	blogAdmin := service.User.GetBlogAdmin(blogID)
	if b3Key != blogAdmin.B3Key {
		result.Code = -1
		result.Msg = "B3 key not match, ignored add article"

		return
	}

	article := service.Article.ConsoleGetArticle(articleId)
	if nil == article {
		result.Code = -1
		result.Msg = "not found article [ID=" + fmt.Sprintf("%d", articleId) + "] to update"

		return
	}

	article.Title = requestArticle["articleTitle"].(string)
	article.Tags = requestArticle["articleTags"].(string)
	article.Content = requestArticle["articleContent"].(string)

	if err := service.Article.UpdateArticle(article); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}

	if err := service.Article.UpdatePushedAt(article); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}
