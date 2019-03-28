package stm

import "bytes"

// NewBufferAdapter returns the created the BufferAdapter's pointer
func NewBufferAdapter() *BufferAdapter {
	adapter := &BufferAdapter{}
	return adapter
}

// BufferAdapter provides implementation for the Adapter interface.
type BufferAdapter struct {
	bufs []*bytes.Buffer // TODO: contains with filename
}

// Bytes gets written content.
func (adp *BufferAdapter) Bytes() [][]byte {
	bufs := make([][]byte, len(adp.bufs))

	for i, buf := range adp.bufs {
		bufs[i] = buf.Bytes()
	}
	return bufs
}

// Write will create sitemap xml file into the file systems.
func (adp *BufferAdapter) Write(loc *Location, data []byte) {
	adp.bufs = append(adp.bufs, bytes.NewBuffer(data))
}
