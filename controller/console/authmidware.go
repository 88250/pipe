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
