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

package console

import (
	"net/http"
	"strconv"
	"time"

	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

func BlogSwitchAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	blogID, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Code = -1

		return
	}

	session := util.GetSession(c)
	userID := session.UID

	userBlogs := service.User.GetUserBlogs(userID)
	if 1 > len(userBlogs) {
		result.Code = -1
		result.Msg = "switch blog failed"

		return
	}

	role := -1
	for _, userBlog := range userBlogs {
		if userBlog.ID == uint64(blogID) {
			role = userBlog.UserRole

			break
		}
	}

	if -1 == role {
		result.Code = -1
		result.Msg = "switch blog failed"

		return
	}

	result.Data = role

	session.URole = role
	session.BID = uint64(blogID)
	session.Save(c)
}

func CheckVersion(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	rhyResult := map[string]interface{}{}
	request := gorequest.New()
	_, _, errs := request.Get("http://rhythm.b3log.org/version/pipe/latest/"+util.Version).
		Set("user-agent", util.UserAgent).Timeout(30*time.Second).
		Retry(3, 5*time.Second, http.StatusInternalServerError).EndStruct(&rhyResult)
	if nil != errs {
		result.Code = -1
		result.Msg = errs[0].Error()

		return
	}

	data := map[string]interface{}{}
	data["version"] = rhyResult["pipeVersion"]
	data["download"] = rhyResult["pipeDownload"]
	result.Data = data
}
