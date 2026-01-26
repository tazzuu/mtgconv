package core

// base object class for holding global configs to be passed throughout the program
// this should generally be set via the cli interface
// and maps back to cli settings
type Config struct {
	Debug          bool   // enable debug entrypoint for dev testing
	Verbose        bool   // enable verbose logging
	PrintVersion   bool   // print version and quit
	CompatibilityMode bool // enable compatibility mode for modified output
	Version        string // the current version of the program
	OutputFilename string // output file name
	OutputFormat   OutputFormat
	UserAgent      string // user agent token string for web requests
	UrlString      string // user supplied URL to query for decklist
	AutoFilename bool // automatically create an output filename
}
