package stm

import "regexp"

// GzipPtn determines gzip file.
var GzipPtn = regexp.MustCompile(".gz$")

// Adapter provides interface for writes some kind of sitemap.
type Adapter interface {
	Write(loc *Location, data []byte)
}
