package stm

import (
	"time"

	"github.com/beevik/etree"
)

// NewSitemapIndexURL and NewSitemapURL are almost the same behavior.
func NewSitemapIndexURL(opts *Options, url URL) SitemapURL {
	return &sitemapIndexURL{opts: opts, data: url}
}

// sitemapIndexURL and sitemapURL are almost the same behavior.
type sitemapIndexURL struct {
	opts *Options
	data URL
}

// XML and sitemapIndexURL.XML are almost the same behavior.
func (su *sitemapIndexURL) XML() []byte {
	doc := etree.NewDocument()
	sitemap := doc.CreateElement("sitemap")

	SetBuilderElementValue(sitemap, su.data, "loc")

	if _, ok := SetBuilderElementValue(sitemap, su.data, "lastmod"); !ok {
		lastmod := sitemap.CreateElement("lastmod")
		lastmod.SetText(time.Now().Format(time.RFC3339))
	}

	if su.opts.pretty {
		doc.Indent(2)
	}
	buf := poolBuffer.Get()
	doc.WriteTo(buf)

	bytes := buf.Bytes()
	poolBuffer.Put(buf)

	return bytes
}
