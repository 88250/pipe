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

package console

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

func GetCommentsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	commentModels, pagination := service.Comment.ConsoleGetComments(util.GetPage(c), session.BID)
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, session.BID)

	comments := []*ConsoleComment{}
	for _, commentModel := range commentModels {
		article := service.Article.ConsoleGetArticle(commentModel.ArticleID)
		articleAuthor := service.User.GetUser(article.AuthorID)
		consoleArticleAuthor := &ConsoleAuthor{
			URL:       blogURLSetting.Value + util.PathAuthors + "/" + articleAuthor.Name,
			Name:      articleAuthor.Name,
			AvatarURL: articleAuthor.AvatarURL,
		}

		commentAuthor := service.User.GetUser(commentModel.AuthorID)
		commentAuthorBlog := service.User.GetOwnBlog(commentModel.AuthorID)
		blogURLSetting = service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, commentAuthorBlog.ID)
		author := &ConsoleAuthor{
			URL:       blogURLSetting.Value + util.PathAuthors + "/" + commentAuthor.Name,
			Name:      commentAuthor.Name,
			AvatarURL: commentAuthor.AvatarURL,
		}

		mdResult := util.Markdown(commentModel.Content)
		comment := &ConsoleComment{
			ID:            commentModel.ID,
			Author:        author,
			ArticleAuthor: consoleArticleAuthor,
			CreatedAt:     commentModel.CreatedAt.Format("2006-01-02"),
			Title:         article.Title,
			Content:       template.HTML(mdResult.ContentHTML),
			URL:           session.BURL + "/todo comment path",
		}

		comments = append(comments, comment)
	}

	data := map[string]interface{}{}
	data["comments"] = comments
	data["pagination"] = pagination
	result.Data = data
}

func RemoveCommentAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.Atoi(idArg)
	if nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}

	if err := service.Comment.RemoveComment(uint(id)); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

func RemoveCommentsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "parses batch remove comments request failed"

		return
	}

	ids := arg["ids"].([]interface{})
	for _, id := range ids {
		if err := service.Comment.RemoveComment(uint(id.(float64))); nil != err {
			logger.Errorf("remove comment failed: " + err.Error())
		}
	}
}
