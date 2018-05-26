package stm

import (
	"compress/gzip"
	"log"
	"os"
)

// NewFileAdapter returns the created the FileAdapter's pointer
func NewFileAdapter() *FileAdapter {
	adapter := &FileAdapter{}
	return adapter
}

// FileAdapter provides implementation for the Adapter interface.
type FileAdapter struct{}

// Write will create sitemap xml file into the file systems.
func (adp *FileAdapter) Write(loc *Location, data []byte) {
	dir := loc.Directory()
	fi, err := os.Stat(dir)
	if err != nil {
		_ = os.MkdirAll(dir, 0755)
	} else if !fi.IsDir() {
		log.Fatalf("[F] %s should be a directory", dir)
	}

	file, _ := os.OpenFile(loc.Path(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	fi, err = file.Stat()
	if err != nil {
		log.Fatalf("[F] %s file not exists", loc.Path())
	} else if !fi.Mode().IsRegular() {
		log.Fatalf("[F] %s should be a filename", loc.Path())
	}

	if GzipPtn.MatchString(loc.Path()) {
		adp.gzip(file, data)
	} else {
		adp.plain(file, data)
	}
}

// gzip will create sitemap file as a gzip.
func (adp *FileAdapter) gzip(file *os.File, data []byte) {
	gz := gzip.NewWriter(file)
	defer gz.Close()
	gz.Write(data)
}

// plain will create uncompressed file.
func (adp *FileAdapter) plain(file *os.File, data []byte) {
	file.Write(data)
	defer file.Close()
}
