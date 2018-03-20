// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2018, b3log.org
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

	requestCmt := arg["comment"].(map[string]interface{})
	articleId, err := strconv.ParseUint(requestCmt["articleId"].(string), 10, 64)
	if nil != err {
		result.Code = -1
		result.Msg = "parses add comment request failed"

		return
	}

	logger.Info("add a comment come from sym: %+v", arg)

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
}
