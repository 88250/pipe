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
	loginAuthURL := "https://ld246.com/login?goto=" + referer + "/api/login/callback"
	state := gulu.Rand.String(16)
	states[state] = referer
	path := loginAuthURL + "&state=" + state + "&v=" + util.Version
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

	accessToken := c.Query("access_token")
	userInfo := util.HacPaiUserInfo(accessToken)

	userId := userInfo["userId"].(string)
	userName := userInfo["userName"].(string)
	avatar := userInfo["avatar"].(string)
	user := service.User.GetUserByGitHubId(userId)
	if nil == user {
		if !service.Init.Inited() {
			user = &model.User{
				Name:      userName,
				AvatarURL: avatar,
				B3Key:     userName,
				GithubId:  userId,
			}

			if err := service.Init.InitPlatform(user); nil != err {
				logger.Errorf("init platform via community login failed: " + err.Error())
				c.Status(http.StatusInternalServerError)
				return
			}
		} else {
			user = &model.User{
				Name:      userName,
				AvatarURL: avatar,
				B3Key:     userName,
				GithubId:  userId,
			}

			if err := service.Init.InitBlog(user); nil != err {
				logger.Errorf("init blog via community login failed: " + err.Error())
				c.Status(http.StatusInternalServerError)
				return
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
