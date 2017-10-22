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

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

func addCommentAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	blogAdminVal, _ := c.Get("blogAdmin")
	blogAdmin := blogAdminVal.(*model.User)

	sessionData := util.GetSession(c)
	if nil == sessionData {
		result.Code = -1
		result.Msg = "Please login before comment"

		return
	}

	comment := &model.Comment{
		AuthorID: sessionData.UID,
		BlogID:   blogAdmin.BlogID,
	}
	if err := c.BindJSON(comment); nil != err {
		result.Code = -1
		result.Msg = "parses add comment request failed"

		return
	}

	if err := service.Comment.AddComment(comment); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}
