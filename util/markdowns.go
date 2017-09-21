// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
// Copyright (C) 2017, b3log.org
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
	"crypto/md5"

	"github.com/bluele/gcache"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var markdownCache = gcache.New(1024).LRU().Build()

func Markdown(mdText string) string {
	mdTextBytes := []byte(mdText)

	digest := md5.New()
	digest.Write(mdTextBytes)
	key := string(digest.Sum(nil))

	ret, err := markdownCache.Get(key)
	if nil == err {
		return ret.(string)
	}

	unsafe := blackfriday.MarkdownCommon(mdTextBytes)
	ret = string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))

	markdownCache.Set(key, ret)

	return ret.(string)
}
