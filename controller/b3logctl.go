// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package controller

import (
	"net/http"
	"strconv"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

// addSymArticleAction adds an article come from Sym. Sees https://ld246.com/article/1457158841475 for more details.
func addSymArticleAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = util.CodeErr
		result.Msg = "parses add article request failed"

		return
	}

	logger.Infof("Add an article from Sym: %+v", arg)

	client := arg["client"].(map[string]interface{})
	b3Key := client["userB3Key"].(string)
	articleAuthorName := client["userName"].(string)
	articleAuthor := service.User.GetUserByName(articleAuthorName)
	if articleAuthor.B3Key != b3Key {
		result.Code = util.CodeErr
		result.Msg = "wrong B3 Key"

		return
	}

	requestArticle := arg["article"].(map[string]interface{})
	articleId, err := strconv.ParseUint(requestArticle["id"].(string), 10, 64)
	if nil != err {
		result.Code = util.CodeErr
		result.Msg = "parses add article request failed"

		return
	}

	article := service.Article.ConsoleGetArticle(articleId)
	if nil == article {
		blogID := getBlogID(c)
		article = &model.Article{
			BlogID:      blogID,
			AuthorID:    articleAuthor.ID,
			Title:       requestArticle["title"].(string),
			Tags:        requestArticle["tags"].(string),
			Content:     requestArticle["content"].(string),
			Commentable: true,
		}
		article.ID = articleId

		if err := service.Article.AddArticle(article); nil != err {
			result.Code = util.CodeErr
			result.Msg = err.Error()

			return
		}
	} else {
		article.Title = requestArticle["title"].(string)
		article.Tags = requestArticle["tags"].(string)
		article.Content = requestArticle["content"].(string)

		if err := service.Article.UpdateArticle(article); nil != err {
			result.Code = util.CodeErr
			result.Msg = err.Error()

			return
		}
	}

	if err := service.Article.UpdatePushedAt(article); nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()
	}
}
