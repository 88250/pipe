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
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/b3log/gulu"
)

// Logger
var logger = gulu.Log.NewLogger(os.Stdout)

// RandAvatarData returns random avatar image byte array data from Gravatar (http://www.gravatar.com).
// Sees https://github.com/b3log/pipe/issues/131 for more details.
func RandAvatarData() (ret []byte) {
	modes := []string{"identicon", "monsterid", "wavatar", "retro", "robohash"}
	d := modes[rand.Intn(len(modes))]
	h := gulu.Rand.String(16)

	http.DefaultClient.Timeout = 2 * time.Second
	response, err := http.Get("http://www.gravatar.com/avatar/" + h + "?s=256&d=" + d)
	if nil != err {
		logger.Error("generate random avatar from Gavatar failed: " + err.Error())

		return nil
	}
	defer response.Body.Close()
	ret, err = ioutil.ReadAll(response.Body)
	if nil != err {
		logger.Error("generate random avatar from Gavatar failed: " + err.Error())

		return nil
	}

	return
}
