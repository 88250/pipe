package stm

import (
	"errors"
	"fmt"
	"time"

	"github.com/beevik/etree"
	"github.com/fatih/structs"
)

// URLModel is specific sample model for valuedate.
// http://www.sitemaps.org/protocol.html
// https://support.google.com/webmasters/answer/178636
type URLModel struct {
	Priority   float64                `valid:"float,length(0.0|1.0)"`
	Changefreq string                 `valid:"alpha(always|hourly|daily|weekly|monthly|yearly|never)"`
	Lastmod    time.Time              `valid:"-"`
	Expires    time.Time              `valid:"-"`
	Host       string                 `valid:"ipv4"`
	Loc        string                 `valid:"url"`
	Image      string                 `valid:"url"`
	Video      string                 `valid:"url"`
	Tag        string                 `valid:""`
	Geo        string                 `valid:""`
	News       string                 `valid:"-"`
	Mobile     bool                   `valid:"-"`
	Alternate  string                 `valid:"-"`
	Alternates map[string]interface{} `valid:"-"`
	Pagemap    map[string]interface{} `valid:"-"`
}

// fieldnames []string{"priority" "changefreq" "lastmod" "expires" "host" "images"
// "video" "geo" "news" "videos" "mobile" "alternate" "alternates" "pagemap"}
var fieldnames = ToLowerString(structs.Names(&URLModel{}))

// NewSitemapURL returns the created the SitemapURL's pointer
// and it validates URL types error.
func NewSitemapURL(opts *Options, url URL) (SitemapURL, error) {
	smu := &sitemapURL{opts: opts, data: url}
	err := smu.validate()
	return smu, err
}

// sitemapURL provides xml validator and xml builder.
type sitemapURL struct {
	opts *Options
	data URL
}

// validate is checking correct keys and checks the existence.
// TODO: Will create value's validator
func (su *sitemapURL) validate() error {
	var key string
	var invalid bool
	var locOk, hostOk bool

	for _, value := range su.data {
		key = value[0].(string)
		switch key {
		case "loc":
			locOk = true
		case "host":
			hostOk = true
		}

		invalid = true
		for _, name := range fieldnames {
			if key == name {
				invalid = false
				break
			}
		}
		if invalid {
			break
		}
	}

	if invalid {
		msg := fmt.Sprintf("Unknown map's key `%s` in URL type", key)
		return errors.New(msg)
	}
	if !locOk {
		msg := fmt.Sprintf("URL type must have `loc` map's key")
		return errors.New(msg)
	}
	if !hostOk {
		msg := fmt.Sprintf("URL type must have `host` map's key")
		return errors.New(msg)
	}
	return nil
}

// XML is building xml.
func (su *sitemapURL) XML() []byte {
	doc := etree.NewDocument()
	url := doc.CreateElement("url")

	SetBuilderElementValue(url, su.data.URLJoinBy("loc", "host", "loc"), "loc")
	if _, ok := SetBuilderElementValue(url, su.data, "lastmod"); !ok {
		lastmod := url.CreateElement("lastmod")
		lastmod.SetText(time.Now().Format(time.RFC3339))
	}
	if _, ok := SetBuilderElementValue(url, su.data, "changefreq"); !ok {
		changefreq := url.CreateElement("changefreq")
		changefreq.SetText("weekly")
	}
	if _, ok := SetBuilderElementValue(url, su.data, "priority"); !ok {
		priority := url.CreateElement("priority")
		priority.SetText("0.5")
	}
	SetBuilderElementValue(url, su.data, "expires")
	SetBuilderElementValue(url, su.data, "mobile")
	SetBuilderElementValue(url, su.data, "news")
	SetBuilderElementValue(url, su.data, "video")
	SetBuilderElementValue(url, su.data, "image")
	SetBuilderElementValue(url, su.data, "geo")

	if su.opts.pretty {
		doc.Indent(2)
	}
	buf := poolBuffer.Get()
	doc.WriteTo(buf)

	bytes := buf.Bytes()
	poolBuffer.Put(buf)

	return bytes
}
