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
	"crypto/md5"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/88250/lute"
	"github.com/bluele/gcache"
	"github.com/microcosm-cc/bluemonday"
)

var markdownCache = gcache.New(1024).LRU().Build()

// MarkdownResult represents markdown result.
type MarkdownResult struct {
	ContentHTML  string
	AbstractText string
	ThumbURL     string
}

// Markdown process the specified markdown text to HTML.
func Markdown(mdText string) *MarkdownResult {
	mdText = strings.Replace(mdText, "\r\n", "\n", -1)

	digest := md5.New()
	digest.Write([]byte(mdText))
	key := string(digest.Sum(nil))

	cached, err := markdownCache.Get(key)
	if nil == err {
		return cached.(*MarkdownResult)
	}

	luteEngine := lute.New()
	unsafe, err := luteEngine.MarkdownStr("", mdText)
	if nil != err {
		return &MarkdownResult{
			ContentHTML:  err.Error(),
			AbstractText: err.Error(),
			ThumbURL:     "",
		}
	}
	contentHTML := unsafe
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(contentHTML))
	doc.Find("img").Each(func(i int, ele *goquery.Selection) {
		src, _ := ele.Attr("src")
		ele.SetAttr("data-src", src)
		ele.RemoveAttr("src")
	})

	contentHTML, _ = doc.Find("body").Html()
	contentHTML = bluemonday.UGCPolicy().AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code").
		AllowAttrs("data-src").OnElements("img").
		AllowAttrs("class", "target", "id", "style", "align").Globally().
		AllowAttrs("src", "width", "height", "border", "marginwidth", "marginheight").OnElements("iframe").
		AllowAttrs("controls", "src").OnElements("audio").
		AllowAttrs("color").OnElements("font").
		AllowAttrs("controls", "src", "width", "height").OnElements("video").
		AllowAttrs("src", "media", "type").OnElements("source").
		AllowAttrs("width", "height", "data", "type").OnElements("object").
		AllowAttrs("name", "value").OnElements("param").
		AllowAttrs("src", "type", "width", "height", "wmode", "allowNetworking").OnElements("embed").
		Sanitize(contentHTML)

	text := doc.Text()
	var runes []rune
	for i, w := 0, 0; i < len(text); i += w {
		runeValue, width := utf8.DecodeRuneInString(text[i:])
		w = width

		if unicode.IsSpace(runeValue) {
			continue
		}

		runes = append(runes, runeValue)
		if 200 < len(runes) {
			break
		}
	}

	selection := doc.Find("img").First()
	thumbnailURL, _ := selection.Attr("src")
	if "" == thumbnailURL {
		thumbnailURL, _ = selection.Attr("data-src")
	}
	abstractText := strings.TrimSpace(runesToString(runes))
	abstractText = luteEngine.Space(abstractText)
	abstractText = strings.Replace(abstractText, "<", "&lt;", -1)
	abstractText = strings.Replace(abstractText, ">", "&gt;", -1)

	ret := &MarkdownResult{
		ContentHTML:  contentHTML,
		AbstractText: abstractText,
		ThumbURL:     thumbnailURL,
	}
	markdownCache.Set(key, ret)

	return ret
}

func runesToString(runes []rune) (ret string) {
	for _, v := range runes {
		ret += string(v)
	}

	return
}
