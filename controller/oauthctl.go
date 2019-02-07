// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2019, b3log.org & hacpai.com
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
	"fmt"
	"net/http"
	"strings"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

var states = map[string]string{}

// redirectGitHubLoginAction redirects to GitHub auth page.
func redirectGitHubLoginAction(c *gin.Context) {
	referer := c.Request.URL.Query().Get("referer")
	if "" == referer {
		referer = model.Conf.Server
	}
	if strings.HasSuffix(referer, "/") {
		referer = referer[:len(referer)-1]
	}
	state := util.RandString(16) + referer
	states[state] = state
	path := "https://github.com/login/oauth/authorize" + "?client_id=af7df3c80f26af88a8b3&state=" + state + "&scope=public_repo,user"

	c.Redirect(http.StatusSeeOther, path)
}

func githubCallbackAction(c *gin.Context) {
	state := c.Query("state")
	if _, exist := states[state]; !exist {
		c.Status(http.StatusForbidden)

		return
	}
	delete(states, state)

	referer := state[16:]

	accessToken := c.Query("ak")
	githubUser := util.GitHubUserInfo(accessToken)
	if nil == githubUser {
		c.Status(http.StatusForbidden)

		return
	}

	githubId := fmt.Sprintf("%v", githubUser["userId"])
	userName := githubUser["userName"].(string)
	user := service.User.GetUserByGitHubId(githubId)
	if nil == user {
		if !service.Init.Inited() {
			user = &model.User{
				Name:      userName,
				Password:  util.RandString(8),
				AvatarURL: githubUser["userAvatarURL"].(string),
				B3Key:     githubId,
				GithubId:  githubId,
			}

			if err := service.Init.InitPlatform(user); nil != err {
				logger.Errorf("init platform via github login failed: " + err.Error())
				c.Status(http.StatusInternalServerError)

				return
			}
		} else {
			user = service.User.GetUserByName(userName)
			if nil == user {
				if !model.Conf.OpenRegister {
					c.Status(http.StatusForbidden)

					return
				}

				user = &model.User{
					Name:      userName,
					Password:  util.RandString(8),
					AvatarURL: githubUser["userAvatarURL"].(string),
					B3Key:     githubId,
					GithubId:  githubId,
				}

				if err := service.Init.InitBlog(user); nil != err {
					logger.Errorf("init blog via github login failed: " + err.Error())
					c.Status(http.StatusInternalServerError)

					return
				}
			} else {
				user.GithubId = githubId
				user.B3Key = githubId
				service.User.UpdateUser(user)
			}
		}
	}

	ownBlog := service.User.GetOwnBlog(user.ID)
	session := &util.SessionData{
		UID:     user.ID,
		UName:   user.Name,
		UB3Key:  user.B3Key,
		UAvatar: user.AvatarURL,
		URole:   ownBlog.UserRole,
		BID:     ownBlog.ID,
		BURL:    ownBlog.URL,
	}
	if err := session.Save(c); nil != err {
		logger.Errorf("saves session failed: " + err.Error())
		c.Status(http.StatusInternalServerError)
	}

	c.Redirect(http.StatusSeeOther, referer)
}
