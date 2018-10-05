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

package util

import (
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

// GitHubUserInfo returns GitHub user info specified by the given access token.
func GitHubUserInfo(accessToken string) (ret map[string]interface{}) {
	response, data, errors := gorequest.New().Get("https://api.github.com/user?access_token=" + accessToken).Timeout(7 * time.Second).
		Set("User-Agent", "Pipe; +https://github.com/b3log/pipe").EndStruct(&ret)
	if nil != errors || http.StatusOK != response.StatusCode {
		logger.Errorf("get github user info failed: %+v, %s", errors, data)

		return nil
	}

	return
}
