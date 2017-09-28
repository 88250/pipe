// Solo.go - A small and beautiful blogging platform written in golang.
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

	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	NameOrEmail    string `json:"nameOrEmail" binding:"required"`
	PasswordHashed string `json:"passwordHashed" binding:"required"`
}

func loginCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	reqData := &loginRequest{}
	if err := c.BindJSON(reqData); nil != err {
		result.Code = -1
		result.Msg = "parses login request failed"

		return
	}

	user := service.User.GetUserByNameOrEmail(reqData.NameOrEmail)
	if nil == user {
		result.Code = -1
		result.Msg = "login failed"

		return
	}

	if user.Password != reqData.PasswordHashed {
		result.Code = -1
		result.Msg = "login failed"

		return
	}

	session := sessions.Default(c)
	session.Set("id", user.ID)
	session.Set("name", user.Name)
	session.Set("role", user.Role)
	session.Save()
}

func logoutCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := sessions.Default(c)
	session.Options(sessions.Options{
		MaxAge: -1,
	})
	session.Save()
}
