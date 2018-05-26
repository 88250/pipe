package stm

// NewOptions returns the created the Options's pointer
func NewOptions() *Options {
	// Default values
	return &Options{
		defaultHost:  "http://www.example.com",
		sitemapsHost: "", // http://s3.amazonaws.com/sitemap-generator/,
		publicPath:   "public/",
		sitemapsPath: "sitemaps/",
		filename:     "sitemap",
		verbose:      true,
		compress:     true,
		pretty:       false,
		adp:          NewFileAdapter(),
	}
}

// Options exists for the Sitemap struct.
type Options struct {
	defaultHost  string
	sitemapsHost string
	publicPath   string
	sitemapsPath string
	filename     string
	verbose      bool
	compress     bool
	pretty       bool
	adp          Adapter
	nmr          *Namer
	loc          *Location
}

// SetDefaultHost sets that arg from Sitemap.Finalize method
func (opts *Options) SetDefaultHost(host string) {
	opts.defaultHost = host
}

// SetSitemapsHost sets that arg from Sitemap.SetSitemapsHost method
func (opts *Options) SetSitemapsHost(host string) {
	opts.sitemapsHost = host
}

// SetSitemapsPath sets that arg from Sitemap.SetSitemapsPath method.
func (opts *Options) SetSitemapsPath(path string) {
	opts.sitemapsPath = path
}

// SetPublicPath sets that arg from Sitemap.SetPublicPath method
func (opts *Options) SetPublicPath(path string) {
	opts.publicPath = path
}

// SetFilename sets that arg from Sitemap.SetFilename method
func (opts *Options) SetFilename(filename string) {
	opts.filename = filename
}

// SetVerbose sets that arg from Sitemap.SetVerbose method
func (opts *Options) SetVerbose(verbose bool) {
	opts.verbose = verbose
}

// SetCompress sets that arg from Sitemap.SetCompress method
func (opts *Options) SetCompress(compress bool) {
	opts.compress = compress
}

// SetPretty option sets pretty option to Options struct which allows pretty formatting to output files.
func (opts *Options) SetPretty(pretty bool) {
	opts.pretty = pretty
}

// SetAdapter sets that arg from Sitemap.SetAdapter method
func (opts *Options) SetAdapter(adp Adapter) {
	opts.adp = adp
}

// SitemapsHost sets that arg from Sitemap.SitemapsHost method
func (opts *Options) SitemapsHost() string {
	if opts.sitemapsHost != "" {
		return opts.sitemapsHost
	}
	return opts.defaultHost
}

// Location returns the Location's pointer with
// set option to arguments for Builderfile struct.
func (opts *Options) Location() *Location {
	return NewLocation(opts)
}

// IndexLocation returns the Location's pointer with
// set option to arguments for BuilderIndexfile struct.
func (opts *Options) IndexLocation() *Location {
	o := opts.Clone()
	o.nmr = NewNamer(&NOpts{base: opts.filename})
	return NewLocation(o)
}

// Namer returns Namer's pointer cache. If didn't create that yet,
// It also returns created Namer's pointer.
func (opts *Options) Namer() *Namer {
	if opts.nmr == nil {
		opts.nmr = NewNamer(&NOpts{base: opts.filename, zero: 1, start: 2})
	}
	return opts.nmr
}

// Clone method returns it copied myself.
func (opts *Options) Clone() *Options {
	o := *opts
	return &o
}
