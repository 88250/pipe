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
	"crypto/tls"
	"net/http"
	"strconv"
	"time"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

// BlogSwitchAction switches blog.
func BlogSwitchAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	blogID, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Code = util.CodeErr

		return
	}

	session := util.GetSession(c)
	userID := session.UID

	userBlogs := service.User.GetUserBlogs(userID)
	if 1 > len(userBlogs) {
		result.Code = util.CodeErr
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
		result.Code = util.CodeErr
		result.Msg = "switch blog failed"

		return
	}

	result.Data = role

	session.URole = role
	session.BID = uint64(blogID)
	session.Save(c)
}

// CheckVersionAction checks version.
func CheckVersionAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	rhyResult := map[string]interface{}{}
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	_, _, errs := request.Get("https://rhythm.b3log.org/version/pipe/latest/"+model.Version).
		Set("User-Agent", model.UserAgent).Timeout(30*time.Second).
		Retry(3, 5*time.Second).EndStruct(&rhyResult)
	if nil != errs {
		result.Code = util.CodeErr
		result.Msg = errs[0].Error()

		return
	}

	data := map[string]interface{}{}
	data["version"] = rhyResult["pipeVersion"]
	data["download"] = rhyResult["pipeDownload"]
	result.Data = data
}
