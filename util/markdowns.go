// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2019, b3log.org & hacpai.com
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
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/bluele/gcache"
	"github.com/hackebrot/turtle"
	"github.com/microcosm-cc/bluemonday"
	"github.com/vinta/pangu"
	"gopkg.in/russross/blackfriday.v2"
)

var markdownCache = gcache.New(1024).LRU().Build()

// MarkdownResult represents markdown result.
type MarkdownResult struct {
	ContentHTML  string
	AbstractText string
	ThumbURL     string
}

var MarkedAvailable = false

// LoadMarkdown loads markdown process engine.
func LoadMarkdown() {
	request, err := http.NewRequest("POST", "http://localhost:8250", strings.NewReader("旧日的足迹"))
	if nil != err {
		logger.Info("[markdown-http] is not available, uses built-in [blackfriday] for markdown processing")

		return
	}
	http.DefaultClient.Timeout = 2 * time.Second
	response, err := http.DefaultClient.Do(request)
	if nil != err {
		logger.Info("[markdown-http] is not available, uses built-in [blackfriday] for markdown processing")

		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if nil != err {
		logger.Info("[markdown-http] is not available, uses built-in [blackfriday] for markdown processing")

		return
	}

	content := string(data)
	MarkedAvailable = "<p>旧日的足迹</p>\n" == content
	if MarkedAvailable {
		logger.Debug("[markdown-http] is available, uses it for markdown processing")
	} else {
		logger.Debug("[markdown-http] is not available, uses built-in [blackfriday] for markdown processing")
	}
}

func marked(mdText string) []byte {
	request, err := http.NewRequest("POST", "http://localhost:8250", strings.NewReader(mdText))
	if nil != err {
		logger.Info("marked failed: " + err.Error())

		return []byte("")
	}
	http.DefaultClient.Timeout = time.Second
	response, err := http.DefaultClient.Do(request)
	if nil != err {
		logger.Warnf("[markdown-http] failed [err=" + err.Error() + "], uses built-in [blackfriday] instead")

		return bf(mdText)
	}
	defer response.Body.Close()
	ret, err := ioutil.ReadAll(response.Body)
	if nil != err {
		logger.Info("marked failed: " + err.Error() + ", try to use built-in md engine instead")

		return bf(mdText)
	}

	return ret
}

func bf(mdText string) []byte {
	return blackfriday.Run([]byte(mdText))
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

	mdText = emojify(mdText)
	var unsafe []byte
	if MarkedAvailable {
		unsafe = marked(mdText)
	} else {
		unsafe = bf(mdText)
	}
	contentHTML := string(unsafe)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(contentHTML))
	doc.Find("img").Each(func(i int, ele *goquery.Selection) {
		src, _ := ele.Attr("src")
		ele.SetAttr("data-src", src)
		ele.RemoveAttr("src")
	})

	doc.Find("*").Contents().FilterFunction(func(i int, ele *goquery.Selection) bool {
		if "#text" != goquery.NodeName(ele) {
			return false
		}
		parent := goquery.NodeName(ele.Parent())

		return "span" != parent && "code" != parent && "pre" != parent
	}).Each(func(i int, ele *goquery.Selection) {
		text := ele.Text()
		text = pangu.SpacingText(text)
		ele.ReplaceWithHtml(text)
	})

	if !MarkedAvailable {
		doc.Find("code").Each(func(i int, ele *goquery.Selection) {
			code, err := ele.Html()
			if nil != err {
				logger.Errorf("get element [%+v]' HTML failed: %s", ele, err)
			} else {
				code = strings.Replace(code, "<", "&lt;", -1)
				code = strings.Replace(code, ">", "&gt;", -1)
				ele.SetHtml(code)
			}
		})
	}

	contentHTML, _ = doc.Find("body").Html()
	contentHTML = bluemonday.UGCPolicy().AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code").
		AllowAttrs("data-src").OnElements("img").
		AllowAttrs("class", "target", "id", "style").Globally().
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
	abstractText = pangu.SpacingText(abstractText)
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

var emojiRegx = regexp.MustCompile(":[\\w]+:")
var emojiImages = []string{"c.png", "d.png", "e50a.png", "f.png", "g.png", "huaji.gif", "doge.png", "i.png", "j.png", "k.png", "octocat.png", "r.png", "trollface.png", "u.png"}

func emojiImg(emojiASCII string) string {
	for _, img := range emojiImages {
		if emojiASCII == img[:strings.Index(img, ".")] {
			return "<img class=\"emoji\" src=\"https://cdn.jsdelivr.net/npm/vditor/src/assets/emoji/" + img + "\">"
		}
	}

	return ":" + emojiASCII + ":"
}

func emojify(text string) string {
	return emojiRegx.ReplaceAllStringFunc(text, func(emojiASCII string) string {
		emojiASCII = strings.Replace(emojiASCII, ":", "", -1)
		emoji := turtle.Emojis[emojiASCII]
		if nil == emoji {
			return emojiImg(emojiASCII)
		}

		return emoji.Char
	})
}
