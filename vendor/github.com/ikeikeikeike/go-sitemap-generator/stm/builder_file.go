package stm

import (
	"bytes"
	"log"
)

// builderFileError is implementation for the BuilderError interface.
type builderFileError struct {
	error
	full bool
}

// FullError returns true if a sitemap xml had been limit size.
func (e *builderFileError) FullError() bool {
	return e.full
}

// NewBuilderFile returns the created the BuilderFile's pointer
func NewBuilderFile(opts *Options, loc *Location) *BuilderFile {
	b := &BuilderFile{opts: opts, loc: loc}
	b.clear()
	return b
}

// BuilderFile provides implementation for the Builder interface.
type BuilderFile struct {
	opts    *Options
	loc     *Location
	content []byte
	linkcnt int
	newscnt int
}

// Add method joins old bytes with creates bytes by it calls from Sitemap.Add method.
func (b *BuilderFile) Add(url interface{}) BuilderError {
	u := MergeMap(url.(URL),
		URL{{"host", b.loc.opts.defaultHost}},
	)

	b.linkcnt++

	smu, err := NewSitemapURL(b.opts, u)
	if err != nil {
		log.Fatalf("[F] Sitemap: %s", err)
	}

	bytes := smu.XML()

	if !b.isFileCanFit(bytes) {
		return &builderFileError{error: err, full: true}
	}

	b.content = append(b.content, bytes...)

	return nil
}

// isFileCanFit checks bytes to bigger than consts values.
func (b *BuilderFile) isFileCanFit(bytes []byte) bool {
	r := len(append(b.content, bytes...)) < MaxSitemapFilesize
	r = r && b.linkcnt < MaxSitemapLinks
	return r && b.newscnt < MaxSitemapNews
}

// clear will initialize xml content.
func (b *BuilderFile) clear() {
	b.content = make([]byte, 0, MaxSitemapFilesize)
}

// Content will return pooled bytes on content attribute.
func (b *BuilderFile) Content() []byte {
	return b.content
}

// XMLContent will return an XML of the sitemap built
func (b *BuilderFile) XMLContent() []byte {
	c := bytes.Join(bytes.Fields(XMLHeader), []byte(" "))
	c = append(append(c, b.Content()...), XMLFooter...)

	return c
}

// Write will write pooled bytes with header and footer to
// Location path for output sitemap file.
func (b *BuilderFile) Write() {
	b.loc.ReserveName()

	c := b.XMLContent()

	b.loc.Write(c, b.linkcnt)
	b.clear()
}
