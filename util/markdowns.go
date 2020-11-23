// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package util

import (
	"crypto/md5"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/88250/lute"
	"github.com/PuerkitoBio/goquery"
	"github.com/bluele/gcache"
	"github.com/microcosm-cc/bluemonday"
)

var markdownCache = gcache.New(1024).LRU().Expiration(30 * time.Minute).Build()

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
	contentHTML := luteEngine.MarkdownStr("", mdText)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(contentHTML))
	doc.Find("img").Each(func(i int, ele *goquery.Selection) {
		src, _ := ele.Attr("src")
		if Uploaded(src) && !strings.Contains(src, ".gif") && !strings.Contains(src, "imageView") {
			src += "?imageView2/2/w/1280/format/jpg/interlace/1/q/100"
		}
		ele.SetAttr("data-src", src)
		ele.RemoveAttr("src")
	})

	contentHTML, _ = doc.Find("body").Html()
	contentHTML = bluemonday.UGCPolicy().
		AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code").
		AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("div").
		AllowAttrs("data-code").OnElements("div").
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
		AllowElements("kbd").
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
