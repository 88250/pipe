// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-present, b3log.org
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
	"net/http"

	"github.com/88250/gulu"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

// UpdateAccountAction updates an account.
func UpdateAccountAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = util.CodeErr
		result.Msg = "parses update account request failed"

		return
	}

	b3Key := arg["b3key"].(string)
	avatarURL := arg["avatarURL"].(string)

	session := util.GetSession(c)
	user := service.User.GetUserByName(session.UName)
	user.B3Key = b3Key
	user.AvatarURL = avatarURL
	if err := service.User.UpdateUser(user); nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()

		return
	}
	session.UB3Key = b3Key
	session.UAvatar = avatarURL
	session.Save(c)
}

// GetAccountAction gets an account.
func GetAccountAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	data := map[string]interface{}{}
	data["name"] = session.UName
	data["avatarURL"] = session.UAvatar
	data["b3Key"] = session.UB3Key

	result.Data = data
}
