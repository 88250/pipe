// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

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
