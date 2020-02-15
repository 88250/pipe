// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package controller

import (
	"net/http"
	"net/url"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

var states = map[string]string{}

// redirectLoginAction redirects to HacPai auth page.
func redirectLoginAction(c *gin.Context) {
	referer := c.Request.URL.Query().Get("referer")
	if "" == referer {
		referer = model.Conf.Server
	}
	u, err := url.Parse(referer)
	if nil != err {
		referer = model.Conf.Server
	} else {
		referer = u.Scheme + "://" + u.Host
	}
	loginAuthURL := "https://hacpai.com/login?goto=" + referer + "/api/login/callback"
	state := gulu.Rand.String(16)
	states[state] = referer
	path := loginAuthURL + "?state=" + state
	c.Redirect(http.StatusSeeOther, path)
}

func loginCallbackAction(c *gin.Context) {
	state := c.Query("state")
	referer := states[state]
	if "" == referer {
		c.Status(http.StatusBadRequest)

		return
	}
	delete(states, state)

	githubId := c.Query("userId")
	userName := c.Query("userName")
	avatar := c.Query("avatar")
	user := service.User.GetUserByGitHubId(githubId)
	if nil == user {
		if !service.Init.Inited() {
			user = &model.User{
				Name:      userName,
				AvatarURL: avatar,
				B3Key:     userName,
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
				user = &model.User{
					Name:      userName,
					AvatarURL: avatar,
					B3Key:     userName,
					GithubId:  githubId,
				}

				if err := service.Init.InitBlog(user); nil != err {
					logger.Errorf("init blog via github login failed: " + err.Error())
					c.Status(http.StatusInternalServerError)

					return
				}
			} else {
				user.GithubId = githubId
				user.AvatarURL = avatar
				if err := service.User.UpdateUser(user); nil != err {
					logger.Errorf("update user failed: " + err.Error())
					c.Status(http.StatusInternalServerError)

					return
				}
			}
		}
	} else {
		user.Name = userName
		user.AvatarURL = avatar
		if err := service.User.UpdateUser(user); nil != err {
			logger.Errorf("update user failed: " + err.Error())
			c.Status(http.StatusInternalServerError)

			return
		}
	}

	ownBlog := service.User.GetOwnBlog(user.ID)
	if nil == ownBlog {
		logger.Warnf("can not get user by name [" + userName + "]")
		c.Status(http.StatusNotFound)

		return
	}

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
