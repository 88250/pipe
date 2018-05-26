package stm

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

// NewLocation returns created the Location's pointer
func NewLocation(opts *Options) *Location {
	loc := &Location{
		opts: opts,
	}
	return loc
}

// Location provides sitemap's path and filename on file systems
// and it provides proxy for Adapter interface also.
type Location struct {
	opts     *Options
	nmr      *Namer
	filename string
}

// Directory returns path to combine publicPath and sitemapsPath on file systems.
// It also indicates where sitemap files are.
func (loc *Location) Directory() string {
	return filepath.Join(
		loc.opts.publicPath,
		loc.opts.sitemapsPath,
	)
}

// Path returns path to combine publicPath, sitemapsPath and Filename on file systems.
// It also indicates where sitemap name is.
func (loc *Location) Path() string {
	return filepath.Join(
		loc.opts.publicPath,
		loc.opts.sitemapsPath,
		loc.Filename(),
	)
}

// PathInPublic returns path to combine sitemapsPath and Filename on website.
// It also indicates where url file path is.
func (loc *Location) PathInPublic() string {
	return filepath.Join(
		loc.opts.sitemapsPath,
		loc.Filename(),
	)
}

// URL returns path to combine SitemapsHost, sitemapsPath and
// Filename on website with it uses ResolveReference.
func (loc *Location) URL() string {
	base, _ := url.Parse(loc.opts.SitemapsHost())

	for _, ref := range []string{
		loc.opts.sitemapsPath + "/", loc.Filename(),
	} {
		base, _ = base.Parse(ref)
	}

	return base.String()
}

// Filesize returns file size this struct has.
func (loc *Location) Filesize() int64 {
	f, _ := os.Open(loc.Path())
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return 0
	}

	return fi.Size()
}

// reGzip determines gzip file.
var reGzip = regexp.MustCompile(`\.gz$`)

// Namer returns the Namer's pointer that Options struct has.
func (loc *Location) Namer() *Namer {
	return loc.opts.Namer()
}

// Filename returns sitemap filename.
func (loc *Location) Filename() string {
	nmr := loc.Namer()
	if loc.filename == "" && nmr == nil {
		log.Fatal("[F] No filename or namer set")
	}

	if loc.filename == "" {
		loc.filename = nmr.String()

		if !loc.opts.compress {
			newName := reGzip.ReplaceAllString(loc.filename, "")
			loc.filename = newName
		}
	}
	return loc.filename
}

// ReserveName returns that sets filename if this struct didn't keep filename and
// it returns reserved filename if this struct keeps filename also.
func (loc *Location) ReserveName() string {
	nmr := loc.Namer()
	if nmr != nil {
		loc.Filename()
		nmr.Next()
	}

	return loc.filename
}

// IsReservedName confirms that keeps filename on Location.filename.
func (loc *Location) IsReservedName() bool {
	if loc.filename == "" {
		return false
	}
	return true
}

// IsVerbose returns boolean about verbosed summary.
func (loc *Location) IsVerbose() bool {
	return loc.opts.verbose
}

// Write writes sitemap and index files that used from Adapter interface.
func (loc *Location) Write(data []byte, linkCount int) {

	loc.opts.adp.Write(loc, data)
	if !loc.IsVerbose() {
		return
	}

	output := loc.Summary(linkCount)
	if output != "" {
		println(output)
	}
}

// Summary outputs to generated file summary for console.
func (loc *Location) Summary(linkCount int) string {
	nmr := loc.Namer()
	if nmr.IsStart() {
		return ""
	}

	out := fmt.Sprintf("%s '%d' links",
		loc.PathInPublic(), linkCount)

	size := loc.Filesize()
	if size <= 0 {
		return out
	}

	return fmt.Sprintf("%s / %d bytes", out, size)
}
