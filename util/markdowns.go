// Solo.go - A small and beautiful blogging platform written in golang.
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
	"regexp"
	"strings"

	"github.com/bluele/gcache"
	"github.com/hackebrot/turtle"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	log "github.com/sirupsen/logrus"
)

var markdownCache = gcache.New(1024).LRU().Build()

func Markdown(mdText string) string {
	mdText = emojify(mdText)
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

var emojiRegx = regexp.MustCompile(":[a-z_]+:")

func emojify(text string) string {
	return emojiRegx.ReplaceAllStringFunc(text, func(emojiASCII string) string {
		emojiASCII = strings.Replace(emojiASCII, ":", "", -1)
		emoji := turtle.Emojis[emojiASCII]
		if nil == emoji {
			log.Warn("not found [" + emojiASCII + "]")

			return emojiASCII
		}

		return emoji.Char
	})
}
