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

package util

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

// GitHubUserInfo returns GitHub user info specified by the given access token.
func GitHubUserInfo(accessToken string) (ret map[string]interface{}) {
	result := map[string]interface{}{}
	response, data, errors := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		Get(HacPaiURL+"/github/user?ak="+accessToken).Timeout(7*time.Second).
		Set("User-Agent", "Pipe; +https://github.com/b3log/pipe").EndStruct(&result)
	if nil != errors || http.StatusOK != response.StatusCode {
		logger.Errorf("get github user info failed: %+v, %s", errors, data)

		return nil
	}

	if 0 != result["sc"].(float64) {
		return nil
	}

	return result["data"].(map[string]interface{})
}
