package stm

import (
	"fmt"
	"log"
)

// NewNamer returns created the Namer's pointer
func NewNamer(opts *NOpts) *Namer {
	if opts.extension == "" {
		opts.extension = ".xml.gz"
	}

	namer := &Namer{opts: opts}
	namer.Reset()
	return namer
}

// NOpts Namer's option
type NOpts struct {
	base      string // filename base
	zero      int
	extension string
	start     int
}

// Namer provides sitemap's filename as a file number counter.
type Namer struct {
	count int
	opts  *NOpts
}

// String returns that combines filename base and file extension.
func (n *Namer) String() string {
	ext := n.opts.extension
	if n.count == 0 {
		return fmt.Sprintf("%s%s", n.opts.base, ext)
	}
	return fmt.Sprintf("%s%d%s", n.opts.base, n.count, ext)
}

// Reset will initialize to zero value on Namer's counter.
func (n *Namer) Reset() {
	n.count = n.opts.zero
}

// IsStart confirms that this struct has zero value.
func (n *Namer) IsStart() bool {
	return n.count == n.opts.zero
}

// Next is going to go to next index for filename.
func (n *Namer) Next() *Namer {
	if n.IsStart() {
		n.count = n.opts.start
	} else {
		n.count++
	}
	return n
}

// Previous is going to go to previous index for filename.
func (n *Namer) Previous() *Namer {
	if n.IsStart() {
		log.Fatal("[F] Already at the start of the series")
	}
	if n.count <= n.opts.start {
		n.count = n.opts.zero
	} else {
		n.count--
	}
	return n
}
