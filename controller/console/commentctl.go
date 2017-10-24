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

	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func GetCommentsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	sessionData := util.GetSession(c)
	commentModels, pagination := service.Comment.ConsoleGetComments(c.GetInt("p"), sessionData.BID)

	comments := []*ConsoleComment{}
	for _, commentModel := range commentModels {
		article := service.Article.ConsoleGetArticle(commentModel.ArticleID)
		if nil == article {
			log.Errorf("not found comment [id=%d]'s article", commentModel.ID)

			continue
		}
		articleAuthor := service.User.GetUser(article.AuthorID)
		if nil == articleAuthor {
			log.Errorf("not found article [id=%d]'s author", article.ID)

			continue
		}
		consoleArticleAuthor := &ConsoleAuthor{
			Name:      articleAuthor.Name,
			AvatarURL: articleAuthor.AvatarURL,
		}

		author := &ConsoleAuthor{
			Name:      commentModel.AuthorName,
			AvatarURL: commentModel.AuthorAvatarURL,
		}

		comment := &ConsoleComment{
			ID:            commentModel.ID,
			Author:        author,
			ArticleAuthor: consoleArticleAuthor,
			CreatedAt:     commentModel.CreatedAt.Format("2006-01-02"),
			Title:         article.Title,
			Content:       template.HTML(util.Markdown(commentModel.Content)),
			URL:           sessionData.BURL + "/todo comment path",
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
