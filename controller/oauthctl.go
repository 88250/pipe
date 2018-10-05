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

package controller

import (
	"net/http"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

var states map[string]string

// redirectGitHubLoginAction redirects to GitHub auth page.
func redirectGitHubLoginAction(c *gin.Context) {
	state := model.Conf.Server
	states[state] = state
	path := "https://github.com/login/oauth/authorize" + "?client_id=af7df3c80f26af88a8b3&state=" + state + "&scope=public_repo,user"

	c.Redirect(http.StatusSeeOther, path)
}

func githubCallback(c *gin.Context) {
	state := c.Query("state")

	if _, exist := states[state]; !exist {
		c.Status(http.StatusForbidden)

		return
	}

	delete(states, state)

	accessToken := c.Query("ak")
	githubUser := util.GitHubUserInfo(accessToken)
	if nil == githubUser {
		c.Status(http.StatusForbidden)

		return
	}

	logger.Infof("TODO github oauth login for user [%+v]", githubUser)
}
