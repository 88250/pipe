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
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

// LoginCheck checks login or not.
func LoginCheck(c *gin.Context) {
	session := util.GetSession(c)
	if 0 == session.UID {
		result := gulu.Ret.NewResult()
		result.Code = util.CodeAuthErr
		result.Msg = "unauthenticated request"
		c.AbortWithStatusJSON(http.StatusOK, result)

		return
	}

	c.Next()
}
